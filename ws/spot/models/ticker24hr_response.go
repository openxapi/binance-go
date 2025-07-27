package models

import (
	"encoding/json"
)

// Ticker24hrResponseRateLimitsItem represents a nested object structure
type Ticker24hrResponseRateLimitsItem struct {
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

// Ticker24hrResponseResult represents a nested object structure
type Ticker24hrResponseResult struct {
	// askPrice property
	AskPrice string `json:"askPrice,omitempty"`
	// askQty property
	AskQty string `json:"askQty,omitempty"`
	// bidPrice property
	BidPrice string `json:"bidPrice,omitempty"`
	// bidQty property
	BidQty string `json:"bidQty,omitempty"`
	// closeTime property
	CloseTime int64 `json:"closeTime,omitempty"`
	// count property
	Count int64 `json:"count,omitempty"`
	// firstId property
	FirstId int64 `json:"firstId,omitempty"`
	// highPrice property
	HighPrice string `json:"highPrice,omitempty"`
	// lastId property
	LastId int64 `json:"lastId,omitempty"`
	// lastPrice property
	LastPrice string `json:"lastPrice,omitempty"`
	// lastQty property
	LastQty string `json:"lastQty,omitempty"`
	// lowPrice property
	LowPrice string `json:"lowPrice,omitempty"`
	// openPrice property
	OpenPrice string `json:"openPrice,omitempty"`
	// openTime property
	OpenTime int64 `json:"openTime,omitempty"`
	// prevClosePrice property
	PrevClosePrice string `json:"prevClosePrice,omitempty"`
	// priceChange property
	PriceChange string `json:"priceChange,omitempty"`
	// priceChangePercent property
	PriceChangePercent string `json:"priceChangePercent,omitempty"`
	// quoteVolume property
	QuoteVolume string `json:"quoteVolume,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// volume property
	Volume string `json:"volume,omitempty"`
	// weightedAvgPrice property
	WeightedAvgPrice string `json:"weightedAvgPrice,omitempty"`
}

// Ticker24hrResponse - Receive response from ticker.24hr
// Message name: 24hr ticker price change statistics Response
type Ticker24hrResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []Ticker24hrResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *Ticker24hrResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of Ticker24hrResponse
func (s Ticker24hrResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


