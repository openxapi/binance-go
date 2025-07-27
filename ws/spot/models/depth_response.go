package models

import (
	"encoding/json"
)

// DepthResponseRateLimitsItem represents a nested object structure
type DepthResponseRateLimitsItem struct {
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

// DepthResponseResult represents a nested object structure
type DepthResponseResult struct {
	// asks property
	Asks [][]string `json:"asks,omitempty"`
	// bids property
	Bids [][]string `json:"bids,omitempty"`
	// lastUpdateId property
	LastUpdateId int64 `json:"lastUpdateId,omitempty"`
}

// DepthResponse - Receive response from depth
// Message name: Order book Response
type DepthResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []DepthResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *DepthResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of DepthResponse
func (s DepthResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


