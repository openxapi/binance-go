package models

import (
	"encoding/json"
)

// AccountBalanceRequestParams contains the parameters for AccountBalanceRequest
type AccountBalanceRequestParams struct {
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

// AccountBalanceRequest - Send a account.balance request
// Message name: Futures Account Balance(USER_DATA) Request
type AccountBalanceRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *AccountBalanceRequestParams `json:"params,omitempty"`
}

// NewAccountBalanceRequest creates a new AccountBalanceRequest instance
func NewAccountBalanceRequest() *AccountBalanceRequest {
	return &AccountBalanceRequest{}
}

// String returns string representation of AccountBalanceRequest
func (s AccountBalanceRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *AccountBalanceRequest) SetParams(value AccountBalanceRequestParams) *AccountBalanceRequest {
	r.Params = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *AccountBalanceRequest) SetRecvWindow(value int64) *AccountBalanceRequest {
	if r.Params == nil {
		r.Params = &AccountBalanceRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *AccountBalanceRequest) SetTimestamp(value int64) *AccountBalanceRequest {
	if r.Params == nil {
		r.Params = &AccountBalanceRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


