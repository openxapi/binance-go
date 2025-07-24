package models

import (
	"encoding/json"
)

// TickerByUnderlyingEvent represents TickerByUnderlyingEvent
// 24-hour ticker statistics for options by underlying asset and expiration date.
// Stream name pattern: {underlyingAsset}@ticker@{expirationDate}
// Example: ETH@ticker@220930
// Update frequency: 1000ms (1 second)
// Only symbols with changed ticker info are sent.
// 
type TickerByUnderlyingEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time (timestamp)
	EventTime int64 `json:"E,omitempty"`
	// Transaction time (timestamp)
	TransactionTime int64 `json:"T,omitempty"`
	// Option symbol
	Symbol string `json:"s,omitempty"`
	// 24-hour opening price
	O string `json:"o,omitempty"`
	// Highest price
	HighestPrice string `json:"h,omitempty"`
	// Lowest price
	LowestPrice string `json:"l,omitempty"`
	// Latest price
	LatestPrice string `json:"c,omitempty"`
	// Trading volume (in contracts)
	TradingVolume string `json:"V,omitempty"`
	// Trade amount (in quote asset)
	TradeAmount string `json:"A,omitempty"`
	// Price change percent
	PriceChangePercent string `json:"P,omitempty"`
	// Price change
	PriceChange string `json:"p,omitempty"`
	// Volume of last completed trade (in contracts)
	VolumeOfLastCompletedTrade string `json:"Q,omitempty"`
	// First trade ID
	FirstTradeId string `json:"F,omitempty"`
	// Last trade ID
	LastTradeId string `json:"L,omitempty"`
	// Number of trades
	NumberOfTrades int `json:"n,omitempty"`
	// Best buy price
	BestBuyPrice string `json:"bo,omitempty"`
	// Best sell price
	BestSellPrice string `json:"ao,omitempty"`
	// Best buy quantity
	BestBuyQuantity string `json:"bq,omitempty"`
	// Best sell quantity
	BestSellQuantity string `json:"aq,omitempty"`
	// Buy implied volatility
	BuyImpliedVolatility string `json:"b,omitempty"`
	// Sell implied volatility
	SellImpliedVolatility string `json:"a,omitempty"`
	// Delta
	Delta string `json:"d,omitempty"`
	// Theta
	Theta string `json:"t,omitempty"`
	// Gamma
	Gamma string `json:"g,omitempty"`
	// Vega
	Vega string `json:"v,omitempty"`
	// Implied volatility
	ImpliedVolatility string `json:"vo,omitempty"`
	// Mark price
	MarkPrice string `json:"mp,omitempty"`
	// Buy maximum price
	BuyMaximumPrice string `json:"hl,omitempty"`
	// Sell minimum price
	SellMinimumPrice string `json:"ll,omitempty"`
	// Estimated exercise price
	EstimatedExercisePrice string `json:"eep,omitempty"`
}

// String returns string representation of TickerByUnderlyingEvent
func (s TickerByUnderlyingEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


