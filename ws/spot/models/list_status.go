package models

import (
	"encoding/json"
)

// ListStatusEvent represents a nested object structure
type ListStatusEvent struct {
	// List Client Order ID
	ListClientOrderID string `json:"C,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// List Order Status
	ListOrderStatus string `json:"L,omitempty"`
	// Array of orders in the list
	ArrayOfOrdersInTheList []ListStatusEventArrayItem `json:"O,omitempty"`
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

// ListStatusEventArrayItem represents a nested object structure
type ListStatusEventArrayItem struct {
	// ClientOrderId (example: "AJYsMjErWJesZvqlJCTUgL")
	ClientOrderId string `json:"c,omitempty"`
	// OrderId (example: 17)
	OrderId int64 `json:"i,omitempty"`
	// Symbol (example: "ETHBTC")
	Symbol string `json:"s,omitempty"`
}

// ListStatus - Sent with executionReport for order lists (OCO orders)
// Message name: List Status Event
type ListStatus struct {
	Event *ListStatusEvent `json:"event,omitempty"`
}

// String returns string representation of ListStatus
func (s ListStatus) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ListStatus
func (s ListStatus) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "liststatus"
}

// GetEventTime returns the event timestamp for ListStatus
func (s ListStatus) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


