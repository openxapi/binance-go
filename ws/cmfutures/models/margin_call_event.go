package models

import (
	"encoding/json"
)

// MarginCallEventEvent represents a nested object structure
type MarginCallEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Cross Wallet Balance (only for cross position)
	CrossWalletBalance string `json:"cw,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Account Alias
	AccountAlias string `json:"i,omitempty"`
	// Position(s) of Margin Call
	PositionofMarginCall []MarginCallEventEventPItem `json:"p,omitempty"`
}

// MarginCallEventEventPItem represents a nested object structure
type MarginCallEventEventPItem struct {
	// Isolated Wallet (example: "0")
	IsolatedWallet string `json:"iw,omitempty"`
	// Maintenance Margin Required (example: "1.614445")
	MaintenanceMarginRequired string `json:"mm,omitempty"`
	// Mark Price (example: "187.17127")
	MarkPrice string `json:"mp,omitempty"`
	// Margin Type (example: "cross")
	MarginType string `json:"mt,omitempty"`
	// Position Amount (example: "1.327")
	PositionAmount string `json:"pa,omitempty"`
	// Position Side (example: "LONG")
	PositionSide string `json:"ps,omitempty"`
	// Symbol (example: "ETHUSD_PERP")
	Symbol string `json:"s,omitempty"`
	// Unrealized PnL (example: "-1.166074")
	UnrealizedPnL string `json:"up,omitempty"`
}

// MarginCallEvent - Sent when account is under margin call
// Message name: Margin Call Event
type MarginCallEvent struct {
	Event *MarginCallEventEvent `json:"event,omitempty"`
}

// String returns string representation of MarginCallEvent
func (s MarginCallEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for MarginCallEvent
func (s MarginCallEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "margincallevent"
}

// GetEventTime returns the event timestamp for MarginCallEvent
func (s MarginCallEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


