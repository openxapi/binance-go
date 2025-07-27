package models

import (
	"encoding/json"
)

// TradeEvent represents TradeEvent
type TradeEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Trade ID
	TradeId int64 `json:"t,omitempty"`
	// Price
	Price string `json:"p,omitempty"`
	// Quantity
	Quantity string `json:"q,omitempty"`
	// Buyer order ID
	BuyerOrderId int64 `json:"b,omitempty"`
	// Seller order ID
	SellerOrderId int64 `json:"a,omitempty"`
	// Trade time
	TradeTime int64 `json:"T,omitempty"`
	// Is the buyer the market maker
	IsBuyerMaker bool `json:"m,omitempty"`
	// Ignore
	Ignore bool `json:"M,omitempty"`
}

// String returns string representation of TradeEvent
func (s TradeEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


