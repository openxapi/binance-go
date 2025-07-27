package models

import (
	"encoding/json"
)

// OrderListPlaceResponseRateLimitsItem represents a nested object structure
type OrderListPlaceResponseRateLimitsItem struct {
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

// OrderListPlaceResponseResult represents a nested object structure
type OrderListPlaceResponseResult struct {
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
	OrderReports []OrderListPlaceResponseResultOrderReportsItem `json:"orderReports,omitempty"`
	// orders property
	Orders []OrderListPlaceResponseResultOrdersItem `json:"orders,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// transactionTime property
	TransactionTime int64 `json:"transactionTime,omitempty"`
}

// OrderListPlaceResponseResultOrderReportsItem represents a nested object structure
type OrderListPlaceResponseResultOrderReportsItem struct {
	// clientOrderId property (example: "BqtFCj5odMoWtSqGk2X9tU")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property (example: "0.00000000")
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property (example: "0.00000000")
	ExecutedQty string `json:"executedQty,omitempty"`
	// orderId property (example: 12569138901)
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property (example: 1274512)
	OrderListId int64 `json:"orderListId,omitempty"`
	// origQty property (example: "0.00650000")
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property (example: "0.000000")
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// price property (example: "23410.00000000")
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property (example: "NONE")
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// side property (example: "SELL")
	Side string `json:"side,omitempty"`
	// status property (example: "NEW")
	Status string `json:"status,omitempty"`
	// stopPrice property (example: "23405.00000000")
	StopPrice string `json:"stopPrice,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property (example: "GTC")
	TimeInForce string `json:"timeInForce,omitempty"`
	// transactTime property (example: 1660801713793)
	TransactTime int64 `json:"transactTime,omitempty"`
	// type property (example: "STOP_LOSS_LIMIT")
	Type string `json:"type,omitempty"`
	// workingTime property (example: -1)
	WorkingTime int64 `json:"workingTime,omitempty"`
}

// OrderListPlaceResponseResultOrdersItem represents a nested object structure
type OrderListPlaceResponseResultOrdersItem struct {
	// clientOrderId property (example: "BqtFCj5odMoWtSqGk2X9tU")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// orderId property (example: 12569138901)
	OrderId int64 `json:"orderId,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
}

// OrderListPlaceResponse - Receive response from orderList.place
// Message name: Place new OCO - Deprecated (TRADE) Response
type OrderListPlaceResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderListPlaceResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *OrderListPlaceResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderListPlaceResponse
func (s OrderListPlaceResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


