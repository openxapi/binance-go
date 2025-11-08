package models

// MarginBalanceUpdateEvent represents global message '#/components/messages/marginBalanceUpdateEvent'
type MarginBalanceUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	Asset string `json:"a,omitempty"` // Asset
	BalanceDelta string `json:"d,omitempty"` // Balance Delta
	EventUpdateID int `json:"U,omitempty"` // Event Update ID
	ClearTime int64 `json:"T,omitempty"` // Clear Time (milliseconds)
}


