package models

import (
	"encoding/json"
)

// RiskLevelChangeEvent - Sent when user's position risk ratio is too high
// Message name: Risk Level Change Event
type RiskLevelChangeEvent struct {
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

// String returns string representation of RiskLevelChangeEvent
func (s RiskLevelChangeEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for RiskLevelChangeEvent
func (s RiskLevelChangeEvent) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "risklevelchangeevent"
}

// GetEventTime returns the event timestamp for RiskLevelChangeEvent
func (s RiskLevelChangeEvent) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


