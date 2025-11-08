package models

// DiffDepthEvent represents global message '#/components/messages/diffDepthEvent'
type DiffDepthEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	FirstUpdateIDInEvent int64 `json:"U"` // First update ID in event
	FinalUpdateIDInEvent int64 `json:"u"` // Final update ID in event
	BidsToBeUpdated [][]string `json:"b"` // Bids to be updated
	AsksToBeUpdated [][]string `json:"a"` // Asks to be updated
}


