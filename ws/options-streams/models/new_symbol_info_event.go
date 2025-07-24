package models

import (
	"encoding/json"
)

// NewSymbolInfoEvent represents NewSymbolInfoEvent
// New option symbol information event (option_pair stream).
// Provides real-time updates about new option symbol listings with comprehensive contract details.
// Stream name: option_pair
// Update frequency: 50ms
// 
type NewSymbolInfoEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time (timestamp)
	EventTime int64 `json:"E,omitempty"`
	// Underlying index
	UnderlyingIndex string `json:"u,omitempty"`
	// Quotation asset
	QuotationAsset string `json:"qa,omitempty"`
	// Trading pair name (option symbol)
	Symbol string `json:"s,omitempty"`
	// Conversion ratio (contract's underlying asset representation ratio)
	ConversionRatio float64 `json:"unit,omitempty"`
	// Minimum trade volume
	MinimumTradeVolume string `json:"mq,omitempty"`
	// Option type
	OptionType string `json:"d,omitempty"`
	// Strike price
	StrikePrice string `json:"sp,omitempty"`
	// Expiration time (timestamp)
	ExpirationTime int64 `json:"ed,omitempty"`
}

// String returns string representation of NewSymbolInfoEvent
func (s NewSymbolInfoEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


