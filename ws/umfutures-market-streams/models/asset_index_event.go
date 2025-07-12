package models

import (
	"encoding/json"
)

// AssetIndexEvent represents AssetIndexEvent
type AssetIndexEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Asset symbol
	Symbol string `json:"s,omitempty"`
	// Asset index
	Interval string `json:"i,omitempty"`
}

// String returns string representation of AssetIndexEvent
func (s AssetIndexEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


