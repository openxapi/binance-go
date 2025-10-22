package models

// MarkPriceKlineEvent represents global message '#/components/messages/markPriceKlineEvent'
type MarkPriceKlineEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	KlineData struct {
		KlineStartTime int64 `json:"t"` // Kline start time
		KlineCloseTime int64 `json:"T"` // Kline close time
		Symbol string `json:"s"` // Symbol
		Interval string `json:"i"` // Interval
		FirstUpdateID int64 `json:"f,omitempty"` // First update ID (ignore for mark price kline)
		LastUpdateID int64 `json:"L,omitempty"` // Last update ID (ignore for mark price kline)
		OpenPrice string `json:"o"` // Open price
		ClosePrice string `json:"c"` // Close price
		HighPrice string `json:"h"` // High price
		LowPrice string `json:"l"` // Low price
		Volume string `json:"v,omitempty"` // Volume (ignore for mark price kline)
		NumberOfBasicData int `json:"n"` // Number of basic data
		IsThisKlineClosed bool `json:"x"` // Is this kline closed
		QuoteAssetVolume string `json:"q,omitempty"` // Quote asset volume (ignore for mark price kline)
		TakerBuyBaseAssetVolume string `json:"V,omitempty"` // Taker buy base asset volume (ignore for mark price kline)
		TakerBuyQuoteAssetVolume string `json:"Q,omitempty"` // Taker buy quote asset volume (ignore for mark price kline)
		Ignore string `json:"B,omitempty"` // ignore
	} `json:"k"` // Kline data
}


