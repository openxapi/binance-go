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

// CombinedMarketStreamChannel represents connection and handlers for channel 'combinedMarketStream'
type CombinedMarketStreamChannel struct {
	client       *Client
	isConnected  bool
	addrTemplate string
	mu           sync.RWMutex
	// handler maps keyed by message name
	msgHandlers  map[string]func(context.Context, []byte) error
}

// NewCombinedMarketStreamChannel constructs a channel bound to a client
func NewCombinedMarketStreamChannel(client *Client) *CombinedMarketStreamChannel {
	return &CombinedMarketStreamChannel{
		client:       client,
		addrTemplate: "/stream?streams={streams}",
		msgHandlers:  make(map[string]func(context.Context, []byte) error),
	}
}

// Connect resolves the channel address and establishes a WebSocket connection
func (ch *CombinedMarketStreamChannel) Connect(ctx context.Context, streams string) error {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	if ch.isConnected {
		return fmt.Errorf("channel already connected")
	}
	base := ch.client.serverManager.GetActiveServerURL()
	if base == "" {
		return fmt.Errorf("no active server configured")
	}
	path := ch.addrTemplate
	path = strings.ReplaceAll(path, "{streams}", streams)
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
	full := strings.TrimRight(base, "/") + path
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
	// register handlers and start shared read loop
	ch.client.RegisterHandlers("combinedMarketStream", ch.msgHandlers)
	ch.client.ensureReadLoop(ctx)
	ch.isConnected = true
	return nil
}

// Disconnect tears down channel handlers and cancels the client's read loop
func (ch *CombinedMarketStreamChannel) Disconnect(ctx context.Context) error {
	// Stop the read loop and underlying connection early to avoid handler lock contention
	ch.client.StopReadLoop()
	// Wait for the read loop to exit or the context to cancel
	if err := ch.client.Wait(ctx); err != nil && err != context.Canceled { return err }
	// Remove handlers for this channel
	ch.client.handlersMu.Lock()
	delete(ch.client.handlers, "combinedMarketStream")
	ch.client.handlersMu.Unlock()
	// Mark channel as disconnected
	ch.mu.Lock()
	ch.isConnected = false
	ch.mu.Unlock()
	return nil
}

