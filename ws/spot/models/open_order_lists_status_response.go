package models

import (
	"encoding/json"
)

// OpenOrderListsStatusResponseRateLimitsItem represents a nested object structure
type OpenOrderListsStatusResponseRateLimitsItem struct {
	// count property (example: 6)
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

// OpenOrderListsStatusResponseResultItem represents a nested object structure
type OpenOrderListsStatusResponseResultItem struct {
	// contingencyType property (example: "OCO")
	ContingencyType string `json:"contingencyType,omitempty"`
	// listClientOrderId property (example: "08985fedd9ea2cf6b28996")
	ListClientOrderId string `json:"listClientOrderId,omitempty"`
	// listOrderStatus property (example: "EXECUTING")
	ListOrderStatus string `json:"listOrderStatus,omitempty"`
	// listStatusType property (example: "EXEC_STARTED")
	ListStatusType string `json:"listStatusType,omitempty"`
	// orderListId property (example: 0)
	OrderListId int64 `json:"orderListId,omitempty"`
	// orders property
	Orders []OpenOrderListsStatusResponseResultItemOrdersItem `json:"orders,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// transactionTime property (example: 1660801713793)
	TransactionTime int64 `json:"transactionTime,omitempty"`
}

// OpenOrderListsStatusResponseResultItemOrdersItem represents a nested object structure
type OpenOrderListsStatusResponseResultItemOrdersItem struct {
	// clientOrderId property (example: "CUhLgTXnX5n2c0gWiLpV4d")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// orderId property (example: 4)
	OrderId int64 `json:"orderId,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
}

// OpenOrderListsStatusResponse - Receive response from openOrderLists.status
// Message name: Current open Order lists (USER_DATA) Response
type OpenOrderListsStatusResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OpenOrderListsStatusResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []OpenOrderListsStatusResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OpenOrderListsStatusResponse
func (s OpenOrderListsStatusResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


