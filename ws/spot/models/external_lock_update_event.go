package models

// ExternalLockUpdateEvent represents global message '#/components/messages/externalLockUpdateEvent'
type ExternalLockUpdateEvent struct {
	Event struct {
		EventTime int64 `json:"E"` // Event Time (milliseconds)
		TransactionTime int64 `json:"T"` // Transaction Time (milliseconds)
		Asset string `json:"a"` // Asset
		Delta string `json:"d"` // Delta
		EventType string `json:"e"` // Event Type
	} `json:"event,omitempty"`
	SubscriptionId int64 `json:"subscriptionId,omitempty"` // subscriptionId
}


