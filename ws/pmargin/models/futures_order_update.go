package models

import (
	"encoding/json"
)

// FuturesOrderUpdateOrderDetails represents a nested object structure
type FuturesOrderUpdateOrderDetails struct {
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Client Order Id
	ClientOrderId string `json:"c,omitempty"`
	// Side
	Side string `json:"S,omitempty"`
	// Order Type
	OrderType string `json:"o,omitempty"`
	// Time in Force
	TimeInForce string `json:"f,omitempty"`
	// Original Quantity
	OriginalQuantity string `json:"q,omitempty"`
	// Original Price
	OriginalPrice string `json:"p,omitempty"`
	// Average Price
	AveragePrice string `json:"ap,omitempty"`
	// Stop Price
	StopPrice string `json:"sp,omitempty"`
	// Execution Type
	ExecutionType string `json:"x,omitempty"`
	// Order Status
	OrderStatus string `json:"X,omitempty"`
	// Order Id
	OrderId int64 `json:"i,omitempty"`
	// Order Last Filled Quantity
	OrderLastFilledQuantity string `json:"l,omitempty"`
	// Order Filled Accumulated Quantity
	OrderFilledAccumulatedQuantity string `json:"z,omitempty"`
	// Last Filled Price
	LastFilledPrice string `json:"L,omitempty"`
	// Commission Asset
	CommissionAsset string `json:"N,omitempty"`
	// Commission
	Commission string `json:"n,omitempty"`
}

// FuturesOrderUpdate - Futures order trade update event
// Message name: Futures Order Update Event
type FuturesOrderUpdate struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event business unit
	EventBusinessUnit string `json:"fs,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Account Alias
	AccountAlias string `json:"i,omitempty"`
	// Order details
	OrderDetails *FuturesOrderUpdateOrderDetails `json:"o,omitempty"`
}

// String returns string representation of FuturesOrderUpdate
func (s FuturesOrderUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for FuturesOrderUpdate
func (s FuturesOrderUpdate) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "futuresorderupdate"
}

// GetEventTime returns the event timestamp for FuturesOrderUpdate
func (s FuturesOrderUpdate) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


