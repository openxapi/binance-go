package models

// AccountUpdateEvent represents global message '#/components/messages/accountUpdateEvent'
type AccountUpdateEvent struct {
	Event struct {
		BalancesArray []struct {
			Asset string `json:"a,omitempty"` // Asset
			Free string `json:"f,omitempty"` // Free
			Locked string `json:"l,omitempty"` // Locked
		} `json:"B"` // Balances Array
		EventTime int64 `json:"E"` // Event Time (milliseconds)
		EventType string `json:"e"` // Event type
		TimeOfLastAccountUpdate int64 `json:"u"` // Time of last account update (milliseconds)
	} `json:"event"`
	SubscriptionId int64 `json:"subscriptionId"` // subscriptionId
}


