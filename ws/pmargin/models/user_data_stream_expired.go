package models

import (
	"encoding/json"
)

// UserDataStreamExpired - Sent when listen key has expired
// Message name: User Data Stream Expired Event
type UserDataStreamExpired struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
}

// String returns string representation of UserDataStreamExpired
func (s UserDataStreamExpired) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for UserDataStreamExpired
func (s UserDataStreamExpired) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "userdatastreamexpired"
}

// GetEventTime returns the event timestamp for UserDataStreamExpired
func (s UserDataStreamExpired) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


