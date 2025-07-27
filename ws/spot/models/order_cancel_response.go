package models

import (
	"encoding/json"
)

// OrderCancelResponseRateLimitsItem represents a nested object structure
type OrderCancelResponseRateLimitsItem struct {
	// count property (example: 1)
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

// OrderCancelResponseResult represents a nested object structure
type OrderCancelResponseResult struct {
	// clientOrderId property
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property
	ExecutedQty string `json:"executedQty,omitempty"`
	// icebergQty property
	IcebergQty string `json:"icebergQty,omitempty"`
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
	// stopPrice property
	StopPrice string `json:"stopPrice,omitempty"`
	// strategyId property
	StrategyId int64 `json:"strategyId,omitempty"`
	// strategyType property
	StrategyType int64 `json:"strategyType,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property
	TimeInForce string `json:"timeInForce,omitempty"`
	// trailingDelta property
	TrailingDelta int64 `json:"trailingDelta,omitempty"`
	// transactTime property
	TransactTime int64 `json:"transactTime,omitempty"`
	// type property
	Type string `json:"type,omitempty"`
}

// OrderCancelResponse - Receive response from order.cancel
// Message name: Cancel order (TRADE) Response
type OrderCancelResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderCancelResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *OrderCancelResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderCancelResponse
func (s OrderCancelResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


