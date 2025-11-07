package cmfuturesstreams

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
	"net/url"
	"github.com/gorilla/websocket"
	"github.com/openxapi/binance-go/ws/cmfutures-streams/models"
)

// MarketStreamChannel represents connection and handlers for channel 'marketStream'
type MarketStreamChannel struct {
	client       *Client
	isConnected  bool
	addrTemplate string
	mu           sync.RWMutex
	// handler maps keyed by message name
	msgHandlers  map[string]func(context.Context, []byte) error
}

func (ch *MarketStreamChannel) setHandlerLocked(key string, fn func(context.Context, []byte) error) {
	if ch.msgHandlers == nil {
		ch.msgHandlers = make(map[string]func(context.Context, []byte) error)
	}
	if fn != nil {
		ch.msgHandlers[key] = fn
		return
	}
	delete(ch.msgHandlers, key)
}

func (ch *MarketStreamChannel) applyHandlers() {
	ch.mu.RLock()
	snapshot := make(map[string]func(context.Context, []byte) error, len(ch.msgHandlers))
	for k, v := range ch.msgHandlers {
		snapshot[k] = v
	}
	ch.mu.RUnlock()
	ch.client.RegisterHandlers("marketStream", snapshot)
}

// NewMarketStreamChannel constructs a channel bound to a client
func NewMarketStreamChannel(client *Client) *MarketStreamChannel {
	return &MarketStreamChannel{
		client:       client,
		addrTemplate: "/ws/{streamName}",
		msgHandlers:  make(map[string]func(context.Context, []byte) error),
	}
}

// Connect resolves the channel address and establishes a WebSocket connection
func (ch *MarketStreamChannel) Connect(ctx context.Context, streamName string) error {
	ch.mu.Lock()
	if ch.isConnected {
		ch.mu.Unlock()
		return fmt.Errorf("channel already connected")
	}
	templatePath := ch.addrTemplate
	ch.mu.Unlock()
	base := ch.client.GetCurrentURL()
	if base == "" {
		return fmt.Errorf("no active server configured")
	}
	path := templatePath
	path = strings.ReplaceAll(path, "{streamName}", streamName)
	if i := strings.Index(path, "?"); i >= 0 {
		basePath := path[:i]
		q := path[i+1:]
		parts := strings.Split(q, "&")
		kept := make([]string, 0, len(parts))
		for _, p := range parts {
			if p == "" { continue }
			kv := strings.SplitN(p, "=", 2)
			if len(kv) == 2 && kv[1] == "" { continue }
			kept = append(kept, p)
		}
		if len(kept) > 0 {
			path = basePath + "?" + strings.Join(kept, "&")
		} else {
			path = basePath
		}
	}
	baseClean := strings.TrimRight(base, "/")
	var query string
	if idx := strings.Index(path, "?"); idx >= 0 {
		query = path[idx:]
		path = path[:idx]
	}
	if path != "" {
		path = "/" + strings.Trim(path, "/")
		if path == "/" {
			path = ""
		}
	}
	full := baseClean + path + query
	ch.client.connMu.RLock()
	clConn := ch.client.conn
	ch.client.connMu.RUnlock()
	if clConn == nil {
		u, err := url.Parse(full)
		if err != nil { return fmt.Errorf("invalid URL: %w", err) }
		dialer := websocket.DefaultDialer
		dialer.HandshakeTimeout = 10 * time.Second
		conn, _, err := dialer.DialContext(ctx, u.String(), nil)
		if err != nil { return fmt.Errorf("websocket dial failed: %w", err) }
		ch.client.connMu.Lock()
		ch.client.conn = conn
		ch.client.isConnected = true
		ch.client.connMu.Unlock()
	}
	ch.mu.Lock()
	ch.isConnected = true
	ch.mu.Unlock()
	ch.applyHandlers()
	ch.client.ensureReadLoop(ctx)
	return nil
}

