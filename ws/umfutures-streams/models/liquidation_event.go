package models

import (
	"encoding/json"
)

// LiquidationEvent represents LiquidationEvent
type LiquidationEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Liquidation order
	OpenPrice interface{} `json:"o,omitempty"`
}

// String returns string representation of LiquidationEvent
func (s LiquidationEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


