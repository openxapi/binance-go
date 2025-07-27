package models

import (
	"encoding/json"
)

// ExternalLockUpdateEventEvent represents a nested object structure
type ExternalLockUpdateEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Asset
	Asset string `json:"a,omitempty"`
	// Delta
	Delta string `json:"d,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
}

// ExternalLockUpdateEvent - Sent when spot wallet balance is locked/unlocked by external system
// Message name: External Lock Update Event
type ExternalLockUpdateEvent struct {
	Event *ExternalLockUpdateEventEvent `json:"event,omitempty"`
}

// String returns string representation of ExternalLockUpdateEvent
func (s ExternalLockUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ExternalLockUpdateEvent
func (s ExternalLockUpdateEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "externallockupdateevent"
}

// GetEventTime returns the event timestamp for ExternalLockUpdateEvent
func (s ExternalLockUpdateEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


