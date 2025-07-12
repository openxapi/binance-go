package models

import (
	"encoding/json"
)

// TickerBookResponseRateLimitsItem represents a nested object structure
type TickerBookResponseRateLimitsItem struct {
	// count property (example: 2)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 2400)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// TickerBookResponseResult represents a nested object structure
type TickerBookResponseResult struct {
	// askPrice property
	AskPrice string `json:"askPrice,omitempty"`
	// askQty property
	AskQty string `json:"askQty,omitempty"`
	// bidPrice property
	BidPrice string `json:"bidPrice,omitempty"`
	// bidQty property
	BidQty string `json:"bidQty,omitempty"`
	// lastUpdateId property
	LastUpdateId int64 `json:"lastUpdateId,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// time property
	Time int64 `json:"time,omitempty"`
}

// TickerBookResponse - Receive response from ticker.book
// Message name: Symbol Order Book Ticker Response
type TickerBookResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []TickerBookResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *TickerBookResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of TickerBookResponse
func (s TickerBookResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


