package models

// TickerByUnderlyingEventItem is the item type for TickerByUnderlyingEvent
type TickerByUnderlyingEventItem struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time (timestamp)
	TransactionTime int64 `json:"T"` // Transaction time (timestamp)
	OptionSymbol string `json:"s"` // Option symbol
	X24HourOpeningPrice string `json:"o"` // 24-hour opening price
	HighestPrice string `json:"h"` // Highest price
	LowestPrice string `json:"l"` // Lowest price
	LatestPrice string `json:"c"` // Latest price
	TradingVolume string `json:"V"` // Trading volume (in contracts)
	TradeAmount string `json:"A"` // Trade amount (in quote asset)
	PriceChangePercent string `json:"P"` // Price change percent
	PriceChange string `json:"p"` // Price change
	VolumeOfLastCompletedTrade string `json:"Q"` // Volume of last completed trade (in contracts)
	FirstTradeID string `json:"F"` // First trade ID
	LastTradeID string `json:"L"` // Last trade ID
	NumberOfTrades int `json:"n"` // Number of trades
	BestBuyPrice string `json:"bo"` // Best buy price
	BestSellPrice string `json:"ao"` // Best sell price
	BestBuyQuantity string `json:"bq"` // Best buy quantity
	BestSellQuantity string `json:"aq"` // Best sell quantity
	BuyImpliedVolatility string `json:"b"` // Buy implied volatility
	SellImpliedVolatility string `json:"a"` // Sell implied volatility
	Delta string `json:"d"` // Delta
	Theta string `json:"t"` // Theta
	Gamma string `json:"g"` // Gamma
	Vega string `json:"v"` // Vega
	ImpliedVolatility string `json:"vo"` // Implied volatility
	MarkPrice string `json:"mp"` // Mark price
	BuyMaximumPrice string `json:"hl,omitempty"` // Buy maximum price
	SellMinimumPrice string `json:"ll,omitempty"` // Sell minimum price
	EstimatedExercisePrice string `json:"eep"` // Estimated exercise price
}

// TickerByUnderlyingEvent is an array of TickerByUnderlyingEventItem
type TickerByUnderlyingEvent []TickerByUnderlyingEventItem

