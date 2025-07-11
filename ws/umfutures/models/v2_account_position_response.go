package models

import (
	"encoding/json"
)

// V2AccountPositionResponseRateLimitsItem represents a nested object structure
type V2AccountPositionResponseRateLimitsItem struct {
	// count property (example: 20)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 2400)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// V2AccountPositionResponseResultItem represents a nested object structure
type V2AccountPositionResponseResultItem struct {
	// adl property (example: 0)
	Adl int64 `json:"adl,omitempty"`
	// askNotional property (example: "0")
	AskNotional string `json:"askNotional,omitempty"`
	// bidNotional property (example: "0")
	BidNotional string `json:"bidNotional,omitempty"`
	// breakEvenPrice property (example: "0.0")
	BreakEvenPrice string `json:"breakEvenPrice,omitempty"`
	// entryPrice property (example: "0.00000")
	EntryPrice string `json:"entryPrice,omitempty"`
	// initialMargin property (example: "0")
	InitialMargin string `json:"initialMargin,omitempty"`
	// isolatedMargin property (example: "0.00000000")
	IsolatedMargin string `json:"isolatedMargin,omitempty"`
	// isolatedWallet property (example: "0")
	IsolatedWallet string `json:"isolatedWallet,omitempty"`
	// liquidationPrice property (example: "0")
	LiquidationPrice string `json:"liquidationPrice,omitempty"`
	// maintMargin property (example: "0")
	MaintMargin string `json:"maintMargin,omitempty"`
	// marginAsset property (example: "USDT")
	MarginAsset string `json:"marginAsset,omitempty"`
	// markPrice property (example: "6679.50671178")
	MarkPrice string `json:"markPrice,omitempty"`
	// notional property (example: "0")
	Notional string `json:"notional,omitempty"`
	// openOrderInitialMargin property (example: "0")
	OpenOrderInitialMargin string `json:"openOrderInitialMargin,omitempty"`
	// positionAmt property (example: "1.000")
	PositionAmt string `json:"positionAmt,omitempty"`
	// positionInitialMargin property (example: "0")
	PositionInitialMargin string `json:"positionInitialMargin,omitempty"`
	// positionSide property (example: "BOTH")
	PositionSide string `json:"positionSide,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// unrealizedProfit property (example: "0.00000000")
	UnrealizedProfit string `json:"unrealizedProfit,omitempty"`
	// updateTime property (example: 0)
	UpdateTime int64 `json:"updateTime,omitempty"`
}

// V2AccountPositionResponse - Receive response from v2/account.position
// Message name: Position Information V2 (USER_DATA) Response
type V2AccountPositionResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []V2AccountPositionResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []V2AccountPositionResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of V2AccountPositionResponse
func (s V2AccountPositionResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


