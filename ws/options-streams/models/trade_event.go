package models

import (
	"encoding/json"
)

// TradeEvent represents TradeEvent
// Real-time trade information for options contracts.
// Stream name pattern: {symbol}@trade or {underlyingAsset}@trade
// Example: BTC-200630-9000-P@trade or ETH@trade
// Update frequency: 50ms
// Provides raw trade information for specific symbols or underlying assets.
// 
type TradeEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event timestamp
	EventTime int64 `json:"E,omitempty"`
	// Option trading symbol
	Symbol string `json:"s,omitempty"`
	// Trade ID
	TradeId int64 `json:"t,omitempty"`
	// Price
	Price string `json:"p,omitempty"`
	// Quantity (can be negative)
	Quantity string `json:"q,omitempty"`
	// Buy order ID
	BuyerOrderId int64 `json:"b,omitempty"`
	// Sell order ID
	SellerOrderId int64 `json:"a,omitempty"`
	// Trade completed timestamp
	TradeTime int64 `json:"T,omitempty"`
	// Trade direction (-1 indicates direction)
	TradeDirection string `json:"S,omitempty"`
	// Trade type
	TradeType string `json:"X,omitempty"`
}

// String returns string representation of TradeEvent
func (s TradeEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


