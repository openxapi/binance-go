package models

import (
	"encoding/json"
)

// ListenKeyExpiredEvent represents a nested object structure
type ListenKeyExpiredEvent struct {
	// Event time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Event type
	EventType string `json:"e,omitempty"`
	// The expired listen key
	TheExpiredListenKey string `json:"listenKey,omitempty"`
}

// ListenKeyExpired - Sent when the listen key expires
// Message name: Listen Key Expired Event
type ListenKeyExpired struct {
	Event *ListenKeyExpiredEvent `json:"event,omitempty"`
}

// String returns string representation of ListenKeyExpired
func (s ListenKeyExpired) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ListenKeyExpired
func (s ListenKeyExpired) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "listenkeyexpired"
}

// GetEventTime returns the event timestamp for ListenKeyExpired
func (s ListenKeyExpired) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


