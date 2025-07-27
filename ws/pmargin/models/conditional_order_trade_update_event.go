package models

import (
	"encoding/json"
)

// ConditionalOrderTradeUpdateEventStrategyOrderDetails represents a nested object structure
type ConditionalOrderTradeUpdateEventStrategyOrderDetails struct {
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Strategy Client Order Id
	StrategyClientOrderId string `json:"c,omitempty"`
	// Strategy ID
	StrategyID int64 `json:"si,omitempty"`
	// Side
	Side string `json:"S,omitempty"`
	// Strategy Type
	StrategyType string `json:"st,omitempty"`
	// Time in Force
	TimeInForce string `json:"f,omitempty"`
	// Quantity
	Quantity string `json:"q,omitempty"`
	// Price
	Price string `json:"p,omitempty"`
	// Stop Price
	StopPrice string `json:"sp,omitempty"`
	// Strategy Order Status
	StrategyOrderStatus string `json:"os,omitempty"`
	// Order book Time (milliseconds)
	OrderBookTime int64 `json:"T,omitempty"`
	// Order update Time (milliseconds)
	OrderUpdateTime int64 `json:"ut,omitempty"`
	// Is this reduce only
	IsThisReduceOnly bool `json:"R,omitempty"`
	// Stop Price Working Type
	StopPriceWorkingType string `json:"wt,omitempty"`
	// Position Side
	PositionSide string `json:"ps,omitempty"`
	// If Close-All
	IfCloseAll bool `json:"cp,omitempty"`
	// Activation Price
	ActivationPrice string `json:"AP,omitempty"`
	// Callback Rate
	CallbackRate string `json:"cr,omitempty"`
	// Order Id
	OrderId int64 `json:"i,omitempty"`
	// STP mode
	STPMode string `json:"V,omitempty"`
	// Good Till Date
	GoodTillDate int64 `json:"gtd,omitempty"`
}

// ConditionalOrderTradeUpdateEvent - Sent when new order created or order status changed
// Message name: Conditional Order Trade Update Event
type ConditionalOrderTradeUpdateEvent struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Event business unit
	EventBusinessUnit string `json:"fs,omitempty"`
	// Strategy order details
	StrategyOrderDetails *ConditionalOrderTradeUpdateEventStrategyOrderDetails `json:"so,omitempty"`
}

// String returns string representation of ConditionalOrderTradeUpdateEvent
func (s ConditionalOrderTradeUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ConditionalOrderTradeUpdateEvent
func (s ConditionalOrderTradeUpdateEvent) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "conditionalordertradeupdateevent"
}

// GetEventTime returns the event timestamp for ConditionalOrderTradeUpdateEvent
func (s ConditionalOrderTradeUpdateEvent) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


