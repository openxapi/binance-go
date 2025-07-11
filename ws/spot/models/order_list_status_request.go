package models

import (
	"encoding/json"
)

// OrderListStatusRequestParams contains the parameters for OrderListStatusRequest
type OrderListStatusRequestParams struct {
	ApiKey *string `json:"apiKey,omitempty"`
	// Query order list by orderListId.orderListId or origClientOrderId must be provided.
	OrderListId *int64 `json:"orderListId,omitempty"`
	// Query order list by listClientOrderId.orderListId or origClientOrderId must be provided.
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature *string `json:"signature,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// OrderListStatusRequest - Send a orderList.status request
// Message name: Query Order list (USER_DATA) Request
type OrderListStatusRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderListStatusRequestParams `json:"params,omitempty"`
}

// NewOrderListStatusRequest creates a new OrderListStatusRequest instance
func NewOrderListStatusRequest() *OrderListStatusRequest {
	return &OrderListStatusRequest{}
}

// String returns string representation of OrderListStatusRequest
func (s OrderListStatusRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderListStatusRequest) SetParams(value OrderListStatusRequestParams) *OrderListStatusRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OrderListStatusRequest) SetApiKey(value string) *OrderListStatusRequest {
	if r.Params == nil {
		r.Params = &OrderListStatusRequestParams{}
	}
	r.Params.ApiKey = &value
	return r
}

// SetOrderListId sets the orderListId parameter and returns the struct for method chaining
func (r *OrderListStatusRequest) SetOrderListId(value int64) *OrderListStatusRequest {
	if r.Params == nil {
		r.Params = &OrderListStatusRequestParams{}
	}
	r.Params.OrderListId = &value
	return r
}

// SetOrigClientOrderId sets the origClientOrderId parameter and returns the struct for method chaining
func (r *OrderListStatusRequest) SetOrigClientOrderId(value string) *OrderListStatusRequest {
	if r.Params == nil {
		r.Params = &OrderListStatusRequestParams{}
	}
	r.Params.OrigClientOrderId = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderListStatusRequest) SetRecvWindow(value int64) *OrderListStatusRequest {
	if r.Params == nil {
		r.Params = &OrderListStatusRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderListStatusRequest) SetSignature(value string) *OrderListStatusRequest {
	if r.Params == nil {
		r.Params = &OrderListStatusRequestParams{}
	}
	r.Params.Signature = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderListStatusRequest) SetTimestamp(value int64) *OrderListStatusRequest {
	if r.Params == nil {
		r.Params = &OrderListStatusRequestParams{}
	}
	r.Params.Timestamp = &value
	return r
}


