package models

import (
	"encoding/json"
)

// AccountStatusRequestParams contains the parameters for AccountStatusRequest
type AccountStatusRequestParams struct {
	ApiKey string `json:"apiKey"`
	// When set to true, emits only the non-zero balances of an account. Default value: false
	OmitZeroBalances *bool `json:"omitZeroBalances,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	Timestamp int64 `json:"timestamp"`
}

// AccountStatusRequest - Send a account.status request
// Message name: Account information (USER_DATA) Request
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

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *AccountStatusRequest) SetApiKey(value string) *AccountStatusRequest {
	if r.Params == nil {
		r.Params = &AccountStatusRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetOmitZeroBalances sets the omitZeroBalances parameter and returns the struct for method chaining
func (r *AccountStatusRequest) SetOmitZeroBalances(value bool) *AccountStatusRequest {
	if r.Params == nil {
		r.Params = &AccountStatusRequestParams{}
	}
	r.Params.OmitZeroBalances = &value
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

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *AccountStatusRequest) SetSignature(value string) *AccountStatusRequest {
	if r.Params == nil {
		r.Params = &AccountStatusRequestParams{}
	}
	r.Params.Signature = value
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


