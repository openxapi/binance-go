package models

import (
	"encoding/json"
)

// ConditionalOrderTriggerRejectEvent represents a nested object structure
type ConditionalOrderTriggerRejectEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Message Send Time (milliseconds)
	MessageSendTime int64 `json:"T,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Order Reject
	OrderReject *ConditionalOrderTriggerRejectEventOrderReject `json:"or,omitempty"`
}

// ConditionalOrderTriggerRejectEventOrderReject represents a nested object structure
type ConditionalOrderTriggerRejectEventOrderReject struct {
	// Order ID
	OrderID int64 `json:"i,omitempty"`
	// Reject Reason
	RejectReason string `json:"r,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
}

// ConditionalOrderTriggerReject - Sent when conditional order trigger is rejected
// Message name: Conditional Order Trigger Reject Event
type ConditionalOrderTriggerReject struct {
	Event *ConditionalOrderTriggerRejectEvent `json:"event,omitempty"`
}

// String returns string representation of ConditionalOrderTriggerReject
func (s ConditionalOrderTriggerReject) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ConditionalOrderTriggerReject
func (s ConditionalOrderTriggerReject) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "conditionalordertriggerreject"
}

// GetEventTime returns the event timestamp for ConditionalOrderTriggerReject
func (s ConditionalOrderTriggerReject) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


