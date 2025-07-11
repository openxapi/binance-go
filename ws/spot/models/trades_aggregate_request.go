package models

import (
	"encoding/json"
)

// TradesAggregateRequestParams contains the parameters for TradesAggregateRequest
type TradesAggregateRequestParams struct {
	EndTime *int64 `json:"endTime,omitempty"`
	// Aggregate trade ID to begin at
	FromId *int64 `json:"fromId,omitempty"`
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	Symbol string `json:"symbol"`
}

// TradesAggregateRequest - Send a trades.aggregate request
// Message name: Aggregate trades Request
type TradesAggregateRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *TradesAggregateRequestParams `json:"params,omitempty"`
}

// NewTradesAggregateRequest creates a new TradesAggregateRequest instance
func NewTradesAggregateRequest() *TradesAggregateRequest {
	return &TradesAggregateRequest{}
}

// String returns string representation of TradesAggregateRequest
func (s TradesAggregateRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *TradesAggregateRequest) SetParams(value TradesAggregateRequestParams) *TradesAggregateRequest {
	r.Params = &value
	return r
}

// SetEndTime sets the endTime parameter and returns the struct for method chaining
func (r *TradesAggregateRequest) SetEndTime(value int64) *TradesAggregateRequest {
	if r.Params == nil {
		r.Params = &TradesAggregateRequestParams{}
	}
	r.Params.EndTime = &value
	return r
}

// SetFromId sets the fromId parameter and returns the struct for method chaining
func (r *TradesAggregateRequest) SetFromId(value int64) *TradesAggregateRequest {
	if r.Params == nil {
		r.Params = &TradesAggregateRequestParams{}
	}
	r.Params.FromId = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *TradesAggregateRequest) SetLimit(value int64) *TradesAggregateRequest {
	if r.Params == nil {
		r.Params = &TradesAggregateRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetStartTime sets the startTime parameter and returns the struct for method chaining
func (r *TradesAggregateRequest) SetStartTime(value int64) *TradesAggregateRequest {
	if r.Params == nil {
		r.Params = &TradesAggregateRequestParams{}
	}
	r.Params.StartTime = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *TradesAggregateRequest) SetSymbol(value string) *TradesAggregateRequest {
	if r.Params == nil {
		r.Params = &TradesAggregateRequestParams{}
	}
	r.Params.Symbol = value
	return r
}


