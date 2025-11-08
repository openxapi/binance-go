package models

// IndexPriceEvent represents global message '#/components/messages/indexPriceEvent'
type IndexPriceEvent struct {
	EventType string `json:"e"` // Event type
	EventTimestamp int64 `json:"E"` // Event timestamp
	UnderlyingSymbol string `json:"s"` // Underlying symbol
	IndexPrice string `json:"p"` // Index price
}


