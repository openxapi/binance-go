package models

import (
	"encoding/json"
)

// OrderListCancelRequestParams contains the parameters for OrderListCancelRequest
type OrderListCancelRequestParams struct {
	ApiKey string `json:"apiKey"`
	// Cancel order list by listClientId
	ListClientOrderId *string `json:"listClientOrderId,omitempty"`
	// New ID for the canceled order list. Automatically generated if not sent
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	// Cancel order list by orderListId
	OrderListId int64 `json:"orderListId"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// OrderListCancelRequest - Send a orderList.cancel request
// Message name: Cancel Order list (TRADE) Request
type OrderListCancelRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderListCancelRequestParams `json:"params,omitempty"`
}

// NewOrderListCancelRequest creates a new OrderListCancelRequest instance
func NewOrderListCancelRequest() *OrderListCancelRequest {
	return &OrderListCancelRequest{}
}

// String returns string representation of OrderListCancelRequest
func (s OrderListCancelRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderListCancelRequest) SetParams(value OrderListCancelRequestParams) *OrderListCancelRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OrderListCancelRequest) SetApiKey(value string) *OrderListCancelRequest {
	if r.Params == nil {
		r.Params = &OrderListCancelRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetListClientOrderId sets the listClientOrderId parameter and returns the struct for method chaining
func (r *OrderListCancelRequest) SetListClientOrderId(value string) *OrderListCancelRequest {
	if r.Params == nil {
		r.Params = &OrderListCancelRequestParams{}
	}
	r.Params.ListClientOrderId = &value
	return r
}

// SetNewClientOrderId sets the newClientOrderId parameter and returns the struct for method chaining
func (r *OrderListCancelRequest) SetNewClientOrderId(value string) *OrderListCancelRequest {
	if r.Params == nil {
		r.Params = &OrderListCancelRequestParams{}
	}
	r.Params.NewClientOrderId = &value
	return r
}

// SetOrderListId sets the orderListId parameter and returns the struct for method chaining
func (r *OrderListCancelRequest) SetOrderListId(value int64) *OrderListCancelRequest {
	if r.Params == nil {
		r.Params = &OrderListCancelRequestParams{}
	}
	r.Params.OrderListId = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderListCancelRequest) SetRecvWindow(value int64) *OrderListCancelRequest {
	if r.Params == nil {
		r.Params = &OrderListCancelRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderListCancelRequest) SetSignature(value string) *OrderListCancelRequest {
	if r.Params == nil {
		r.Params = &OrderListCancelRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderListCancelRequest) SetSymbol(value string) *OrderListCancelRequest {
	if r.Params == nil {
		r.Params = &OrderListCancelRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderListCancelRequest) SetTimestamp(value int64) *OrderListCancelRequest {
	if r.Params == nil {
		r.Params = &OrderListCancelRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


