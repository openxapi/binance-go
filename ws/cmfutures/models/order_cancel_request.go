package models

import (
	"encoding/json"
)

// OrderCancelRequestParams contains the parameters for OrderCancelRequest
type OrderCancelRequestParams struct {
	OrderId *int64 `json:"orderId,omitempty"`
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"`
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// OrderCancelRequest - Send a order.cancel request
// Message name: Cancel Order (TRADE) Request
type OrderCancelRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderCancelRequestParams `json:"params,omitempty"`
}

// NewOrderCancelRequest creates a new OrderCancelRequest instance
func NewOrderCancelRequest() *OrderCancelRequest {
	return &OrderCancelRequest{}
}

// String returns string representation of OrderCancelRequest
func (s OrderCancelRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderCancelRequest) SetParams(value OrderCancelRequestParams) *OrderCancelRequest {
	r.Params = &value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetOrderId(value int64) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.OrderId = &value
	return r
}

// SetOrigClientOrderId sets the origClientOrderId parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetOrigClientOrderId(value string) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.OrigClientOrderId = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetRecvWindow(value int64) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetSymbol(value string) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetTimestamp(value int64) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


