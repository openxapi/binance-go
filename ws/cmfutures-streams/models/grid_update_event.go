package models

// GridUpdateEvent represents global message '#/components/messages/gridUpdateEvent'
type GridUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	GridUpdate struct {
		StrategyID int64 `json:"si,omitempty"` // Strategy ID
		StrategyType string `json:"st,omitempty"` // Strategy Type
		StrategyStatus string `json:"ss,omitempty"` // Strategy Status
		Symbol string `json:"s,omitempty"` // Symbol
		RealizedPNL string `json:"r,omitempty"` // Realized PNL
		UnmatchedAveragePrice string `json:"up,omitempty"` // Unmatched Average Price
		UnmatchedQuantity string `json:"uq,omitempty"` // Unmatched Quantity
		UnmatchedFee string `json:"uf,omitempty"` // Unmatched Fee
		MatchedPNL string `json:"mp,omitempty"` // Matched PNL
		UpdateTime int64 `json:"ut,omitempty"` // Update Time (milliseconds)
	} `json:"gu,omitempty"` // Grid Update
}


