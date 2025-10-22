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

// UserDataStreamChannel represents connection and handlers for channel 'userDataStream'
type UserDataStreamChannel struct {
	client       *Client
	isConnected  bool
	addrTemplate string
	mu           sync.RWMutex
	// handler maps keyed by message name
	msgHandlers  map[string]func(context.Context, []byte) error
}

// NewUserDataStreamChannel constructs a channel bound to a client
func NewUserDataStreamChannel(client *Client) *UserDataStreamChannel {
	return &UserDataStreamChannel{
		client:       client,
		addrTemplate: "/ws/{listenKey}",
		msgHandlers:  make(map[string]func(context.Context, []byte) error),
	}
}

// Connect resolves the channel address and establishes a WebSocket connection
func (ch *UserDataStreamChannel) Connect(ctx context.Context, listenKey string) error {
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
	path = strings.ReplaceAll(path, "{listenKey}", listenKey)
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
	ch.client.RegisterHandlers("userDataStream", ch.msgHandlers)
	ch.client.ensureReadLoop(ctx)
	ch.isConnected = true
	return nil
}

// Disconnect tears down channel handlers and cancels the client's read loop
func (ch *UserDataStreamChannel) Disconnect(ctx context.Context) error {
	// Stop the read loop and underlying connection early to avoid handler lock contention
	ch.client.StopReadLoop()
	// Wait for the read loop to exit or the context to cancel
	if err := ch.client.Wait(ctx); err != nil && err != context.Canceled { return err }
	// Remove handlers for this channel
	ch.client.handlersMu.Lock()
	delete(ch.client.handlers, "userDataStream")
	ch.client.handlersMu.Unlock()
	// Mark channel as disconnected
	ch.mu.Lock()
	ch.isConnected = false
	ch.mu.Unlock()
	return nil
}

// Start sends a message for operation 'userDataStreamStart' on userDataStream
func (ch *UserDataStreamChannel) Start(ctx context.Context, req *models.UserDataStreamStartRequest, handler *func(context.Context, *models.UserDataStreamStartResponse) error) error {
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
			var v models.UserDataStreamStartResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			return (*handler)(ctx, &v)
		})
	}
	// Apply const constraints from schema
	req.Method = "userDataStream.start"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Ping sends a message for operation 'userDataStreamPing' on userDataStream
func (ch *UserDataStreamChannel) Ping(ctx context.Context, req *models.UserDataStreamPingRequest, handler *func(context.Context, *models.UserDataStreamPingResponse) error) error {
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
			var v models.UserDataStreamPingResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			return (*handler)(ctx, &v)
		})
	}
	// Apply const constraints from schema
	req.Method = "userDataStream.ping"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Stop sends a message for operation 'userDataStreamStop' on userDataStream
func (ch *UserDataStreamChannel) Stop(ctx context.Context, req *models.UserDataStreamStopRequest, handler *func(context.Context, *models.UserDataStreamStopResponse) error) error {
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
			var v models.UserDataStreamStopResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			if handler == nil || *handler == nil { return nil }
			return (*handler)(ctx, &v)
		})
	}
	// Apply const constraints from schema
	req.Method = "userDataStream.stop"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// HandleListenKeyExpiredEvent registers a handler for message 'Listen Key Expired Event' on userDataStream
func (ch *UserDataStreamChannel) HandleListenKeyExpiredEvent(fn func(context.Context, *models.ListenKeyExpiredEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:listenKeyExpired"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "listenKeyExpired" { return fmt.Errorf("unexpected event type") }
		var v models.ListenKeyExpiredEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *UserDataStreamChannel) UnregisterListenKeyExpiredEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:listenKeyExpired")
	ch.client.handlersMu.Unlock()
}

// HandleAccountUpdateEvent registers a handler for message 'Account Update Event' on userDataStream
func (ch *UserDataStreamChannel) HandleAccountUpdateEvent(fn func(context.Context, *models.AccountUpdateEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:ACCOUNT_UPDATE"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "ACCOUNT_UPDATE" { return fmt.Errorf("unexpected event type") }
		var v models.AccountUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *UserDataStreamChannel) UnregisterAccountUpdateEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:ACCOUNT_UPDATE")
	ch.client.handlersMu.Unlock()
}

// HandleMarginCallEvent registers a handler for message 'Margin Call Event' on userDataStream
func (ch *UserDataStreamChannel) HandleMarginCallEvent(fn func(context.Context, *models.MarginCallEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:MARGIN_CALL"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "MARGIN_CALL" { return fmt.Errorf("unexpected event type") }
		var v models.MarginCallEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *UserDataStreamChannel) UnregisterMarginCallEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:MARGIN_CALL")
	ch.client.handlersMu.Unlock()
}

// HandleOrderTradeUpdateEvent registers a handler for message 'Order Trade Update Event' on userDataStream
func (ch *UserDataStreamChannel) HandleOrderTradeUpdateEvent(fn func(context.Context, *models.OrderTradeUpdateEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:ORDER_TRADE_UPDATE"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "ORDER_TRADE_UPDATE" { return fmt.Errorf("unexpected event type") }
		var v models.OrderTradeUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *UserDataStreamChannel) UnregisterOrderTradeUpdateEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:ORDER_TRADE_UPDATE")
	ch.client.handlersMu.Unlock()
}

// HandleAccountConfigUpdateEvent registers a handler for message 'Account Configuration Update Event' on userDataStream
func (ch *UserDataStreamChannel) HandleAccountConfigUpdateEvent(fn func(context.Context, *models.AccountConfigUpdateEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:ACCOUNT_CONFIG_UPDATE"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "ACCOUNT_CONFIG_UPDATE" { return fmt.Errorf("unexpected event type") }
		var v models.AccountConfigUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *UserDataStreamChannel) UnregisterAccountConfigUpdateEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:ACCOUNT_CONFIG_UPDATE")
	ch.client.handlersMu.Unlock()
}

// HandleStrategyUpdateEvent registers a handler for message 'Strategy Update Event' on userDataStream
func (ch *UserDataStreamChannel) HandleStrategyUpdateEvent(fn func(context.Context, *models.StrategyUpdateEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:STRATEGY_UPDATE"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "STRATEGY_UPDATE" { return fmt.Errorf("unexpected event type") }
		var v models.StrategyUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *UserDataStreamChannel) UnregisterStrategyUpdateEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:STRATEGY_UPDATE")
	ch.client.handlersMu.Unlock()
}

// HandleGridUpdateEvent registers a handler for message 'Grid Update Event' on userDataStream
func (ch *UserDataStreamChannel) HandleGridUpdateEvent(fn func(context.Context, *models.GridUpdateEvent) error) {
	if fn == nil { return }
	if ch.msgHandlers == nil { ch.msgHandlers = make(map[string]func(context.Context, []byte) error) }
	ch.client.handlersMu.Lock()
	ch.msgHandlers["evt:GRID_UPDATE"] = func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "GRID_UPDATE" { return fmt.Errorf("unexpected event type") }
		var v models.GridUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.client.handlersMu.Unlock()
}

func (ch *UserDataStreamChannel) UnregisterGridUpdateEvent() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "evt:GRID_UPDATE")
	ch.client.handlersMu.Unlock()
}

// HandleErrorMessage registers a handler for message 'Error Message' on userDataStream
func (ch *UserDataStreamChannel) HandleErrorMessage(fn func(context.Context, *models.ErrorMessage) error) {
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

func (ch *UserDataStreamChannel) UnregisterErrorMessage() {
	ch.client.handlersMu.Lock()
	delete(ch.msgHandlers, "error")
	ch.client.handlersMu.Unlock()
}


