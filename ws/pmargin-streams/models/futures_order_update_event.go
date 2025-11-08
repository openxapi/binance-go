package models

// FuturesOrderUpdateEvent represents global message '#/components/messages/futuresOrderUpdateEvent'
type FuturesOrderUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventBusinessUnit string `json:"fs,omitempty"` // Event business unit
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	AccountAlias string `json:"i,omitempty"` // Account Alias
	OrderDetails struct {
		Symbol string `json:"s,omitempty"` // Symbol
		ClientOrderId string `json:"c,omitempty"` // Client Order Id
		Side string `json:"S,omitempty"` // Side
		OrderType string `json:"o,omitempty"` // Order Type
		TimeInForce string `json:"f,omitempty"` // Time in Force
		OriginalQuantity string `json:"q,omitempty"` // Original Quantity
		OriginalPrice string `json:"p,omitempty"` // Original Price
		AveragePrice string `json:"ap,omitempty"` // Average Price
		StopPrice string `json:"sp,omitempty"` // Stop Price
		ExecutionType string `json:"x,omitempty"` // Execution Type
		OrderStatus string `json:"X,omitempty"` // Order Status
		OrderId int `json:"i,omitempty"` // Order Id
		OrderLastFilledQuantity string `json:"l,omitempty"` // Order Last Filled Quantity
		OrderFilledAccumulatedQuantity string `json:"z,omitempty"` // Order Filled Accumulated Quantity
		LastFilledPrice string `json:"L,omitempty"` // Last Filled Price
		CommissionAsset string `json:"N,omitempty"` // Commission Asset
		Commission string `json:"n,omitempty"` // Commission
		OrderTradeTime int64 `json:"T,omitempty"` // Order Trade Time (milliseconds)
		TradeId int64 `json:"t,omitempty"` // Trade Id
		BidsNotional string `json:"b,omitempty"` // Bids Notional
		AskNotional string `json:"a,omitempty"` // Ask Notional
		IsThisTradeTheMakerSide bool `json:"m,omitempty"` // Is this trade the maker side
		IsThisReduceOnly bool `json:"R,omitempty"` // Is this reduce only
		PositionSide string `json:"ps,omitempty"` // Position Side
		RealizedProfit string `json:"rp,omitempty"` // Realized Profit
		StratgyType string `json:"st,omitempty"` // Stratgy Type
		StrategyId int64 `json:"si,omitempty"` // Strategy Id
		SelfTradePreventionMode string `json:"V,omitempty"` // Self Trade Prevention Mode
		GoodTillDate int64 `json:"gtd,omitempty"` // Good Till Date
	} `json:"o,omitempty"` // Order details
}


