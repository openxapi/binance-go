package models

import (
	"encoding/json"
)

// GridUpdateEventEvent represents a nested object structure
type GridUpdateEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Grid Update
	GridUpdate *GridUpdateEventEventGridUpdate `json:"gu,omitempty"`
}

// GridUpdateEventEventGridUpdate represents a nested object structure
type GridUpdateEventEventGridUpdate struct {
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

// GridUpdateEvent - Sent when grid strategy updates
// Message name: Grid Update Event
type GridUpdateEvent struct {
	Event *GridUpdateEventEvent `json:"event,omitempty"`
}

// String returns string representation of GridUpdateEvent
func (s GridUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for GridUpdateEvent
func (s GridUpdateEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "gridupdateevent"
}

// GetEventTime returns the event timestamp for GridUpdateEvent
func (s GridUpdateEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


