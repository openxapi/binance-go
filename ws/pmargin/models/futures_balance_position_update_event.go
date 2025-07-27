package models

import (
	"encoding/json"
)

// FuturesBalancePositionUpdateEventAccountUpdateDetails represents a nested object structure
type FuturesBalancePositionUpdateEventAccountUpdateDetails struct {
	// Event Reason Type
	EventReasonType string `json:"m,omitempty"`
	// Balances Array
	BalancesArray []FuturesBalancePositionUpdateEventAccountUpdateDetailsBItem `json:"B,omitempty"`
	// Positions Array
	PositionsArray []FuturesBalancePositionUpdateEventAccountUpdateDetailsPItem `json:"P,omitempty"`
}

// FuturesBalancePositionUpdateEventAccountUpdateDetailsBItem represents a nested object structure
type FuturesBalancePositionUpdateEventAccountUpdateDetailsBItem struct {
	// Asset (example: "USDT")
	Asset string `json:"a,omitempty"`
	// Wallet Balance (example: "122624.12345678")
	WalletBalance string `json:"wb,omitempty"`
	// Cross Wallet Balance (example: "100.12345678")
	CrossWalletBalance string `json:"cw,omitempty"`
	// Balance Change (example: "50.12345678")
	BalanceChange string `json:"bc,omitempty"`
}

// FuturesBalancePositionUpdateEventAccountUpdateDetailsPItem represents a nested object structure
type FuturesBalancePositionUpdateEventAccountUpdateDetailsPItem struct {
	// Symbol (example: "BTCUSDT")
	Symbol string `json:"s,omitempty"`
	// Position Amount (example: "0")
	PositionAmount string `json:"pa,omitempty"`
	// Entry Price (example: "0.00000")
	EntryPrice string `json:"ep,omitempty"`
	// Accumulated Realized (example: "200")
	AccumulatedRealized string `json:"cr,omitempty"`
	// Unrealized PnL (example: "0")
	UnrealizedPnL string `json:"up,omitempty"`
	// Position Side (example: "BOTH")
	PositionSide string `json:"ps,omitempty"`
	// Breakeven Price (example: "0.00000")
	BreakevenPrice string `json:"bep,omitempty"`
}

// FuturesBalancePositionUpdateEvent - Futures balance and position update event
// Message name: Futures Balance and Position Update Event
type FuturesBalancePositionUpdateEvent struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event business unit
	EventBusinessUnit string `json:"fs,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Account Alias
	AccountAlias string `json:"i,omitempty"`
	// Account update details
	AccountUpdateDetails *FuturesBalancePositionUpdateEventAccountUpdateDetails `json:"a,omitempty"`
}

// String returns string representation of FuturesBalancePositionUpdateEvent
func (s FuturesBalancePositionUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for FuturesBalancePositionUpdateEvent
func (s FuturesBalancePositionUpdateEvent) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "futuresbalancepositionupdateevent"
}

// GetEventTime returns the event timestamp for FuturesBalancePositionUpdateEvent
func (s FuturesBalancePositionUpdateEvent) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


