package models

import (
	"encoding/json"
)

// ConditionalOrderTradeUpdateStrategyOrderDetails represents a nested object structure
type ConditionalOrderTradeUpdateStrategyOrderDetails struct {
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

// ConditionalOrderTradeUpdate - Sent when new order created or order status changed
// Message name: Conditional Order Trade Update Event
type ConditionalOrderTradeUpdate struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Event business unit
	EventBusinessUnit string `json:"fs,omitempty"`
	// Strategy order details
	StrategyOrderDetails *ConditionalOrderTradeUpdateStrategyOrderDetails `json:"so,omitempty"`
}

// String returns string representation of ConditionalOrderTradeUpdate
func (s ConditionalOrderTradeUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ConditionalOrderTradeUpdate
func (s ConditionalOrderTradeUpdate) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "conditionalordertradeupdate"
}

// GetEventTime returns the event timestamp for ConditionalOrderTradeUpdate
func (s ConditionalOrderTradeUpdate) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


