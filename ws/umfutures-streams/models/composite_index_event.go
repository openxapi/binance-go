package models

// CompositeIndexEvent represents global message '#/components/messages/compositeIndexEvent'
type CompositeIndexEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	Price string `json:"p"` // Price
	AssetType string `json:"C"` // Asset type
	Composition []struct {
		BaseAsset string `json:"b"` // Base asset
		QuoteAsset string `json:"q"` // Quote asset
		WeightInQuantity string `json:"w"` // Weight in quantity
		WeightInPercentage string `json:"W"` // Weight in percentage
		IndexPrice string `json:"i"` // Index price
	} `json:"c"` // Composition
}


