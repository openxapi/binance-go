package models

import (
	"encoding/json"
)

// AvgPriceResponseRateLimitsItem represents a nested object structure
type AvgPriceResponseRateLimitsItem struct {
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

// AvgPriceResponseResult represents a nested object structure
type AvgPriceResponseResult struct {
	// closeTime property
	CloseTime int64 `json:"closeTime,omitempty"`
	// mins property
	Mins int64 `json:"mins,omitempty"`
	// price property
	Price string `json:"price,omitempty"`
}

// AvgPriceResponse - Receive response from avgPrice
// Message name: Current average price Response
type AvgPriceResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []AvgPriceResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *AvgPriceResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of AvgPriceResponse
func (s AvgPriceResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


