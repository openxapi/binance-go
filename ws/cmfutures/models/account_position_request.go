package models

import (
	"encoding/json"
)

// AccountPositionRequestParams contains the parameters for AccountPositionRequest
type AccountPositionRequestParams struct {
	MarginAsset *string `json:"marginAsset,omitempty"`
	Pair *string `json:"pair,omitempty"`
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

// AccountPositionRequest - Send a account.position request
// Message name: Position Information(USER_DATA) Request
type AccountPositionRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *AccountPositionRequestParams `json:"params,omitempty"`
}

// NewAccountPositionRequest creates a new AccountPositionRequest instance
func NewAccountPositionRequest() *AccountPositionRequest {
	return &AccountPositionRequest{}
}

// String returns string representation of AccountPositionRequest
func (s AccountPositionRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *AccountPositionRequest) SetParams(value AccountPositionRequestParams) *AccountPositionRequest {
	r.Params = &value
	return r
}

// SetMarginAsset sets the marginAsset parameter and returns the struct for method chaining
func (r *AccountPositionRequest) SetMarginAsset(value string) *AccountPositionRequest {
	if r.Params == nil {
		r.Params = &AccountPositionRequestParams{}
	}
	r.Params.MarginAsset = &value
	return r
}

// SetPair sets the pair parameter and returns the struct for method chaining
func (r *AccountPositionRequest) SetPair(value string) *AccountPositionRequest {
	if r.Params == nil {
		r.Params = &AccountPositionRequestParams{}
	}
	r.Params.Pair = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *AccountPositionRequest) SetRecvWindow(value int64) *AccountPositionRequest {
	if r.Params == nil {
		r.Params = &AccountPositionRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *AccountPositionRequest) SetTimestamp(value int64) *AccountPositionRequest {
	if r.Params == nil {
		r.Params = &AccountPositionRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


