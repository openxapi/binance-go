package models

import (
	"encoding/json"
)

// TimeResponseRateLimitsItem represents a nested object structure
type TimeResponseRateLimitsItem struct {
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

// TimeResponseResult represents a nested object structure
type TimeResponseResult struct {
	// serverTime property
	ServerTime int64 `json:"serverTime,omitempty"`
}

// TimeResponse - Receive response from time
// Message name: Check server time Response
type TimeResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []TimeResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *TimeResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of TimeResponse
func (s TimeResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


