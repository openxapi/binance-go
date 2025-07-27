package models

import (
	"encoding/json"
)

// V2AccountStatusRequestParams contains the parameters for V2AccountStatusRequest
type V2AccountStatusRequestParams struct {
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

// V2AccountStatusRequest - Send a v2/account.status request
// Message name: Account Information V2(USER_DATA) Request
type V2AccountStatusRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *V2AccountStatusRequestParams `json:"params,omitempty"`
}

// NewV2AccountStatusRequest creates a new V2AccountStatusRequest instance
func NewV2AccountStatusRequest() *V2AccountStatusRequest {
	return &V2AccountStatusRequest{}
}

// String returns string representation of V2AccountStatusRequest
func (s V2AccountStatusRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *V2AccountStatusRequest) SetParams(value V2AccountStatusRequestParams) *V2AccountStatusRequest {
	r.Params = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *V2AccountStatusRequest) SetRecvWindow(value int64) *V2AccountStatusRequest {
	if r.Params == nil {
		r.Params = &V2AccountStatusRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *V2AccountStatusRequest) SetTimestamp(value int64) *V2AccountStatusRequest {
	if r.Params == nil {
		r.Params = &V2AccountStatusRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


