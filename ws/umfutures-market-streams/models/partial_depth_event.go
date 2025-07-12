package models

import (
	"encoding/json"
)

// PartialDepthEvent represents PartialDepthEvent
type PartialDepthEvent struct {
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
	// Final update ID in last stream
	FinalUpdateIdInLastStream int64 `json:"pu,omitempty"`
	// Bids
	BuyerOrderId [][]string `json:"b,omitempty"`
	// Asks
	AggregateTradeId [][]string `json:"a,omitempty"`
}

// String returns string representation of PartialDepthEvent
func (s PartialDepthEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


