package models

import (
	"encoding/json"
)

// OrderStatusResponseRateLimitsItem represents a nested object structure
type OrderStatusResponseRateLimitsItem struct {
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

// OrderStatusResponseResult represents a nested object structure
type OrderStatusResponseResult struct {
	// clientOrderId property
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property
	ExecutedQty string `json:"executedQty,omitempty"`
	// icebergQty property
	IcebergQty string `json:"icebergQty,omitempty"`
	// isWorking property
	IsWorking bool `json:"isWorking,omitempty"`
	// orderId property
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property
	OrderListId int64 `json:"orderListId,omitempty"`
	// origQty property
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// preventedMatchId property
	PreventedMatchId int64 `json:"preventedMatchId,omitempty"`
	// preventedQuantity property
	PreventedQuantity string `json:"preventedQuantity,omitempty"`
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
	// time property
	Time int64 `json:"time,omitempty"`
	// timeInForce property
	TimeInForce string `json:"timeInForce,omitempty"`
	// trailingDelta property
	TrailingDelta int64 `json:"trailingDelta,omitempty"`
	// trailingTime property
	TrailingTime int64 `json:"trailingTime,omitempty"`
	// type property
	Type string `json:"type,omitempty"`
	// updateTime property
	UpdateTime int64 `json:"updateTime,omitempty"`
	// workingTime property
	WorkingTime int64 `json:"workingTime,omitempty"`
}

// OrderStatusResponse - Receive response from order.status
// Message name: Query order (USER_DATA) Response
type OrderStatusResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderStatusResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *OrderStatusResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderStatusResponse
func (s OrderStatusResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


