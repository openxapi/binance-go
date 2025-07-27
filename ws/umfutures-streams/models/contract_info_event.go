package models

import (
	"encoding/json"
)

// ContractInfoEvent represents ContractInfoEvent
type ContractInfoEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Contract status
	ContractStatus string `json:"cs,omitempty"`
	// Pair
	Pair string `json:"ps,omitempty"`
	// Base asset
	BaseAsset string `json:"bs,omitempty"`
	// Quote asset
	QuoteAsset string `json:"qs,omitempty"`
	// Trigger protect
	TriggerProtect int64 `json:"ts,omitempty"`
	// Margin asset
	MarginAsset string `json:"ma,omitempty"`
	// Settlement price
	SettlementPrice string `json:"sp,omitempty"`
	// Settlement fee
	SettlementFee string `json:"sf,omitempty"`
}

// String returns string representation of ContractInfoEvent
func (s ContractInfoEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


