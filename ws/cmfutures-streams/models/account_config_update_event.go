package models

// AccountConfigUpdateEvent represents global message '#/components/messages/accountConfigUpdateEvent'
type AccountConfigUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	AccountConfigurationForTradePair struct {
		Symbol string `json:"s,omitempty"` // Symbol
		Leverage int `json:"l,omitempty"` // Leverage
	} `json:"ac,omitempty"` // Account Configuration for Trade Pair
}


