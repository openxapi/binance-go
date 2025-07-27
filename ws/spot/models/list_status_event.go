package models

import (
	"encoding/json"
)

// ListStatusEventEvent represents a nested object structure
type ListStatusEventEvent struct {
	// List Client Order ID
	ListClientOrderID string `json:"C,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// List Order Status
	ListOrderStatus string `json:"L,omitempty"`
	// Array of orders in the list
	ArrayOfOrdersInTheList []ListStatusEventEventOItem `json:"O,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Contingency Type
	ContingencyType string `json:"c,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// OrderListId
	OrderListId int64 `json:"g,omitempty"`
	// List Status Type
	ListStatusType string `json:"l,omitempty"`
	// List Reject Reason
	ListRejectReason string `json:"r,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
}

// ListStatusEventEventOItem represents a nested object structure
type ListStatusEventEventOItem struct {
	// ClientOrderId (example: "AJYsMjErWJesZvqlJCTUgL")
	ClientOrderId string `json:"c,omitempty"`
	// OrderId (example: 17)
	OrderId int64 `json:"i,omitempty"`
	// Symbol (example: "ETHBTC")
	Symbol string `json:"s,omitempty"`
}

// ListStatusEvent - Sent with executionReport for order lists (OCO orders)
// Message name: List Status Event
type ListStatusEvent struct {
	Event *ListStatusEventEvent `json:"event,omitempty"`
}

// String returns string representation of ListStatusEvent
func (s ListStatusEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ListStatusEvent
func (s ListStatusEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "liststatusevent"
}

// GetEventTime returns the event timestamp for ListStatusEvent
func (s ListStatusEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


