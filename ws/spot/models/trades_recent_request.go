package models

import (
	"encoding/json"
)

// TradesRecentRequestParams contains the parameters for TradesRecentRequest
type TradesRecentRequestParams struct {
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	Symbol string `json:"symbol"`
}

// TradesRecentRequest - Send a trades.recent request
// Message name: Recent trades Request
type TradesRecentRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *TradesRecentRequestParams `json:"params,omitempty"`
}

// NewTradesRecentRequest creates a new TradesRecentRequest instance
func NewTradesRecentRequest() *TradesRecentRequest {
	return &TradesRecentRequest{}
}

// String returns string representation of TradesRecentRequest
func (s TradesRecentRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *TradesRecentRequest) SetParams(value TradesRecentRequestParams) *TradesRecentRequest {
	r.Params = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *TradesRecentRequest) SetLimit(value int64) *TradesRecentRequest {
	if r.Params == nil {
		r.Params = &TradesRecentRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *TradesRecentRequest) SetSymbol(value string) *TradesRecentRequest {
	if r.Params == nil {
		r.Params = &TradesRecentRequestParams{}
	}
	r.Params.Symbol = value
	return r
}


