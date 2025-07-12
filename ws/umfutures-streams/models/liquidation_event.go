package models

import (
	"encoding/json"
)

// LiquidationEvent represents LiquidationEvent
type LiquidationEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Liquidation order
	OpenPrice *LiquidationEventOpenPrice `json:"o,omitempty"`
}

// LiquidationEventOpenPrice represents the openprice details
type LiquidationEventOpenPrice struct {
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Side
	Side string `json:"S,omitempty"`
	// Order type
	OpenPrice string `json:"o,omitempty"`
	// Time in force
	FirstTradeId string `json:"f,omitempty"`
	// Original quantity
	Quantity string `json:"q,omitempty"`
	// Price
	Price string `json:"p,omitempty"`
	// Average price
	AveragePrice string `json:"ap,omitempty"`
	// Order status
	OrderStatus string `json:"X,omitempty"`
	// Last filled quantity
	LastTradeId string `json:"l,omitempty"`
	// Accumulated filled quantity
	AccumulatedFilledQuantity string `json:"z,omitempty"`
	// Trade time
	TradeTime int64 `json:"T,omitempty"`
}

// String returns string representation of LiquidationEvent
func (s LiquidationEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


