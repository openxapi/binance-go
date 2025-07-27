package models

import (
	"encoding/json"
)

// KlinesResponseRateLimitsItem represents a nested object structure
type KlinesResponseRateLimitsItem struct {
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

// KlinesResponse - Receive response from klines
// Message name: Klines Response
type KlinesResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []KlinesResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result [][]interface{} `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of KlinesResponse
func (s KlinesResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