// Disconnect tears down channel handlers and cancels the client's read loop
func (ch *MarketStreamChannel) Disconnect(ctx context.Context) error {
	// Stop the read loop and underlying connection early to avoid handler lock contention
	ch.client.StopReadLoop()
	// Wait for the read loop to exit or the context to cancel
	if err := ch.client.Wait(ctx); err != nil && err != context.Canceled { return err }
	// Remove handlers for this channel
	ch.client.handlersMu.Lock()
	delete(ch.client.handlers, "marketStream")
	ch.client.handlersMu.Unlock()
	// Mark channel as disconnected
	ch.mu.Lock()
	ch.isConnected = false
	ch.mu.Unlock()
	return nil
}

// Subscribe sends a message for operation 'marketStreamSubscribe' on marketStream
func (ch *MarketStreamChannel) Subscribe(ctx context.Context, req *models.SubscribeRequest, handler *func(context.Context, *models.SubscribeResponse, error) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {
			var envelope map[string]json.RawMessage
			if err := json.Unmarshal(b, &envelope); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			if rawErr, ok := envelope["error"]; ok && len(rawErr) > 0 {
				var errMsg models.ErrorMessage
				if err := json.Unmarshal(b, &errMsg); err != nil { return err }
				return (*handler)(ctx, nil, &errMsg)
			}

			if _, ok := envelope["result"]; !ok { return fmt.Errorf("no result field") }
			var v models.SubscribeResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "SUBSCRIBE"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Unsubscribe sends a message for operation 'marketStreamUnsubscribe' on marketStream
func (ch *MarketStreamChannel) Unsubscribe(ctx context.Context, req *models.UnsubscribeRequest, handler *func(context.Context, *models.UnsubscribeResponse, error) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {
			var envelope map[string]json.RawMessage
			if err := json.Unmarshal(b, &envelope); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			if rawErr, ok := envelope["error"]; ok && len(rawErr) > 0 {
				var errMsg models.ErrorMessage
				if err := json.Unmarshal(b, &errMsg); err != nil { return err }
				return (*handler)(ctx, nil, &errMsg)
			}

			if _, ok := envelope["result"]; !ok { return fmt.Errorf("no result field") }
			var v models.UnsubscribeResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "UNSUBSCRIBE"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// ListSubscriptions sends a message for operation 'marketStreamListSubscriptions' on marketStream
func (ch *MarketStreamChannel) ListSubscriptions(ctx context.Context, req *models.ListSubscriptionsRequest, handler *func(context.Context, *models.ListSubscriptionsResponse, error) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {
			var envelope map[string]json.RawMessage
			if err := json.Unmarshal(b, &envelope); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			if rawErr, ok := envelope["error"]; ok && len(rawErr) > 0 {
				var errMsg models.ErrorMessage
				if err := json.Unmarshal(b, &errMsg); err != nil { return err }
				return (*handler)(ctx, nil, &errMsg)
			}

			if v, ok := envelope["result"]; !ok || len(v) == 0 || v[0] != '[' { return fmt.Errorf("not array result") }
			var v models.ListSubscriptionsResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "LIST_SUBSCRIPTIONS"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// SetProperty sends a message for operation 'marketStreamSetProperty' on marketStream
func (ch *MarketStreamChannel) SetProperty(ctx context.Context, req *models.SetPropertyRequest, handler *func(context.Context, *models.SetPropertyResponse, error) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {
			var envelope map[string]json.RawMessage
			if err := json.Unmarshal(b, &envelope); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			if rawErr, ok := envelope["error"]; ok && len(rawErr) > 0 {
				var errMsg models.ErrorMessage
				if err := json.Unmarshal(b, &errMsg); err != nil { return err }
				return (*handler)(ctx, nil, &errMsg)
			}

			if _, ok := envelope["result"]; !ok { return fmt.Errorf("no result field") }
			var v models.SetPropertyResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "SET_PROPERTY"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// GetProperty sends a message for operation 'marketStreamGetProperty' on marketStream
func (ch *MarketStreamChannel) GetProperty(ctx context.Context, req *models.GetPropertyRequest, handler *func(context.Context, *models.GetPropertyResponse, error) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {
			var envelope map[string]json.RawMessage
			if err := json.Unmarshal(b, &envelope); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			if rawErr, ok := envelope["error"]; ok && len(rawErr) > 0 {
				var errMsg models.ErrorMessage
				if err := json.Unmarshal(b, &errMsg); err != nil { return err }
				return (*handler)(ctx, nil, &errMsg)
			}

			if _, ok := envelope["result"]; !ok { return fmt.Errorf("no result field") }
			var v models.GetPropertyResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "GET_PROPERTY"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Patterns:
//   - {symbol}@aggTrade
// Examples:
//   - btcusd_perp@aggTrade
// Update speeds:
//   - 100ms
// HandleAggregateTradeEvent registers a handler for message 'Aggregate Trade Event' on marketStream
func (ch *MarketStreamChannel) HandleAggregateTradeEvent(fn func(context.Context, *models.AggregateTradeEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "aggTrade" { return fmt.Errorf("unexpected event type") }
		var v models.AggregateTradeEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:aggTrade", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterAggregateTradeEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:aggTrade", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {pair}@indexPrice
//   - {pair}@indexPrice@{indexPriceInterval}
// Examples:
//   - btcusd@indexPrice
//   - btcusd@indexPrice@1s
// Update speeds:
//   - 3000ms
//   - 1000ms
// HandleIndexPriceEvent registers a handler for message 'Index Price Event' on marketStream
func (ch *MarketStreamChannel) HandleIndexPriceEvent(fn func(context.Context, *models.IndexPriceEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "indexPriceUpdate" { return fmt.Errorf("unexpected event type") }
		var v models.IndexPriceEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:indexPriceUpdate", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterIndexPriceEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:indexPriceUpdate", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@markPrice
//   - {symbol}@markPrice@{markPriceInterval}
// Examples:
//   - btcusd_perp@markPrice
//   - btcusd_200626@markPrice@1s
// Update speeds:
//   - 3000ms
//   - 1000ms
// HandleMarkPriceEvent registers a handler for message 'Mark Price Event' on marketStream
func (ch *MarketStreamChannel) HandleMarkPriceEvent(fn func(context.Context, *models.MarkPriceEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "markPriceUpdate" { return fmt.Errorf("unexpected event type") }
		var v models.MarkPriceEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:markPriceUpdate", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterMarkPriceEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:markPriceUpdate", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - !markPrice@arr
//   - !markPrice@arr@{markPriceInterval}
// Examples:
//   - !markPrice@arr
//   - !markPrice@arr@3s
// Update speeds:
//   - 3000ms
//   - 1000ms
// HandleAllMarkPricesEvent registers a handler for message 'All Mark Prices Event' on marketStream
func (ch *MarketStreamChannel) HandleAllMarkPricesEvent(fn func(context.Context, *models.AllMarkPricesEvent) error) {
	if fn == nil { return }
	ch.mu.Lock()
	ch.setHandlerLocked("evt:markPriceUpdate:array", func(ctx context.Context, b []byte) error {
		var v models.AllMarkPricesEvent

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "markPriceUpdate" { return fmt.Errorf("unexpected event type") }
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	})
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterAllMarkPricesEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:markPriceUpdate:array", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@kline_{interval}
// Examples:
//   - btcusd_200626@kline_1m
// Update speeds:
//   - 250ms
// HandleKlineEvent registers a handler for message 'Kline Event' on marketStream
func (ch *MarketStreamChannel) HandleKlineEvent(fn func(context.Context, *models.KlineEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "kline" { return fmt.Errorf("unexpected event type") }
		var v models.KlineEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:kline", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterKlineEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:kline", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {pair}_{contractType}@continuousKline_{interval}
// Examples:
//   - btcusd_next_quarter@continuousKline_1m
// Update speeds:
//   - 250ms
// HandleContinuousKlineEvent registers a handler for message 'Continuous Kline Event' on marketStream
func (ch *MarketStreamChannel) HandleContinuousKlineEvent(fn func(context.Context, *models.ContinuousKlineEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "continuous_kline" { return fmt.Errorf("unexpected event type") }
		var v models.ContinuousKlineEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:continuous_kline", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterContinuousKlineEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:continuous_kline", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {pair}@indexPriceKline_{interval}
// Examples:
//   - btcusd@indexPriceKline_1h
// Update speeds:
//   - 250ms
// HandleIndexKlineEvent registers a handler for message 'Index Kline Event' on marketStream
func (ch *MarketStreamChannel) HandleIndexKlineEvent(fn func(context.Context, *models.IndexKlineEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "indexPrice_kline" { return fmt.Errorf("unexpected event type") }
		var v models.IndexKlineEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:indexPrice_kline", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterIndexKlineEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:indexPrice_kline", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@markPriceKline_{interval}
// Examples:
//   - btcusd_perp@markPriceKline_4h
// Update speeds:
//   - 250ms
// HandleMarkPriceKlineEvent registers a handler for message 'Mark Price Kline Event' on marketStream
func (ch *MarketStreamChannel) HandleMarkPriceKlineEvent(fn func(context.Context, *models.MarkPriceKlineEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "markPrice_kline" { return fmt.Errorf("unexpected event type") }
		var v models.MarkPriceKlineEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:markPrice_kline", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterMarkPriceKlineEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:markPrice_kline", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@miniTicker
// Examples:
//   - btcusd_perp@miniTicker
// Update speeds:
//   - 500ms
// HandleMiniTickerEvent registers a handler for message 'Mini Ticker Event' on marketStream
func (ch *MarketStreamChannel) HandleMiniTickerEvent(fn func(context.Context, *models.MiniTickerEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "24hrMiniTicker" { return fmt.Errorf("unexpected event type") }
		var v models.MiniTickerEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrMiniTicker", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterMiniTickerEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrMiniTicker", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - !miniTicker@arr
// Examples:
//   - !miniTicker@arr
// HandleAllMiniTickersEvent registers a handler for message 'All Mini Tickers Event' on marketStream
func (ch *MarketStreamChannel) HandleAllMiniTickersEvent(fn func(context.Context, *models.AllMiniTickersEvent) error) {
	if fn == nil { return }
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrMiniTicker:array", func(ctx context.Context, b []byte) error {
		var v models.AllMiniTickersEvent

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "24hrMiniTicker" { return fmt.Errorf("unexpected event type") }
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	})
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterAllMiniTickersEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrMiniTicker:array", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@ticker
// Examples:
//   - btcusd_perp@ticker
// Update speeds:
//   - 500ms
// HandleTickerEvent registers a handler for message 'Ticker Event' on marketStream
func (ch *MarketStreamChannel) HandleTickerEvent(fn func(context.Context, *models.TickerEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "24hrTicker" { return fmt.Errorf("unexpected event type") }
		var v models.TickerEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrTicker", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterTickerEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrTicker", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - !ticker@arr
// Examples:
//   - !ticker@arr
// HandleAllTickersEvent registers a handler for message 'All Tickers Event' on marketStream
func (ch *MarketStreamChannel) HandleAllTickersEvent(fn func(context.Context, *models.AllTickersEvent) error) {
	if fn == nil { return }
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrTicker:array", func(ctx context.Context, b []byte) error {
		var v models.AllTickersEvent

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "24hrTicker" { return fmt.Errorf("unexpected event type") }
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	})
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterAllTickersEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrTicker:array", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@bookTicker
// Examples:
//   - btcusd_perp@bookTicker
// HandleBookTickerEvent registers a handler for message 'Book Ticker Event' on marketStream
func (ch *MarketStreamChannel) HandleBookTickerEvent(fn func(context.Context, *models.BookTickerEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "bookTicker" { return fmt.Errorf("unexpected event type") }
		var v models.BookTickerEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:bookTicker", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterBookTickerEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:bookTicker", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - !bookTicker
// Examples:
//   - !bookTicker
// HandleAllBookTickersEvent registers a handler for message 'All Book Tickers Event' on marketStream
func (ch *MarketStreamChannel) HandleAllBookTickersEvent(fn func(context.Context, *models.AllBookTickersEvent) error) {
	if fn == nil { return }
	ch.mu.Lock()
	ch.setHandlerLocked("evt:bookTicker:array", func(ctx context.Context, b []byte) error {
		var v models.AllBookTickersEvent

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "bookTicker" { return fmt.Errorf("unexpected event type") }
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	})
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterAllBookTickersEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:bookTicker:array", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@forceOrder
// Examples:
//   - btcusd_perp@forceOrder
// Update speeds:
//   - 1000ms
// HandleLiquidationEvent registers a handler for message 'Liquidation Event' on marketStream
func (ch *MarketStreamChannel) HandleLiquidationEvent(fn func(context.Context, *models.LiquidationEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "forceOrder" { return fmt.Errorf("unexpected event type") }
		var v models.LiquidationEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:forceOrder", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterLiquidationEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:forceOrder", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - !forceOrder@arr
// Examples:
//   - !forceOrder@arr
// Update speeds:
//   - 1000ms
// HandleAllLiquidationsEvent registers a handler for message 'All Liquidations Event' on marketStream
func (ch *MarketStreamChannel) HandleAllLiquidationsEvent(fn func(context.Context, *models.AllLiquidationsEvent) error) {
	if fn == nil { return }
	ch.mu.Lock()
	ch.setHandlerLocked("evt:forceOrder:array", func(ctx context.Context, b []byte) error {
		var v models.AllLiquidationsEvent

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "forceOrder" { return fmt.Errorf("unexpected event type") }
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	})
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterAllLiquidationsEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:forceOrder:array", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - !contractInfo
// Examples:
//   - !contractInfo
// HandleContractInfoEvent registers a handler for message 'Contract Info Event' on marketStream
func (ch *MarketStreamChannel) HandleContractInfoEvent(fn func(context.Context, *models.ContractInfoEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "contractInfo" { return fmt.Errorf("unexpected event type") }
		var v models.ContractInfoEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:contractInfo", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterContractInfoEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:contractInfo", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@depth{levels}
//   - {symbol}@depth{levels}@{speed}
// Examples:
//   - btcusd_perp@depth5@100ms
// Update speeds:
//   - 100ms
//   - 250ms
//   - 500ms
// HandlePartialDepthEvent registers a handler for message 'Partial Depth Event' on marketStream
func (ch *MarketStreamChannel) HandlePartialDepthEvent(fn func(context.Context, *models.PartialDepthEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "depthUpdate" { return fmt.Errorf("unexpected event type") }
		var v models.PartialDepthEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:depthUpdate", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterPartialDepthEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:depthUpdate", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@depth
//   - {symbol}@depth@{speed}
// Examples:
//   - btcusd_perp@depth@100ms
// Update speeds:
//   - 100ms
//   - 250ms
//   - 500ms
// HandleDiffDepthEvent registers a handler for message 'Diff Depth Event' on marketStream
func (ch *MarketStreamChannel) HandleDiffDepthEvent(fn func(context.Context, *models.DiffDepthEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "depthUpdate" { return fmt.Errorf("unexpected event type") }
		var v models.DiffDepthEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:depthUpdate", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterDiffDepthEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:depthUpdate", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}


