package models

import (
	"encoding/json"
)

// TradesRecentResponseRateLimitsItem represents a nested object structure
type TradesRecentResponseRateLimitsItem struct {
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

// TradesRecentResponseResultItem represents a nested object structure
type TradesRecentResponseResultItem struct {
	// id property (example: 194686783)
	Id int64 `json:"id,omitempty"`
	// isBestMatch property (example: true)
	IsBestMatch bool `json:"isBestMatch,omitempty"`
	// isBuyerMaker property (example: true)
	IsBuyerMaker bool `json:"isBuyerMaker,omitempty"`
	// price property (example: "0.01361000")
	Price string `json:"price,omitempty"`
	// qty property (example: "0.01400000")
	Qty string `json:"qty,omitempty"`
	// quoteQty property (example: "0.00019054")
	QuoteQty string `json:"quoteQty,omitempty"`
	// time property (example: 1660009530807)
	Time int64 `json:"time,omitempty"`
}

// TradesRecentResponse - Receive response from trades.recent
// Message name: Recent trades Response
type TradesRecentResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []TradesRecentResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []TradesRecentResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of TradesRecentResponse
func (s TradesRecentResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


