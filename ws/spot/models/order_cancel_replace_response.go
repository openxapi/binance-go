package models

import (
	"encoding/json"
)

// OrderCancelReplaceResponseRateLimitsItem represents a nested object structure
type OrderCancelReplaceResponseRateLimitsItem struct {
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

// OrderCancelReplaceResponseResult represents a nested object structure
type OrderCancelReplaceResponseResult struct {
	// cancelResponse property
	CancelResponse *OrderCancelReplaceResponseResultCancelResponse `json:"cancelResponse,omitempty"`
	// cancelResult property
	CancelResult string `json:"cancelResult,omitempty"`
	// newOrderResponse property
	NewOrderResponse *OrderCancelReplaceResponseResultNewOrderResponse `json:"newOrderResponse,omitempty"`
	// newOrderResult property
	NewOrderResult string `json:"newOrderResult,omitempty"`
}

// OrderCancelReplaceResponseResultCancelResponse represents a nested object structure
type OrderCancelReplaceResponseResultCancelResponse struct {
	// clientOrderId property
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property
	ExecutedQty string `json:"executedQty,omitempty"`
	// orderId property
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property
	OrderListId int64 `json:"orderListId,omitempty"`
	// origClientOrderId property
	OrigClientOrderId string `json:"origClientOrderId,omitempty"`
	// origQty property
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// price property
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// side property
	Side string `json:"side,omitempty"`
	// status property
	Status string `json:"status,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property
	TimeInForce string `json:"timeInForce,omitempty"`
	// transactTime property
	TransactTime int64 `json:"transactTime,omitempty"`
	// type property
	Type string `json:"type,omitempty"`
}

// OrderCancelReplaceResponseResultNewOrderResponse represents a nested object structure
type OrderCancelReplaceResponseResultNewOrderResponse struct {
	// clientOrderId property
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property
	ExecutedQty string `json:"executedQty,omitempty"`
	// orderId property
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property
	OrderListId int64 `json:"orderListId,omitempty"`
	// origQty property
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// price property
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// side property
	Side string `json:"side,omitempty"`
	// status property
	Status string `json:"status,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property
	TimeInForce string `json:"timeInForce,omitempty"`
	// transactTime property
	TransactTime int64 `json:"transactTime,omitempty"`
	// type property
	Type string `json:"type,omitempty"`
}

// OrderCancelReplaceResponse - Receive response from order.cancelReplace
// Message name: Cancel and replace order (TRADE) Response
type OrderCancelReplaceResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderCancelReplaceResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *OrderCancelReplaceResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderCancelReplaceResponse
func (s OrderCancelReplaceResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