// Subscribe sends a message for operation 'combinedMarketStreamSubscribe' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) Subscribe(ctx context.Context, req *models.SubscribeRequest, handler *func(context.Context, *models.SubscribeResponse) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {

		var probe map[string]json.RawMessage
		if err := json.Unmarshal(b, &probe); err != nil { return err }
		if _, ok := probe["result"]; !ok { return fmt.Errorf("no result field") }
			var v models.SubscribeResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			return (*handler)(ctx, &v)
		})
	}
	// Apply const constraints from schema
	req.Method = "SUBSCRIBE"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Unsubscribe sends a message for operation 'combinedMarketStreamUnsubscribe' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) Unsubscribe(ctx context.Context, req *models.UnsubscribeRequest, handler *func(context.Context, *models.UnsubscribeResponse) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {

		var probe map[string]json.RawMessage
		if err := json.Unmarshal(b, &probe); err != nil { return err }
		if _, ok := probe["result"]; !ok { return fmt.Errorf("no result field") }
			var v models.UnsubscribeResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			return (*handler)(ctx, &v)
		})
	}
	// Apply const constraints from schema
	req.Method = "UNSUBSCRIBE"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// ListSubscriptions sends a message for operation 'combinedMarketStreamListSubscriptions' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) ListSubscriptions(ctx context.Context, req *models.ListSubscriptionsRequest, handler *func(context.Context, *models.ListSubscriptionsResponse) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {

		var probe map[string]json.RawMessage
		if err := json.Unmarshal(b, &probe); err != nil { return err }
		if v, ok := probe["result"]; !ok || len(v) == 0 || v[0] != '[' { return fmt.Errorf("not array result") }
			var v models.ListSubscriptionsResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			return (*handler)(ctx, &v)
		})
	}
	// Apply const constraints from schema
	req.Method = "LIST_SUBSCRIPTIONS"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// SetProperty sends a message for operation 'combinedMarketStreamSetProperty' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) SetProperty(ctx context.Context, req *models.SetPropertyRequest, handler *func(context.Context, *models.SetPropertyResponse) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {

		var probe map[string]json.RawMessage
		if err := json.Unmarshal(b, &probe); err != nil { return err }
		if _, ok := probe["result"]; !ok { return fmt.Errorf("no result field") }
			var v models.SetPropertyResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			return (*handler)(ctx, &v)
		})
	}
	// Apply const constraints from schema
	req.Method = "SET_PROPERTY"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// GetProperty sends a message for operation 'combinedMarketStreamGetProperty' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) GetProperty(ctx context.Context, req *models.GetPropertyRequest, handler *func(context.Context, *models.GetPropertyResponse) error) error {
	ch.client.connMu.RLock()
	conn := ch.client.conn
	ch.client.connMu.RUnlock()
	if conn == nil { return fmt.Errorf("not connected") }
	if handler != nil && *handler != nil {
		idStr := fmt.Sprintf("%v", req.Id)
		ch.client.pendingByID.Store(idStr, func(ctx context.Context, b []byte) error {

		var probe map[string]interface{}
		if err := json.Unmarshal(b, &probe); err != nil { return err }
		if _, ok := probe["result"]; !ok { return fmt.Errorf("no result field") }
			var v models.GetPropertyResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			return (*handler)(ctx, &v)
		})
	}
	// Apply const constraints from schema
	req.Method = "GET_PROPERTY"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// HandleCombinedMarketStreamEvent registers a handler for message 'Combined Stream Data' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleCombinedMarketStreamEvent(fn func(context.Context, *models.CombinedMarketStreamEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["wrap:combined"] = func(ctx context.Context, b []byte) error {

		var probe map[string]json.RawMessage
		if err := json.Unmarshal(b, &probe); err != nil { return err }
		if _, ok := probe["stream"]; !ok { return fmt.Errorf("not combined wrapper") }
		if _, ok := probe["data"]; !ok { return fmt.Errorf("not combined wrapper") }
		var v models.CombinedMarketStreamEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterCombinedMarketStreamEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "wrap:combined")
	ch.client.handlersMu.Unlock()
}

// HandleErrorMessage registers a handler for message 'Error Message' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleErrorMessage(fn func(context.Context, *models.ErrorMessage) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["error"] = func(ctx context.Context, b []byte) error {

		var probe map[string]json.RawMessage
		if err := json.Unmarshal(b, &probe); err != nil { return err }
		if _, ok := probe["error"]; !ok { return fmt.Errorf("not error message") }
		var v models.ErrorMessage
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterErrorMessage() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "error")
	ch.client.handlersMu.Unlock()
}

// Patterns:
//   - option_pair
// Update speeds:
//   - 50ms
// HandleNewSymbolInfoEvent registers a handler for unwrapped event 'newSymbolInfoEvent' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleNewSymbolInfoEvent(fn func(context.Context, *models.NewSymbolInfoEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:OPTION_PAIR"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "OPTION_PAIR" { return fmt.Errorf("unexpected event type") }
		var v models.NewSymbolInfoEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterNewSymbolInfoEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:OPTION_PAIR")
	ch.client.handlersMu.Unlock()
}

// Patterns:
//   - {underlyingAsset}@openInterest@{expirationDate}
// Examples:
//   - ETH@openInterest@221125
// Update speeds:
//   - 60s
// HandleOpenInterestEvent registers a handler for unwrapped event 'openInterestEvent' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleOpenInterestEvent(fn func(context.Context, *models.OpenInterestEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:openInterest:array"] = func(ctx context.Context, b []byte) error {

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "openInterest" { return fmt.Errorf("unexpected event type") }
		var v models.OpenInterestEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterOpenInterestEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:openInterest:array")
	ch.client.handlersMu.Unlock()
}

