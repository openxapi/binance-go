package models

import (
	"encoding/json"
)

// OrderCancelRequestParams contains the parameters for OrderCancelRequest
type OrderCancelRequestParams struct {
	ApiKey string `json:"apiKey"`
	// Supported values: ONLY_NEW - Cancel will succeed if the order status is NEW. ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED.
	CancelRestrictions *string `json:"cancelRestrictions,omitempty"`
	// New ID for the canceled order. Automatically generated if not sent
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	// Cancel order by orderId
	OrderId int64 `json:"orderId"`
	// Cancel order by clientOrderId
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// OrderCancelRequest - Send a order.cancel request
// Message name: Cancel order (TRADE) Request
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

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetApiKey(value string) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetCancelRestrictions sets the cancelRestrictions parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetCancelRestrictions(value string) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.CancelRestrictions = &value
	return r
}

// SetNewClientOrderId sets the newClientOrderId parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetNewClientOrderId(value string) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.NewClientOrderId = &value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetOrderId(value int64) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.OrderId = value
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

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderCancelRequest) SetSignature(value string) *OrderCancelRequest {
	if r.Params == nil {
		r.Params = &OrderCancelRequestParams{}
	}
	r.Params.Signature = value
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


