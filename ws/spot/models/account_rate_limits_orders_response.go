package models

import (
	"encoding/json"
)

// AccountRateLimitsOrdersResponseRateLimitsItem represents a nested object structure
type AccountRateLimitsOrdersResponseRateLimitsItem struct {
	// count property (example: 40)
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

// AccountRateLimitsOrdersResponseResultItem represents a nested object structure
type AccountRateLimitsOrdersResponseResultItem struct {
	// count property (example: 0)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "SECOND")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 10)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 50)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "ORDERS")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// AccountRateLimitsOrdersResponse - Receive response from account.rateLimits.orders
// Message name: Unfilled Order Count (USER_DATA) Response
type AccountRateLimitsOrdersResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []AccountRateLimitsOrdersResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []AccountRateLimitsOrdersResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of AccountRateLimitsOrdersResponse
func (s AccountRateLimitsOrdersResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


