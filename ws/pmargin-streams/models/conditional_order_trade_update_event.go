package models

// ConditionalOrderTradeUpdateEvent represents global message '#/components/messages/conditionalOrderTradeUpdateEvent'
type ConditionalOrderTradeUpdateEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	EventBusinessUnit string `json:"fs,omitempty"` // Event business unit
	StrategyOrderDetails struct {
		Symbol string `json:"s,omitempty"` // Symbol
		StrategyClientOrderId string `json:"c,omitempty"` // Strategy Client Order Id
		StrategyID int `json:"si,omitempty"` // Strategy ID
		Side string `json:"S,omitempty"` // Side
		StrategyType string `json:"st,omitempty"` // Strategy Type
		TimeInForce string `json:"f,omitempty"` // Time in Force
		Quantity string `json:"q,omitempty"` // Quantity
		Price string `json:"p,omitempty"` // Price
		StopPrice string `json:"sp,omitempty"` // Stop Price
		StrategyOrderStatus string `json:"os,omitempty"` // Strategy Order Status
		OrderBookTime int64 `json:"T,omitempty"` // Order book Time (milliseconds)
		OrderUpdateTime int64 `json:"ut,omitempty"` // Order update Time (milliseconds)
		IsThisReduceOnly bool `json:"R,omitempty"` // Is this reduce only
		StopPriceWorkingType string `json:"wt,omitempty"` // Stop Price Working Type
		PositionSide string `json:"ps,omitempty"` // Position Side
		IfCloseAll bool `json:"cp,omitempty"` // If Close-All
		ActivationPrice string `json:"AP,omitempty"` // Activation Price
		CallbackRate string `json:"cr,omitempty"` // Callback Rate
		OrderId int `json:"i,omitempty"` // Order Id
		STPMode string `json:"V,omitempty"` // STP mode
		GoodTillDate int `json:"gtd,omitempty"` // Good Till Date
	} `json:"so,omitempty"` // Strategy order details
}


