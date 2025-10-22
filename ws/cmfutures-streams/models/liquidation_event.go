package models

// LiquidationEvent represents global message '#/components/messages/liquidationEvent'
type LiquidationEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	LiquidationOrderDetails struct {
		Symbol string `json:"s"` // Symbol
		Pair string `json:"ps"` // Pair
		Side string `json:"S"` // Side
		OrderType string `json:"o"` // Order type
		TimeInForce string `json:"f"` // Time in force
		OriginalQuantity string `json:"q"` // Original quantity
		Price string `json:"p"` // Price
		AveragePrice string `json:"ap"` // Average price
		OrderStatus string `json:"X"` // Order status
		OrderLastFilledQuantity string `json:"l"` // Order Last Filled Quantity
		OrderFilledAccumulatedQuantity string `json:"z"` // Order Filled Accumulated Quantity
		OrderTradeTime int64 `json:"T"` // Order Trade Time
	} `json:"o"` // Liquidation order details
}


