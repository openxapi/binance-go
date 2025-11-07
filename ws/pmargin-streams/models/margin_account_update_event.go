package models

// MarginAccountUpdateEvent represents global message '#/components/messages/marginAccountUpdateEvent'
type MarginAccountUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	TimeOfLastAccountUpdate int64 `json:"u,omitempty"` // Time of last account update (milliseconds)
	UpdateID int64 `json:"U,omitempty"` // Update ID
	BalancesArray []struct {
		Asset string `json:"a,omitempty"` // Asset
		FreeBalance string `json:"f,omitempty"` // Free balance
		LockedBalance string `json:"l,omitempty"` // Locked balance
	} `json:"B,omitempty"` // Balances Array
}


