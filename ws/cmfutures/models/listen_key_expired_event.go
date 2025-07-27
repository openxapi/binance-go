package models

import (
	"encoding/json"
)

// ListenKeyExpiredEventEvent represents a nested object structure
type ListenKeyExpiredEventEvent struct {
	// Event time (timestamp)
	EventTime int64 `json:"E,omitempty"`
	// Event type
	EventType string `json:"e,omitempty"`
	// The expired listen key
	TheExpiredListenKey string `json:"listenKey,omitempty"`
}

// ListenKeyExpiredEvent - Sent when the listen key expires
// Message name: Listen Key Expired Event
type ListenKeyExpiredEvent struct {
	Event *ListenKeyExpiredEventEvent `json:"event,omitempty"`
}

// String returns string representation of ListenKeyExpiredEvent
func (s ListenKeyExpiredEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ListenKeyExpiredEvent
func (s ListenKeyExpiredEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "listenkeyexpiredevent"
}

// GetEventTime returns the event timestamp for ListenKeyExpiredEvent
func (s ListenKeyExpiredEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


