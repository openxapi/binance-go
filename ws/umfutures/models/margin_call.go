package models

import (
	"encoding/json"
)

// MarginCallEvent represents a nested object structure
type MarginCallEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Cross Wallet Balance (only for crossed position)
	CrossWalletBalance string `json:"cw,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Position(s) of Margin Call
	PositionofMarginCall []MarginCallEventPItem `json:"p,omitempty"`
}

// MarginCallEventPItem represents a nested object structure
type MarginCallEventPItem struct {
	// Isolated Wallet (example: "0")
	IsolatedWallet string `json:"iw,omitempty"`
	// Maintenance Margin Required (example: "1.614445")
	MaintenanceMarginRequired string `json:"mm,omitempty"`
	// Mark Price (example: "187.17127")
	MarkPrice string `json:"mp,omitempty"`
	// Margin Type (example: "crossed")
	MarginType string `json:"mt,omitempty"`
	// Position Amount (example: "1.327")
	PositionAmount string `json:"pa,omitempty"`
	// Position Side (example: "LONG")
	PositionSide string `json:"ps,omitempty"`
	// Symbol (example: "ETHUSDT")
	Symbol string `json:"s,omitempty"`
	// Unrealized PnL (example: "-1.166074")
	UnrealizedPnL string `json:"up,omitempty"`
}

// MarginCall - Sent when account is under margin call
// Message name: Margin Call Event
type MarginCall struct {
	Event *MarginCallEvent `json:"event,omitempty"`
}

// String returns string representation of MarginCall
func (s MarginCall) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for MarginCall
func (s MarginCall) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "margincall"
}

// GetEventTime returns the event timestamp for MarginCall
func (s MarginCall) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


