package models

import (
	"encoding/json"
)

// GridUpdateEvent represents a nested object structure
type GridUpdateEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Grid Update
	GridUpdate *GridUpdateEventGridUpdate `json:"gu,omitempty"`
}

// GridUpdateEventGridUpdate represents a nested object structure
type GridUpdateEventGridUpdate struct {
	// Additional properties
	AdditionalProperties string `json:"mp,omitempty"`
	// Realized PNL
	RealizedPNL string `json:"r,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Strategy ID
	StrategyID int64 `json:"si,omitempty"`
	// Strategy Status
	StrategyStatus string `json:"ss,omitempty"`
	// Strategy Type
	StrategyType string `json:"st,omitempty"`
	// Unmatched Fee
	UnmatchedFee string `json:"uf,omitempty"`
	// Unmatched Average Price
	UnmatchedAveragePrice string `json:"up,omitempty"`
	// Unmatched Qty
	UnmatchedQty string `json:"uq,omitempty"`
}

// GridUpdate - Sent when grid strategy updates
// Message name: Grid Update Event
type GridUpdate struct {
	Event *GridUpdateEvent `json:"event,omitempty"`
}

// String returns string representation of GridUpdate
func (s GridUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for GridUpdate
func (s GridUpdate) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "gridupdate"
}

// GetEventTime returns the event timestamp for GridUpdate
func (s GridUpdate) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


