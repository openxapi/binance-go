package models

import (
	"encoding/json"
)

// RiskLevelChange - Sent when user's position risk ratio is too high
// Message name: Risk Level Change Event
type RiskLevelChange struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// UniMMR level
	UniMMRLevel string `json:"u,omitempty"`
	// Risk level status
	RiskLevelStatus string `json:"s,omitempty"`
	// Account equity in USD
	AccountEquityInUSD string `json:"eq,omitempty"`
	// Actual equity without collateral rate in USD
	ActualEquityWithoutCollateralRateInUSD string `json:"ae,omitempty"`
	// Total maintenance margin in USD
	TotalMaintenanceMarginInUSD string `json:"m,omitempty"`
}

// String returns string representation of RiskLevelChange
func (s RiskLevelChange) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for RiskLevelChange
func (s RiskLevelChange) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "risklevelchange"
}

// GetEventTime returns the event timestamp for RiskLevelChange
func (s RiskLevelChange) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


