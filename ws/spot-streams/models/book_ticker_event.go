package models

import (
	"encoding/json"
)

// BookTickerEvent represents BookTickerEvent
type BookTickerEvent struct {
	// Order book updateId
	UpdateId int64 `json:"u,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Best bid price
	BestBidPrice string `json:"b,omitempty"`
	// Best bid quantity
	BestBidQuantity string `json:"B,omitempty"`
	// Best ask price
	BestAskPrice string `json:"a,omitempty"`
	// Best ask quantity
	BestAskQuantity string `json:"A,omitempty"`
}

// String returns string representation of BookTickerEvent
func (s BookTickerEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


