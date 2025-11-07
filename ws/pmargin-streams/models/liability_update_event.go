package models

// LiabilityUpdateEvent represents global message '#/components/messages/liabilityUpdateEvent'
type LiabilityUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	Asset string `json:"a,omitempty"` // Asset
	Type string `json:"t,omitempty"` // Type
	TransactionID int64 `json:"T,omitempty"` // Transaction ID
	Principal string `json:"p,omitempty"` // Principal
	Interest string `json:"i,omitempty"` // Interest
	TotalLiability string `json:"l,omitempty"` // Total Liability
}


