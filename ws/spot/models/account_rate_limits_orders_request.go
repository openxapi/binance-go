package models

import (
	"encoding/json"
)

// AccountRateLimitsOrdersRequestParams contains the parameters for AccountRateLimitsOrdersRequest
type AccountRateLimitsOrdersRequestParams struct {
	ApiKey string `json:"apiKey"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	Timestamp int64 `json:"timestamp"`
}

// AccountRateLimitsOrdersRequest - Send a account.rateLimits.orders request
// Message name: Unfilled Order Count (USER_DATA) Request
type AccountRateLimitsOrdersRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *AccountRateLimitsOrdersRequestParams `json:"params,omitempty"`
}

// NewAccountRateLimitsOrdersRequest creates a new AccountRateLimitsOrdersRequest instance
func NewAccountRateLimitsOrdersRequest() *AccountRateLimitsOrdersRequest {
	return &AccountRateLimitsOrdersRequest{}
}

// String returns string representation of AccountRateLimitsOrdersRequest
func (s AccountRateLimitsOrdersRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *AccountRateLimitsOrdersRequest) SetParams(value AccountRateLimitsOrdersRequestParams) *AccountRateLimitsOrdersRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *AccountRateLimitsOrdersRequest) SetApiKey(value string) *AccountRateLimitsOrdersRequest {
	if r.Params == nil {
		r.Params = &AccountRateLimitsOrdersRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *AccountRateLimitsOrdersRequest) SetRecvWindow(value int64) *AccountRateLimitsOrdersRequest {
	if r.Params == nil {
		r.Params = &AccountRateLimitsOrdersRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *AccountRateLimitsOrdersRequest) SetSignature(value string) *AccountRateLimitsOrdersRequest {
	if r.Params == nil {
		r.Params = &AccountRateLimitsOrdersRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *AccountRateLimitsOrdersRequest) SetTimestamp(value int64) *AccountRateLimitsOrdersRequest {
	if r.Params == nil {
		r.Params = &AccountRateLimitsOrdersRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


