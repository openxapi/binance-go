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
	// limit property (example: 6000)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// AccountStatusResponseResult represents a nested object structure
type AccountStatusResponseResult struct {
	// accountType property
	AccountType string `json:"accountType,omitempty"`
	// balances property
	Balances []AccountStatusResponseResultBalancesItem `json:"balances,omitempty"`
	// brokered property
	Brokered bool `json:"brokered,omitempty"`
	// buyerCommission property
	BuyerCommission int64 `json:"buyerCommission,omitempty"`
	// canDeposit property
	CanDeposit bool `json:"canDeposit,omitempty"`
	// canTrade property
	CanTrade bool `json:"canTrade,omitempty"`
	// canWithdraw property
	CanWithdraw bool `json:"canWithdraw,omitempty"`
	// commissionRates property
	CommissionRates *AccountStatusResponseResultCommissionRates `json:"commissionRates,omitempty"`
	// makerCommission property
	MakerCommission int64 `json:"makerCommission,omitempty"`
	// permissions property
	Permissions []string `json:"permissions,omitempty"`
	// preventSor property
	PreventSor bool `json:"preventSor,omitempty"`
	// requireSelfTradePrevention property
	RequireSelfTradePrevention bool `json:"requireSelfTradePrevention,omitempty"`
	// sellerCommission property
	SellerCommission int64 `json:"sellerCommission,omitempty"`
	// takerCommission property
	TakerCommission int64 `json:"takerCommission,omitempty"`
	// uid property
	Uid int64 `json:"uid,omitempty"`
	// updateTime property
	UpdateTime int64 `json:"updateTime,omitempty"`
}

// AccountStatusResponseResultBalancesItem represents a nested object structure
type AccountStatusResponseResultBalancesItem struct {
	// asset property (example: "BNB")
	Asset string `json:"asset,omitempty"`
	// free property (example: "0.00000000")
	Free string `json:"free,omitempty"`
	// locked property (example: "0.00000000")
	Locked string `json:"locked,omitempty"`
}

// AccountStatusResponseResultCommissionRates represents a nested object structure
type AccountStatusResponseResultCommissionRates struct {
	// buyer property
	Buyer string `json:"buyer,omitempty"`
	// maker property
	Maker string `json:"maker,omitempty"`
	// seller property
	Seller string `json:"seller,omitempty"`
	// taker property
	Taker string `json:"taker,omitempty"`
}

// AccountStatusResponse - Receive response from account.status
// Message name: Account information (USER_DATA) Response
type AccountStatusResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []AccountStatusResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *AccountStatusResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of AccountStatusResponse
func (s AccountStatusResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


