package models

// PartialDepthEvent represents global message '#/components/messages/partialDepthEvent'
type PartialDepthEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	TransactionTime int64 `json:"T"` // Transaction time
	Symbol string `json:"s"` // Symbol
	FirstUpdateIDInEvent int64 `json:"U"` // First update ID in event
	FinalUpdateIDInEvent int64 `json:"u"` // Final update ID in event
	FinalUpdateIDInLastStream int64 `json:"pu"` // Final update ID in last stream
	Bids [][]string `json:"b"` // Bids
	Asks [][]string `json:"a"` // Asks
}


