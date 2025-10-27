package models

// KlineEvent represents global message '#/components/messages/klineEvent'
type KlineEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	KlineData struct {
		KlineStartTime int64 `json:"t"` // Kline start time
		KlineCloseTime int64 `json:"T"` // Kline close time
		Symbol string `json:"s"` // Symbol
		Interval string `json:"i"` // Interval
		FirstTradeID int64 `json:"f"` // First trade ID
		LastTradeID int64 `json:"L"` // Last trade ID
		OpenPrice string `json:"o"` // Open price
		ClosePrice string `json:"c"` // Close price
		HighPrice string `json:"h"` // High price
		LowPrice string `json:"l"` // Low price
		BaseAssetVolume string `json:"v"` // Base asset volume
		NumberOfTrades int `json:"n"` // Number of trades
		IsThisKlineClosed bool `json:"x"` // Is this kline closed
		QuoteAssetVolume string `json:"q"` // Quote asset volume
		TakerBuyBaseAssetVolume string `json:"V"` // Taker buy base asset volume
		TakerBuyQuoteAssetVolume string `json:"Q"` // Taker buy quote asset volume
		Ignore string `json:"B"` // Ignore
	} `json:"k"` // Kline data
}


