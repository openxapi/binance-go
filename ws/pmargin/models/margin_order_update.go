package models

import (
	"encoding/json"
)

// MarginOrderUpdate - Margin order execution report event
// Message name: Margin Order Update Event
type MarginOrderUpdate struct {
	// Event Type
	EventType string `json:"e,omitempty"`
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Client Order ID
	ClientOrderID string `json:"c,omitempty"`
	// Side
	Side string `json:"S,omitempty"`
	// Order Type
	OrderType string `json:"o,omitempty"`
	// Time in Force
	TimeInForce string `json:"f,omitempty"`
	// Order Quantity
	OrderQuantity string `json:"q,omitempty"`
	// Order Price
	OrderPrice string `json:"p,omitempty"`
	// Stop Price
	StopPrice string `json:"P,omitempty"`
	// Trailing Delta
	TrailingDelta int64 `json:"d,omitempty"`
	// Iceberg Quantity
	IcebergQuantity string `json:"F,omitempty"`
	// OrderListId
	OrderListId int64 `json:"g,omitempty"`
	// Original Client Order ID
	OriginalClientOrderID string `json:"C,omitempty"`
	// Current Execution Type
	CurrentExecutionType string `json:"x,omitempty"`
	// Current Order Status
	CurrentOrderStatus string `json:"X,omitempty"`
	// Order Reject Reason
	OrderRejectReason string `json:"r,omitempty"`
	// Order ID
	OrderID int64 `json:"i,omitempty"`
	// Last Executed Quantity
	LastExecutedQuantity string `json:"l,omitempty"`
	// Cumulative Filled Quantity
	CumulativeFilledQuantity string `json:"z,omitempty"`
	// Last Executed Price
	LastExecutedPrice string `json:"L,omitempty"`
	// Commission Amount
	CommissionAmount string `json:"n,omitempty"`
	// Commission Asset
	CommissionAsset string `json:"N,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Trade ID
	TradeID int64 `json:"t,omitempty"`
	// Prevented Match Id
	PreventedMatchId int64 `json:"v,omitempty"`
	// UpdateId
	UpdateId int64 `json:"I,omitempty"`
	// Is Working
	IsWorking bool `json:"w,omitempty"`
	// Is Maker
	IsMaker bool `json:"m,omitempty"`
	// Ignore
	Ignore bool `json:"M,omitempty"`
	// Order Creation Time (milliseconds)
	OrderCreationTime int64 `json:"O,omitempty"`
	// Cumulative Quote Asset Transacted Quantity
	CumulativeQuoteAssetTransactedQuantity string `json:"Z,omitempty"`
	// Last Quote Asset Transacted Quantity
	LastQuoteAssetTransactedQuantity string `json:"Y,omitempty"`
	// Quote Order Quantity
	QuoteOrderQuantity string `json:"Q,omitempty"`
	// Trailing Time (milliseconds)
	TrailingTime int64 `json:"D,omitempty"`
	// Strategy ID
	StrategyID int64 `json:"j,omitempty"`
	// Strategy Type
	StrategyType int64 `json:"J,omitempty"`
	// Working Time (milliseconds)
	WorkingTime int64 `json:"W,omitempty"`
	// Self Trade Prevention Mode
	SelfTradePreventionMode string `json:"V,omitempty"`
	// Count of Order Update
	CountOfOrderUpdate int64 `json:"u,omitempty"`
	// Order Update ID
	OrderUpdateID int64 `json:"U,omitempty"`
	// Allowed Self Trade Prevention Mode
	AllowedSelfTradePreventionMode string `json:"A,omitempty"`
	// Prevented Quantity
	PreventedQuantity string `json:"B,omitempty"`
}

// String returns string representation of MarginOrderUpdate
func (s MarginOrderUpdate) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for MarginOrderUpdate
func (s MarginOrderUpdate) GetEventType() string {
	if s.EventType != "" {
		return s.EventType
	}
	return "marginorderupdate"
}

// GetEventTime returns the event timestamp for MarginOrderUpdate
func (s MarginOrderUpdate) GetEventTime() int64 {
	if s.EventTime != 0 {
		return s.EventTime
	}
	return 0
}


