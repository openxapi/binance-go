package models

import (
	"encoding/json"
)

// BalanceUpdateEvent represents a nested object structure
type BalanceUpdateEvent struct {
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

// BalanceUpdate - Occurs during deposits, withdrawals, or transfers
// Message name: Balance Update Event
type BalanceUpdate struct {
	Event BalanceUpdateEvent `json:"event,omitempty"`
}

// String returns string representation of BalanceUpdate
func (s BalanceUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for BalanceUpdate
func (s BalanceUpdate) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "balanceupdate"
}

// GetEventTime returns the event timestamp for BalanceUpdate
func (s BalanceUpdate) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


