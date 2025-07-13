package models

import (
	"encoding/json"
)

// KlineEvent represents KlineEvent
type KlineEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Kline data
	Kline *KlineEventKline `json:"k,omitempty"`
}

// KlineEventKline represents the kline details
type KlineEventKline struct {
	// Kline start time
	KlineStartTime int64 `json:"t,omitempty"`
	// Kline close time
	KlineCloseTime int64 `json:"T,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Interval
	Interval string `json:"i,omitempty"`
	// First trade ID
	FirstTradeId int64 `json:"f,omitempty"`
	// Last trade ID
	LastTradeId int64 `json:"L,omitempty"`
	// Open price
	OpenPrice string `json:"o,omitempty"`
	// Close price
	ClosePrice string `json:"c,omitempty"`
	// High price
	HighPrice string `json:"h,omitempty"`
	// Low price
	LowPrice string `json:"l,omitempty"`
	// Base asset volume
	BaseAssetVolume string `json:"v,omitempty"`
	// Number of trades
	NumberOfTrades int `json:"n,omitempty"`
	// Is this kline closed
	IsKlineClosed bool `json:"x,omitempty"`
	// Quote asset volume
	QuoteAssetVolume string `json:"q,omitempty"`
	// Taker buy base asset volume
	TakerBuyBaseAssetVolume string `json:"V,omitempty"`
	// Taker buy quote asset volume
	TakerBuyQuoteAssetVolume string `json:"Q,omitempty"`
	// Ignore field
	IgnoreField string `json:"B,omitempty"`
}

// String returns string representation of KlineEvent
func (s KlineEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


