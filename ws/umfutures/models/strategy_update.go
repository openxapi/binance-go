package models

import (
	"encoding/json"
)

// StrategyUpdateEvent represents a nested object structure
type StrategyUpdateEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Strategy Update
	StrategyUpdate *StrategyUpdateEventStrategyUpdate `json:"su,omitempty"`
}

// StrategyUpdateEventStrategyUpdate represents a nested object structure
type StrategyUpdateEventStrategyUpdate struct {
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

// StrategyUpdate - Sent when strategy status changes
// Message name: Strategy Update Event
type StrategyUpdate struct {
	Event *StrategyUpdateEvent `json:"event,omitempty"`
}

// String returns string representation of StrategyUpdate
func (s StrategyUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for StrategyUpdate
func (s StrategyUpdate) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "strategyupdate"
}

// GetEventTime returns the event timestamp for StrategyUpdate
func (s StrategyUpdate) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


