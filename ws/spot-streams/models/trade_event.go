package models

// TradeEvent represents global message '#/components/messages/tradeEvent'
type TradeEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	TradeID int64 `json:"t"` // Trade ID
	Price string `json:"p"` // Price
	Quantity string `json:"q"` // Quantity
	BuyerOrderID int64 `json:"b"` // Buyer order ID
	SellerOrderID int64 `json:"a"` // Seller order ID
	TradeTime int64 `json:"T"` // Trade time
	IsTheBuyerTheMarketMaker bool `json:"m"` // Is the buyer the market maker
	Ignore bool `json:"M"` // Ignore
}


