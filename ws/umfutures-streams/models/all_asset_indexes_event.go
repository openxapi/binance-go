package models

// AllAssetIndexesEventItem is the item type for AllAssetIndexesEvent
type AllAssetIndexesEventItem struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	AssetIndexSymbol string `json:"s"` // asset index symbol
	IndexPrice string `json:"i"` // index price
	BidBuffer string `json:"b"` // bid buffer
	AskBuffer string `json:"a"` // ask buffer
	BidRate string `json:"B"` // bid rate
	AskRate string `json:"A"` // ask rate
	AutoExchangeBidBuffer string `json:"q"` // auto exchange bid buffer
	AutoExchangeAskBuffer string `json:"g"` // auto exchange ask buffer
	AutoExchangeBidRate string `json:"Q"` // auto exchange bid rate
	AutoExchangeAskRate string `json:"G"` // auto exchange ask rate
}

// AllAssetIndexesEvent is an array of AllAssetIndexesEventItem
type AllAssetIndexesEvent []AllAssetIndexesEventItem

