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

// UserDataStreamChannel represents connection and handlers for channel 'userDataStream'
type UserDataStreamChannel struct {
	client       *Client
	isConnected  bool
	addrTemplate string
	mu           sync.RWMutex
	// handler maps keyed by message name
	msgHandlers  map[string]func(context.Context, []byte) error
}

func (ch *UserDataStreamChannel) setHandlerLocked(key string, fn func(context.Context, []byte) error) {
	if ch.msgHandlers == nil {
		ch.msgHandlers = make(map[string]func(context.Context, []byte) error)
	}
	if fn != nil {
		ch.msgHandlers[key] = fn
		return
	}
	delete(ch.msgHandlers, key)
}

func (ch *UserDataStreamChannel) applyHandlers() {
	ch.mu.RLock()
	snapshot := make(map[string]func(context.Context, []byte) error, len(ch.msgHandlers))
	for k, v := range ch.msgHandlers {
		snapshot[k] = v
	}
	ch.mu.RUnlock()
	ch.client.RegisterHandlers("userDataStream", snapshot)
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

// HandleAccountUpdateEvent registers a handler for message 'Account Update Event' on userDataStream
func (ch *UserDataStreamChannel) HandleAccountUpdateEvent(fn func(context.Context, *models.AccountUpdateEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "ACCOUNT_UPDATE" { return fmt.Errorf("unexpected event type") }
		var v models.AccountUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:ACCOUNT_UPDATE", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *UserDataStreamChannel) UnregisterAccountUpdateEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:ACCOUNT_UPDATE", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// HandleOrderTradeUpdateEvent registers a handler for message 'Order Trade Update Event' on userDataStream
func (ch *UserDataStreamChannel) HandleOrderTradeUpdateEvent(fn func(context.Context, *models.OrderTradeUpdateEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "ORDER_TRADE_UPDATE" { return fmt.Errorf("unexpected event type") }
		var v models.OrderTradeUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:ORDER_TRADE_UPDATE", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *UserDataStreamChannel) UnregisterOrderTradeUpdateEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:ORDER_TRADE_UPDATE", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// HandleRiskLevelChangeEvent registers a handler for message 'Risk Level Change Event' on userDataStream
func (ch *UserDataStreamChannel) HandleRiskLevelChangeEvent(fn func(context.Context, *models.RiskLevelChangeEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "RISK_LEVEL_CHANGE" { return fmt.Errorf("unexpected event type") }
		var v models.RiskLevelChangeEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:RISK_LEVEL_CHANGE", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *UserDataStreamChannel) UnregisterRiskLevelChangeEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:RISK_LEVEL_CHANGE", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}


