package models

import (
	"encoding/json"
)

// CompositeIndexEvent represents CompositeIndexEvent
type CompositeIndexEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Price
	Price string `json:"p,omitempty"`
	// Asset type
	CloseTime string `json:"C,omitempty"`
	// Composition
	ClosePrice []interface{} `json:"c,omitempty"`
}

// String returns string representation of CompositeIndexEvent
func (s CompositeIndexEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


