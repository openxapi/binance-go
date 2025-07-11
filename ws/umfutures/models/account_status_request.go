package models

import (
	"encoding/json"
)

// AccountStatusRequestParams contains the parameters for AccountStatusRequest
type AccountStatusRequestParams struct {
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

// AccountStatusRequest - Send a account.status request
// Message name: Account Information(USER_DATA) Request
type AccountStatusRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *AccountStatusRequestParams `json:"params,omitempty"`
}

// NewAccountStatusRequest creates a new AccountStatusRequest instance
func NewAccountStatusRequest() *AccountStatusRequest {
	return &AccountStatusRequest{}
}

// String returns string representation of AccountStatusRequest
func (s AccountStatusRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *AccountStatusRequest) SetParams(value AccountStatusRequestParams) *AccountStatusRequest {
	r.Params = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *AccountStatusRequest) SetRecvWindow(value int64) *AccountStatusRequest {
	if r.Params == nil {
		r.Params = &AccountStatusRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *AccountStatusRequest) SetTimestamp(value int64) *AccountStatusRequest {
	if r.Params == nil {
		r.Params = &AccountStatusRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


