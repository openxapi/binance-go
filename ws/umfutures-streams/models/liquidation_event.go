package models

// LiquidationEvent represents global message '#/components/messages/liquidationEvent'
type LiquidationEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	LiquidationOrder struct {
		Symbol string `json:"s"` // Symbol
		Side string `json:"S"` // Side
		OrderType string `json:"o"` // Order type
		TimeInForce string `json:"f"` // Time in force
		OriginalQuantity string `json:"q"` // Original quantity
		Price string `json:"p"` // Price
		AveragePrice string `json:"ap"` // Average price
		OrderStatus string `json:"X"` // Order status
		LastFilledQuantity string `json:"l"` // Last filled quantity
		AccumulatedFilledQuantity string `json:"z"` // Accumulated filled quantity
		TradeTime int64 `json:"T"` // Trade time
	} `json:"o"` // Liquidation order
}


