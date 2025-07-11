package models

import (
	"encoding/json"
)

// AccountPositionResponseRateLimitsItem represents a nested object structure
type AccountPositionResponseRateLimitsItem struct {
	// count property (example: 10)
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

// AccountPositionResponseResultItem represents a nested object structure
type AccountPositionResponseResultItem struct {
	// breakEvenPrice property (example: "0.00000000")
	BreakEvenPrice string `json:"breakEvenPrice,omitempty"`
	// entryPrice property (example: "0.00000000")
	EntryPrice string `json:"entryPrice,omitempty"`
	// isAutoAddMargin property (example: "false")
	IsAutoAddMargin string `json:"isAutoAddMargin,omitempty"`
	// isolatedMargin property (example: "0.00000000")
	IsolatedMargin string `json:"isolatedMargin,omitempty"`
	// isolatedWallet property (example: "0")
	IsolatedWallet string `json:"isolatedWallet,omitempty"`
	// leverage property (example: "7")
	Leverage string `json:"leverage,omitempty"`
	// liquidationPrice property (example: "0")
	LiquidationPrice string `json:"liquidationPrice,omitempty"`
	// marginType property (example: "cross")
	MarginType string `json:"marginType,omitempty"`
	// markPrice property (example: "62297.60417296")
	MarkPrice string `json:"markPrice,omitempty"`
	// maxQty property (example: "100")
	MaxQty string `json:"maxQty,omitempty"`
	// notionalValue property (example: "0")
	NotionalValue string `json:"notionalValue,omitempty"`
	// positionAmt property (example: "0")
	PositionAmt string `json:"positionAmt,omitempty"`
	// positionSide property (example: "BOTH")
	PositionSide string `json:"positionSide,omitempty"`
	// symbol property (example: "BTCUSD_PERP")
	Symbol string `json:"symbol,omitempty"`
	// unRealizedProfit property (example: "0.00000000")
	UnRealizedProfit string `json:"unRealizedProfit,omitempty"`
	// updateTime property (example: 1726731195634)
	UpdateTime int64 `json:"updateTime,omitempty"`
}

// AccountPositionResponse - Receive response from account.position
// Message name: Position Information(USER_DATA) Response
type AccountPositionResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []AccountPositionResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []AccountPositionResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of AccountPositionResponse
func (s AccountPositionResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


