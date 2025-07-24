package models

import (
	"encoding/json"
)

// PartialDepthEvent represents PartialDepthEvent
// Partial order book depth information for options contracts.
// Stream name patterns:
// - {symbol}@depth{levels} (e.g., BTC-200630-9000-P@depth10)
// - {symbol}@depth{levels}@{speed} (e.g., BTC-200630-9000-P@depth10@100ms)
// 
// Supported depth levels: 10, 20, 50, 100
// Supported update speeds: 100ms, 500ms (default), 1000ms
// 
type PartialDepthEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time (timestamp)
	EventTime int64 `json:"E,omitempty"`
	// Transaction time (timestamp)
	TransactionTime int64 `json:"T,omitempty"`
	// Option symbol
	Symbol string `json:"s,omitempty"`
	// Update ID in event
	FinalUpdateId int64 `json:"u,omitempty"`
	// Previous update ID
	PrevFinalUpdateId int64 `json:"pu,omitempty"`
	// Buy orders (bids)
	Bids [][]string `json:"b,omitempty"`
	// Sell orders (asks)
	Asks [][]string `json:"a,omitempty"`
}

// String returns string representation of PartialDepthEvent
func (s PartialDepthEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


