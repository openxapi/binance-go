package models

// BookTickerEvent represents global message '#/components/messages/bookTickerEvent'
type BookTickerEvent struct {
	EventType string `json:"e"` // Event type
	OrderBookUpdateId int64 `json:"u"` // Order book updateId
	Symbol string `json:"s"` // Symbol
	BestBidPrice string `json:"b"` // Best bid price
	BestBidQuantity string `json:"B"` // Best bid quantity
	BestAskPrice string `json:"a"` // Best ask price
	BestAskQuantity string `json:"A"` // Best ask quantity
	TransactionTime int64 `json:"T"` // Transaction time
	EventTime int64 `json:"E"` // Event time
}


