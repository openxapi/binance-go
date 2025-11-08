package umfutures

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
	"net/url"
	"github.com/gorilla/websocket"
	"github.com/openxapi/binance-go/ws/umfutures/models"
)

// UmfuturesChannel represents connection and handlers for channel 'umfutures'
type UmfuturesChannel struct {
	client       *Client
	isConnected  bool
	addrTemplate string
	mu           sync.RWMutex
	// handler maps keyed by message name
	msgHandlers  map[string]func(context.Context, []byte) error
}

func (ch *UmfuturesChannel) setHandlerLocked(key string, fn func(context.Context, []byte) error) {
	if ch.msgHandlers == nil {
		ch.msgHandlers = make(map[string]func(context.Context, []byte) error)
	}
	if fn != nil {
		ch.msgHandlers[key] = fn
		return
	}
	delete(ch.msgHandlers, key)
}

func (ch *UmfuturesChannel) applyHandlers() {
	ch.mu.RLock()
	snapshot := make(map[string]func(context.Context, []byte) error, len(ch.msgHandlers))
	for k, v := range ch.msgHandlers {
		snapshot[k] = v
	}
	ch.mu.RUnlock()
	ch.client.RegisterHandlers("umfutures", snapshot)
}

// NewUmfuturesChannel constructs a channel bound to a client
func NewUmfuturesChannel(client *Client) *UmfuturesChannel {
	return &UmfuturesChannel{
		client:       client,
		addrTemplate: "/",
		msgHandlers:  make(map[string]func(context.Context, []byte) error),
	}
}

// Connect resolves the channel address and establishes a WebSocket connection
func (ch *UmfuturesChannel) Connect(ctx context.Context) error {
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
func (ch *UmfuturesChannel) Disconnect(ctx context.Context) error {
	// Stop the read loop and underlying connection early to avoid handler lock contention
	ch.client.StopReadLoop()
	// Wait for the read loop to exit or the context to cancel
	if err := ch.client.Wait(ctx); err != nil && err != context.Canceled { return err }
	// Remove handlers for this channel
	ch.client.handlersMu.Lock()
	delete(ch.client.handlers, "umfutures")
	ch.client.handlersMu.Unlock()
	// Mark channel as disconnected
	ch.mu.Lock()
	ch.isConnected = false
	ch.mu.Unlock()
	return nil
}

// AccountBalance sends a message for operation 'accountBalance' on umfutures
func (ch *UmfuturesChannel) AccountBalance(ctx context.Context, req *models.AccountBalanceRequest, handler *func(context.Context, *models.AccountBalanceResponse, error) error) error {
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
			var v models.AccountBalanceResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "account.balance"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// AccountPosition sends a message for operation 'accountPosition' on umfutures
func (ch *UmfuturesChannel) AccountPosition(ctx context.Context, req *models.AccountPositionRequest, handler *func(context.Context, *models.AccountPositionResponse, error) error) error {
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
			var v models.AccountPositionResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "account.position"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// AccountStatus sends a message for operation 'accountStatus' on umfutures
func (ch *UmfuturesChannel) AccountStatus(ctx context.Context, req *models.AccountStatusRequest, handler *func(context.Context, *models.AccountStatusResponse, error) error) error {
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
			var v models.AccountStatusResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "account.status"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Depth sends a message for operation 'depth' on umfutures
func (ch *UmfuturesChannel) Depth(ctx context.Context, req *models.DepthRequest, handler *func(context.Context, *models.DepthResponse, error) error) error {
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
			var v models.DepthResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "depth"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderCancel sends a message for operation 'orderCancel' on umfutures
func (ch *UmfuturesChannel) OrderCancel(ctx context.Context, req *models.OrderCancelRequest, handler *func(context.Context, *models.OrderCancelResponse, error) error) error {
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
			var v models.OrderCancelResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "order.cancel"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderModify sends a message for operation 'orderModify' on umfutures
func (ch *UmfuturesChannel) OrderModify(ctx context.Context, req *models.OrderModifyRequest, handler *func(context.Context, *models.OrderModifyResponse, error) error) error {
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
			var v models.OrderModifyResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "order.modify"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderPlace sends a message for operation 'orderPlace' on umfutures
func (ch *UmfuturesChannel) OrderPlace(ctx context.Context, req *models.OrderPlaceRequest, handler *func(context.Context, *models.OrderPlaceResponse, error) error) error {
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
			var v models.OrderPlaceResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "order.place"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderStatus sends a message for operation 'orderStatus' on umfutures
func (ch *UmfuturesChannel) OrderStatus(ctx context.Context, req *models.OrderStatusRequest, handler *func(context.Context, *models.OrderStatusResponse, error) error) error {
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
			var v models.OrderStatusResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "order.status"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// SessionLogon sends a message for operation 'sessionLogon' on umfutures
func (ch *UmfuturesChannel) SessionLogon(ctx context.Context, req *models.SessionLogonRequest, handler *func(context.Context, *models.SessionLogonResponse, error) error) error {
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
			var v models.SessionLogonResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "session.logon"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// SessionLogout sends a message for operation 'sessionLogout' on umfutures
func (ch *UmfuturesChannel) SessionLogout(ctx context.Context, req *models.SessionLogoutRequest, handler *func(context.Context, *models.SessionLogoutResponse, error) error) error {
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
			var v models.SessionLogoutResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "session.logout"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// SessionStatus sends a message for operation 'sessionStatus' on umfutures
func (ch *UmfuturesChannel) SessionStatus(ctx context.Context, req *models.SessionStatusRequest, handler *func(context.Context, *models.SessionStatusResponse, error) error) error {
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
			var v models.SessionStatusResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "session.status"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// TickerBook sends a message for operation 'tickerBook' on umfutures
func (ch *UmfuturesChannel) TickerBook(ctx context.Context, req *models.TickerBookRequest, handler *func(context.Context, *models.TickerBookResponse, error) error) error {
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
			var v models.TickerBookResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "ticker.book"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// TickerPrice sends a message for operation 'tickerPrice' on umfutures
func (ch *UmfuturesChannel) TickerPrice(ctx context.Context, req *models.TickerPriceRequest, handler *func(context.Context, *models.TickerPriceResponse, error) error) error {
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
			var v models.TickerPriceResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "ticker.price"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// V2AccountBalance sends a message for operation 'v2AccountBalance' on umfutures
func (ch *UmfuturesChannel) V2AccountBalance(ctx context.Context, req *models.V2AccountBalanceRequest, handler *func(context.Context, *models.V2AccountBalanceResponse, error) error) error {
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
			var v models.V2AccountBalanceResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "v2/account.balance"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// V2AccountPosition sends a message for operation 'v2AccountPosition' on umfutures
func (ch *UmfuturesChannel) V2AccountPosition(ctx context.Context, req *models.V2AccountPositionRequest, handler *func(context.Context, *models.V2AccountPositionResponse, error) error) error {
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
			var v models.V2AccountPositionResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "v2/account.position"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// V2AccountStatus sends a message for operation 'v2AccountStatus' on umfutures
func (ch *UmfuturesChannel) V2AccountStatus(ctx context.Context, req *models.V2AccountStatusRequest, handler *func(context.Context, *models.V2AccountStatusResponse, error) error) error {
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
			var v models.V2AccountStatusResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "v2/account.status"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}


