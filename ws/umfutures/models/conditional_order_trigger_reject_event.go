package models

import (
	"encoding/json"
)

// ConditionalOrderTriggerRejectEventEvent represents a nested object structure
type ConditionalOrderTriggerRejectEventEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Message Send Time (milliseconds)
	MessageSendTime int64 `json:"T,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Order Reject
	OrderReject *ConditionalOrderTriggerRejectEventEventOrderReject `json:"or,omitempty"`
}

// ConditionalOrderTriggerRejectEventEventOrderReject represents a nested object structure
type ConditionalOrderTriggerRejectEventEventOrderReject struct {
	// Order ID
	OrderID int64 `json:"i,omitempty"`
	// Reject Reason
	RejectReason string `json:"r,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
}

// ConditionalOrderTriggerRejectEvent - Sent when conditional order trigger is rejected
// Message name: Conditional Order Trigger Reject Event
type ConditionalOrderTriggerRejectEvent struct {
	Event *ConditionalOrderTriggerRejectEventEvent `json:"event,omitempty"`
}

// String returns string representation of ConditionalOrderTriggerRejectEvent
func (s ConditionalOrderTriggerRejectEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ConditionalOrderTriggerRejectEvent
func (s ConditionalOrderTriggerRejectEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "conditionalordertriggerrejectevent"
}

// GetEventTime returns the event timestamp for ConditionalOrderTriggerRejectEvent
func (s ConditionalOrderTriggerRejectEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


