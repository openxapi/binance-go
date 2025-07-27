package models

import (
	"encoding/json"
)

// EventStreamTerminatedEventEvent represents a nested object structure
type EventStreamTerminatedEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
}

// EventStreamTerminatedEvent - Sent when User Data Stream is stopped (WebSocket API only)
// Message name: Event Stream Terminated Event
type EventStreamTerminatedEvent struct {
	Event *EventStreamTerminatedEventEvent `json:"event,omitempty"`
}

// String returns string representation of EventStreamTerminatedEvent
func (s EventStreamTerminatedEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for EventStreamTerminatedEvent
func (s EventStreamTerminatedEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "eventstreamterminatedevent"
}

// GetEventTime returns the event timestamp for EventStreamTerminatedEvent
func (s EventStreamTerminatedEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


