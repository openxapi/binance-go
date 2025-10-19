package models

// StrategyUpdateEvent represents global message '#/components/messages/strategyUpdateEvent'
type StrategyUpdateEvent struct {
	Event struct {
		EventType string `json:"e,omitempty"` // Event Type
		TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
		EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
		StrategyUpdate struct {
			StrategyID int64 `json:"si,omitempty"` // Strategy ID
			StrategyType string `json:"st,omitempty"` // Strategy Type
			StrategyStatus string `json:"ss,omitempty"` // Strategy Status
			Symbol string `json:"s,omitempty"` // Symbol
			OperationCode int `json:"c,omitempty"` // Operation Code
			UpdateTime int64 `json:"ut,omitempty"` // Update Time (milliseconds)
		} `json:"su,omitempty"` // Strategy Update
	} `json:"event,omitempty"`
}


