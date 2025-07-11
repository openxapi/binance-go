package models

import (
	"encoding/json"
)

// OpenOrderListsStatusRequestParams contains the parameters for OpenOrderListsStatusRequest
type OpenOrderListsStatusRequestParams struct {
	ApiKey string `json:"apiKey"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	Timestamp int64 `json:"timestamp"`
}

// OpenOrderListsStatusRequest - Send a openOrderLists.status request
// Message name: Current open Order lists (USER_DATA) Request
type OpenOrderListsStatusRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OpenOrderListsStatusRequestParams `json:"params,omitempty"`
}

// NewOpenOrderListsStatusRequest creates a new OpenOrderListsStatusRequest instance
func NewOpenOrderListsStatusRequest() *OpenOrderListsStatusRequest {
	return &OpenOrderListsStatusRequest{}
}

// String returns string representation of OpenOrderListsStatusRequest
func (s OpenOrderListsStatusRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OpenOrderListsStatusRequest) SetParams(value OpenOrderListsStatusRequestParams) *OpenOrderListsStatusRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OpenOrderListsStatusRequest) SetApiKey(value string) *OpenOrderListsStatusRequest {
	if r.Params == nil {
		r.Params = &OpenOrderListsStatusRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OpenOrderListsStatusRequest) SetRecvWindow(value int64) *OpenOrderListsStatusRequest {
	if r.Params == nil {
		r.Params = &OpenOrderListsStatusRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OpenOrderListsStatusRequest) SetSignature(value string) *OpenOrderListsStatusRequest {
	if r.Params == nil {
		r.Params = &OpenOrderListsStatusRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OpenOrderListsStatusRequest) SetTimestamp(value int64) *OpenOrderListsStatusRequest {
	if r.Params == nil {
		r.Params = &OpenOrderListsStatusRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


