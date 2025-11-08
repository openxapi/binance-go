package models

// FuturesBalancePositionUpdateEvent represents global message '#/components/messages/futuresBalancePositionUpdateEvent'
type FuturesBalancePositionUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventBusinessUnit string `json:"fs,omitempty"` // Event business unit
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	AccountAlias string `json:"i,omitempty"` // Account Alias
	AccountUpdateDetails struct {
		EventReasonType string `json:"m,omitempty"` // Event Reason Type
		Balances []struct {
			Asset string `json:"a,omitempty"` // Asset
			WalletBalance string `json:"wb,omitempty"` // Wallet Balance
			CrossWalletBalance string `json:"cw,omitempty"` // Cross Wallet Balance
			BalanceChange string `json:"bc,omitempty"` // Balance Change
		} `json:"B,omitempty"` // Balances
		Positions []struct {
			Symbol string `json:"s,omitempty"` // Symbol
			PositionAmount string `json:"pa,omitempty"` // Position Amount
			EntryPrice string `json:"ep,omitempty"` // Entry Price
			AccumulatedRealized string `json:"cr,omitempty"` // Accumulated Realized
			UnrealizedPnL string `json:"up,omitempty"` // Unrealized PnL
			PositionSide string `json:"ps,omitempty"` // Position Side
			BreakevenPrice string `json:"bep,omitempty"` // Breakeven Price
		} `json:"P,omitempty"` // Positions
	} `json:"a,omitempty"` // Account update details
}


