package models

import (
	"encoding/json"
)

// ExecutionReportEventEvent represents a nested object structure
type ExecutionReportEventEvent struct {
	// Prevented Quantity
	PreventedQuantity string `json:"A,omitempty"`
	// Last Prevented Quantity
	LastPreventedQuantity string `json:"B,omitempty"`
	// Original client order ID (for canceled orders)
	OriginalClientOrderID string `json:"C,omitempty"`
	// Counter Symbol
	CounterSymbol string `json:"Cs,omitempty"`
	// Trailing Time (milliseconds)
	TrailingTime int64 `json:"D,omitempty"`
	// Event time (milliseconds)
	EventTime int64 `json:"E,omitempty"`
	// Iceberg quantity
	IcebergQuantity string `json:"F,omitempty"`
	// Execution Id
	ExecutionId int64 `json:"I,omitempty"`
	// Strategy Type (if provided upon order placement)
	StrategyType int64 `json:"J,omitempty"`
	// Last executed price
	LastExecutedPrice string `json:"L,omitempty"`
	// Ignore
	Ignore bool `json:"M,omitempty"`
	// Commission asset
	CommissionAsset string `json:"N,omitempty"`
	// Order creation time (milliseconds)
	OrderCreationTime int64 `json:"O,omitempty"`
	// Stop price
	StopPrice string `json:"P,omitempty"`
	// Quote Order Quantity
	QuoteOrderQuantity string `json:"Q,omitempty"`
	// Side
	Side string `json:"S,omitempty"`
	// Transaction time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"`
	// Counter Order Id
	CounterOrderId int64 `json:"U,omitempty"`
	// SelfTradePreventionMode
	SelfTradePreventionMode string `json:"V,omitempty"`
	// Working Time (visible if order is on the book)
	WorkingTime int64 `json:"W,omitempty"`
	// Current order status
	CurrentOrderStatus string `json:"X,omitempty"`
	// Last quote asset transacted quantity
	LastQuoteAssetTransactedQuantity string `json:"Y,omitempty"`
	// Cumulative quote asset transacted quantity
	CumulativeQuoteAssetTransactedQuantity string `json:"Z,omitempty"`
	// Allocation ID
	AllocationID int64 `json:"a,omitempty"`
	// Match Type (for orders with allocations)
	MatchType string `json:"b,omitempty"`
	// Client order ID
	ClientOrderID string `json:"c,omitempty"`
	// Trailing Delta (for trailing stop orders)
	TrailingDelta int64 `json:"d,omitempty"`
	// Event type
	EventType string `json:"e,omitempty"`
	// Time in force
	TimeInForce string `json:"f,omitempty"`
	// OrderListId
	OrderListId int64 `json:"g,omitempty"`
	// Order ID
	OrderID int64 `json:"i,omitempty"`
	// Strategy Id (if provided upon order placement)
	StrategyId int64 `json:"j,omitempty"`
	// Working Floor (for orders with potential allocations)
	WorkingFloor string `json:"k,omitempty"`
	// Last executed quantity
	LastExecutedQuantity string `json:"l,omitempty"`
	// Maker side
	MakerSide bool `json:"m,omitempty"`
	// Commission amount
	CommissionAmount string `json:"n,omitempty"`
	// Order type
	OrderType string `json:"o,omitempty"`
	// Order price
	OrderPrice string `json:"p,omitempty"`
	// Prevented Execution Price
	PreventedExecutionPrice string `json:"pL,omitempty"`
	// Prevented Execution Quote Qty
	PreventedExecutionQuoteQty string `json:"pY,omitempty"`
	// Prevented Execution Quantity
	PreventedExecutionQuantity string `json:"pl,omitempty"`
	// Order quantity
	OrderQuantity string `json:"q,omitempty"`
	// Order reject reason
	OrderRejectReason string `json:"r,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Trade ID
	TradeID int64 `json:"t,omitempty"`
	// Trade Group Id
	TradeGroupId int64 `json:"u,omitempty"`
	// UsedSor (for orders that used SOR)
	UsedSor bool `json:"uS,omitempty"`
	// Prevented Match Id (visible if order expired due to STP)
	PreventedMatchId int64 `json:"v,omitempty"`
	// Order on the book
	OrderOnTheBook bool `json:"w,omitempty"`
	// Current execution type
	CurrentExecutionType string `json:"x,omitempty"`
	// Cumulative filled quantity
	CumulativeFilledQuantity string `json:"z,omitempty"`
}

// ExecutionReportEvent - Orders are updated with executionReport event
// Message name: Execution Report Event
type ExecutionReportEvent struct {
	Event *ExecutionReportEventEvent `json:"event,omitempty"`
}

// String returns string representation of ExecutionReportEvent
func (s ExecutionReportEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// GetEventType returns the event type for ExecutionReportEvent
func (s ExecutionReportEvent) GetEventType() string {
	if s.Event.EventType != "" {
		return s.Event.EventType
	}
	return "executionreportevent"
}

// GetEventTime returns the event timestamp for ExecutionReportEvent
func (s ExecutionReportEvent) GetEventTime() int64 {
	if s.Event.EventTime != 0 {
		return s.Event.EventTime
	}
	return 0
}


