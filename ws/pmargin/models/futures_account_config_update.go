package models

import (
	"encoding/json"
)

// FuturesAccountConfigUpdateAccountConfigurationDetails represents a nested object structure
type FuturesAccountConfigUpdateAccountConfigurationDetails struct {
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Leverage
	Leverage int64 `json:"l,omitempty"`
}

// FuturesAccountConfigUpdate - Futures account configuration update event
// Message name: Futures Account Configuration Update Event
type FuturesAccountConfigUpdate struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event business unit
	EventBusinessUnit string `json:"fs,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Account configuration details
	AccountConfigurationDetails *FuturesAccountConfigUpdateAccountConfigurationDetails `json:"ac,omitempty"`
}

// String returns string representation of FuturesAccountConfigUpdate
func (s FuturesAccountConfigUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for FuturesAccountConfigUpdate
func (s FuturesAccountConfigUpdate) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "futuresaccountconfigupdate"
}

// GetEventTime returns the event timestamp for FuturesAccountConfigUpdate
func (s FuturesAccountConfigUpdate) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


