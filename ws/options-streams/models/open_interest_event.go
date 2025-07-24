package models

import (
	"encoding/json"
)

// OpenInterestEvent represents OpenInterestEvent
// Open interest information event for options contracts.
// Stream name pattern: {underlyingAsset}@openInterest@{expirationDate}
// Example: ETH@openInterest@221125
// Update frequency: 60s (once per minute)
// 
type OpenInterestEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event timestamp
	EventTime int64 `json:"E,omitempty"`
	// Option symbol
	Symbol string `json:"s,omitempty"`
	// Open interest in contracts
	OpenInterestInContracts string `json:"o,omitempty"`
	// Open interest in USDT
	OpenInterestInUsdt string `json:"h,omitempty"`
}

// String returns string representation of OpenInterestEvent
func (s OpenInterestEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


