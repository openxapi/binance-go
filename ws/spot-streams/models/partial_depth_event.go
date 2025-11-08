package models

// PartialDepthEvent represents global message '#/components/messages/partialDepthEvent'
type PartialDepthEvent struct {
	LastUpdateID int64 `json:"lastUpdateId"` // Last update ID
	Bids [][]string `json:"bids"` // Bids
	Asks [][]string `json:"asks"` // Asks
}


