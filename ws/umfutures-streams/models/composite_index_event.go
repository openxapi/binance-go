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
	AssetType string `json:"C,omitempty"`
	// Composition
	Composition []CompositeIndexEventCompositionItem `json:"c,omitempty"`
}

// CompositeIndexEventCompositionItem represents the composition item details
type CompositeIndexEventCompositionItem struct {
	// Base asset
	BaseAsset string `json:"b,omitempty"`
	// Quote asset
	QuoteAsset string `json:"q,omitempty"`
	// Weight in quantity
	WeightInQuantity string `json:"w,omitempty"`
	// Weight in percentage
	WeightInPercentage string `json:"W,omitempty"`
	// Index price
	IndexPrice string `json:"i,omitempty"`
}

// String returns string representation of CompositeIndexEvent
func (s CompositeIndexEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


