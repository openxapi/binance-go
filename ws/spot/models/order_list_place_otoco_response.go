package models

import (
	"encoding/json"
)

// OrderListPlaceOtocoResponseRateLimitsItem represents a nested object structure
type OrderListPlaceOtocoResponseRateLimitsItem struct {
	// count property (example: 18)
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

// OrderListPlaceOtocoResponseResult represents a nested object structure
type OrderListPlaceOtocoResponseResult struct {
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
	OrderReports []OrderListPlaceOtocoResponseResultOrderReportsItem `json:"orderReports,omitempty"`
	// orders property
	Orders []OrderListPlaceOtocoResponseResultOrdersItem `json:"orders,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// transactionTime property
	TransactionTime int64 `json:"transactionTime,omitempty"`
}

// OrderListPlaceOtocoResponseResultOrderReportsItem represents a nested object structure
type OrderListPlaceOtocoResponseResultOrderReportsItem struct {
	// clientOrderId property (example: "OVQOpKwfmPCfaBTD0n7e7H")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property (example: "0.000000")
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property (example: "0.000000")
	ExecutedQty string `json:"executedQty,omitempty"`
	// orderId property (example: 23)
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property (example: 629)
	OrderListId int64 `json:"orderListId,omitempty"`
	// origQty property (example: "1.000000")
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property (example: "0.000000")
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// price property (example: "1.500000")
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property (example: "NONE")
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// side property (example: "BUY")
	Side string `json:"side,omitempty"`
	// status property (example: "NEW")
	Status string `json:"status,omitempty"`
	// symbol property (example: "LTCBNB")
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property (example: "GTC")
	TimeInForce string `json:"timeInForce,omitempty"`
	// transactTime property (example: 1712544408537)
	TransactTime int64 `json:"transactTime,omitempty"`
	// type property (example: "LIMIT")
	Type string `json:"type,omitempty"`
	// workingTime property (example: 1712544408537)
	WorkingTime int64 `json:"workingTime,omitempty"`
}

// OrderListPlaceOtocoResponseResultOrdersItem represents a nested object structure
type OrderListPlaceOtocoResponseResultOrdersItem struct {
	// clientOrderId property (example: "OVQOpKwfmPCfaBTD0n7e7H")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// orderId property (example: 23)
	OrderId int64 `json:"orderId,omitempty"`
	// symbol property (example: "LTCBNB")
	Symbol string `json:"symbol,omitempty"`
}

// OrderListPlaceOtocoResponse - Receive response from orderList.place.otoco
// Message name: Place new Order list - OTOCO (TRADE) Response
type OrderListPlaceOtocoResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderListPlaceOtocoResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result OrderListPlaceOtocoResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderListPlaceOtocoResponse
func (s OrderListPlaceOtocoResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


