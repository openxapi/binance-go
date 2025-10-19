package models

// AccountConfigUpdateEvent represents global message '#/components/messages/accountConfigUpdateEvent'
type AccountConfigUpdateEvent struct {
	Event struct {
		EventType string `json:"e,omitempty"` // Event Type
		TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
		EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
		AccountConfigurationUpdate struct {
			Symbol string `json:"s,omitempty"` // Symbol
			Leverage int `json:"l,omitempty"` // Leverage
		} `json:"ac,omitempty"` // Account Configuration Update
		AccountConfigurationUpdate2 struct {
			MultiAssetsMode bool `json:"j,omitempty"` // Multi-Assets Mode
		} `json:"ai,omitempty"` // Account Configuration Update
	} `json:"event,omitempty"`
}


