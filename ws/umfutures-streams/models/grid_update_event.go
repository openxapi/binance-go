package models

// GridUpdateEvent represents global message '#/components/messages/gridUpdateEvent'
type GridUpdateEvent struct {
	Event struct {
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
			UnmatchedQty string `json:"uq,omitempty"` // Unmatched Qty
			UnmatchedFee string `json:"uf,omitempty"` // Unmatched Fee
			AdditionalProperties string `json:"mp,omitempty"` // Additional properties
			UpdateTime int64 `json:"ut,omitempty"` // Update Time (milliseconds)
		} `json:"gu,omitempty"` // Grid Update
	} `json:"event,omitempty"`
}


