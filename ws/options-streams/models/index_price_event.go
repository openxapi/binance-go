package models

import (
	"encoding/json"
)

// IndexPriceEvent represents IndexPriceEvent
// Index price information event for underlying assets.
// Stream name pattern: {symbol}@index
// Example: ETHUSDT@index
// Update frequency: 1000ms (1 second)
// 
type IndexPriceEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event timestamp
	EventTime int64 `json:"E,omitempty"`
	// Underlying symbol
	Symbol string `json:"s,omitempty"`
	// Index price
	IndexPrice string `json:"p,omitempty"`
}

// String returns string representation of IndexPriceEvent
func (s IndexPriceEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


