package models

// TradeLiteEvent represents global message '#/components/messages/tradeLiteEvent'
type TradeLiteEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
	Symbol string `json:"s,omitempty"` // Symbol
	OriginalQuantity string `json:"q,omitempty"` // Original Quantity
	OriginalPrice string `json:"p,omitempty"` // Original Price
	MakerSide bool `json:"m,omitempty"` // Maker side
	ClientOrderId string `json:"c,omitempty"` // Client Order Id
	Side string `json:"S,omitempty"` // Side
	LastFilledPrice string `json:"L,omitempty"` // Last Filled Price
	OrderLastFilledQuantity string `json:"l,omitempty"` // Order Last Filled Quantity
	TradeID int64 `json:"t,omitempty"` // Trade ID
	OrderID int64 `json:"i,omitempty"` // Order ID
}


