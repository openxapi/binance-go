package models

import (
	"encoding/json"
)

// OrderListStatusResponseRateLimitsItem represents a nested object structure
type OrderListStatusResponseRateLimitsItem struct {
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

// OrderListStatusResponseResult represents a nested object structure
type OrderListStatusResponseResult struct {
	// contingencyType property
	ContingencyType string `json:"contingencyType,omitempty"`
	// listClientOrderId property
	ListClientOrderId string `json:"listClientOrderId,omitempty"`
	// listOrderStatus property
	ListOrderStatus string `json:"listOrderStatus,omitempty"`
	// listStatusType property
	ListStatusType string `json:"listStatusType,omitempty"`
	// orderListId property
	OrderListId int64 `json:"orderListId,omitempty"`
	// orders property
	Orders []OrderListStatusResponseResultOrdersItem `json:"orders,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// transactionTime property
	TransactionTime int64 `json:"transactionTime,omitempty"`
}

// OrderListStatusResponseResultOrdersItem represents a nested object structure
type OrderListStatusResponseResultOrdersItem struct {
	// clientOrderId property (example: "BqtFCj5odMoWtSqGk2X9tU")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// orderId property (example: 12569138901)
	OrderId int64 `json:"orderId,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
}

// OrderListStatusResponse - Receive response from orderList.status
// Message name: Query Order list (USER_DATA) Response
type OrderListStatusResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderListStatusResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result OrderListStatusResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderListStatusResponse
func (s OrderListStatusResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


