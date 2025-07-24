package models

import (
	"encoding/json"
)

// AggTrade represents Aggregate Trade
type AggTrade struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Aggregate trade ID
	AggregateTradeId int64 `json:"a,omitempty"`
	// Price
	Price string `json:"p,omitempty"`
	// Quantity
	Quantity string `json:"q,omitempty"`
	// First trade ID
	FirstTradeId int64 `json:"f,omitempty"`
	// Last trade ID
	LastTradeId int64 `json:"l,omitempty"`
	// Trade time
	TradeTime int64 `json:"T,omitempty"`
	// Is the buyer the market maker
	IsBuyerMaker bool `json:"m,omitempty"`
}

// String returns string representation of AggTrade
func (s AggTrade) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewAggTrade creates a new AggTrade instance
func NewAggTrade() *AggTrade {
	return &AggTrade{}
}


