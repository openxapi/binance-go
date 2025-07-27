package models

import (
	"encoding/json"
)

// OrderStatusRequestParams contains the parameters for OrderStatusRequest
type OrderStatusRequestParams struct {
	OrderId *int64 `json:"orderId,omitempty"`
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"`
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// OrderStatusRequest - Send a order.status request
// Message name: Query Order (USER_DATA) Request
type OrderStatusRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderStatusRequestParams `json:"params,omitempty"`
}

// NewOrderStatusRequest creates a new OrderStatusRequest instance
func NewOrderStatusRequest() *OrderStatusRequest {
	return &OrderStatusRequest{}
}

// String returns string representation of OrderStatusRequest
func (s OrderStatusRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderStatusRequest) SetParams(value OrderStatusRequestParams) *OrderStatusRequest {
	r.Params = &value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *OrderStatusRequest) SetOrderId(value int64) *OrderStatusRequest {
	if r.Params == nil {
		r.Params = &OrderStatusRequestParams{}
	}
	r.Params.OrderId = &value
	return r
}

// SetOrigClientOrderId sets the origClientOrderId parameter and returns the struct for method chaining
func (r *OrderStatusRequest) SetOrigClientOrderId(value string) *OrderStatusRequest {
	if r.Params == nil {
		r.Params = &OrderStatusRequestParams{}
	}
	r.Params.OrigClientOrderId = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderStatusRequest) SetRecvWindow(value int64) *OrderStatusRequest {
	if r.Params == nil {
		r.Params = &OrderStatusRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderStatusRequest) SetSymbol(value string) *OrderStatusRequest {
	if r.Params == nil {
		r.Params = &OrderStatusRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderStatusRequest) SetTimestamp(value int64) *OrderStatusRequest {
	if r.Params == nil {
		r.Params = &OrderStatusRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


