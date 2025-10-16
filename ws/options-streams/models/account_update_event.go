package models

// AccountUpdateEvent represents global message '#/components/messages/accountUpdateEvent'
type AccountUpdateEvent struct {
	EventType string `json:"e"` // Event Type
	EventTime int64 `json:"E"` // Event Time (milliseconds)
	UserID float64 `json:"uid"` // User ID
	AccountBalanceArray []struct {
		MarginAsset string `json:"a,omitempty"` // Margin asset
		AccountBalance string `json:"b,omitempty"` // Account balance
		PositionValue string `json:"m,omitempty"` // Position value
		UnrealizedProfitLoss string `json:"u,omitempty"` // Unrealized profit/loss
		InitialMargin string `json:"i,omitempty"` // Initial margin
		MaintenanceMargin string `json:"M,omitempty"` // Maintenance margin
		PositiveUnrealizedProfitForLongPosition float64 `json:"U,omitempty"` // Positive unrealized profit for long position
	} `json:"B,omitempty"` // Account Balance Array
	PositionsArray []struct {
		ContractSymbol string `json:"s,omitempty"` // Contract symbol
		CurrentPositions string `json:"c,omitempty"` // Current positions
		PositionsThatCanBeReduced string `json:"r,omitempty"` // Positions that can be reduced
		PositionValue string `json:"p,omitempty"` // Position value
		AverageEntryPrice string `json:"a,omitempty"` // Average entry price
	} `json:"P,omitempty"` // Positions Array
	GreeksArray []struct {
		Underlying string `json:"ui,omitempty"` // Underlying
		Delta float64 `json:"d,omitempty"` // Delta
		Gamma float64 `json:"g,omitempty"` // Gamma
		Theta float64 `json:"t,omitempty"` // Theta
		Vega float64 `json:"v,omitempty"` // Vega
	} `json:"G,omitempty"` // Greeks Array
}


