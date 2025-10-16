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

// MarketStreamsChannel represents connection and handlers for channel 'marketStreams'
type MarketStreamsChannel struct {
	client       *Client
	isConnected  bool
	addrTemplate string
	mu           sync.RWMutex
	// handler maps keyed by message name
	msgHandlers  map[string]func(context.Context, []byte) error
}

// NewMarketStreamsChannel constructs a channel bound to a client
func NewMarketStreamsChannel(client *Client) *MarketStreamsChannel {
	return &MarketStreamsChannel{
		client:       client,
		addrTemplate: "/ws/{streamName}",
		msgHandlers:  make(map[string]func(context.Context, []byte) error),
	}
}

// Connect resolves the channel address and establishes a WebSocket connection
func (ch *MarketStreamsChannel) Connect(ctx context.Context, streamName string) error {
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
	ch.client.RegisterHandlers("marketStreams", ch.msgHandlers)
	ch.client.ensureReadLoop(ctx)
	ch.isConnected = true
	return nil
}

// Disconnect tears down channel handlers and cancels the client's read loop
func (ch *MarketStreamsChannel) Disconnect(ctx context.Context) error {
	// Stop the read loop and underlying connection early to avoid handler lock contention
	ch.client.StopReadLoop()
	// Wait for the read loop to exit or the context to cancel
	if err := ch.client.Wait(ctx); err != nil && err != context.Canceled { return err }
	// Remove handlers for this channel
	ch.client.handlersMu.Lock()
	delete(ch.client.handlers, "marketStreams")
	ch.client.handlersMu.Unlock()
	// Mark channel as disconnected
	ch.mu.Lock()
	ch.isConnected = false
	ch.mu.Unlock()
	return nil
}

