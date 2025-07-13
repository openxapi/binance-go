package models

import (
	"encoding/json"
)

// OutboundAccountPositionEvent represents a nested object structure
type OutboundAccountPositionEvent struct {
	// Balances Array
	BalancesArray []OutboundAccountPositionEventBItem `json:"B,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Event type
	EventType string `json:"e,omitempty"`
	// Time of last account update (milliseconds)
	TimeOfLastAccountUpdate int64 `json:"u,omitempty"`
}

// OutboundAccountPositionEventBItem represents a nested object structure
type OutboundAccountPositionEventBItem struct {
	// Asset (example: "ETH")
	Asset string `json:"a,omitempty"`
	// Free (example: "10000.000000")
	Free string `json:"f,omitempty"`
	// Locked (example: "0.000000")
	Locked string `json:"l,omitempty"`
}

// OutboundAccountPosition - Sent when account balance changes
// Message name: Outbound Account Position Event
type OutboundAccountPosition struct {
	Event *OutboundAccountPositionEvent `json:"event,omitempty"`
}

// String returns string representation of OutboundAccountPosition
func (s OutboundAccountPosition) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for OutboundAccountPosition
func (s OutboundAccountPosition) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "outboundaccountposition"
}

// GetEventTime returns the event timestamp for OutboundAccountPosition
func (s OutboundAccountPosition) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


