package models

import (
	"encoding/json"
)

// AvgPriceEvent represents AvgPriceEvent
type AvgPriceEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Average price interval (minutes)
	Interval string `json:"i,omitempty"`
	// Average price
	WeightedAveragePrice string `json:"w,omitempty"`
	// Last trade time
	TradeTime int64 `json:"T,omitempty"`
}

// String returns string representation of AvgPriceEvent
func (s AvgPriceEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


