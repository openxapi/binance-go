package models

import (
	"encoding/json"
)

// SubscribeRequest represents SubscribeRequest
// Request to subscribe to one or more streams. Stream names should follow
// Binance naming conventions (lowercase, with @ separators).
// 
// Supported stream patterns:
// - Symbol streams: {symbol}@{streamType} (e.g., "btcusdt@aggTrade")
// - All market streams: !{streamType}@arr (e.g., "!ticker@arr")
// - Depth streams: {symbol}@depth{levels}{speed} (e.g., "btcusdt@depth5@100ms")
// - Kline streams: {symbol}@kline_{interval} (e.g., "btcusdt@kline_1m")
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


