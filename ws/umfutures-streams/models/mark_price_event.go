package models

import (
	"encoding/json"
)

// MarkPriceEvent represents MarkPriceEvent
type MarkPriceEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Mark price
	MarkPrice string `json:"p,omitempty"`
	// Estimated settle price (only for quarterly contract)
	EstimatedSettlePrice string `json:"P,omitempty"`
	// Index price
	IndexPrice string `json:"i,omitempty"`
	// Funding rate
	FundingRate string `json:"r,omitempty"`
	// Next funding time
	NextFundingTime int64 `json:"T,omitempty"`
}

// String returns string representation of MarkPriceEvent
func (s MarkPriceEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


