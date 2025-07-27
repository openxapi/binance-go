package models

import (
	"encoding/json"
)

// AccountUpdateEventEvent represents a nested object structure
type AccountUpdateEventEvent struct {
	// Account Balance Array
	AccountBalanceArray []AccountUpdateEventEventBItem `json:"B,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Greeks Array
	GreeksArray []AccountUpdateEventEventGItem `json:"G,omitempty"`
	// Positions Array
	PositionsArray []AccountUpdateEventEventPItem `json:"P,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// User ID
	UserID float64 `json:"uid,omitempty"`
}

// AccountUpdateEventEventBItem represents a nested object structure
type AccountUpdateEventEventBItem struct {
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

// AccountUpdateEventEventGItem represents a nested object structure
type AccountUpdateEventEventGItem struct {
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

// AccountUpdateEventEventPItem represents a nested object structure
type AccountUpdateEventEventPItem struct {
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

// AccountUpdateEvent - Sent when account balance or position changes
// Message name: Account Update Event
type AccountUpdateEvent struct {
	Event *AccountUpdateEventEvent `json:"event,omitempty"`
}

// String returns string representation of AccountUpdateEvent
func (s AccountUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for AccountUpdateEvent
func (s AccountUpdateEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "accountupdateevent"
}

// GetEventTime returns the event timestamp for AccountUpdateEvent
func (s AccountUpdateEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


