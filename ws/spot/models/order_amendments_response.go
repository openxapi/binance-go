package models

import (
	"encoding/json"
)

// OrderAmendmentsResponseRateLimitsItem represents a nested object structure
type OrderAmendmentsResponseRateLimitsItem struct {
	// count property (example: 4)
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

// OrderAmendmentsResponseResultItem represents a nested object structure
type OrderAmendmentsResponseResultItem struct {
	// executionId property (example: 60)
	ExecutionId int64 `json:"executionId,omitempty"`
	// newClientOrderId property (example: "xbxXh5SSwaHS7oUEOCI88B")
	NewClientOrderId string `json:"newClientOrderId,omitempty"`
	// newQty property (example: "5.00000000")
	NewQty string `json:"newQty,omitempty"`
	// orderId property (example: 23)
	OrderId int64 `json:"orderId,omitempty"`
	// origClientOrderId property (example: "my_pending_order")
	OrigClientOrderId string `json:"origClientOrderId,omitempty"`
	// origQty property (example: "7.00000000")
	OrigQty string `json:"origQty,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// time property (example: 1741924229819)
	Time int64 `json:"time,omitempty"`
}

// OrderAmendmentsResponse - Receive response from order.amendments
// Message name: Query Order Amendments (USER_DATA) Response
type OrderAmendmentsResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderAmendmentsResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []OrderAmendmentsResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderAmendmentsResponse
func (s OrderAmendmentsResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


