package models

import (
	"encoding/json"
)

// CombinedData represents Combined Stream Data
type CombinedData struct {
	// Original stream name that generated this event
	StreamName string `json:"stream,omitempty"`
	// Original stream event data (unwrapped). This is exactly the same data
	// that would be received when connecting to the individual stream directly.
	// The event type can be determined by examining the 'stream' field.
	// 
	// For all-market streams (!ticker@arr, !miniTicker@arr), the data will be
	// an array of events instead of a single event object.
	// 
	StreamData interface{} `json:"data,omitempty"`
}

// String returns string representation of CombinedData
func (s CombinedData) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewCombinedData creates a new CombinedData instance
func NewCombinedData() *CombinedData {
	return &CombinedData{}
}


