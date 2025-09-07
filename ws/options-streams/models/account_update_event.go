package models

import (
	"encoding/json"
)

// AccountUpdateEvent represents AccountUpdateEvent
// Account Data Update - sent when account balance or position changes
type AccountUpdateEvent struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// User ID
	UserId float64 `json:"uid,omitempty"`
	// Account Balance Array
	AccountBalanceArray []AccountUpdateEventAccountBalanceArrayItem `json:"B,omitempty"`
	// Positions Array
	PositionsArray []AccountUpdateEventPositionsArrayItem `json:"P,omitempty"`
	// Greeks Array
	GreeksArray []AccountUpdateEventGreeksArrayItem `json:"G,omitempty"`
}

// AccountUpdateEventAccountBalanceArrayItem represents the accountbalancearray item details
type AccountUpdateEventAccountBalanceArrayItem struct {
	// Margin asset
	MarginAsset string `json:"a,omitempty"`
	// Account balance
	AccountBalance string `json:"b,omitempty"`
	// Position value
	PositionValue string `json:"m,omitempty"`
	// Unrealized profit/loss
	UnrealizedProfitloss string `json:"u,omitempty"`
	// Initial margin
	InitialMargin string `json:"i,omitempty"`
	// Maintenance margin
	MaintenanceMargin string `json:"M,omitempty"`
	// Positive unrealized profit for long position
	PositiveUnrealizedProfitForLongPosition float64 `json:"U,omitempty"`
}

// AccountUpdateEventPositionsArrayItem represents the positionsarray item details
type AccountUpdateEventPositionsArrayItem struct {
	// Contract symbol
	Symbol string `json:"s,omitempty"`
	// Current positions
	CurrentPositions string `json:"c,omitempty"`
	// Positions that can be reduced
	PositionsThatCanBeReduced string `json:"r,omitempty"`
	// Position value
	PositionValue string `json:"p,omitempty"`
	// Average entry price
	AverageEntryPrice string `json:"a,omitempty"`
}

// AccountUpdateEventGreeksArrayItem represents the greeksarray item details
type AccountUpdateEventGreeksArrayItem struct {
	// Underlying
	Underlying string `json:"ui,omitempty"`
	// Delta
	Delta float64 `json:"d,omitempty"`
	// Gamma
	Gamma float64 `json:"g,omitempty"`
	// Theta
	Theta float64 `json:"t,omitempty"`
	// Vega
	Vega float64 `json:"v,omitempty"`
}

// String returns string representation of AccountUpdateEvent
func (s AccountUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


