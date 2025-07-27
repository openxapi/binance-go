package models

import (
	"encoding/json"
)

// RollingWindowTickerEvent represents RollingWindowTickerEvent
type RollingWindowTickerEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Price change
	PriceChange string `json:"p,omitempty"`
	// Price change percent
	PriceChangePercent string `json:"P,omitempty"`
	// Open price
	OpenPrice string `json:"o,omitempty"`
	// High price
	HighPrice string `json:"h,omitempty"`
	// Low price
	LowPrice string `json:"l,omitempty"`
	// Last price
	LastPrice string `json:"c,omitempty"`
	// Weighted average price
	WeightedAveragePrice string `json:"w,omitempty"`
	// Total traded base asset volume
	TotalTradedBaseAssetVolume string `json:"v,omitempty"`
	// Total traded quote asset volume
	TotalTradedQuoteAssetVolume string `json:"q,omitempty"`
	// Statistics open time
	OpenTime int64 `json:"O,omitempty"`
	// Statistics close time
	CloseTime int64 `json:"C,omitempty"`
	// First trade ID
	FirstTradeId int64 `json:"F,omitempty"`
	// Last trade ID
	LastTradeId int64 `json:"L,omitempty"`
	// Total number of trades
	TotalNumberOfTrades int `json:"n,omitempty"`
}

// String returns string representation of RollingWindowTickerEvent
func (s RollingWindowTickerEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


