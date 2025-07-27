package models

import (
	"encoding/json"
)

// TradesAggregateResponseRateLimitsItem represents a nested object structure
type TradesAggregateResponseRateLimitsItem struct {
	// count property (example: 2)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 6000)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// TradesAggregateResponseResultItem represents a nested object structure
type TradesAggregateResponseResultItem struct {
	// M property (example: true)
	M bool `json:"M,omitempty"`
	// T property (example: 1565877971222)
	T int64 `json:"T,omitempty"`
	// a property (example: 50000000)
	A int64 `json:"a,omitempty"`
	// f property (example: 59120167)
	F int64 `json:"f,omitempty"`
	// l property (example: 59120170)
	L int64 `json:"l,omitempty"`
	// m property (example: true)
	M2 bool `json:"m,omitempty"`
	// p property (example: "0.00274100")
	P string `json:"p,omitempty"`
	// q property (example: "57.19000000")
	Q string `json:"q,omitempty"`
}

// TradesAggregateResponse - Receive response from trades.aggregate
// Message name: Aggregate trades Response
type TradesAggregateResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []TradesAggregateResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []TradesAggregateResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of TradesAggregateResponse
func (s TradesAggregateResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


