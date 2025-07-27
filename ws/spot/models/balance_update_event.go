package models

import (
	"encoding/json"
)

// BalanceUpdateEventEvent represents a nested object structure
type BalanceUpdateEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Clear Time (milliseconds)
	ClearTime int64 `json:"T,omitempty"`
	// Asset
	Asset string `json:"a,omitempty"`
	// Balance Delta
	BalanceDelta string `json:"d,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
}

// BalanceUpdateEvent - Occurs during deposits, withdrawals, or transfers
// Message name: Balance Update Event
type BalanceUpdateEvent struct {
	Event *BalanceUpdateEventEvent `json:"event,omitempty"`
}

// String returns string representation of BalanceUpdateEvent
func (s BalanceUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for BalanceUpdateEvent
func (s BalanceUpdateEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "balanceupdateevent"
}

// GetEventTime returns the event timestamp for BalanceUpdateEvent
func (s BalanceUpdateEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


