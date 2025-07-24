package models

import (
	"encoding/json"
)

// OpenOrderLossOItem represents a nested object structure
type OpenOrderLossOItem struct {
	// Asset (example: "BUSD")
	Asset string `json:"a,omitempty"`
	// Loss Amount (negative number) (example: "-0.1232313")
	LossAmount string `json:"o,omitempty"`
}

// OpenOrderLoss - Cross margin order margin stream updates
// Message name: Open Order Loss Update Event
type OpenOrderLoss struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Array of loss updates
	ArrayOfLossUpdates []OpenOrderLossOItem `json:"O,omitempty"`
}

// String returns string representation of OpenOrderLoss
func (s OpenOrderLoss) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for OpenOrderLoss
func (s OpenOrderLoss) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "openorderloss"
}

// GetEventTime returns the event timestamp for OpenOrderLoss
func (s OpenOrderLoss) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


