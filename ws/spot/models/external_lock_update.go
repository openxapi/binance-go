package models

import (
	"encoding/json"
)

// ExternalLockUpdateEvent represents a nested object structure
type ExternalLockUpdateEvent struct {
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

// ExternalLockUpdate - Sent when spot wallet balance is locked/unlocked by external system
// Message name: External Lock Update Event
type ExternalLockUpdate struct {
	Event ExternalLockUpdateEvent `json:"event,omitempty"`
}

// String returns string representation of ExternalLockUpdate
func (s ExternalLockUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ExternalLockUpdate
func (s ExternalLockUpdate) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "externallockupdate"
}

// GetEventTime returns the event timestamp for ExternalLockUpdate
func (s ExternalLockUpdate) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


