package models

import (
	"encoding/json"
)

// LiabilityUpdate - Margin liability update event
// Message name: Liability Update Event
type LiabilityUpdate struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Asset
	Asset string `json:"a,omitempty"`
	// Type
	Type string `json:"t,omitempty"`
	// Transaction ID
	TransactionID int64 `json:"T,omitempty"`
	// Principal
	Principal string `json:"p,omitempty"`
	// Interest
	Interest string `json:"i,omitempty"`
	// Total Liability
	TotalLiability string `json:"l,omitempty"`
}

// String returns string representation of LiabilityUpdate
func (s LiabilityUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for LiabilityUpdate
func (s LiabilityUpdate) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "liabilityupdate"
}

// GetEventTime returns the event timestamp for LiabilityUpdate
func (s LiabilityUpdate) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


