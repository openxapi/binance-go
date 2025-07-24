package models

import (
	"encoding/json"
)

// AccountUpdateEvent represents a nested object structure
type AccountUpdateEvent struct {
	// Account Balance Array
	AccountBalanceArray []AccountUpdateEventBItem `json:"B,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Greeks Array
	GreeksArray []AccountUpdateEventGItem `json:"G,omitempty"`
	// Positions Array
	PositionsArray []AccountUpdateEventPItem `json:"P,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// User ID
	UserID float64 `json:"uid,omitempty"`
}

// AccountUpdateEventBItem represents a nested object structure
type AccountUpdateEventBItem struct {
	// Maintenance margin (example: "100.00000000")
	MaintenanceMargin string `json:"M,omitempty"`
	// Positive unrealized profit for long position (example: 25.5)
	PositiveUnrealizedProfitForLongPosition float64 `json:"U,omitempty"`
	// Margin asset (example: "USDT")
	MarginAsset string `json:"a,omitempty"`
	// Account balance (example: "1000.00000000")
	AccountBalance string `json:"b,omitempty"`
	// Initial margin (example: "200.00000000")
	InitialMargin string `json:"i,omitempty"`
	// Position value (example: "500.00000000")
	PositionValue string `json:"m,omitempty"`
	// Unrealized profit/loss (example: "50.00000000")
	UnrealizedProfitLoss string `json:"u,omitempty"`
}

// AccountUpdateEventGItem represents a nested object structure
type AccountUpdateEventGItem struct {
	// Delta (example: 0.5)
	Delta float64 `json:"d,omitempty"`
	// Gamma (example: 0.01)
	Gamma float64 `json:"g,omitempty"`
	// Theta (example: -0.1)
	Theta float64 `json:"t,omitempty"`
	// Underlying (example: "BTCUSDT")
	Underlying string `json:"ui,omitempty"`
	// Vega (example: 0.05)
	Vega float64 `json:"v,omitempty"`
}

// AccountUpdateEventPItem represents a nested object structure
type AccountUpdateEventPItem struct {
	// Average entry price (example: "100.00000000")
	AverageEntryPrice string `json:"a,omitempty"`
	// Current positions (example: "10")
	CurrentPositions string `json:"c,omitempty"`
	// Position value (example: "1000.00000000")
	PositionValue string `json:"p,omitempty"`
	// Positions that can be reduced (example: "5")
	PositionsThatCanBeReduced string `json:"r,omitempty"`
	// Contract symbol (example: "BTC-220930-18000-C")
	ContractSymbol string `json:"s,omitempty"`
}

// AccountUpdate - Sent when account balance or position changes
// Message name: Account Update Event
type AccountUpdate struct {
	Event *AccountUpdateEvent `json:"event,omitempty"`
}

// String returns string representation of AccountUpdate
func (s AccountUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for AccountUpdate
func (s AccountUpdate) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "accountupdate"
}

// GetEventTime returns the event timestamp for AccountUpdate
func (s AccountUpdate) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


