package models

import (
	"encoding/json"
)

// UiKlinesResponseRateLimitsItem represents a nested object structure
type UiKlinesResponseRateLimitsItem struct {
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

// UiKlinesResponse - Receive response from uiKlines
// Message name: UI Klines Response
type UiKlinesResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []UiKlinesResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result [][]interface{} `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of UiKlinesResponse
func (s UiKlinesResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


