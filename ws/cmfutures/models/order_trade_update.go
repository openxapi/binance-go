package models

import (
	"encoding/json"
)

// OrderTradeUpdateEvent represents a nested object structure
type OrderTradeUpdateEvent struct {
	// Event Time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Transaction Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Event Type
	EventType string `json:"e,omitempty"`
	// Account Alias
	AccountAlias string `json:"i,omitempty"`
	// Order details
	OrderDetails *OrderTradeUpdateEventOrderDetails `json:"o,omitempty"`
}

// OrderTradeUpdateEventOrderDetails represents a nested object structure
type OrderTradeUpdateEventOrderDetails struct {
	// Last Filled Price
	LastFilledPrice string `json:"L,omitempty"`
	// Commission Asset
	CommissionAsset string `json:"N,omitempty"`
	// Is this reduce only
	IsThisReduceOnly bool `json:"R,omitempty"`
	// Side
	Side string `json:"S,omitempty"`
	// Order Trade Time (milliseconds)
	OrderTradeTime int64 `json:"T,omitempty"`
	// Order Status
	OrderStatus string `json:"X,omitempty"`
	// Ask Notional
	AskNotional string `json:"a,omitempty"`
	// Average Price
	AveragePrice string `json:"ap,omitempty"`
	// Bids Notional
	BidsNotional string `json:"b,omitempty"`
	// Client Order ID
	ClientOrderID string `json:"c,omitempty"`
	// If Close-All
	IfCloseAll bool `json:"cp,omitempty"`
	// Time in Force
	TimeInForce string `json:"f,omitempty"`
	// Order ID
	OrderID int64 `json:"i,omitempty"`
	// Order Last Filled Quantity
	OrderLastFilledQuantity string `json:"l,omitempty"`
	// Maker side
	MakerSide bool `json:"m,omitempty"`
	// Commission
	Commission string `json:"n,omitempty"`
	// Order Type
	OrderType string `json:"o,omitempty"`
	// Original Order Type
	OriginalOrderType string `json:"ot,omitempty"`
	// Original Price
	OriginalPrice string `json:"p,omitempty"`
	// Protect position
	ProtectPosition bool `json:"pP,omitempty"`
	// Position Side
	PositionSide string `json:"ps,omitempty"`
	// Original Quantity
	OriginalQuantity string `json:"q,omitempty"`
	// Realized Profit of the trade
	RealizedProfitOfTheTrade string `json:"rp,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Ignore
	Ignore int64 `json:"si,omitempty"`
	// Stop Price
	StopPrice string `json:"sp,omitempty"`
	// Ignore
	Ignore2 int64 `json:"ss,omitempty"`
	// Trade ID
	TradeID int64 `json:"t,omitempty"`
	// Stop Price Working Type
	StopPriceWorkingType string `json:"wt,omitempty"`
	// Execution Type
	ExecutionType string `json:"x,omitempty"`
	// Order Filled Accumulated Quantity
	OrderFilledAccumulatedQuantity string `json:"z,omitempty"`
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


