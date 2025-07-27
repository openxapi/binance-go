package models

import (
	"encoding/json"
)

// CombinedStreamEvent represents CombinedStreamEvent
// Wrapper for combined stream data. When connecting to combined streams,
// all events are wrapped with this structure to identify the source stream.
// 
// The "stream" field contains the original stream name, and "data" contains
// the original event data that would be received from an individual stream.
// 
type CombinedStreamEvent struct {
	// Original stream name that generated this event
	StreamName string `json:"stream,omitempty"`
	// Original stream event data (unwrapped). This is exactly the same data
	// that would be received when connecting to the individual stream directly.
	// The event type can be determined by examining the "stream" field.
	// 
	StreamData interface{} `json:"data,omitempty"`
}

// String returns string representation of CombinedStreamEvent
func (s CombinedStreamEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


