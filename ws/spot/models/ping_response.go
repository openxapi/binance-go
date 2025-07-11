package models

import (
	"encoding/json"
)

// PingResponseRateLimitsItem represents a nested object structure
type PingResponseRateLimitsItem struct {
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

// PingResponseResult represents a nested object structure
type PingResponseResult struct {
}

// PingResponse - Receive response from ping
// Message name: Test connectivity Response
type PingResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []PingResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result interface{} `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of PingResponse
func (s PingResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


