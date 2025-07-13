package models

import (
	"encoding/json"
)

// AccountUpdateEvent represents a nested object structure
type AccountUpdateEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Account information
	AccountInformation *AccountUpdateEventAccountInformation `json:"a,omitempty"`
	// Event type
	EventType string `json:"e,omitempty"`
}

// AccountUpdateEventAccountInformation represents a nested object structure
type AccountUpdateEventAccountInformation struct {
	// Balance updates
	BalanceUpdates []AccountUpdateEventAccountInformationBItem `json:"B,omitempty"`
	// Position updates
	PositionUpdates []AccountUpdateEventAccountInformationPItem `json:"P,omitempty"`
	// Event reason type
	EventReasonType string `json:"m,omitempty"`
}

// AccountUpdateEventAccountInformationBItem represents a nested object structure
type AccountUpdateEventAccountInformationBItem struct {
	// Asset name (example: "USDT")
	AssetName string `json:"a,omitempty"`
	// Balance Change except PnL and Commission (example: "-49.12345678")
	BalanceChangeExceptPnLAndCommission string `json:"bc,omitempty"`
	// Cross Wallet Balance (example: "100.12345678")
	CrossWalletBalance string `json:"cw,omitempty"`
	// Wallet Balance (example: "122624.12345678")
	WalletBalance string `json:"wb,omitempty"`
}

// AccountUpdateEventAccountInformationPItem represents a nested object structure
type AccountUpdateEventAccountInformationPItem struct {
	// Breakeven Price (example: "0.0")
	BreakevenPrice string `json:"bep,omitempty"`
	// Accumulated Realized (example: "200")
	AccumulatedRealized string `json:"cr,omitempty"`
	// Entry Price (example: "0.00000")
	EntryPrice string `json:"ep,omitempty"`
	// Isolated Wallet (example: "0.00000000")
	IsolatedWallet string `json:"iw,omitempty"`
	// Margin Type (example: "isolated")
	MarginType string `json:"mt,omitempty"`
	// Position Amount (example: "0")
	PositionAmount string `json:"pa,omitempty"`
	// Position Side (example: "BOTH")
	PositionSide string `json:"ps,omitempty"`
	// Symbol (example: "BTCUSDT")
	Symbol string `json:"s,omitempty"`
	// Unrealized PnL (example: "0.00000000")
	UnrealizedPnL string `json:"up,omitempty"`
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


