package models

import (
	"encoding/json"
)

// MyPreventedMatchesResponseRateLimitsItem represents a nested object structure
type MyPreventedMatchesResponseRateLimitsItem struct {
	// count property (example: 20)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 6000)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// MyPreventedMatchesResponseResultItem represents a nested object structure
type MyPreventedMatchesResponseResultItem struct {
	// makerOrderId property (example: 3)
	MakerOrderId int64 `json:"makerOrderId,omitempty"`
	// makerPreventedQuantity property (example: "1.300000")
	MakerPreventedQuantity string `json:"makerPreventedQuantity,omitempty"`
	// makerSymbol property (example: "BTCUSDT")
	MakerSymbol string `json:"makerSymbol,omitempty"`
	// preventedMatchId property (example: 1)
	PreventedMatchId int64 `json:"preventedMatchId,omitempty"`
	// price property (example: "1.100000")
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property (example: "EXPIRE_MAKER")
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// takerOrderId property (example: 5)
	TakerOrderId int64 `json:"takerOrderId,omitempty"`
	// tradeGroupId property (example: 1)
	TradeGroupId int64 `json:"tradeGroupId,omitempty"`
	// transactTime property (example: 1669101687094)
	TransactTime int64 `json:"transactTime,omitempty"`
}

// MyPreventedMatchesResponse - Receive response from myPreventedMatches
// Message name: Account prevented matches (USER_DATA) Response
type MyPreventedMatchesResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []MyPreventedMatchesResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []MyPreventedMatchesResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of MyPreventedMatchesResponse
func (s MyPreventedMatchesResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


