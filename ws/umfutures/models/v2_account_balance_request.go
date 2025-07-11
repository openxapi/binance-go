package models

import (
	"encoding/json"
)

// V2AccountBalanceRequestParams contains the parameters for V2AccountBalanceRequest
type V2AccountBalanceRequestParams struct {
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

// V2AccountBalanceRequest - Send a v2/account.balance request
// Message name: Futures Account Balance V2(USER_DATA) Request
type V2AccountBalanceRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *V2AccountBalanceRequestParams `json:"params,omitempty"`
}

// NewV2AccountBalanceRequest creates a new V2AccountBalanceRequest instance
func NewV2AccountBalanceRequest() *V2AccountBalanceRequest {
	return &V2AccountBalanceRequest{}
}

// String returns string representation of V2AccountBalanceRequest
func (s V2AccountBalanceRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *V2AccountBalanceRequest) SetParams(value V2AccountBalanceRequestParams) *V2AccountBalanceRequest {
	r.Params = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *V2AccountBalanceRequest) SetRecvWindow(value int64) *V2AccountBalanceRequest {
	if r.Params == nil {
		r.Params = &V2AccountBalanceRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *V2AccountBalanceRequest) SetTimestamp(value int64) *V2AccountBalanceRequest {
	if r.Params == nil {
		r.Params = &V2AccountBalanceRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


