package models

import (
	"encoding/json"
)

// IndexKlineEvent represents IndexKlineEvent
type IndexKlineEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Pair
	Pair string `json:"ps,omitempty"`
	// Kline data
	Kline *IndexKlineEventKline `json:"k,omitempty"`
}

// IndexKlineEventKline represents the kline details
type IndexKlineEventKline struct {
	// Kline start time
	KlineStartTime int64 `json:"t,omitempty"`
	// Kline close time
	KlineCloseTime int64 `json:"T,omitempty"`
	// Interval
	Interval string `json:"i,omitempty"`
	// Open price
	OpenPrice string `json:"o,omitempty"`
	// Close price
	ClosePrice string `json:"c,omitempty"`
	// High price
	HighPrice string `json:"h,omitempty"`
	// Low price
	LowPrice string `json:"l,omitempty"`
	// Number of basic data points
	NumberOfTrades int `json:"n,omitempty"`
	// Is this kline closed
	IsKlineClosed bool `json:"x,omitempty"`
}

// String returns string representation of IndexKlineEvent
func (s IndexKlineEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