// SubscribeToMarketStreams sends a message for operation 'subscribeToMarketStreams' on marketStreams
func (ch *MarketStreamsChannel) SubscribeToMarketStreams(ctx context.Context, req *models.SubscribeRequest, handler *func(context.Context, *models.SubscriptionResponse) error) error {
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
			var v models.SubscriptionResponse
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

// UnsubscribeFromMarketStreams sends a message for operation 'unsubscribeFromMarketStreams' on marketStreams
func (ch *MarketStreamsChannel) UnsubscribeFromMarketStreams(ctx context.Context, req *models.UnsubscribeRequest, handler *func(context.Context, *models.UnsubscriptionResponse) error) error {
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
			var v models.UnsubscriptionResponse
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

// ListSubscriptionsFromMarketStreams sends a message for operation 'listSubscriptionsFromMarketStreams' on marketStreams
func (ch *MarketStreamsChannel) ListSubscriptionsFromMarketStreams(ctx context.Context, req *models.ListSubscriptionsRequest, handler *func(context.Context, *models.ListSubscriptionsResponse) error) error {
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

// SetPropertyOnMarketStreams sends a message for operation 'setPropertyOnMarketStreams' on marketStreams
func (ch *MarketStreamsChannel) SetPropertyOnMarketStreams(ctx context.Context, req *models.SetPropertyRequest, handler *func(context.Context, *models.SetPropertyResponse) error) error {
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

// GetPropertyFromMarketStreams sends a message for operation 'getPropertyFromMarketStreams' on marketStreams
func (ch *MarketStreamsChannel) GetPropertyFromMarketStreams(ctx context.Context, req *models.GetPropertyRequest, handler *func(context.Context, *models.GetPropertyResponse) error) error {
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

// Patterns:
//   - option_pair
// Update speeds:
//   - 50ms
// HandleNewSymbolInfoEvent registers a handler for message 'New Symbol Info Event' on marketStreams
func (ch *MarketStreamsChannel) HandleNewSymbolInfoEvent(fn func(context.Context, *models.NewSymbolInfoEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:OPTION_PAIR"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		if v, ok := typ["e"].(string); !ok || v != "OPTION_PAIR" { return fmt.Errorf("unexpected event type") }
		var v models.NewSymbolInfoEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *MarketStreamsChannel) UnregisterNewSymbolInfoEvent() {
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
// HandleOpenInterestEvent registers a handler for message 'Open Interest Event' on marketStreams
func (ch *MarketStreamsChannel) HandleOpenInterestEvent(fn func(context.Context, *models.OpenInterestEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:openInterest:array"] = func(ctx context.Context, b []byte) error {

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		if v, ok := typ["e"].(string); !ok || v != "openInterest" { return fmt.Errorf("unexpected event type") }
		var v models.OpenInterestEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *MarketStreamsChannel) UnregisterOpenInterestEvent() {
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
// HandleMarkPriceEvent registers a handler for message 'Mark Price Event' on marketStreams
func (ch *MarketStreamsChannel) HandleMarkPriceEvent(fn func(context.Context, *models.MarkPriceEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:markPrice:array"] = func(ctx context.Context, b []byte) error {

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		if v, ok := typ["e"].(string); !ok || v != "markPrice" { return fmt.Errorf("unexpected event type") }
		var v models.MarkPriceEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *MarketStreamsChannel) UnregisterMarkPriceEvent() {
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
// HandleKlineEvent registers a handler for message 'Kline Event' on marketStreams
func (ch *MarketStreamsChannel) HandleKlineEvent(fn func(context.Context, *models.KlineEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:kline"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		if v, ok := typ["e"].(string); !ok || v != "kline" { return fmt.Errorf("unexpected event type") }
		var v models.KlineEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *MarketStreamsChannel) UnregisterKlineEvent() {
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
// HandleTickerByUnderlyingEvent registers a handler for message 'Ticker by Underlying Event' on marketStreams
func (ch *MarketStreamsChannel) HandleTickerByUnderlyingEvent(fn func(context.Context, *models.TickerByUnderlyingEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:24hrTicker:array"] = func(ctx context.Context, b []byte) error {

		var arr []json.RawMessage
		if err := json.Unmarshal(b, &arr); err != nil { return err }
		if len(arr) == 0 { return fmt.Errorf("empty array") }
		var typ map[string]interface{}
		if err := json.Unmarshal(arr[0], &typ); err != nil { return err }
		if v, ok := typ["e"].(string); !ok || v != "24hrTicker" { return fmt.Errorf("unexpected event type") }
		var v models.TickerByUnderlyingEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *MarketStreamsChannel) UnregisterTickerByUnderlyingEvent() {
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
// HandleIndexPriceEvent registers a handler for message 'Index Price Event' on marketStreams
func (ch *MarketStreamsChannel) HandleIndexPriceEvent(fn func(context.Context, *models.IndexPriceEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:index"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		if v, ok := typ["e"].(string); !ok || v != "index" { return fmt.Errorf("unexpected event type") }
		var v models.IndexPriceEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *MarketStreamsChannel) UnregisterIndexPriceEvent() {
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
// HandleTickerEvent registers a handler for message 'Ticker Event' on marketStreams
func (ch *MarketStreamsChannel) HandleTickerEvent(fn func(context.Context, *models.TickerEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:24hrTicker"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		if v, ok := typ["e"].(string); !ok || v != "24hrTicker" { return fmt.Errorf("unexpected event type") }
		var v models.TickerEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *MarketStreamsChannel) UnregisterTickerEvent() {
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
// HandleTradeEvent registers a handler for message 'Trade Event' on marketStreams
func (ch *MarketStreamsChannel) HandleTradeEvent(fn func(context.Context, *models.TradeEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:trade"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		if v, ok := typ["e"].(string); !ok || v != "trade" { return fmt.Errorf("unexpected event type") }
		var v models.TradeEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *MarketStreamsChannel) UnregisterTradeEvent() {
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
// HandlePartialDepthEvent registers a handler for message 'Partial Depth Event' on marketStreams
func (ch *MarketStreamsChannel) HandlePartialDepthEvent(fn func(context.Context, *models.PartialDepthEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:depth"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		if v, ok := typ["e"].(string); !ok || v != "depth" { return fmt.Errorf("unexpected event type") }
		var v models.PartialDepthEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *MarketStreamsChannel) UnregisterPartialDepthEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:depth")
	ch.client.handlersMu.Unlock()
}

// HandleErrorMessage registers a handler for message 'Error Message' on marketStreams
func (ch *MarketStreamsChannel) HandleErrorMessage(fn func(context.Context, *models.ErrorMessage) error) {
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

func (ch *MarketStreamsChannel) UnregisterErrorMessage() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "error")
	ch.client.handlersMu.Unlock()
}


