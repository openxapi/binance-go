package models

import (
	"encoding/json"
)

// AccountStatusResponseRateLimitsItem represents a nested object structure
type AccountStatusResponseRateLimitsItem struct {
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

// AccountStatusResponseResult represents a nested object structure
type AccountStatusResponseResult struct {
	// assets property
	Assets []AccountStatusResponseResultAssetsItem `json:"assets,omitempty"`
	// availableBalance property
	AvailableBalance string `json:"availableBalance,omitempty"`
	// canDeposit property
	CanDeposit bool `json:"canDeposit,omitempty"`
	// canTrade property
	CanTrade bool `json:"canTrade,omitempty"`
	// canWithdraw property
	CanWithdraw bool `json:"canWithdraw,omitempty"`
	// feeTier property
	FeeTier int64 `json:"feeTier,omitempty"`
	// maxWithdrawAmount property
	MaxWithdrawAmount string `json:"maxWithdrawAmount,omitempty"`
	// multiAssetsMargin property
	MultiAssetsMargin bool `json:"multiAssetsMargin,omitempty"`
	// positions property
	Positions []AccountStatusResponseResultPositionsItem `json:"positions,omitempty"`
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
	// tradeGroupId property
	TradeGroupId int64 `json:"tradeGroupId,omitempty"`
	// updateTime property
	UpdateTime int64 `json:"updateTime,omitempty"`
}

// AccountStatusResponseResultAssetsItem represents a nested object structure
type AccountStatusResponseResultAssetsItem struct {
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
	// marginAvailable property (example: true)
	MarginAvailable bool `json:"marginAvailable,omitempty"`
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

// AccountStatusResponseResultPositionsItem represents a nested object structure
type AccountStatusResponseResultPositionsItem struct {
	// askNotional property (example: "0")
	AskNotional string `json:"askNotional,omitempty"`
	// bidNotional property (example: "0")
	BidNotional string `json:"bidNotional,omitempty"`
	// entryPrice property (example: "0.00000")
	EntryPrice string `json:"entryPrice,omitempty"`
	// initialMargin property (example: "0")
	InitialMargin string `json:"initialMargin,omitempty"`
	// isolated property (example: true)
	Isolated bool `json:"isolated,omitempty"`
	// leverage property (example: "100")
	Leverage string `json:"leverage,omitempty"`
	// maintMargin property (example: "0")
	MaintMargin string `json:"maintMargin,omitempty"`
	// maxNotional property (example: "250000")
	MaxNotional string `json:"maxNotional,omitempty"`
	// openOrderInitialMargin property (example: "0")
	OpenOrderInitialMargin string `json:"openOrderInitialMargin,omitempty"`
	// positionAmt property (example: "0")
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

// AccountStatusResponse - Receive response from account.status
// Message name: Account Information(USER_DATA) Response
type AccountStatusResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []AccountStatusResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result AccountStatusResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of AccountStatusResponse
func (s AccountStatusResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


