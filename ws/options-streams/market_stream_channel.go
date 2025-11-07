package optionsstreams

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
	"net/url"
	"github.com/gorilla/websocket"
	"github.com/openxapi/binance-go/ws/options-streams/models"
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
//   - option_pair
// Update speeds:
//   - 50ms
// HandleNewSymbolInfoEvent registers a handler for message 'New Symbol Info Event' on marketStream
func (ch *MarketStreamChannel) HandleNewSymbolInfoEvent(fn func(context.Context, *models.NewSymbolInfoEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "option_pair" { return fmt.Errorf("unexpected event type") }
		var v models.NewSymbolInfoEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:option_pair", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterNewSymbolInfoEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:option_pair", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {underlyingAsset}@openInterest@{expirationDate}
// Examples:
//   - ETH@openInterest@221125
// Update speeds:
//   - 60s
// HandleOpenInterestEvent registers a handler for message 'Open Interest Event' on marketStream
func (ch *MarketStreamChannel) HandleOpenInterestEvent(fn func(context.Context, *models.OpenInterestEvent) error) {
	if fn == nil { return }
	ch.mu.Lock()
	ch.setHandlerLocked("evt:openInterest:array", func(ctx context.Context, b []byte) error {
		var v models.OpenInterestEvent

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "openInterest" { return fmt.Errorf("unexpected event type") }
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	})
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterOpenInterestEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:openInterest:array", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {underlyingAsset}@markPrice
// Examples:
//   - ETH@markPrice
// Update speeds:
//   - 1000ms
// HandleMarkPriceEvent registers a handler for message 'Mark Price Event' on marketStream
func (ch *MarketStreamChannel) HandleMarkPriceEvent(fn func(context.Context, *models.MarkPriceEvent) error) {
	if fn == nil { return }
	ch.mu.Lock()
	ch.setHandlerLocked("evt:markPrice:array", func(ctx context.Context, b []byte) error {
		var v models.MarkPriceEvent

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "markPrice" { return fmt.Errorf("unexpected event type") }
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	})
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterMarkPriceEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:markPrice:array", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@kline_{interval}
// Examples:
//   - BTC-200630-9000-P@kline_1m
// Update speeds:
//   - 1000ms
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
//   - {underlyingAsset}@ticker@{expirationDate}
// Examples:
//   - ETH@ticker@220930
// Update speeds:
//   - 1000ms
// HandleTickerByUnderlyingEvent registers a handler for message 'Ticker by Underlying Event' on marketStream
func (ch *MarketStreamChannel) HandleTickerByUnderlyingEvent(fn func(context.Context, *models.TickerByUnderlyingEvent) error) {
	if fn == nil { return }
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrTicker:array", func(ctx context.Context, b []byte) error {
		var v models.TickerByUnderlyingEvent

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

func (ch *MarketStreamChannel) UnregisterTickerByUnderlyingEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:24hrTicker:array", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@index
// Examples:
//   - ETHUSDT@index
// Update speeds:
//   - 1000ms
// HandleIndexPriceEvent registers a handler for message 'Index Price Event' on marketStream
func (ch *MarketStreamChannel) HandleIndexPriceEvent(fn func(context.Context, *models.IndexPriceEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "index" { return fmt.Errorf("unexpected event type") }
		var v models.IndexPriceEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:index", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterIndexPriceEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:index", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@ticker
// Examples:
//   - BTC-210630-9000-P@ticker
// Update speeds:
//   - 1000ms
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
//   - {symbol}@trade
//   - {underlyingAsset}@trade
// Examples:
//   - BTC-210630-9000-P@trade
//   - ETH@trade
// Update speeds:
//   - 50ms
// HandleTradeEvent registers a handler for message 'Trade Event' on marketStream
func (ch *MarketStreamChannel) HandleTradeEvent(fn func(context.Context, *models.TradeEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "trade" { return fmt.Errorf("unexpected event type") }
		var v models.TradeEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:trade", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterTradeEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:trade", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// Patterns:
//   - {symbol}@depth{levels}
//   - {symbol}@depth{levels}@{speed}
// Examples:
//   - BTC-210630-9000-P@depth10
//   - BTC-210630-9000-P@depth10@100ms
// Update speeds:
//   - 100ms
//   - 1000ms
//   - 500ms
// HandlePartialDepthEvent registers a handler for message 'Partial Depth Event' on marketStream
func (ch *MarketStreamChannel) HandlePartialDepthEvent(fn func(context.Context, *models.PartialDepthEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "depth" { return fmt.Errorf("unexpected event type") }
		var v models.PartialDepthEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:depth", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *MarketStreamChannel) UnregisterPartialDepthEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:depth", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}


