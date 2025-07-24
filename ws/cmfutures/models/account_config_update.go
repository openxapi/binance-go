package models

import (
	"encoding/json"
)

// AccountConfigUpdateEvent represents a nested object structure
type AccountConfigUpdateEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Account Configuration for Trade Pair
	AccountConfigurationForTradePair *AccountConfigUpdateEventAccountConfigurationForTradePair `json:"ac,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Account Alias
	AccountAlias string `json:"i,omitempty"`
}

// AccountConfigUpdateEventAccountConfigurationForTradePair represents a nested object structure
type AccountConfigUpdateEventAccountConfigurationForTradePair struct {
	// Leverage
	Leverage int64 `json:"l,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
}

// AccountConfigUpdate - Sent when account configuration changes (leverage update)
// Message name: Account Configuration Update Event
type AccountConfigUpdate struct {
	Event *AccountConfigUpdateEvent `json:"event,omitempty"`
}

// String returns string representation of AccountConfigUpdate
func (s AccountConfigUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for AccountConfigUpdate
func (s AccountConfigUpdate) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "accountconfigupdate"
}

// GetEventTime returns the event timestamp for AccountConfigUpdate
func (s AccountConfigUpdate) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


