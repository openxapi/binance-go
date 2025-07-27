package models

import (
	"encoding/json"
)

// UserDataStreamExpiredEvent - Sent when listen key has expired
// Message name: User Data Stream Expired Event
type UserDataStreamExpiredEvent struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
}

// String returns string representation of UserDataStreamExpiredEvent
func (s UserDataStreamExpiredEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for UserDataStreamExpiredEvent
func (s UserDataStreamExpiredEvent) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "userdatastreamexpiredevent"
}

// GetEventTime returns the event timestamp for UserDataStreamExpiredEvent
func (s UserDataStreamExpiredEvent) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


