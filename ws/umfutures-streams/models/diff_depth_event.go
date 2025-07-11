package models

import (
	"encoding/json"
)

// DiffDepthEvent represents DiffDepthEvent
type DiffDepthEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Transaction time
	TradeTime int64 `json:"T,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// First update ID in event
	FirstUpdateId int64 `json:"U,omitempty"`
	// Final update ID in event
	UpdateId int64 `json:"u,omitempty"`
	// Previous final update ID in event
	PreviousFinalUpdateIdInEvent int64 `json:"pu,omitempty"`
	// Bids to be updated
	BuyerOrderId [][]string `json:"b,omitempty"`
	// Asks to be updated
	AggregateTradeId [][]string `json:"a,omitempty"`
}

// String returns string representation of DiffDepthEvent
func (s DiffDepthEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


