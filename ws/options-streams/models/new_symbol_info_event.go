package models

// NewSymbolInfoEvent represents global message '#/components/messages/newSymbolInfoEvent'
type NewSymbolInfoEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time (timestamp)
	UnderlyingIndex string `json:"u"` // Underlying index
	QuotationAsset string `json:"qa"` // Quotation asset
	TradingPairName string `json:"s"` // Trading pair name (option symbol)
	ConversionRatio float64 `json:"unit"` // Conversion ratio (contract's underlying asset representation ratio)
	MinimumTradeVolume string `json:"mq"` // Minimum trade volume
	OptionType string `json:"d"` // Option type
	StrikePrice string `json:"sp"` // Strike price
	ExpirationTime int64 `json:"ed"` // Expiration time (timestamp)
}


