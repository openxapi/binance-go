package spot

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
	"net/url"
	"github.com/gorilla/websocket"
	"github.com/openxapi/binance-go/ws/spot/models"
)

// SpotChannel represents connection and handlers for channel 'spot'
type SpotChannel struct {
	client       *Client
	isConnected  bool
	addrTemplate string
	mu           sync.RWMutex
	// handler maps keyed by message name
	msgHandlers  map[string]func(context.Context, []byte) error
}

func (ch *SpotChannel) setHandlerLocked(key string, fn func(context.Context, []byte) error) {
	if ch.msgHandlers == nil {
		ch.msgHandlers = make(map[string]func(context.Context, []byte) error)
	}
	if fn != nil {
		ch.msgHandlers[key] = fn
		return
	}
	delete(ch.msgHandlers, key)
}

func (ch *SpotChannel) applyHandlers() {
	ch.mu.RLock()
	snapshot := make(map[string]func(context.Context, []byte) error, len(ch.msgHandlers))
	for k, v := range ch.msgHandlers {
		snapshot[k] = v
	}
	ch.mu.RUnlock()
	ch.client.RegisterHandlers("spot", snapshot)
}

// NewSpotChannel constructs a channel bound to a client
func NewSpotChannel(client *Client) *SpotChannel {
	return &SpotChannel{
		client:       client,
		addrTemplate: "/",
		msgHandlers:  make(map[string]func(context.Context, []byte) error),
	}
}

