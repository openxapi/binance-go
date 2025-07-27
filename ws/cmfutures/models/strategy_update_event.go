package models

import (
	"encoding/json"
)

// StrategyUpdateEventEvent represents a nested object structure
type StrategyUpdateEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Strategy Update
	StrategyUpdate *StrategyUpdateEventEventStrategyUpdate `json:"su,omitempty"`
}

// StrategyUpdateEventEventStrategyUpdate represents a nested object structure
type StrategyUpdateEventEventStrategyUpdate struct {
	// Operation Code
	OperationCode int64 `json:"c,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Strategy ID
	StrategyID int64 `json:"si,omitempty"`
	// Strategy Status
	StrategyStatus string `json:"ss,omitempty"`
	// Strategy Type
	StrategyType string `json:"st,omitempty"`
	// Update Time (milliseconds)
	UpdateTime int64 `json:"ut,omitempty"`
}

// StrategyUpdateEvent - Sent when strategy status changes
// Message name: Strategy Update Event
type StrategyUpdateEvent struct {
	Event *StrategyUpdateEventEvent `json:"event,omitempty"`
}

// String returns string representation of StrategyUpdateEvent
func (s StrategyUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for StrategyUpdateEvent
func (s StrategyUpdateEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "strategyupdateevent"
}

// GetEventTime returns the event timestamp for StrategyUpdateEvent
func (s StrategyUpdateEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


