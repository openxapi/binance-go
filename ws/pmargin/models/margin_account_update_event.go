package models

import (
	"encoding/json"
)

// MarginAccountUpdateEventBItem represents a nested object structure
type MarginAccountUpdateEventBItem struct {
	// Asset (example: "ETH")
	Asset string `json:"a,omitempty"`
	// Free balance (example: "10000.000000")
	FreeBalance string `json:"f,omitempty"`
	// Locked balance (example: "0.000000")
	LockedBalance string `json:"l,omitempty"`
}

// MarginAccountUpdateEvent - Sent when account balance has changed
// Message name: Margin Account Update Event
type MarginAccountUpdateEvent struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Time of last account update (milliseconds)
	TimeOfLastAccountUpdate int64 `json:"u,omitempty"`
	// Update ID
	UpdateID int64 `json:"U,omitempty"`
	// Balances Array
	BalancesArray []MarginAccountUpdateEventBItem `json:"B,omitempty"`
}

// String returns string representation of MarginAccountUpdateEvent
func (s MarginAccountUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for MarginAccountUpdateEvent
func (s MarginAccountUpdateEvent) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "marginaccountupdateevent"
}

// GetEventTime returns the event timestamp for MarginAccountUpdateEvent
func (s MarginAccountUpdateEvent) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


