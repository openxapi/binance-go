package models

import (
	"encoding/json"
)

// V2AccountPositionRequestParams contains the parameters for V2AccountPositionRequest
type V2AccountPositionRequestParams struct {
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

// V2AccountPositionRequest - Send a v2/account.position request
// Message name: Position Information V2 (USER_DATA) Request
type V2AccountPositionRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *V2AccountPositionRequestParams `json:"params,omitempty"`
}

// NewV2AccountPositionRequest creates a new V2AccountPositionRequest instance
func NewV2AccountPositionRequest() *V2AccountPositionRequest {
	return &V2AccountPositionRequest{}
}

// String returns string representation of V2AccountPositionRequest
func (s V2AccountPositionRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *V2AccountPositionRequest) SetParams(value V2AccountPositionRequestParams) *V2AccountPositionRequest {
	r.Params = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *V2AccountPositionRequest) SetRecvWindow(value int64) *V2AccountPositionRequest {
	if r.Params == nil {
		r.Params = &V2AccountPositionRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *V2AccountPositionRequest) SetSymbol(value string) *V2AccountPositionRequest {
	if r.Params == nil {
		r.Params = &V2AccountPositionRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *V2AccountPositionRequest) SetTimestamp(value int64) *V2AccountPositionRequest {
	if r.Params == nil {
		r.Params = &V2AccountPositionRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


