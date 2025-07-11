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
	Kline interface{} `json:"k,omitempty"`
}

// String returns string representation of KlineEvent
func (s KlineEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


