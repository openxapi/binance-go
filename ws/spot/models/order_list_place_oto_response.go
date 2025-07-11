package models

import (
	"encoding/json"
)

// OrderListPlaceOtoResponseRateLimitsItem represents a nested object structure
type OrderListPlaceOtoResponseRateLimitsItem struct {
	// count property (example: 10)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 10000000)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "ORDERS")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// OrderListPlaceOtoResponseResult represents a nested object structure
type OrderListPlaceOtoResponseResult struct {
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
	// orderReports property
	OrderReports []OrderListPlaceOtoResponseResultOrderReportsItem `json:"orderReports,omitempty"`
	// orders property
	Orders []OrderListPlaceOtoResponseResultOrdersItem `json:"orders,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// transactionTime property
	TransactionTime int64 `json:"transactionTime,omitempty"`
}

// OrderListPlaceOtoResponseResultOrderReportsItem represents a nested object structure
type OrderListPlaceOtoResponseResultOrderReportsItem struct {
	// clientOrderId property (example: "YiAUtM9yJjl1a2jXHSp9Ny")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property (example: "0.000000")
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property (example: "0.000000")
	ExecutedQty string `json:"executedQty,omitempty"`
	// orderId property (example: 13)
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property (example: 626)
	OrderListId int64 `json:"orderListId,omitempty"`
	// origQty property (example: "1.000000")
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property (example: "0.000000")
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// price property (example: "1.000000")
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property (example: "NONE")
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// side property (example: "SELL")
	Side string `json:"side,omitempty"`
	// status property (example: "NEW")
	Status string `json:"status,omitempty"`
	// symbol property (example: "LTCBNB")
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property (example: "GTC")
	TimeInForce string `json:"timeInForce,omitempty"`
	// transactTime property (example: 1712544395981)
	TransactTime int64 `json:"transactTime,omitempty"`
	// type property (example: "LIMIT")
	Type string `json:"type,omitempty"`
	// workingTime property (example: 1712544395981)
	WorkingTime int64 `json:"workingTime,omitempty"`
}

// OrderListPlaceOtoResponseResultOrdersItem represents a nested object structure
type OrderListPlaceOtoResponseResultOrdersItem struct {
	// clientOrderId property (example: "YiAUtM9yJjl1a2jXHSp9Ny")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// orderId property (example: 13)
	OrderId int64 `json:"orderId,omitempty"`
	// symbol property (example: "LTCBNB")
	Symbol string `json:"symbol,omitempty"`
}

// OrderListPlaceOtoResponse - Receive response from orderList.place.oto
// Message name: Place new Order list - OTO (TRADE) Response
type OrderListPlaceOtoResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderListPlaceOtoResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result OrderListPlaceOtoResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderListPlaceOtoResponse
func (s OrderListPlaceOtoResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


