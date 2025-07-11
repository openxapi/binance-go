package models

import (
	"encoding/json"
)

// OpenOrdersStatusRequestParams contains the parameters for OpenOrdersStatusRequest
type OpenOrdersStatusRequestParams struct {
	ApiKey string `json:"apiKey"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	// If omitted, open orders for all symbols are returned
	Symbol *string `json:"symbol,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

// OpenOrdersStatusRequest - Send a openOrders.status request
// Message name: Current open orders (USER_DATA) Request
type OpenOrdersStatusRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OpenOrdersStatusRequestParams `json:"params,omitempty"`
}

// NewOpenOrdersStatusRequest creates a new OpenOrdersStatusRequest instance
func NewOpenOrdersStatusRequest() *OpenOrdersStatusRequest {
	return &OpenOrdersStatusRequest{}
}

// String returns string representation of OpenOrdersStatusRequest
func (s OpenOrdersStatusRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OpenOrdersStatusRequest) SetParams(value OpenOrdersStatusRequestParams) *OpenOrdersStatusRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OpenOrdersStatusRequest) SetApiKey(value string) *OpenOrdersStatusRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersStatusRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OpenOrdersStatusRequest) SetRecvWindow(value int64) *OpenOrdersStatusRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersStatusRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OpenOrdersStatusRequest) SetSignature(value string) *OpenOrdersStatusRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersStatusRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OpenOrdersStatusRequest) SetSymbol(value string) *OpenOrdersStatusRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersStatusRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OpenOrdersStatusRequest) SetTimestamp(value int64) *OpenOrdersStatusRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersStatusRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


