package models

import (
	"encoding/json"
)

// CompositeIndexEvent represents CompositeIndexEvent
type CompositeIndexEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Price
	Price string `json:"p,omitempty"`
	// Asset type
	CloseTime string `json:"C,omitempty"`
	// Composition
	ClosePrice []CompositeIndexEventClosePriceItem `json:"c,omitempty"`
}

// CompositeIndexEventClosePriceItem represents the closeprice item details
type CompositeIndexEventClosePriceItem struct {
	// Base asset
	BuyerOrderId string `json:"b,omitempty"`
	// Quote asset
	Quantity string `json:"q,omitempty"`
	// Weight in quantity
	WeightedAveragePrice string `json:"w,omitempty"`
	// Weight in percentage
	WeightInPercentage string `json:"W,omitempty"`
	// Index price
	Interval string `json:"i,omitempty"`
}

// String returns string representation of CompositeIndexEvent
func (s CompositeIndexEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