// Connect resolves the channel address and establishes a WebSocket connection
func (ch *SpotChannel) Connect(ctx context.Context) error {
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
func (ch *SpotChannel) Disconnect(ctx context.Context) error {
	// Stop the read loop and underlying connection early to avoid handler lock contention
	ch.client.StopReadLoop()
	// Wait for the read loop to exit or the context to cancel
	if err := ch.client.Wait(ctx); err != nil && err != context.Canceled { return err }
	// Remove handlers for this channel
	ch.client.handlersMu.Lock()
	delete(ch.client.handlers, "spot")
	ch.client.handlersMu.Unlock()
	// Mark channel as disconnected
	ch.mu.Lock()
	ch.isConnected = false
	ch.mu.Unlock()
	return nil
}

// AccountCommission sends a message for operation 'accountCommission' on spot
func (ch *SpotChannel) AccountCommission(ctx context.Context, req *models.AccountCommissionRequest, handler *func(context.Context, *models.AccountCommissionResponse, error) error) error {
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
			var v models.AccountCommissionResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "account.commission"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// AccountRateLimitsOrders sends a message for operation 'accountRateLimitsOrders' on spot
func (ch *SpotChannel) AccountRateLimitsOrders(ctx context.Context, req *models.AccountRateLimitsOrdersRequest, handler *func(context.Context, *models.AccountRateLimitsOrdersResponse, error) error) error {
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
			var v models.AccountRateLimitsOrdersResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "account.rateLimits.orders"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// AccountStatus sends a message for operation 'accountStatus' on spot
func (ch *SpotChannel) AccountStatus(ctx context.Context, req *models.AccountStatusRequest, handler *func(context.Context, *models.AccountStatusResponse, error) error) error {
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

// AllOrderLists sends a message for operation 'allOrderLists' on spot
func (ch *SpotChannel) AllOrderLists(ctx context.Context, req *models.AllOrderListsRequest, handler *func(context.Context, *models.AllOrderListsResponse, error) error) error {
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
			var v models.AllOrderListsResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "allOrderLists"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// AllOrders sends a message for operation 'allOrders' on spot
func (ch *SpotChannel) AllOrders(ctx context.Context, req *models.AllOrdersRequest, handler *func(context.Context, *models.AllOrdersResponse, error) error) error {
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
			var v models.AllOrdersResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "allOrders"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// AvgPrice sends a message for operation 'avgPrice' on spot
func (ch *SpotChannel) AvgPrice(ctx context.Context, req *models.AvgPriceRequest, handler *func(context.Context, *models.AvgPriceResponse, error) error) error {
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
			var v models.AvgPriceResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "avgPrice"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Depth sends a message for operation 'depth' on spot
func (ch *SpotChannel) Depth(ctx context.Context, req *models.DepthRequest, handler *func(context.Context, *models.DepthResponse, error) error) error {
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

// ExchangeInfo sends a message for operation 'exchangeInfo' on spot
func (ch *SpotChannel) ExchangeInfo(ctx context.Context, req *models.ExchangeInfoRequest, handler *func(context.Context, *models.ExchangeInfoResponse, error) error) error {
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
			var v models.ExchangeInfoResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "exchangeInfo"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Klines sends a message for operation 'klines' on spot
func (ch *SpotChannel) Klines(ctx context.Context, req *models.KlinesRequest, handler *func(context.Context, *models.KlinesResponse, error) error) error {
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
			var v models.KlinesResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "klines"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// MyAllocations sends a message for operation 'myAllocations' on spot
func (ch *SpotChannel) MyAllocations(ctx context.Context, req *models.MyAllocationsRequest, handler *func(context.Context, *models.MyAllocationsResponse, error) error) error {
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
			var v models.MyAllocationsResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "myAllocations"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// MyFilters sends a message for operation 'myFilters' on spot
func (ch *SpotChannel) MyFilters(ctx context.Context, req *models.MyFiltersRequest, handler *func(context.Context, *models.MyFiltersResponse, error) error) error {
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
			var v models.MyFiltersResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "myFilters"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// MyPreventedMatches sends a message for operation 'myPreventedMatches' on spot
func (ch *SpotChannel) MyPreventedMatches(ctx context.Context, req *models.MyPreventedMatchesRequest, handler *func(context.Context, *models.MyPreventedMatchesResponse, error) error) error {
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
			var v models.MyPreventedMatchesResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "myPreventedMatches"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// MyTrades sends a message for operation 'myTrades' on spot
func (ch *SpotChannel) MyTrades(ctx context.Context, req *models.MyTradesRequest, handler *func(context.Context, *models.MyTradesResponse, error) error) error {
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
			var v models.MyTradesResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "myTrades"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OpenOrderListsStatus sends a message for operation 'openOrderListsStatus' on spot
func (ch *SpotChannel) OpenOrderListsStatus(ctx context.Context, req *models.OpenOrderListsStatusRequest, handler *func(context.Context, *models.OpenOrderListsStatusResponse, error) error) error {
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
			var v models.OpenOrderListsStatusResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "openOrderLists.status"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OpenOrdersCancelAll sends a message for operation 'openOrdersCancelAll' on spot
func (ch *SpotChannel) OpenOrdersCancelAll(ctx context.Context, req *models.OpenOrdersCancelAllRequest, handler *func(context.Context, *models.OpenOrdersCancelAllResponse, error) error) error {
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
			var v models.OpenOrdersCancelAllResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "openOrders.cancelAll"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OpenOrdersStatus sends a message for operation 'openOrdersStatus' on spot
func (ch *SpotChannel) OpenOrdersStatus(ctx context.Context, req *models.OpenOrdersStatusRequest, handler *func(context.Context, *models.OpenOrdersStatusResponse, error) error) error {
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
			var v models.OpenOrdersStatusResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "openOrders.status"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderAmendKeepPriority sends a message for operation 'orderAmendKeepPriority' on spot
func (ch *SpotChannel) OrderAmendKeepPriority(ctx context.Context, req *models.OrderAmendKeepPriorityRequest, handler *func(context.Context, *models.OrderAmendKeepPriorityResponse, error) error) error {
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
			var v models.OrderAmendKeepPriorityResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "order.amend.keepPriority"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderAmendments sends a message for operation 'orderAmendments' on spot
func (ch *SpotChannel) OrderAmendments(ctx context.Context, req *models.OrderAmendmentsRequest, handler *func(context.Context, *models.OrderAmendmentsResponse, error) error) error {
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
			var v models.OrderAmendmentsResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "order.amendments"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderCancel sends a message for operation 'orderCancel' on spot
func (ch *SpotChannel) OrderCancel(ctx context.Context, req *models.OrderCancelRequest, handler *func(context.Context, *models.OrderCancelResponse, error) error) error {
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

// OrderCancelReplace sends a message for operation 'orderCancelReplace' on spot
func (ch *SpotChannel) OrderCancelReplace(ctx context.Context, req *models.OrderCancelReplaceRequest, handler *func(context.Context, *models.OrderCancelReplaceResponse, error) error) error {
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
			var v models.OrderCancelReplaceResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "order.cancelReplace"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderListCancel sends a message for operation 'orderListCancel' on spot
func (ch *SpotChannel) OrderListCancel(ctx context.Context, req *models.OrderListCancelRequest, handler *func(context.Context, *models.OrderListCancelResponse, error) error) error {
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
			var v models.OrderListCancelResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "orderList.cancel"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderListPlace sends a message for operation 'orderListPlace' on spot
func (ch *SpotChannel) OrderListPlace(ctx context.Context, req *models.OrderListPlaceRequest, handler *func(context.Context, *models.OrderListPlaceResponse, error) error) error {
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
			var v models.OrderListPlaceResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "orderList.place"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderListPlaceOco sends a message for operation 'orderListPlaceOco' on spot
func (ch *SpotChannel) OrderListPlaceOco(ctx context.Context, req *models.OrderListPlaceOcoRequest, handler *func(context.Context, *models.OrderListPlaceOcoResponse, error) error) error {
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
			var v models.OrderListPlaceOcoResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "orderList.place.oco"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderListPlaceOto sends a message for operation 'orderListPlaceOto' on spot
func (ch *SpotChannel) OrderListPlaceOto(ctx context.Context, req *models.OrderListPlaceOtoRequest, handler *func(context.Context, *models.OrderListPlaceOtoResponse, error) error) error {
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
			var v models.OrderListPlaceOtoResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "orderList.place.oto"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderListPlaceOtoco sends a message for operation 'orderListPlaceOtoco' on spot
func (ch *SpotChannel) OrderListPlaceOtoco(ctx context.Context, req *models.OrderListPlaceOtocoRequest, handler *func(context.Context, *models.OrderListPlaceOtocoResponse, error) error) error {
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
			var v models.OrderListPlaceOtocoResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "orderList.place.otoco"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderListStatus sends a message for operation 'orderListStatus' on spot
func (ch *SpotChannel) OrderListStatus(ctx context.Context, req *models.OrderListStatusRequest, handler *func(context.Context, *models.OrderListStatusResponse, error) error) error {
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
			var v models.OrderListStatusResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "orderList.status"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// OrderPlace sends a message for operation 'orderPlace' on spot
func (ch *SpotChannel) OrderPlace(ctx context.Context, req *models.OrderPlaceRequest, handler *func(context.Context, *models.OrderPlaceResponse, error) error) error {
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

// OrderStatus sends a message for operation 'orderStatus' on spot
func (ch *SpotChannel) OrderStatus(ctx context.Context, req *models.OrderStatusRequest, handler *func(context.Context, *models.OrderStatusResponse, error) error) error {
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

// OrderTest sends a message for operation 'orderTest' on spot
func (ch *SpotChannel) OrderTest(ctx context.Context, req *models.OrderTestRequest, handler *func(context.Context, *models.OrderTestResponse, error) error) error {
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
			var v models.OrderTestResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "order.test"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Ping sends a message for operation 'ping' on spot
func (ch *SpotChannel) Ping(ctx context.Context, req *models.PingRequest, handler *func(context.Context, *models.PingResponse, error) error) error {
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
			var v models.PingResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "ping"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// SessionLogon sends a message for operation 'sessionLogon' on spot
func (ch *SpotChannel) SessionLogon(ctx context.Context, req *models.SessionLogonRequest, handler *func(context.Context, *models.SessionLogonResponse, error) error) error {
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

// SessionLogout sends a message for operation 'sessionLogout' on spot
func (ch *SpotChannel) SessionLogout(ctx context.Context, req *models.SessionLogoutRequest, handler *func(context.Context, *models.SessionLogoutResponse, error) error) error {
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

// SessionStatus sends a message for operation 'sessionStatus' on spot
func (ch *SpotChannel) SessionStatus(ctx context.Context, req *models.SessionStatusRequest, handler *func(context.Context, *models.SessionStatusResponse, error) error) error {
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

// SessionSubscriptions sends a message for operation 'sessionSubscriptions' on spot
func (ch *SpotChannel) SessionSubscriptions(ctx context.Context, req *models.SessionSubscriptionsRequest, handler *func(context.Context, *models.SessionSubscriptionsResponse, error) error) error {
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
			var v models.SessionSubscriptionsResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "session.subscriptions"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// SorOrderPlace sends a message for operation 'sorOrderPlace' on spot
func (ch *SpotChannel) SorOrderPlace(ctx context.Context, req *models.SorOrderPlaceRequest, handler *func(context.Context, *models.SorOrderPlaceResponse, error) error) error {
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
			var v models.SorOrderPlaceResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "sor.order.place"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// SorOrderTest sends a message for operation 'sorOrderTest' on spot
func (ch *SpotChannel) SorOrderTest(ctx context.Context, req *models.SorOrderTestRequest, handler *func(context.Context, *models.SorOrderTestResponse, error) error) error {
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
			var v models.SorOrderTestResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "sor.order.test"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Ticker sends a message for operation 'ticker' on spot
func (ch *SpotChannel) Ticker(ctx context.Context, req *models.TickerRequest, handler *func(context.Context, *models.TickerResponse, error) error) error {
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
			var v models.TickerResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "ticker"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Ticker24hr sends a message for operation 'ticker24hr' on spot
func (ch *SpotChannel) Ticker24hr(ctx context.Context, req *models.Ticker24hrRequest, handler *func(context.Context, *models.Ticker24hrResponse, error) error) error {
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
			var v models.Ticker24hrResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "ticker.24hr"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// TickerBook sends a message for operation 'tickerBook' on spot
func (ch *SpotChannel) TickerBook(ctx context.Context, req *models.TickerBookRequest, handler *func(context.Context, *models.TickerBookResponse, error) error) error {
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

// TickerPrice sends a message for operation 'tickerPrice' on spot
func (ch *SpotChannel) TickerPrice(ctx context.Context, req *models.TickerPriceRequest, handler *func(context.Context, *models.TickerPriceResponse, error) error) error {
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

// TickerTradingDay sends a message for operation 'tickerTradingDay' on spot
func (ch *SpotChannel) TickerTradingDay(ctx context.Context, req *models.TickerTradingDayRequest, handler *func(context.Context, *models.TickerTradingDayResponse, error) error) error {
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
			var v models.TickerTradingDayResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "ticker.tradingDay"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// Time sends a message for operation 'time' on spot
func (ch *SpotChannel) Time(ctx context.Context, req *models.TimeRequest, handler *func(context.Context, *models.TimeResponse, error) error) error {
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
			var v models.TimeResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "time"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// TradesAggregate sends a message for operation 'tradesAggregate' on spot
func (ch *SpotChannel) TradesAggregate(ctx context.Context, req *models.TradesAggregateRequest, handler *func(context.Context, *models.TradesAggregateResponse, error) error) error {
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
			var v models.TradesAggregateResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "trades.aggregate"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// TradesHistorical sends a message for operation 'tradesHistorical' on spot
func (ch *SpotChannel) TradesHistorical(ctx context.Context, req *models.TradesHistoricalRequest, handler *func(context.Context, *models.TradesHistoricalResponse, error) error) error {
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
			var v models.TradesHistoricalResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "trades.historical"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// TradesRecent sends a message for operation 'tradesRecent' on spot
func (ch *SpotChannel) TradesRecent(ctx context.Context, req *models.TradesRecentRequest, handler *func(context.Context, *models.TradesRecentResponse, error) error) error {
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
			var v models.TradesRecentResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "trades.recent"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// UiKlines sends a message for operation 'uiKlines' on spot
func (ch *SpotChannel) UiKlines(ctx context.Context, req *models.UiKlinesRequest, handler *func(context.Context, *models.UiKlinesResponse, error) error) error {
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
			var v models.UiKlinesResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "uiKlines"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// UserDataStreamSubscribe sends a message for operation 'userDataStreamSubscribe' on spot
func (ch *SpotChannel) UserDataStreamSubscribe(ctx context.Context, req *models.UserDataStreamSubscribeRequest, handler *func(context.Context, *models.UserDataStreamSubscribeResponse, error) error) error {
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
			var v models.UserDataStreamSubscribeResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "userDataStream.subscribe"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// UserDataStreamSubscribeSignature sends a message for operation 'userDataStreamSubscribeSignature' on spot
func (ch *SpotChannel) UserDataStreamSubscribeSignature(ctx context.Context, req *models.UserDataStreamSubscribeSignatureRequest, handler *func(context.Context, *models.UserDataStreamSubscribeSignatureResponse, error) error) error {
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
			var v models.UserDataStreamSubscribeSignatureResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "userDataStream.subscribe.signature"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// UserDataStreamUnsubscribe sends a message for operation 'userDataStreamUnsubscribe' on spot
func (ch *SpotChannel) UserDataStreamUnsubscribe(ctx context.Context, req *models.UserDataStreamUnsubscribeRequest, handler *func(context.Context, *models.UserDataStreamUnsubscribeResponse, error) error) error {
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
			var v models.UserDataStreamUnsubscribeResponse
			if err := json.Unmarshal(b, &v); err != nil { return err }
			return (*handler)(ctx, &v, nil)
		})
	}
	// Apply const constraints from schema
	req.Method = "userDataStream.unsubscribe"
	data, err := json.Marshal(req)
	if err != nil { return fmt.Errorf("marshal request: %w", err) }
	return conn.WriteMessage(websocket.TextMessage, data)
}

// HandleAccountUpdateEvent registers a handler for message 'Outbound Account Position Event' on spot
func (ch *SpotChannel) HandleAccountUpdateEvent(fn func(context.Context, *models.AccountUpdateEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "outboundAccountPosition" { return fmt.Errorf("unexpected event type") }
		var v models.AccountUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:outboundAccountPosition", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *SpotChannel) UnregisterAccountUpdateEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:outboundAccountPosition", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// HandleBalanceUpdateEvent registers a handler for message 'Balance Update Event' on spot
func (ch *SpotChannel) HandleBalanceUpdateEvent(fn func(context.Context, *models.BalanceUpdateEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "balanceUpdate" { return fmt.Errorf("unexpected event type") }
		var v models.BalanceUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:balanceUpdate", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *SpotChannel) UnregisterBalanceUpdateEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:balanceUpdate", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// HandleOrderUpdateEvent registers a handler for message 'Execution Report Event' on spot
func (ch *SpotChannel) HandleOrderUpdateEvent(fn func(context.Context, *models.OrderUpdateEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "executionReport" { return fmt.Errorf("unexpected event type") }
		var v models.OrderUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:executionReport", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *SpotChannel) UnregisterOrderUpdateEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:executionReport", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// HandleListStatusEvent registers a handler for message 'List Status Event' on spot
func (ch *SpotChannel) HandleListStatusEvent(fn func(context.Context, *models.ListStatusEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "listStatus" { return fmt.Errorf("unexpected event type") }
		var v models.ListStatusEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:listStatus", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *SpotChannel) UnregisterListStatusEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:listStatus", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// HandleExternalLockUpdateEvent registers a handler for message 'External Lock Update Event' on spot
func (ch *SpotChannel) HandleExternalLockUpdateEvent(fn func(context.Context, *models.ExternalLockUpdateEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "externalLockUpdate" { return fmt.Errorf("unexpected event type") }
		var v models.ExternalLockUpdateEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:externalLockUpdate", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *SpotChannel) UnregisterExternalLockUpdateEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:externalLockUpdate", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

// HandleEventStreamTerminatedEvent registers a handler for message 'Event Stream Terminated Event' on spot
func (ch *SpotChannel) HandleEventStreamTerminatedEvent(fn func(context.Context, *models.EventStreamTerminatedEvent) error) {
	if fn == nil { return }
	handler := func(ctx context.Context, b []byte) error {

		var typ map[string]interface{}
		if err := json.Unmarshal(b, &typ); err != nil { return err }
		var ev string
		if v, ok := typ["e"].(string); ok { ev = v } else if evobj, ok := typ["event"].(map[string]interface{}); ok { if vv, ok2 := evobj["e"].(string); ok2 { ev = vv } }
		if ev != "eventStreamTerminated" { return fmt.Errorf("unexpected event type") }
		var v models.EventStreamTerminatedEvent
		if err := json.Unmarshal(b, &v); err != nil { return err }
		return fn(ctx, &v)
	}
	ch.mu.Lock()
	ch.setHandlerLocked("evt:eventStreamTerminated", handler)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}

func (ch *SpotChannel) UnregisterEventStreamTerminatedEvent() {
	ch.mu.Lock()
	ch.setHandlerLocked("evt:eventStreamTerminated", nil)
	connected := ch.isConnected
	ch.mu.Unlock()
	if connected { ch.applyHandlers() }
}


