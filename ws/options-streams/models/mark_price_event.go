package models

// MarkPriceEventItem is the item type for MarkPriceEvent
type MarkPriceEventItem struct {
	EventType string `json:"e"` // Event type
	EventTimestamp int64 `json:"E"` // Event timestamp
	OptionSymbol string `json:"s"` // Option symbol
	OptionMarkPrice string `json:"mp"` // Option mark price
}

// MarkPriceEvent is an array of MarkPriceEventItem
type MarkPriceEvent []MarkPriceEventItem

