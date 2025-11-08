package models

// ContinuousKlineEvent represents global message '#/components/messages/continuousKlineEvent'
type ContinuousKlineEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Pair string `json:"ps"` // Pair
	ContractType string `json:"ct"` // Contract type
	KlineData struct {
		KlineStartTime int64 `json:"t"` // Kline start time
		KlineCloseTime int64 `json:"T"` // Kline close time
		Interval string `json:"i"` // Interval
		FirstUpdateID int64 `json:"f"` // First update ID
		LastUpdateID int64 `json:"L"` // Last update ID
		OpenPrice string `json:"o"` // Open price
		ClosePrice string `json:"c"` // Close price
		HighPrice string `json:"h"` // High price
		LowPrice string `json:"l"` // Low price
		Volume string `json:"v"` // Volume
		NumberOfTrades int `json:"n"` // Number of trades
		IsThisKlineClosed bool `json:"x"` // Is this kline closed
		BaseAssetVolume string `json:"q"` // Base asset volume
		TakerBuyVolume string `json:"V"` // Taker buy volume
		TakerBuyBaseAssetVolume string `json:"Q"` // Taker buy base asset volume
		Ignore string `json:"B"` // Ignore
	} `json:"k"` // Kline data
}


