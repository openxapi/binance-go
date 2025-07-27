package models

import (
	"encoding/json"
)

// RiskLevelChangeEventEvent represents a nested object structure
type RiskLevelChangeEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Margin Balance
	MarginBalance string `json:"mb,omitempty"`
	// Maintenance Margin
	MaintenanceMargin string `json:"mm,omitempty"`
	// Risk Level
	RiskLevel string `json:"s,omitempty"`
}

// RiskLevelChangeEvent - Sent when risk level changes (applicable only to VIP and Market Maker accounts)
// Message name: Risk Level Change Event
type RiskLevelChangeEvent struct {
	Event *RiskLevelChangeEventEvent `json:"event,omitempty"`
}

// String returns string representation of RiskLevelChangeEvent
func (s RiskLevelChangeEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for RiskLevelChangeEvent
func (s RiskLevelChangeEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "risklevelchangeevent"
}

// GetEventTime returns the event timestamp for RiskLevelChangeEvent
func (s RiskLevelChangeEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


