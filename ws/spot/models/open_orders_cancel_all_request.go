package models

import (
	"encoding/json"
)

// OpenOrdersCancelAllRequestParams contains the parameters for OpenOrdersCancelAllRequest
type OpenOrdersCancelAllRequestParams struct {
	ApiKey string `json:"apiKey"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// OpenOrdersCancelAllRequest - Send a openOrders.cancelAll request
// Message name: Cancel open orders (TRADE) Request
type OpenOrdersCancelAllRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OpenOrdersCancelAllRequestParams `json:"params,omitempty"`
}

// NewOpenOrdersCancelAllRequest creates a new OpenOrdersCancelAllRequest instance
func NewOpenOrdersCancelAllRequest() *OpenOrdersCancelAllRequest {
	return &OpenOrdersCancelAllRequest{}
}

// String returns string representation of OpenOrdersCancelAllRequest
func (s OpenOrdersCancelAllRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OpenOrdersCancelAllRequest) SetParams(value OpenOrdersCancelAllRequestParams) *OpenOrdersCancelAllRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OpenOrdersCancelAllRequest) SetApiKey(value string) *OpenOrdersCancelAllRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersCancelAllRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OpenOrdersCancelAllRequest) SetRecvWindow(value int64) *OpenOrdersCancelAllRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersCancelAllRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OpenOrdersCancelAllRequest) SetSignature(value string) *OpenOrdersCancelAllRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersCancelAllRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OpenOrdersCancelAllRequest) SetSymbol(value string) *OpenOrdersCancelAllRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersCancelAllRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OpenOrdersCancelAllRequest) SetTimestamp(value int64) *OpenOrdersCancelAllRequest {
	if r.Params == nil {
		r.Params = &OpenOrdersCancelAllRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


