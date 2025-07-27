package models

import (
	"encoding/json"
)

// AccountUpdateEventEvent represents a nested object structure
type AccountUpdateEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Account information
	AccountInformation *AccountUpdateEventEventAccountInformation `json:"a,omitempty"`
	// Event type
	EventType string `json:"e,omitempty"`
	// Account Alias
	AccountAlias string `json:"i,omitempty"`
}

// AccountUpdateEventEventAccountInformation represents a nested object structure
type AccountUpdateEventEventAccountInformation struct {
	// Balance updates
	BalanceUpdates []AccountUpdateEventEventAccountInformationBItem `json:"B,omitempty"`
	// Position updates
	PositionUpdates []AccountUpdateEventEventAccountInformationPItem `json:"P,omitempty"`
	// Event reason type
	EventReasonType string `json:"m,omitempty"`
}

// AccountUpdateEventEventAccountInformationBItem represents a nested object structure
type AccountUpdateEventEventAccountInformationBItem struct {
	// Asset name (example: "BTC")
	AssetName string `json:"a,omitempty"`
	// Balance Change except PnL and Commission (example: "0.00000000")
	BalanceChangeExceptPnLAndCommission string `json:"bc,omitempty"`
	// Cross Wallet Balance (example: "1.00000000")
	CrossWalletBalance string `json:"cw,omitempty"`
	// Wallet Balance (example: "1.00000000")
	WalletBalance string `json:"wb,omitempty"`
}

// AccountUpdateEventEventAccountInformationPItem represents a nested object structure
type AccountUpdateEventEventAccountInformationPItem struct {
	// Breakeven Price (example: "0.0")
	BreakevenPrice string `json:"bep,omitempty"`
	// Accumulated Realized (example: "0")
	AccumulatedRealized string `json:"cr,omitempty"`
	// Entry Price (example: "0.00000")
	EntryPrice string `json:"ep,omitempty"`
	// Isolated Wallet (example: "0.00000000")
	IsolatedWallet string `json:"iw,omitempty"`
	// Margin Type (example: "cross")
	MarginType string `json:"mt,omitempty"`
	// Position Amount (example: "0")
	PositionAmount string `json:"pa,omitempty"`
	// Position Side (example: "BOTH")
	PositionSide string `json:"ps,omitempty"`
	// Symbol (example: "BTCUSD_PERP")
	Symbol string `json:"s,omitempty"`
	// Unrealized PnL (example: "0.00000000")
	UnrealizedPnL string `json:"up,omitempty"`
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


