package models

// IndexPriceEvent represents global message '#/components/messages/indexPriceEvent'
type IndexPriceEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Pair string `json:"i"` // Pair
	IndexPrice string `json:"p"` // Index price
}


