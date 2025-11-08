package models

// TickerEvent represents global message '#/components/messages/tickerEvent'
type TickerEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	Pair string `json:"ps"` // Pair
	PriceChange string `json:"p"` // Price change
	PriceChangePercent string `json:"P"` // Price change percent
	WeightedAveragePrice string `json:"w"` // Weighted average price
	LastPrice string `json:"c"` // Last price
	LastQuantity string `json:"Q"` // Last quantity
	OpenPrice string `json:"o"` // Open price
	HighPrice string `json:"h"` // High price
	LowPrice string `json:"l"` // Low price
	TotalTradedVolume string `json:"v"` // Total traded volume
	TotalTradedBaseAssetVolume string `json:"q"` // Total traded base asset volume
	StatisticsOpenTime int64 `json:"O"` // Statistics open time
	StatisticsCloseTime int64 `json:"C"` // Statistics close time
	FirstTradeID int64 `json:"F"` // First trade ID
	LastTradeID int64 `json:"L"` // Last trade ID
	TotalNumberOfTrades int `json:"n"` // Total number of trades
}


