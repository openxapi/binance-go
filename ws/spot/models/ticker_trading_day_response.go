package models

import (
	"encoding/json"
)

// TickerTradingDayResponseRateLimitsItem represents a nested object structure
type TickerTradingDayResponseRateLimitsItem struct {
	// count property (example: 4)
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

// TickerTradingDayResponseResult represents a nested object structure
type TickerTradingDayResponseResult struct {
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
	// lowPrice property
	LowPrice string `json:"lowPrice,omitempty"`
	// openPrice property
	OpenPrice string `json:"openPrice,omitempty"`
	// openTime property
	OpenTime int64 `json:"openTime,omitempty"`
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

// TickerTradingDayResponse - Receive response from ticker.tradingDay
// Message name: Trading Day Ticker Response
type TickerTradingDayResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []TickerTradingDayResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *TickerTradingDayResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of TickerTradingDayResponse
func (s TickerTradingDayResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


