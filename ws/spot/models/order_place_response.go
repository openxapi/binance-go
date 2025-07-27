package models

import (
	"encoding/json"
)

// OrderPlaceResponseRateLimitsItem represents a nested object structure
type OrderPlaceResponseRateLimitsItem struct {
	// count property (example: 1)
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

// OrderPlaceResponseResult represents a nested object structure
type OrderPlaceResponseResult struct {
	// clientOrderId property
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// orderId property
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property
	OrderListId int64 `json:"orderListId,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// transactTime property
	TransactTime int64 `json:"transactTime,omitempty"`
}

// OrderPlaceResponse - Receive response from order.place
// Message name: Place new order (TRADE) Response
type OrderPlaceResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderPlaceResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *OrderPlaceResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderPlaceResponse
func (s OrderPlaceResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


