package models

import (
	"encoding/json"
)

// TradesHistoricalRequestParams contains the parameters for TradesHistoricalRequest
type TradesHistoricalRequestParams struct {
	// Trade ID to begin at
	FromId *int64 `json:"fromId,omitempty"`
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	Symbol string `json:"symbol"`
}

// TradesHistoricalRequest - Send a trades.historical request
// Message name: Historical trades Request
type TradesHistoricalRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *TradesHistoricalRequestParams `json:"params,omitempty"`
}

// NewTradesHistoricalRequest creates a new TradesHistoricalRequest instance
func NewTradesHistoricalRequest() *TradesHistoricalRequest {
	return &TradesHistoricalRequest{}
}

// String returns string representation of TradesHistoricalRequest
func (s TradesHistoricalRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *TradesHistoricalRequest) SetParams(value TradesHistoricalRequestParams) *TradesHistoricalRequest {
	r.Params = &value
	return r
}

// SetFromId sets the fromId parameter and returns the struct for method chaining
func (r *TradesHistoricalRequest) SetFromId(value int64) *TradesHistoricalRequest {
	if r.Params == nil {
		r.Params = &TradesHistoricalRequestParams{}
	}
	r.Params.FromId = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *TradesHistoricalRequest) SetLimit(value int64) *TradesHistoricalRequest {
	if r.Params == nil {
		r.Params = &TradesHistoricalRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *TradesHistoricalRequest) SetSymbol(value string) *TradesHistoricalRequest {
	if r.Params == nil {
		r.Params = &TradesHistoricalRequestParams{}
	}
	r.Params.Symbol = value
	return r
}


