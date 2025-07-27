package models

import (
	"encoding/json"
)

// SubscribeRequest represents SubscribeRequest
// Request to subscribe to one or more streams. Stream names should follow
// Binance Options naming conventions (UPPERCASE symbols, lowercase stream types).
// 
// Supported stream patterns for Options:
// - Symbol streams: {symbol}@{streamType} (e.g., "BTC-210630-9000-P@ticker")
// - New symbol info: option_pair
// - Open interest: {underlyingAsset}@openInterest@{expirationDate} (e.g., "ETH@openInterest@221125")
// - Mark price: {underlyingAsset}@markPrice (e.g., "ETH@markPrice")
// - Kline streams: {symbol}@kline_{interval} (e.g., "BTC-200630-9000-P@kline_1m")
// - Ticker by underlying: {underlyingAsset}@ticker@{expirationDate} (e.g., "ETH@ticker@220930")
// - Index price: {symbol}@index (e.g., "ETHUSDT@index")
// - Trade streams: {symbol}@trade or {underlyingAsset}@trade
// 
type SubscribeRequest struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Array of stream names to subscribe to
	ArrayOfStreamNamesToSubscribeTo []string `json:"params,omitempty"`
	// Request ID
	RequestId int `json:"id,omitempty"`
}

// String returns string representation of SubscribeRequest
func (s SubscribeRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


