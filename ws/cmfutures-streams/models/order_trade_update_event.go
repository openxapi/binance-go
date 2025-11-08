package models

// OrderTradeUpdateEvent represents global message '#/components/messages/orderTradeUpdateEvent'
type OrderTradeUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	AccountAlias string `json:"i,omitempty"` // Account Alias
	OrderDetails struct {
		Symbol string `json:"s,omitempty"` // Symbol
		ClientOrderID string `json:"c,omitempty"` // Client Order ID
		Side string `json:"S,omitempty"` // Side
		OrderType string `json:"o,omitempty"` // Order Type
		TimeInForce string `json:"f,omitempty"` // Time in Force
		OriginalQuantity string `json:"q,omitempty"` // Original Quantity
		OriginalPrice string `json:"p,omitempty"` // Original Price
		AveragePrice string `json:"ap,omitempty"` // Average Price
		StopPrice string `json:"sp,omitempty"` // Stop Price
		ExecutionType string `json:"x,omitempty"` // Execution Type
		OrderStatus string `json:"X,omitempty"` // Order Status
		OrderID int64 `json:"i,omitempty"` // Order ID
		OrderLastFilledQuantity string `json:"l,omitempty"` // Order Last Filled Quantity
		OrderFilledAccumulatedQuantity string `json:"z,omitempty"` // Order Filled Accumulated Quantity
		LastFilledPrice string `json:"L,omitempty"` // Last Filled Price
		CommissionAsset string `json:"N,omitempty"` // Commission Asset
		Commission string `json:"n,omitempty"` // Commission
		OrderTradeTime int64 `json:"T,omitempty"` // Order Trade Time
		TradeID int64 `json:"t,omitempty"` // Trade ID
		RealizedProfit string `json:"rp,omitempty"` // Realized Profit
		BidQuantityOfBaseAsset string `json:"b,omitempty"` // Bid quantity of base asset
		AskQuantityOfBaseAsset string `json:"a,omitempty"` // Ask quantity of base asset
		IsThisTradeTheMakerSide bool `json:"m,omitempty"` // Is this trade the maker side
		IsReduceOnly bool `json:"R,omitempty"` // Is reduce only
		StopPriceWorkingType string `json:"wt,omitempty"` // Stop Price Working Type
		OriginalOrderType string `json:"ot,omitempty"` // Original Order Type
		PositionSide string `json:"ps,omitempty"` // Position Side
		ConditionalOrder bool `json:"cp,omitempty"` // Conditional Order
		ActivationPrice string `json:"AP,omitempty"` // Activation Price
		CallbackRate string `json:"cr,omitempty"` // Callback Rate
		IsProtected bool `json:"pP,omitempty"` // Is protected
		STPMode string `json:"V,omitempty"` // STP mode
		PriceMatchMode string `json:"pm,omitempty"` // Price match mode
	} `json:"o,omitempty"` // Order details
}


