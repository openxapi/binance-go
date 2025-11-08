package models

// AveragePriceEvent represents global message '#/components/messages/averagePriceEvent'
type AveragePriceEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	AveragePriceInterval string `json:"i"` // Average price interval
	AveragePrice string `json:"w"` // Average price
	LastTradeTime int64 `json:"T"` // Last trade time
}


