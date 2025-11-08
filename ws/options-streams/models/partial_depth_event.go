package models

// PartialDepthEvent represents global message '#/components/messages/partialDepthEvent'
type PartialDepthEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time (timestamp)
	TransactionTime int64 `json:"T"` // Transaction time (timestamp)
	OptionSymbol string `json:"s"` // Option symbol
	UpdateIDInEvent int64 `json:"u"` // Update ID in event
	PreviousUpdateID int64 `json:"pu"` // Previous update ID
	BuyOrders [][]string `json:"b"` // Buy orders (bids)
	SellOrders [][]string `json:"a"` // Sell orders (asks)
}


