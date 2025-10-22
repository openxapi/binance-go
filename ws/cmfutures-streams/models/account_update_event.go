package models

// AccountUpdateEvent represents global message '#/components/messages/accountUpdateEvent'
type AccountUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	AccountAlias string `json:"i,omitempty"` // Account Alias
	UpdateData struct {
		EventReasonType string `json:"m,omitempty"` // Event reason type
		BalanceUpdates []struct {
			Asset string `json:"a,omitempty"` // Asset
			WalletBalance string `json:"wb,omitempty"` // Wallet Balance
			CrossWalletBalance string `json:"cw,omitempty"` // Cross Wallet Balance
			BalanceChangeExceptPnLAndCommission string `json:"bc,omitempty"` // Balance Change except PnL and Commission
		} `json:"B,omitempty"` // Balance updates
		PositionUpdates []struct {
			Symbol string `json:"s,omitempty"` // Symbol
			PositionAmount string `json:"pa,omitempty"` // Position Amount
			EntryPrice string `json:"ep,omitempty"` // Entry Price
			BreakevenPrice string `json:"bep,omitempty"` // Breakeven Price
			AccumulatedRealized string `json:"cr,omitempty"` // Accumulated Realized
			UnrealizedPnL string `json:"up,omitempty"` // Unrealized PnL
			MarginType string `json:"mt,omitempty"` // Margin Type
			IsolatedWallet string `json:"iw,omitempty"` // Isolated Wallet
			PositionSide string `json:"ps,omitempty"` // Position Side
		} `json:"P,omitempty"` // Position updates
	} `json:"a,omitempty"` // Update Data
}


