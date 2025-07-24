package models

import (
	"encoding/json"
)

// MarkPriceKlineEvent represents MarkPriceKlineEvent
type MarkPriceKlineEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Pair
	Pair string `json:"ps,omitempty"`
	// Kline data
	KlineData *MarkPriceKlineEventKline `json:"k,omitempty"`
}

// MarkPriceKlineEventKline represents the kline details
type MarkPriceKlineEventKline struct {
	// Kline start time
	KlineStartTime int64 `json:"t,omitempty"`
	// Kline close time
	NextFundingTime int64 `json:"T,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Interval
	IndexPrice string `json:"i,omitempty"`
	// Open price
	OpenPrice string `json:"o,omitempty"`
	// Close price
	ClosePrice string `json:"c,omitempty"`
	// High price
	HighPrice string `json:"h,omitempty"`
	// Low price
	LowPrice string `json:"l,omitempty"`
	// Number of basic data points
	NumberOfBasicDataPoints int `json:"n,omitempty"`
	// Is this kline closed
	IsThisKlineClosed bool `json:"x,omitempty"`
}

// String returns string representation of MarkPriceKlineEvent
func (s MarkPriceKlineEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


