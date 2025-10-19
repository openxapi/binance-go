package models

// MarginCallEvent represents global message '#/components/messages/marginCallEvent'
type MarginCallEvent struct {
	Event struct {
		EventType string `json:"e,omitempty"` // Event Type
		EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
		CrossWalletBalance string `json:"cw,omitempty"` // Cross Wallet Balance (only for crossed position)
		PositionOfMarginCall []struct {
			Symbol string `json:"s,omitempty"` // Symbol
			PositionSide string `json:"ps,omitempty"` // Position Side
			PositionAmount string `json:"pa,omitempty"` // Position Amount
			MarginType string `json:"mt,omitempty"` // Margin Type
			IsolatedWallet string `json:"iw,omitempty"` // Isolated Wallet
			MarkPrice string `json:"mp,omitempty"` // Mark Price
			UnrealizedPnL string `json:"up,omitempty"` // Unrealized PnL
			MaintenanceMarginRequired string `json:"mm,omitempty"` // Maintenance Margin Required
		} `json:"p,omitempty"` // Position(s) of Margin Call
	} `json:"event,omitempty"`
}


