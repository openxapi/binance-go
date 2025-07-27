package models

import (
	"encoding/json"
)

// IndexPriceEvent represents IndexPriceEvent
type IndexPriceEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Pair
	Pair string `json:"i,omitempty"`
	// Index price
	IndexPrice string `json:"p,omitempty"`
}

// String returns string representation of IndexPriceEvent
func (s IndexPriceEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


