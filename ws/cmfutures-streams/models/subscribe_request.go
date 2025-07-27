package models

import (
	"encoding/json"
)

// SubscribeRequest represents SubscribeRequest
// Request to subscribe to one or more streams. Stream names should follow
// Binance naming conventions (lowercase, with @ separators).
// 
// Supported stream patterns for Coin-M Futures:
// - Symbol streams: {symbol}@{streamType} (e.g., "btcusd_perp@aggTrade")
// - Index streams: {pair}@indexPrice (e.g., "btcusd@indexPrice", "btcusd@indexPrice@1s")
// - Mark price streams: {symbol}@markPrice (e.g., "btcusd_perp@markPrice", "btcusd_200626@markPrice@1s")
// - Mark price for pair: {pair}@markPrice (e.g., "btcusd@markPrice")
// - All market streams: !{streamType}@arr (e.g., "!miniTicker@arr", "!ticker@arr", "!forceOrder@arr")
// - All book tickers: !bookTicker
// - Depth streams: {symbol}@depth{levels}{speed} (e.g., "btcusd_perp@depth5@100ms", "btcusd_200626@depth20")
// - Kline streams: {symbol}@kline_{interval} (e.g., "btcusd_perp@kline_1m", "btcusd_200626@kline_4h")
// - Continuous kline: {pair}_{contractType}@continuousKline_{interval}
// - Index kline: {pair}@indexPriceKline_{interval}
// - Mark price kline: {symbol}@markPriceKline_{interval}
// - Liquidation: {symbol}@forceOrder or !forceOrder@arr
// - Contract info: !contractInfo
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


