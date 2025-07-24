package models

import (
	"encoding/json"
)

// OrderTradeUpdateEvent represents a nested object structure
type OrderTradeUpdateEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Order details array
	OrderDetailsArray []OrderTradeUpdateEventOItem `json:"o,omitempty"`
}

// OrderTradeUpdateEventOItem represents a nested object structure
type OrderTradeUpdateEventOItem struct {
	// Order Status (example: "PARTIALLY_FILLED")
	OrderStatus string `json:"S,omitempty"`
	// Order Create Time (milliseconds) (example: 1657613342918)
	OrderCreateTime int64 `json:"T,omitempty"`
	// Client Order ID (example: "")
	ClientOrderID string `json:"c,omitempty"`
	// Completed Trade Volume (example: "0.1")
	CompletedTradeVolume string `json:"e,omitempty"`
	// Completed Trade Amount (example: "199.3")
	CompletedTradeAmount string `json:"ec,omitempty"`
	// Fee (example: "2")
	Fee string `json:"f,omitempty"`
	// Fills Array
	FillsArray []OrderTradeUpdateEventOItemFiItem `json:"fi,omitempty"`
	// Order ID (example: "4611869636869226548")
	OrderID string `json:"oid,omitempty"`
	// Order Type (example: "LIMIT")
	OrderType string `json:"oty,omitempty"`
	// Order Price (example: "1993")
	OrderPrice string `json:"p,omitempty"`
	// Post Only flag (example: true)
	PostOnlyFlag bool `json:"po,omitempty"`
	// Order Quantity (+ve for BUY, -ve for SELL) (example: "1")
	OrderQuantity string `json:"q,omitempty"`
	// Reduce Only flag (example: false)
	ReduceOnlyFlag bool `json:"r,omitempty"`
	// Symbol (example: "BTC-220930-18000-C")
	Symbol string `json:"s,omitempty"`
	// Not used currently (example: 0)
	NotUsedCurrently int64 `json:"stp,omitempty"`
	// Order Update Time (milliseconds) (example: 1657613342918)
	OrderUpdateTime int64 `json:"t,omitempty"`
	// Time in Force (example: "GTC")
	TimeInForce string `json:"tif,omitempty"`
}

// OrderTradeUpdateEventOItemFiItem represents a nested object structure
type OrderTradeUpdateEventOItemFiItem struct {
	// Trade Time (milliseconds) (example: 1657613774336)
	TradeTime int64 `json:"T,omitempty"`
	// Commission/Rebate (example: "0.0002")
	CommissionRebate string `json:"f,omitempty"`
	// Taker/Maker (example: "TAKER")
	TakerMaker string `json:"m,omitempty"`
	// Trade Price (example: "1993")
	TradePrice string `json:"p,omitempty"`
	// Trade Quantity (example: "0.1")
	TradeQuantity string `json:"q,omitempty"`
	// Trade ID (example: "20")
	TradeID string `json:"t,omitempty"`
}

// OrderTradeUpdate - Orders are updated with orderTradeUpdate event
// Message name: Order Trade Update Event
type OrderTradeUpdate struct {
	Event *OrderTradeUpdateEvent `json:"event,omitempty"`
}

// String returns string representation of OrderTradeUpdate
func (s OrderTradeUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for OrderTradeUpdate
func (s OrderTradeUpdate) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "ordertradeupdate"
}

// GetEventTime returns the event timestamp for OrderTradeUpdate
func (s OrderTradeUpdate) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


