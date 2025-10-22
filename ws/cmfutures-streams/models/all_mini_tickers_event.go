package models

// AllMiniTickersEventItem is the item type for AllMiniTickersEvent
type AllMiniTickersEventItem struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	Pair string `json:"ps"` // Pair
	ClosePrice string `json:"c"` // Close price
	OpenPrice string `json:"o"` // Open price
	HighPrice string `json:"h"` // High price
	LowPrice string `json:"l"` // Low price
	TotalTradedVolume string `json:"v"` // Total traded volume
	TotalTradedBaseAssetVolume string `json:"q"` // Total traded base asset volume
}

// AllMiniTickersEvent is an array of AllMiniTickersEventItem
type AllMiniTickersEvent []AllMiniTickersEventItem

