package models

// FuturesAccountConfigUpdateEvent represents global message '#/components/messages/futuresAccountConfigUpdateEvent'
type FuturesAccountConfigUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventBusinessUnit string `json:"fs,omitempty"` // Event business unit
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	AccountConfigurationDetails struct {
		Symbol string `json:"s,omitempty"` // Symbol
		Leverage int `json:"l,omitempty"` // Leverage
	} `json:"ac,omitempty"` // Account configuration details
}


