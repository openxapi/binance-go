package models

import (
	"encoding/json"
)

// AccountCommissionResponseRateLimitsItem represents a nested object structure
type AccountCommissionResponseRateLimitsItem struct {
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

// AccountCommissionResponseResult represents a nested object structure
type AccountCommissionResponseResult struct {
	// discount property
	Discount *AccountCommissionResponseResultDiscount `json:"discount,omitempty"`
	// standardCommission property
	StandardCommission *AccountCommissionResponseResultStandardCommission `json:"standardCommission,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// taxCommission property
	TaxCommission *AccountCommissionResponseResultTaxCommission `json:"taxCommission,omitempty"`
}

// AccountCommissionResponseResultDiscount represents a nested object structure
type AccountCommissionResponseResultDiscount struct {
	// discount property
	Discount string `json:"discount,omitempty"`
	// discountAsset property
	DiscountAsset string `json:"discountAsset,omitempty"`
	// enabledForAccount property
	EnabledForAccount bool `json:"enabledForAccount,omitempty"`
	// enabledForSymbol property
	EnabledForSymbol bool `json:"enabledForSymbol,omitempty"`
}

// AccountCommissionResponseResultStandardCommission represents a nested object structure
type AccountCommissionResponseResultStandardCommission struct {
	// buyer property
	Buyer string `json:"buyer,omitempty"`
	// maker property
	Maker string `json:"maker,omitempty"`
	// seller property
	Seller string `json:"seller,omitempty"`
	// taker property
	Taker string `json:"taker,omitempty"`
}

// AccountCommissionResponseResultTaxCommission represents a nested object structure
type AccountCommissionResponseResultTaxCommission struct {
	// buyer property
	Buyer string `json:"buyer,omitempty"`
	// maker property
	Maker string `json:"maker,omitempty"`
	// seller property
	Seller string `json:"seller,omitempty"`
	// taker property
	Taker string `json:"taker,omitempty"`
}

// AccountCommissionResponse - Receive response from account.commission
// Message name: Account Commission Rates (USER_DATA) Response
type AccountCommissionResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []AccountCommissionResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *AccountCommissionResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of AccountCommissionResponse
func (s AccountCommissionResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


