package models

import (
	"encoding/json"
)

// OrderTradeUpdateEvent represents OrderTradeUpdateEvent
// Order Update - orders are updated with orderTradeUpdate event
type OrderTradeUpdateEvent struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Order details array
	OrderDetailsArray []OrderTradeUpdateEventOrderDetailsArrayItem `json:"o,omitempty"`
}

// OrderTradeUpdateEventOrderDetailsArrayItem represents the orderdetailsarray item details
type OrderTradeUpdateEventOrderDetailsArrayItem struct {
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Order ID
	OrderId string `json:"oid,omitempty"`
	// Client Order ID
	ClientOrderId string `json:"c,omitempty"`
	// Order Price
	Price string `json:"p,omitempty"`
	// Order Quantity (+ve for BUY, -ve for SELL)
	Quantity string `json:"q,omitempty"`
	// Completed Trade Volume
	EventType string `json:"e,omitempty"`
	// Completed Trade Amount
	CompletedTradeAmount string `json:"ec,omitempty"`
	// Fee
	Fee string `json:"f,omitempty"`
	// Order Create Time (milliseconds)
	TradeTime int64 `json:"T,omitempty"`
	// Order Update Time (milliseconds)
	TradeId int64 `json:"t,omitempty"`
	// Order Type
	OrderType string `json:"oty,omitempty"`
	// Order Status
	OrderStatus string `json:"S,omitempty"`
	// Time in Force
	TimeInForce string `json:"tif,omitempty"`
	// Post Only flag
	PostOnlyFlag bool `json:"po,omitempty"`
	// Reduce Only flag
	ReduceOnlyFlag bool `json:"r,omitempty"`
	// Not used currently
	NotUsedCurrently int `json:"stp,omitempty"`
	// Fills Array
	FillsArray []interface{} `json:"fi,omitempty"`
}

// String returns string representation of OrderTradeUpdateEvent
func (s OrderTradeUpdateEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


