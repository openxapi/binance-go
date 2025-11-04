package models

// OrderUpdateEvent represents global message '#/components/messages/orderUpdateEvent'
type OrderUpdateEvent struct {
	Event struct {
		PreventedQuantity string `json:"A,omitempty"` // Prevented Quantity
		LastPreventedQuantity string `json:"B,omitempty"` // Last Prevented Quantity
		OriginalClientOrderID string `json:"C,omitempty"` // Original client order ID (for canceled orders)
		CounterSymbol string `json:"Cs,omitempty"` // Counter Symbol
		TrailingTime int64 `json:"D,omitempty"` // Trailing Time (milliseconds)
		EventTime int64 `json:"E"` // Event time (milliseconds)
		IcebergQuantity string `json:"F,omitempty"` // Iceberg quantity
		ExecutionId int64 `json:"I,omitempty"` // Execution Id
		StrategyType int64 `json:"J,omitempty"` // Strategy Type (if provided upon order placement)
		LastExecutedPrice string `json:"L,omitempty"` // Last executed price
		Ignore bool `json:"M,omitempty"` // Ignore
		CommissionAsset string `json:"N,omitempty"` // Commission asset
		OrderCreationTime int64 `json:"O,omitempty"` // Order creation time (milliseconds)
		StopPrice string `json:"P,omitempty"` // Stop price
		QuoteOrderQuantity string `json:"Q,omitempty"` // Quote Order Quantity
		Side string `json:"S,omitempty"` // Side
		TransactionTime int64 `json:"T"` // Transaction time (milliseconds)
		CounterOrderId int64 `json:"U,omitempty"` // Counter Order Id
		SelfTradePreventionMode string `json:"V,omitempty"` // SelfTradePreventionMode
		WorkingTime int64 `json:"W,omitempty"` // Working Time (visible if order is on the book)
		CurrentOrderStatus string `json:"X,omitempty"` // Current order status
		LastQuoteAssetTransactedQuantity string `json:"Y,omitempty"` // Last quote asset transacted quantity
		CumulativeQuoteAssetTransactedQuantity string `json:"Z,omitempty"` // Cumulative quote asset transacted quantity
		AllocationID int64 `json:"a,omitempty"` // Allocation ID
		MatchType string `json:"b,omitempty"` // Match Type (for orders with allocations)
		ClientOrderID string `json:"c,omitempty"` // Client order ID
		TrailingDelta int `json:"d,omitempty"` // Trailing Delta (for trailing stop orders)
		EventType string `json:"e"` // Event type
		TimeInForce string `json:"f,omitempty"` // Time in force
		OrderListId int64 `json:"g,omitempty"` // OrderListId
		PeggedOffsetType string `json:"gOT,omitempty"` // Pegged offset Type (for pegged orders)
		PeggedOffsetValue int64 `json:"gOV,omitempty"` // Pegged Offset Value (for pegged orders)
		PeggedPriceType string `json:"gP,omitempty"` // Pegged Price Type (for pegged orders)
		PeggedPrice string `json:"gp,omitempty"` // Pegged Price (for pegged orders)
		OrderID int64 `json:"i,omitempty"` // Order ID
		StrategyId int64 `json:"j,omitempty"` // Strategy Id (if provided upon order placement)
		WorkingFloor string `json:"k,omitempty"` // Working Floor (for orders with potential allocations)
		LastExecutedQuantity string `json:"l,omitempty"` // Last executed quantity
		MakerSide bool `json:"m,omitempty"` // Maker side
		CommissionAmount string `json:"n,omitempty"` // Commission amount
		OrderType string `json:"o,omitempty"` // Order type
		OrderPrice string `json:"p,omitempty"` // Order price
		PreventedExecutionPrice string `json:"pL,omitempty"` // Prevented Execution Price
		PreventedExecutionQuoteQty string `json:"pY,omitempty"` // Prevented Execution Quote Qty
		PreventedExecutionQuantity string `json:"pl,omitempty"` // Prevented Execution Quantity
		OrderQuantity string `json:"q,omitempty"` // Order quantity
		OrderRejectReason string `json:"r,omitempty"` // Order reject reason
		Symbol string `json:"s,omitempty"` // Symbol
		TradeID int64 `json:"t,omitempty"` // Trade ID
		TradeGroupId int64 `json:"u,omitempty"` // Trade Group Id
		UsedSor bool `json:"uS,omitempty"` // UsedSor (for orders that used SOR)
		PreventedMatchId int64 `json:"v,omitempty"` // Prevented Match Id (visible if order expired due to STP)
		OrderOnTheBook bool `json:"w,omitempty"` // Order on the book
		CurrentExecutionType string `json:"x,omitempty"` // Current execution type
		CumulativeFilledQuantity string `json:"z,omitempty"` // Cumulative filled quantity
	} `json:"event"`
	SubscriptionId int64 `json:"subscriptionId"` // subscriptionId
}


