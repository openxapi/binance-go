package models

import (
	"encoding/json"
)

// MarkPriceEvent represents MarkPriceEvent
// Mark price information event for options contracts.
// Stream name pattern: {underlyingAsset}@markPrice
// Example: ETH@markPrice
// Update frequency: 1000ms (1 second)
// Provides mark prices for all option symbols on a specific underlying asset.
// 
type MarkPriceEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event timestamp
	EventTime int64 `json:"E,omitempty"`
	// Option symbol
	Symbol string `json:"s,omitempty"`
	// Option mark price
	OptionMarkPrice string `json:"mp,omitempty"`
}

// String returns string representation of MarkPriceEvent
func (s MarkPriceEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


