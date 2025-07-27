package models

import (
	"encoding/json"
)

// KlineEvent represents KlineEvent
// Kline/candlestick information event for options contracts.
// Stream name pattern: {symbol}@kline_{interval}
// Example: BTC-200630-9000-P@kline_1m
// Update frequency: 1000ms (every second)
// 
// Supported intervals: 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 12h, 1d, 3d, 1w
// 
type KlineEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Option symbol
	Symbol string `json:"s,omitempty"`
	// Kline data
	Kline *KlineEventKline `json:"k,omitempty"`
}

// KlineEventKline represents the kline details
type KlineEventKline struct {
	// Kline start time
	KlineStartTime int64 `json:"t,omitempty"`
	// Kline end time
	KlineCloseTime int64 `json:"T,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Interval
	Interval string `json:"i,omitempty"`
	// First trade ID
	FirstTradeId string `json:"F,omitempty"`
	// Last trade ID
	LastTradeId string `json:"L,omitempty"`
	// Open price
	OpenPrice string `json:"o,omitempty"`
	// Close price
	ClosePrice string `json:"c,omitempty"`
	// High price
	HighPrice string `json:"h,omitempty"`
	// Low price
	LowPrice string `json:"l,omitempty"`
	// Volume (contracts)
	BaseAssetVolume string `json:"v,omitempty"`
	// Number of trades
	NumberOfTrades int `json:"n,omitempty"`
	// Is candle completed
	IsKlineClosed bool `json:"x,omitempty"`
	// Completed trade amount
	QuoteAssetVolume string `json:"q,omitempty"`
	// Taker completed trade volume
	TakerBuyBaseAssetVolume string `json:"V,omitempty"`
	// Taker trade amount
	TakerBuyQuoteAssetVolume string `json:"Q,omitempty"`
}

// String returns string representation of KlineEvent
func (s KlineEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


