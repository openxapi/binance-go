package models

import (
	"encoding/json"
)

// EventStreamTerminatedEvent represents a nested object structure
type EventStreamTerminatedEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
}

// EventStreamTerminated - Sent when User Data Stream is stopped (WebSocket API only)
// Message name: Event Stream Terminated Event
type EventStreamTerminated struct {
	Event EventStreamTerminatedEvent `json:"event,omitempty"`
}

// String returns string representation of EventStreamTerminated
func (s EventStreamTerminated) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for EventStreamTerminated
func (s EventStreamTerminated) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "eventstreamterminated"
}

// GetEventTime returns the event timestamp for EventStreamTerminated
func (s EventStreamTerminated) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


