package models

import (
	"encoding/json"
)

// TickerEvent represents TickerEvent
type TickerEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Price change
	Price string `json:"p,omitempty"`
	// Price change percent
	PriceChangePercent string `json:"P,omitempty"`
	// Weighted average price
	WeightedAveragePrice string `json:"w,omitempty"`
	// First trade(F)-1 price (previous close)
	IsKlineClosed string `json:"x,omitempty"`
	// Last price
	ClosePrice string `json:"c,omitempty"`
	// Last quantity
	TakerBuyQuoteVolume string `json:"Q,omitempty"`
	// Best bid price
	BuyerOrderId string `json:"b,omitempty"`
	// Best bid quantity
	BestBidQuantity string `json:"B,omitempty"`
	// Best ask price
	AggregateTradeId string `json:"a,omitempty"`
	// Best ask quantity
	SellerOrderId string `json:"A,omitempty"`
	// Open price
	OpenPrice string `json:"o,omitempty"`
	// High price
	HighPrice string `json:"h,omitempty"`
	// Low price
	LastTradeId string `json:"l,omitempty"`
	// Total traded base asset volume
	Volume string `json:"v,omitempty"`
	// Total traded quote asset volume
	Quantity string `json:"q,omitempty"`
	// Statistics open time
	OpenTime int64 `json:"O,omitempty"`
	// Statistics close time
	CloseTime int64 `json:"C,omitempty"`
	// First trade ID
	FirstTradeId int64 `json:"F,omitempty"`
	// Last trade ID
	LastTradeId2 int64 `json:"L,omitempty"`
	// Total number of trades
	NumberOfTrades int `json:"n,omitempty"`
}

// String returns string representation of TickerEvent
func (s TickerEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


