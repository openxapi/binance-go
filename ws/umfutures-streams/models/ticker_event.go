package models

import (
	"encoding/json"
)

// TickerEvent represents TickerEvent
type TickerEvent struct {
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
	// Weighted average price
	WeightedAveragePrice string `json:"w,omitempty"`
	// Last price
	LastPrice string `json:"c,omitempty"`
	// Last quantity
	LastQuantity string `json:"Q,omitempty"`
	// Open price
	OpenPrice string `json:"o,omitempty"`
	// High price
	HighPrice string `json:"h,omitempty"`
	// Low price
	LowPrice string `json:"l,omitempty"`
	// Total traded base asset volume
	TotalTradedBaseAssetVolume string `json:"v,omitempty"`
	// Total traded quote asset volume
	Quantity string `json:"q,omitempty"`
	// Statistics open time
	StatisticsOpenTime int64 `json:"O,omitempty"`
	// Statistics close time
	StatisticsCloseTime int64 `json:"C,omitempty"`
	// First trade ID
	FirstTradeId int64 `json:"F,omitempty"`
	// Last trade ID
	LastTradeId int64 `json:"L,omitempty"`
	// Total number of trades
	TotalNumberOfTrades int `json:"n,omitempty"`
}

// String returns string representation of TickerEvent
func (s TickerEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


