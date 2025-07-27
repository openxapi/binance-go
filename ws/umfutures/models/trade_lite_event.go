package models

import (
	"encoding/json"
)

// TradeLiteEventEvent represents a nested object structure
type TradeLiteEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Last Filled Price
	LastFilledPrice string `json:"L,omitempty"`
	// Side
	Side string `json:"S,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Client Order Id
	ClientOrderId string `json:"c,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Order Id
	OrderId int64 `json:"i,omitempty"`
	// Order Last Filled Quantity
	OrderLastFilledQuantity string `json:"l,omitempty"`
	// Maker side
	MakerSide bool `json:"m,omitempty"`
	// Original Price
	OriginalPrice string `json:"p,omitempty"`
	// Original Quantity
	OriginalQuantity string `json:"q,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Trade Id
	TradeId int64 `json:"t,omitempty"`
}

// TradeLiteEvent - Simplified trade execution event
// Message name: Trade Lite Event
type TradeLiteEvent struct {
	Event *TradeLiteEventEvent `json:"event,omitempty"`
}

// String returns string representation of TradeLiteEvent
func (s TradeLiteEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for TradeLiteEvent
func (s TradeLiteEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "tradeliteevent"
}

// GetEventTime returns the event timestamp for TradeLiteEvent
func (s TradeLiteEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


