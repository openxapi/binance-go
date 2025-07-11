package models

import (
	"encoding/json"
)

// AllOrderListsResponseRateLimitsItem represents a nested object structure
type AllOrderListsResponseRateLimitsItem struct {
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

// AllOrderListsResponseResultItem represents a nested object structure
type AllOrderListsResponseResultItem struct {
	// contingencyType property (example: "OCO")
	ContingencyType string `json:"contingencyType,omitempty"`
	// listClientOrderId property (example: "08985fedd9ea2cf6b28996")
	ListClientOrderId string `json:"listClientOrderId,omitempty"`
	// listOrderStatus property (example: "EXECUTING")
	ListOrderStatus string `json:"listOrderStatus,omitempty"`
	// listStatusType property (example: "EXEC_STARTED")
	ListStatusType string `json:"listStatusType,omitempty"`
	// orderListId property (example: 1274512)
	OrderListId int64 `json:"orderListId,omitempty"`
	// orders property
	Orders []AllOrderListsResponseResultItemOrdersItem `json:"orders,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// transactionTime property (example: 1660801713793)
	TransactionTime int64 `json:"transactionTime,omitempty"`
}

// AllOrderListsResponseResultItemOrdersItem represents a nested object structure
type AllOrderListsResponseResultItemOrdersItem struct {
	// clientOrderId property (example: "BqtFCj5odMoWtSqGk2X9tU")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// orderId property (example: 12569138901)
	OrderId int64 `json:"orderId,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
}

// AllOrderListsResponse - Receive response from allOrderLists
// Message name: Account order list history (USER_DATA) Response
type AllOrderListsResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []AllOrderListsResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []AllOrderListsResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of AllOrderListsResponse
func (s AllOrderListsResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


