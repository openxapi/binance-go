package models

import (
	"encoding/json"
)

// FuturesAccountConfigUpdateEventAccountConfigurationDetails represents a nested object structure
type FuturesAccountConfigUpdateEventAccountConfigurationDetails struct {
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Leverage
	Leverage int64 `json:"l,omitempty"`
}

// FuturesAccountConfigUpdateEvent - Futures account configuration update event
// Message name: Futures Account Configuration Update Event
type FuturesAccountConfigUpdateEvent struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event business unit
	EventBusinessUnit string `json:"fs,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Account configuration details
	AccountConfigurationDetails *FuturesAccountConfigUpdateEventAccountConfigurationDetails `json:"ac,omitempty"`
}

// String returns string representation of FuturesAccountConfigUpdateEvent
func (s FuturesAccountConfigUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for FuturesAccountConfigUpdateEvent
func (s FuturesAccountConfigUpdateEvent) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "futuresaccountconfigupdateevent"
}

// GetEventTime returns the event timestamp for FuturesAccountConfigUpdateEvent
func (s FuturesAccountConfigUpdateEvent) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


