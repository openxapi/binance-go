package models

import (
	"encoding/json"
)

// OrderAmendmentsRequestParams contains the parameters for OrderAmendmentsRequest
type OrderAmendmentsRequestParams struct {
	FromExecutionId *int64 `json:"fromExecutionId,omitempty"`
	// Default:500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	OrderId int64 `json:"orderId"`
	// The value cannot be greater than 60000.
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// OrderAmendmentsRequest - Send a order.amendments request
// Message name: Query Order Amendments (USER_DATA) Request
type OrderAmendmentsRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderAmendmentsRequestParams `json:"params,omitempty"`
}

// NewOrderAmendmentsRequest creates a new OrderAmendmentsRequest instance
func NewOrderAmendmentsRequest() *OrderAmendmentsRequest {
	return &OrderAmendmentsRequest{}
}

// String returns string representation of OrderAmendmentsRequest
func (s OrderAmendmentsRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderAmendmentsRequest) SetParams(value OrderAmendmentsRequestParams) *OrderAmendmentsRequest {
	r.Params = &value
	return r
}

// SetFromExecutionId sets the fromExecutionId parameter and returns the struct for method chaining
func (r *OrderAmendmentsRequest) SetFromExecutionId(value int64) *OrderAmendmentsRequest {
	if r.Params == nil {
		r.Params = &OrderAmendmentsRequestParams{}
	}
	r.Params.FromExecutionId = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *OrderAmendmentsRequest) SetLimit(value int64) *OrderAmendmentsRequest {
	if r.Params == nil {
		r.Params = &OrderAmendmentsRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *OrderAmendmentsRequest) SetOrderId(value int64) *OrderAmendmentsRequest {
	if r.Params == nil {
		r.Params = &OrderAmendmentsRequestParams{}
	}
	r.Params.OrderId = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderAmendmentsRequest) SetRecvWindow(value int64) *OrderAmendmentsRequest {
	if r.Params == nil {
		r.Params = &OrderAmendmentsRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderAmendmentsRequest) SetSymbol(value string) *OrderAmendmentsRequest {
	if r.Params == nil {
		r.Params = &OrderAmendmentsRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderAmendmentsRequest) SetTimestamp(value int64) *OrderAmendmentsRequest {
	if r.Params == nil {
		r.Params = &OrderAmendmentsRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


