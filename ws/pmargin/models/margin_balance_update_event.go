package models

import (
	"encoding/json"
)

// MarginBalanceUpdateEvent - Margin balance update event
// Message name: Margin Balance Update Event
type MarginBalanceUpdateEvent struct {
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

// String returns string representation of MarginBalanceUpdateEvent
func (s MarginBalanceUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for MarginBalanceUpdateEvent
func (s MarginBalanceUpdateEvent) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "marginbalanceupdateevent"
}

// GetEventTime returns the event timestamp for MarginBalanceUpdateEvent
func (s MarginBalanceUpdateEvent) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