// Patterns:
//   - {underlyingAsset}@markPrice
// Examples:
//   - ETH@markPrice
// Update speeds:
//   - 1000ms
// HandleMarkPriceEvent registers a handler for unwrapped event 'markPriceEvent' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleMarkPriceEvent(fn func(context.Context, *models.MarkPriceEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:markPrice:array"] = func(ctx context.Context, b []byte) error {

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "markPrice" { return fmt.Errorf("unexpected event type") }
		var v models.MarkPriceEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterMarkPriceEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:markPrice:array")
	ch.client.handlersMu.Unlock()
}

// Patterns:
//   - {symbol}@kline_{interval}
// Examples:
//   - BTC-200630-9000-P@kline_1m
// Update speeds:
//   - 1000ms
// HandleKlineEvent registers a handler for unwrapped event 'klineEvent' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleKlineEvent(fn func(context.Context, *models.KlineEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:kline"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "kline" { return fmt.Errorf("unexpected event type") }
		var v models.KlineEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterKlineEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:kline")
	ch.client.handlersMu.Unlock()
}

// Patterns:
//   - {underlyingAsset}@ticker@{expirationDate}
// Examples:
//   - ETH@ticker@220930
// Update speeds:
//   - 1000ms
// HandleTickerByUnderlyingEvent registers a handler for unwrapped event 'tickerByUnderlyingEvent' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleTickerByUnderlyingEvent(fn func(context.Context, *models.TickerByUnderlyingEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:24hrTicker:array"] = func(ctx context.Context, b []byte) error {

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "24hrTicker" { return fmt.Errorf("unexpected event type") }
		var v models.TickerByUnderlyingEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterTickerByUnderlyingEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:24hrTicker:array")
	ch.client.handlersMu.Unlock()
}

// Patterns:
//   - {symbol}@index
// Examples:
//   - ETHUSDT@index
// Update speeds:
//   - 1000ms
// HandleIndexPriceEvent registers a handler for unwrapped event 'indexPriceEvent' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleIndexPriceEvent(fn func(context.Context, *models.IndexPriceEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:index"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "index" { return fmt.Errorf("unexpected event type") }
		var v models.IndexPriceEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterIndexPriceEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:index")
	ch.client.handlersMu.Unlock()
}

// Patterns:
//   - {symbol}@ticker
// Examples:
//   - BTC-210630-9000-P@ticker
// Update speeds:
//   - 1000ms
// HandleTickerEvent registers a handler for unwrapped event 'tickerEvent' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleTickerEvent(fn func(context.Context, *models.TickerEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:24hrTicker"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "24hrTicker" { return fmt.Errorf("unexpected event type") }
		var v models.TickerEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterTickerEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:24hrTicker")
	ch.client.handlersMu.Unlock()
}

// Patterns:
//   - {symbol}@trade
//   - {underlyingAsset}@trade
// Examples:
//   - BTC-210630-9000-P@trade
//   - ETH@trade
// Update speeds:
//   - 50ms
// HandleTradeEvent registers a handler for unwrapped event 'tradeEvent' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandleTradeEvent(fn func(context.Context, *models.TradeEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:trade"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "trade" { return fmt.Errorf("unexpected event type") }
		var v models.TradeEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterTradeEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:trade")
	ch.client.handlersMu.Unlock()
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
// HandlePartialDepthEvent registers a handler for unwrapped event 'partialDepthEvent' on combinedMarketStream
func (ch *CombinedMarketStreamChannel) HandlePartialDepthEvent(fn func(context.Context, *models.PartialDepthEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:depth"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "depth" { return fmt.Errorf("unexpected event type") }
		var v models.PartialDepthEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *CombinedMarketStreamChannel) UnregisterPartialDepthEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:depth")
	ch.client.handlersMu.Unlock()
}


