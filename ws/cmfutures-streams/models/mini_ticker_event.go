package models

import (
	"encoding/json"
)

// MiniTickerEvent represents MiniTickerEvent
type MiniTickerEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Pair
	Pair string `json:"ps,omitempty"`
	// Close price
	ClosePrice string `json:"c,omitempty"`
	// Open price
	OpenPrice string `json:"o,omitempty"`
	// High price
	HighPrice string `json:"h,omitempty"`
	// Low price
	LowPrice string `json:"l,omitempty"`
	// Total traded volume
	TotalTradedBaseAssetVolume string `json:"v,omitempty"`
	// Total traded base asset volume
	TotalTradedQuoteAssetVolume string `json:"q,omitempty"`
}

// String returns string representation of MiniTickerEvent
func (s MiniTickerEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


