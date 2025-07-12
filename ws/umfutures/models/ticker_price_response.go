package models

import (
	"encoding/json"
)

// TickerPriceResponseRateLimitsItem represents a nested object structure
type TickerPriceResponseRateLimitsItem struct {
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

// TickerPriceResponseResult represents a nested object structure
type TickerPriceResponseResult struct {
	// price property
	Price string `json:"price,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// time property
	Time int64 `json:"time,omitempty"`
}

// TickerPriceResponse - Receive response from ticker.price
// Message name: Symbol Price Ticker Response
type TickerPriceResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []TickerPriceResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *TickerPriceResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of TickerPriceResponse
func (s TickerPriceResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


