package models

import (
	"encoding/json"
)

// MarginBalanceUpdate - Margin balance update event
// Message name: Margin Balance Update Event
type MarginBalanceUpdate struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Asset
	Asset string `json:"a,omitempty"`
	// Balance Delta
	BalanceDelta string `json:"d,omitempty"`
	// Event Update ID
	EventUpdateID int64 `json:"U,omitempty"`
	// Clear Time (milliseconds)
	ClearTime int64 `json:"T,omitempty"`
}

// String returns string representation of MarginBalanceUpdate
func (s MarginBalanceUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for MarginBalanceUpdate
func (s MarginBalanceUpdate) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "marginbalanceupdate"
}

// GetEventTime returns the event timestamp for MarginBalanceUpdate
func (s MarginBalanceUpdate) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


