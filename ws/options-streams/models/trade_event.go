package models

// TradeEvent represents global message '#/components/messages/tradeEvent'
type TradeEvent struct {
	EventType string `json:"e"` // Event type
	EventTimestamp int64 `json:"E"` // Event timestamp
	OptionTradingSymbol string `json:"s"` // Option trading symbol
	TradeID int64 `json:"t"` // Trade ID
	Price string `json:"p"` // Price
	Quantity string `json:"q"` // Quantity (can be negative)
	BuyOrderID int64 `json:"b"` // Buy order ID
	SellOrderID int64 `json:"a"` // Sell order ID
	TradeCompletedTimestamp int64 `json:"T"` // Trade completed timestamp
	TradeDirection string `json:"S"` // Trade direction (-1 indicates direction)
	TradeType string `json:"X"` // Trade type
}


