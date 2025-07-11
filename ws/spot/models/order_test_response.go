package models

import (
	"encoding/json"
)

// OrderTestResponseRateLimitsItem represents a nested object structure
type OrderTestResponseRateLimitsItem struct {
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

// OrderTestResponseResult represents a nested object structure
type OrderTestResponseResult struct {
}

// OrderTestResponse - Receive response from order.test
// Message name: Test new order (TRADE) Response
type OrderTestResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderTestResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result interface{} `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderTestResponse
func (s OrderTestResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


