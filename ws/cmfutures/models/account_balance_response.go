package models

import (
	"encoding/json"
)

// AccountBalanceResponseRateLimitsItem represents a nested object structure
type AccountBalanceResponseRateLimitsItem struct {
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

// AccountBalanceResponseResultItem represents a nested object structure
type AccountBalanceResponseResultItem struct {
	// accountAlias property (example: "fWAuTiuXoCuXmY")
	AccountAlias string `json:"accountAlias,omitempty"`
	// asset property (example: "WLD")
	Asset string `json:"asset,omitempty"`
	// availableBalance property (example: "0.00000000")
	AvailableBalance string `json:"availableBalance,omitempty"`
	// balance property (example: "0.00000000")
	Balance string `json:"balance,omitempty"`
	// crossUnPnl property (example: "0.00000000")
	CrossUnPnl string `json:"crossUnPnl,omitempty"`
	// crossWalletBalance property (example: "0.00000000")
	CrossWalletBalance string `json:"crossWalletBalance,omitempty"`
	// updateTime property (example: 0)
	UpdateTime int64 `json:"updateTime,omitempty"`
	// withdrawAvailable property (example: "0.00000000")
	WithdrawAvailable string `json:"withdrawAvailable,omitempty"`
}

// AccountBalanceResponse - Receive response from account.balance
// Message name: Futures Account Balance(USER_DATA) Response
type AccountBalanceResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []AccountBalanceResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []AccountBalanceResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of AccountBalanceResponse
func (s AccountBalanceResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


