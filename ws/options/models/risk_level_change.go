package models

import (
	"encoding/json"
)

// RiskLevelChangeEvent represents a nested object structure
type RiskLevelChangeEvent struct {
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

// RiskLevelChange - Sent when risk level changes (applicable only to VIP and Market Maker accounts)
// Message name: Risk Level Change Event
type RiskLevelChange struct {
	Event *RiskLevelChangeEvent `json:"event,omitempty"`
}

// String returns string representation of RiskLevelChange
func (s RiskLevelChange) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for RiskLevelChange
func (s RiskLevelChange) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "risklevelchange"
}

// GetEventTime returns the event timestamp for RiskLevelChange
func (s RiskLevelChange) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


