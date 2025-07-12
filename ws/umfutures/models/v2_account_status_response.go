package models

import (
	"encoding/json"
)

// V2AccountStatusResponseRateLimitsItem represents a nested object structure
type V2AccountStatusResponseRateLimitsItem struct {
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

// V2AccountStatusResponseResult represents a nested object structure
type V2AccountStatusResponseResult struct {
	// assets property
	Assets []V2AccountStatusResponseResultAssetsItem `json:"assets,omitempty"`
	// availableBalance property
	AvailableBalance string `json:"availableBalance,omitempty"`
	// maxWithdrawAmount property
	MaxWithdrawAmount string `json:"maxWithdrawAmount,omitempty"`
	// positions property
	Positions []V2AccountStatusResponseResultPositionsItem `json:"positions,omitempty"`
	// totalCrossUnPnl property
	TotalCrossUnPnl string `json:"totalCrossUnPnl,omitempty"`
	// totalCrossWalletBalance property
	TotalCrossWalletBalance string `json:"totalCrossWalletBalance,omitempty"`
	// totalInitialMargin property
	TotalInitialMargin string `json:"totalInitialMargin,omitempty"`
	// totalMaintMargin property
	TotalMaintMargin string `json:"totalMaintMargin,omitempty"`
	// totalMarginBalance property
	TotalMarginBalance string `json:"totalMarginBalance,omitempty"`
	// totalOpenOrderInitialMargin property
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin,omitempty"`
	// totalPositionInitialMargin property
	TotalPositionInitialMargin string `json:"totalPositionInitialMargin,omitempty"`
	// totalUnrealizedProfit property
	TotalUnrealizedProfit string `json:"totalUnrealizedProfit,omitempty"`
	// totalWalletBalance property
	TotalWalletBalance string `json:"totalWalletBalance,omitempty"`
}

// V2AccountStatusResponseResultAssetsItem represents a nested object structure
type V2AccountStatusResponseResultAssetsItem struct {
	// asset property (example: "USDT")
	Asset string `json:"asset,omitempty"`
	// availableBalance property (example: "23.72469206")
	AvailableBalance string `json:"availableBalance,omitempty"`
	// crossUnPnl property (example: "0.00000000")
	CrossUnPnl string `json:"crossUnPnl,omitempty"`
	// crossWalletBalance property (example: "23.72469206")
	CrossWalletBalance string `json:"crossWalletBalance,omitempty"`
	// initialMargin property (example: "0.00000000")
	InitialMargin string `json:"initialMargin,omitempty"`
	// maintMargin property (example: "0.00000000")
	MaintMargin string `json:"maintMargin,omitempty"`
	// marginBalance property (example: "23.72469206")
	MarginBalance string `json:"marginBalance,omitempty"`
	// maxWithdrawAmount property (example: "23.72469206")
	MaxWithdrawAmount string `json:"maxWithdrawAmount,omitempty"`
	// openOrderInitialMargin property (example: "0.00000000")
	OpenOrderInitialMargin string `json:"openOrderInitialMargin,omitempty"`
	// positionInitialMargin property (example: "0.00000000")
	PositionInitialMargin string `json:"positionInitialMargin,omitempty"`
	// unrealizedProfit property (example: "0.00000000")
	UnrealizedProfit string `json:"unrealizedProfit,omitempty"`
	// updateTime property (example: 1625474304765)
	UpdateTime int64 `json:"updateTime,omitempty"`
	// walletBalance property (example: "23.72469206")
	WalletBalance string `json:"walletBalance,omitempty"`
}

// V2AccountStatusResponseResultPositionsItem represents a nested object structure
type V2AccountStatusResponseResultPositionsItem struct {
	// initialMargin property (example: "0")
	InitialMargin string `json:"initialMargin,omitempty"`
	// isolatedMargin property (example: "0.00000000")
	IsolatedMargin string `json:"isolatedMargin,omitempty"`
	// isolatedWallet property (example: "0")
	IsolatedWallet string `json:"isolatedWallet,omitempty"`
	// maintMargin property (example: "0")
	MaintMargin string `json:"maintMargin,omitempty"`
	// notional property (example: "0")
	Notional string `json:"notional,omitempty"`
	// positionAmt property (example: "1.000")
	PositionAmt string `json:"positionAmt,omitempty"`
	// positionSide property (example: "BOTH")
	PositionSide string `json:"positionSide,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// unrealizedProfit property (example: "0.00000000")
	UnrealizedProfit string `json:"unrealizedProfit,omitempty"`
	// updateTime property (example: 0)
	UpdateTime int64 `json:"updateTime,omitempty"`
}

// V2AccountStatusResponse - Receive response from v2/account.status
// Message name: Account Information V2(USER_DATA) Response
type V2AccountStatusResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []V2AccountStatusResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *V2AccountStatusResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of V2AccountStatusResponse
func (s V2AccountStatusResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


