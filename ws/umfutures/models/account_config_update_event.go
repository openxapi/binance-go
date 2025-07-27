package models

import (
	"encoding/json"
)

// AccountConfigUpdateEventEvent represents a nested object structure
type AccountConfigUpdateEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Account Configuration for Trade Pair
	AccountConfigurationForTradePair *AccountConfigUpdateEventEventAccountConfigurationForTradePair `json:"ac,omitempty"`
	// User Account Configuration
	UserAccountConfiguration *AccountConfigUpdateEventEventUserAccountConfiguration `json:"ai,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
}

// AccountConfigUpdateEventEventAccountConfigurationForTradePair represents a nested object structure
type AccountConfigUpdateEventEventAccountConfigurationForTradePair struct {
	// Leverage
	Leverage int64 `json:"l,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
}

// AccountConfigUpdateEventEventUserAccountConfiguration represents a nested object structure
type AccountConfigUpdateEventEventUserAccountConfiguration struct {
	// Multi-Assets Mode
	MultiAssetsMode bool `json:"j,omitempty"`
}

// AccountConfigUpdateEvent - Sent when account configuration changes (leverage update)
// Message name: Account Configuration Update Event
type AccountConfigUpdateEvent struct {
	Event *AccountConfigUpdateEventEvent `json:"event,omitempty"`
}

// String returns string representation of AccountConfigUpdateEvent
func (s AccountConfigUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for AccountConfigUpdateEvent
func (s AccountConfigUpdateEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "accountconfigupdateevent"
}

// GetEventTime returns the event timestamp for AccountConfigUpdateEvent
func (s AccountConfigUpdateEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


