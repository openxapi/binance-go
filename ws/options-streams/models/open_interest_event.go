package models

// OpenInterestEventItem is the item type for OpenInterestEvent
type OpenInterestEventItem struct {
	EventType string `json:"e"` // Event type
	EventTimestamp int64 `json:"E"` // Event timestamp
	OptionSymbol string `json:"s"` // Option symbol
	OpenInterestInContracts string `json:"o"` // Open interest in contracts
	OpenInterestInUSDT string `json:"h"` // Open interest in USDT
}

// OpenInterestEvent is an array of OpenInterestEventItem
type OpenInterestEvent []OpenInterestEventItem

