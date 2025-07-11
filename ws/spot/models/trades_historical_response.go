package models

import (
	"encoding/json"
)

// TradesHistoricalResponseRateLimitsItem represents a nested object structure
type TradesHistoricalResponseRateLimitsItem struct {
	// count property (example: 10)
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

// TradesHistoricalResponseResultItem represents a nested object structure
type TradesHistoricalResponseResultItem struct {
	// id property (example: 0)
	Id int64 `json:"id,omitempty"`
	// isBestMatch property (example: true)
	IsBestMatch bool `json:"isBestMatch,omitempty"`
	// isBuyerMaker property (example: true)
	IsBuyerMaker bool `json:"isBuyerMaker,omitempty"`
	// price property (example: "0.00005000")
	Price string `json:"price,omitempty"`
	// qty property (example: "40.00000000")
	Qty string `json:"qty,omitempty"`
	// quoteQty property (example: "0.00200000")
	QuoteQty string `json:"quoteQty,omitempty"`
	// time property (example: 1500004800376)
	Time int64 `json:"time,omitempty"`
}

// TradesHistoricalResponse - Receive response from trades.historical
// Message name: Historical trades Response
type TradesHistoricalResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []TradesHistoricalResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []TradesHistoricalResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of TradesHistoricalResponse
func (s TradesHistoricalResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


