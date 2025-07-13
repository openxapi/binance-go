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
	LiquidationOrder *LiquidationEventLiquidationOrder `json:"o,omitempty"`
}

// LiquidationEventLiquidationOrder represents the liquidationorder details
type LiquidationEventLiquidationOrder struct {
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Side
	Side string `json:"S,omitempty"`
	// Order type
	OrderType string `json:"o,omitempty"`
	// Time in force
	TimeInForce string `json:"f,omitempty"`
	// Original quantity
	OriginalQuantity string `json:"q,omitempty"`
	// Price
	Price string `json:"p,omitempty"`
	// Average price
	AveragePrice string `json:"ap,omitempty"`
	// Order status
	OrderStatus string `json:"X,omitempty"`
	// Last filled quantity
	LastFilledQuantity string `json:"l,omitempty"`
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


