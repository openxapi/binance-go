package models

import (
	"encoding/json"
)

// SubscribeRequest represents SubscribeRequest
// Request to subscribe to one or more streams. Stream names should follow
// Binance naming conventions (lowercase, with @ separators).
// 
// Supported stream patterns:
// 
// Individual Symbol Streams:
// - Trade: `{symbol}@aggTrade`, `{symbol}@trade`
// - Kline: `{symbol}@kline_{interval}`, `{symbol}@kline_{interval}@+08:00`
// - Ticker: `{symbol}@miniTicker`, `{symbol}@ticker`, `{symbol}@ticker_{window}`
// - Book: `{symbol}@bookTicker`, `{symbol}@avgPrice`
// - Depth: `{symbol}@depth`, `{symbol}@depth{levels}`, with optional @{speed}
// 
// All Market Streams (arrays):
// - `!ticker@arr`, `!miniTicker@arr`, `!ticker_{window}@arr`
// 
type SubscribeRequest struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Array of stream names to subscribe to
	ArrayOfStreamNamesToSubscribeTo []string `json:"params,omitempty"`
	// Request ID
	RequestId string `json:"id,omitempty"`
}

// String returns string representation of SubscribeRequest
func (s SubscribeRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


