package models

// ConditionalOrderTriggerRejectEvent represents global message '#/components/messages/conditionalOrderTriggerRejectEvent'
type ConditionalOrderTriggerRejectEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	MessageSendTime int64 `json:"T,omitempty"` // Message Send Time (milliseconds)
	OrderReject struct {
		Symbol string `json:"s,omitempty"` // Symbol
		OrderID int64 `json:"i,omitempty"` // Order ID
		RejectReason string `json:"r,omitempty"` // Reject Reason
	} `json:"or,omitempty"` // Order Reject
}


