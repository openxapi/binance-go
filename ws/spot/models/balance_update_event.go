package models

// BalanceUpdateEvent represents global message '#/components/messages/balanceUpdateEvent'
type BalanceUpdateEvent struct {
	Event struct {
		EventTime int64 `json:"E"` // Event Time (milliseconds)
		ClearTime int64 `json:"T"` // Clear Time (milliseconds)
		Asset string `json:"a"` // Asset
		BalanceDelta string `json:"d"` // Balance Delta
		EventType string `json:"e"` // Event Type
	} `json:"event"`
	SubscriptionId int64 `json:"subscriptionId"` // subscriptionId
}


