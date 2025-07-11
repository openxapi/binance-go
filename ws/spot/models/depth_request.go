package models

import (
	"encoding/json"
)

// DepthRequestParams contains the parameters for DepthRequest
type DepthRequestParams struct {
	// Default: 100; Maximum: 5000
	Limit *int64 `json:"limit,omitempty"`
	Symbol string `json:"symbol"`
}

// DepthRequest - Send a depth request
// Message name: Order book Request
type DepthRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *DepthRequestParams `json:"params,omitempty"`
}

// NewDepthRequest creates a new DepthRequest instance
func NewDepthRequest() *DepthRequest {
	return &DepthRequest{}
}

// String returns string representation of DepthRequest
func (s DepthRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *DepthRequest) SetParams(value DepthRequestParams) *DepthRequest {
	r.Params = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *DepthRequest) SetLimit(value int64) *DepthRequest {
	if r.Params == nil {
		r.Params = &DepthRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *DepthRequest) SetSymbol(value string) *DepthRequest {
	if r.Params == nil {
		r.Params = &DepthRequestParams{}
	}
	r.Params.Symbol = value
	return r
}


