package models

import (
	"encoding/json"
)

// SorOrderTestResponseRateLimitsItem represents a nested object structure
type SorOrderTestResponseRateLimitsItem struct {
	// count property (example: 1)
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

// SorOrderTestResponseResult represents a nested object structure
type SorOrderTestResponseResult struct {
}

// SorOrderTestResponse - Receive response from sor.order.test
// Message name: Test new order using SOR (TRADE) Response
type SorOrderTestResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []SorOrderTestResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result interface{} `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of SorOrderTestResponse
func (s SorOrderTestResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


