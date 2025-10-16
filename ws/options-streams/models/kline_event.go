package models

// KlineEvent represents global message '#/components/messages/klineEvent'
type KlineEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	OptionSymbol string `json:"s"` // Option symbol
	KlineData struct {
		KlineStartTime int64 `json:"t"` // Kline start time
		KlineEndTime int64 `json:"T"` // Kline end time
		Symbol string `json:"s"` // Symbol
		Interval string `json:"i"` // Interval
		FirstTradeID string `json:"F"` // First trade ID
		LastTradeID string `json:"L"` // Last trade ID
		OpenPrice string `json:"o"` // Open price
		ClosePrice string `json:"c"` // Close price
		HighPrice string `json:"h"` // High price
		LowPrice string `json:"l"` // Low price
		Volume string `json:"v"` // Volume (contracts)
		NumberOfTrades int `json:"n"` // Number of trades
		IsCandleCompleted bool `json:"x"` // Is candle completed
		CompletedTradeAmount string `json:"q"` // Completed trade amount
		TakerCompletedTradeVolume string `json:"V"` // Taker completed trade volume
		TakerTradeAmount string `json:"Q"` // Taker trade amount
	} `json:"k"` // Kline data
}


