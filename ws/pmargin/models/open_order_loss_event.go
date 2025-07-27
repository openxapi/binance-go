package models

import (
	"encoding/json"
)

// OpenOrderLossEventOItem represents a nested object structure
type OpenOrderLossEventOItem struct {
	// Asset (example: "BUSD")
	Asset string `json:"a,omitempty"`
	// Loss Amount (negative number) (example: "-0.1232313")
	LossAmount string `json:"o,omitempty"`
}

// OpenOrderLossEvent - Cross margin order margin stream updates
// Message name: Open Order Loss Update Event
type OpenOrderLossEvent struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Array of loss updates
	ArrayOfLossUpdates []OpenOrderLossEventOItem `json:"O,omitempty"`
}

// String returns string representation of OpenOrderLossEvent
func (s OpenOrderLossEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for OpenOrderLossEvent
func (s OpenOrderLossEvent) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "openorderlossevent"
}

// GetEventTime returns the event timestamp for OpenOrderLossEvent
func (s OpenOrderLossEvent) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


