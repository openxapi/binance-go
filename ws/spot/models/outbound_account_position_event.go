package models

import (
	"encoding/json"
)

// OutboundAccountPositionEventEvent represents a nested object structure
type OutboundAccountPositionEventEvent struct {
	// Balances Array
	BalancesArray []OutboundAccountPositionEventEventBItem `json:"B,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Event type
	EventType string `json:"e,omitempty"`
	// Time of last account update (milliseconds)
	TimeOfLastAccountUpdate int64 `json:"u,omitempty"`
}

// OutboundAccountPositionEventEventBItem represents a nested object structure
type OutboundAccountPositionEventEventBItem struct {
	// Asset (example: "ETH")
	Asset string `json:"a,omitempty"`
	// Free (example: "10000.000000")
	Free string `json:"f,omitempty"`
	// Locked (example: "0.000000")
	Locked string `json:"l,omitempty"`
}

// OutboundAccountPositionEvent - Sent when account balance changes
// Message name: Outbound Account Position Event
type OutboundAccountPositionEvent struct {
	Event *OutboundAccountPositionEventEvent `json:"event,omitempty"`
}

// String returns string representation of OutboundAccountPositionEvent
func (s OutboundAccountPositionEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for OutboundAccountPositionEvent
func (s OutboundAccountPositionEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "outboundaccountpositionevent"
}

// GetEventTime returns the event timestamp for OutboundAccountPositionEvent
func (s OutboundAccountPositionEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


