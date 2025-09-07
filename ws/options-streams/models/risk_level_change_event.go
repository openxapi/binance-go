package models

import (
	"encoding/json"
)

// RiskLevelChangeEvent represents RiskLevelChangeEvent
// Risk Level Change - sent when risk level changes (applicable only to VIP and Market Maker accounts)
type RiskLevelChangeEvent struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Margin Balance
	MarginBalance string `json:"mb,omitempty"`
	// Maintenance Margin
	MaintenanceMargin string `json:"mm,omitempty"`
	// Risk Level
	Symbol string `json:"s,omitempty"`
}

// String returns string representation of RiskLevelChangeEvent
func (s RiskLevelChangeEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


