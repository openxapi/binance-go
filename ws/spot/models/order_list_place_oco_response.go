package models

import (
	"encoding/json"
)

// OrderListPlaceOcoResponseRateLimitsItem represents a nested object structure
type OrderListPlaceOcoResponseRateLimitsItem struct {
	// count property (example: 2)
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

// OrderListPlaceOcoResponseResult represents a nested object structure
type OrderListPlaceOcoResponseResult struct {
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
	OrderReports []OrderListPlaceOcoResponseResultOrderReportsItem `json:"orderReports,omitempty"`
	// orders property
	Orders []OrderListPlaceOcoResponseResultOrdersItem `json:"orders,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// transactionTime property
	TransactionTime int64 `json:"transactionTime,omitempty"`
}

// OrderListPlaceOcoResponseResultOrderReportsItem represents a nested object structure
type OrderListPlaceOcoResponseResultOrderReportsItem struct {
	// clientOrderId property (example: "0m6I4wfxvTUrOBSMUl0OPU")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property (example: "0.00000000")
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property (example: "0.000000")
	ExecutedQty string `json:"executedQty,omitempty"`
	// orderId property (example: 2)
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property (example: 2)
	OrderListId int64 `json:"orderListId,omitempty"`
	// origQty property (example: "1.000000")
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property (example: "0.000000")
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// price property (example: "1.50000000")
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property (example: "NONE")
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// side property (example: "BUY")
	Side string `json:"side,omitempty"`
	// status property (example: "NEW")
	Status string `json:"status,omitempty"`
	// stopPrice property (example: "1.50000001")
	StopPrice string `json:"stopPrice,omitempty"`
	// symbol property (example: "LTCBNB")
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property (example: "GTC")
	TimeInForce string `json:"timeInForce,omitempty"`
	// transactTime property (example: 1711062760648)
	TransactTime int64 `json:"transactTime,omitempty"`
	// type property (example: "STOP_LOSS_LIMIT")
	Type string `json:"type,omitempty"`
	// workingTime property (example: -1)
	WorkingTime int64 `json:"workingTime,omitempty"`
}

// OrderListPlaceOcoResponseResultOrdersItem represents a nested object structure
type OrderListPlaceOcoResponseResultOrdersItem struct {
	// clientOrderId property (example: "0m6I4wfxvTUrOBSMUl0OPU")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// orderId property (example: 2)
	OrderId int64 `json:"orderId,omitempty"`
	// symbol property (example: "LTCBNB")
	Symbol string `json:"symbol,omitempty"`
}

// OrderListPlaceOcoResponse - Receive response from orderList.place.oco
// Message name: Place new Order list - OCO (TRADE) Response
type OrderListPlaceOcoResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderListPlaceOcoResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result OrderListPlaceOcoResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderListPlaceOcoResponse
func (s OrderListPlaceOcoResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


