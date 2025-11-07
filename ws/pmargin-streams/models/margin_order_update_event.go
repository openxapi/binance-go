package models

// MarginOrderUpdateEvent represents global message '#/components/messages/marginOrderUpdateEvent'
type MarginOrderUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	Symbol string `json:"s,omitempty"` // Symbol
	ClientOrderID string `json:"c,omitempty"` // Client Order ID
	Side string `json:"S,omitempty"` // Side
	OrderType string `json:"o,omitempty"` // Order Type
	TimeInForce string `json:"f,omitempty"` // Time in Force
	OrderQuantity string `json:"q,omitempty"` // Order Quantity
	OrderPrice string `json:"p,omitempty"` // Order Price
	StopPrice string `json:"P,omitempty"` // Stop Price
	TrailingDelta int `json:"d,omitempty"` // Trailing Delta
	IcebergQuantity string `json:"F,omitempty"` // Iceberg Quantity
	OrderListId int `json:"g,omitempty"` // OrderListId
	OriginalClientOrderID string `json:"C,omitempty"` // Original Client Order ID
	CurrentExecutionType string `json:"x,omitempty"` // Current Execution Type
	CurrentOrderStatus string `json:"X,omitempty"` // Current Order Status
	OrderRejectReason string `json:"r,omitempty"` // Order Reject Reason
	OrderID int `json:"i,omitempty"` // Order ID
	LastExecutedQuantity string `json:"l,omitempty"` // Last Executed Quantity
	CumulativeFilledQuantity string `json:"z,omitempty"` // Cumulative Filled Quantity
	LastExecutedPrice string `json:"L,omitempty"` // Last Executed Price
	CommissionAmount string `json:"n,omitempty"` // Commission Amount
	CommissionAsset string `json:"N,omitempty"` // Commission Asset
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	TradeID int `json:"t,omitempty"` // Trade ID
	PreventedMatchId int `json:"v,omitempty"` // Prevented Match Id
	UpdateId int `json:"I,omitempty"` // UpdateId
	IsWorking bool `json:"w,omitempty"` // Is Working
	IsMaker bool `json:"m,omitempty"` // Is Maker
	Ignore bool `json:"M,omitempty"` // Ignore
	OrderCreationTime int64 `json:"O,omitempty"` // Order Creation Time (milliseconds)
	CumulativeQuoteAssetTransactedQuantity string `json:"Z,omitempty"` // Cumulative Quote Asset Transacted Quantity
	LastQuoteAssetTransactedQuantity string `json:"Y,omitempty"` // Last Quote Asset Transacted Quantity
	QuoteOrderQuantity string `json:"Q,omitempty"` // Quote Order Quantity
	TrailingTime int64 `json:"D,omitempty"` // Trailing Time (milliseconds)
	StrategyID int `json:"j,omitempty"` // Strategy ID
	StrategyType int `json:"J,omitempty"` // Strategy Type
	WorkingTime int64 `json:"W,omitempty"` // Working Time (milliseconds)
	SelfTradePreventionMode string `json:"V,omitempty"` // Self Trade Prevention Mode
	CountOfOrderUpdate int `json:"u,omitempty"` // Count of Order Update
	OrderUpdateID int `json:"U,omitempty"` // Order Update ID
	AllowedSelfTradePreventionMode string `json:"A,omitempty"` // Allowed Self Trade Prevention Mode
	PreventedQuantity string `json:"B,omitempty"` // Prevented Quantity
}


