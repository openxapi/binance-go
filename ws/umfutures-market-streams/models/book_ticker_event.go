package models

import (
	"encoding/json"
)

// BookTickerEvent represents BookTickerEvent
type BookTickerEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Order book updateId
	UpdateId int64 `json:"u,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Best bid price
	BuyerOrderId string `json:"b,omitempty"`
	// Best bid quantity
	BestBidQuantity string `json:"B,omitempty"`
	// Best ask price
	AggregateTradeId string `json:"a,omitempty"`
	// Best ask quantity
	SellerOrderId string `json:"A,omitempty"`
	// Transaction time
	TradeTime int64 `json:"T,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
}

// String returns string representation of BookTickerEvent
func (s BookTickerEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


