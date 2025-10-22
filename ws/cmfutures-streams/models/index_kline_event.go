package models

// IndexKlineEvent represents global message '#/components/messages/indexKlineEvent'
type IndexKlineEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Pair string `json:"ps"` // Pair
	KlineData struct {
		KlineStartTime int64 `json:"t"` // Kline start time
		KlineCloseTime int64 `json:"T"` // Kline close time
		Symbol string `json:"s,omitempty"` // Symbol (Ignore for index kline)
		Interval string `json:"i"` // Interval
		FirstTradeID int64 `json:"f,omitempty"` // First trade ID (Ignore for index kline)
		LastTradeID int64 `json:"L,omitempty"` // Last trade ID (Ignore for index kline)
		OpenPrice string `json:"o"` // Open price
		ClosePrice string `json:"c"` // Close price
		HighPrice string `json:"h"` // High price
		LowPrice string `json:"l"` // Low price
		BaseAssetVolume string `json:"v,omitempty"` // Base asset volume (Ignore for index kline)
		NumberOfBasicData int `json:"n"` // Number of basic data
		IsThisKlineClosed bool `json:"x"` // Is this kline closed
		QuoteAssetVolume string `json:"q,omitempty"` // Quote asset volume (Ignore for index kline)
		TakerBuyBaseAssetVolume string `json:"V,omitempty"` // Taker buy base asset volume (Ignore for index kline)
		TakerBuyQuoteAssetVolume string `json:"Q,omitempty"` // Taker buy quote asset volume (Ignore for index kline)
		Ignore string `json:"B,omitempty"` // Ignore
	} `json:"k"` // Kline data
}


