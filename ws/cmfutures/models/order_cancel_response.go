package models

import (
	"encoding/json"
)

// OrderCancelResponseRateLimitsItem represents a nested object structure
type OrderCancelResponseRateLimitsItem struct {
	// count property (example: 6)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 2400)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// OrderCancelResponseResult represents a nested object structure
type OrderCancelResponseResult struct {
	// avgPrice property
	AvgPrice string `json:"avgPrice,omitempty"`
	// clientOrderId property
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// closePosition property
	ClosePosition bool `json:"closePosition,omitempty"`
	// cumBase property
	CumBase string `json:"cumBase,omitempty"`
	// cumQty property
	CumQty string `json:"cumQty,omitempty"`
	// executedQty property
	ExecutedQty string `json:"executedQty,omitempty"`
	// orderId property
	OrderId int64 `json:"orderId,omitempty"`
	// origQty property
	OrigQty string `json:"origQty,omitempty"`
	// origType property
	OrigType string `json:"origType,omitempty"`
	// pair property
	Pair string `json:"pair,omitempty"`
	// positionSide property
	PositionSide string `json:"positionSide,omitempty"`
	// price property
	Price string `json:"price,omitempty"`
	// priceProtect property
	PriceProtect bool `json:"priceProtect,omitempty"`
	// reduceOnly property
	ReduceOnly bool `json:"reduceOnly,omitempty"`
	// side property
	Side string `json:"side,omitempty"`
	// status property
	Status string `json:"status,omitempty"`
	// stopPrice property
	StopPrice string `json:"stopPrice,omitempty"`
	// symbol property
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property
	TimeInForce string `json:"timeInForce,omitempty"`
	// type property
	Type string `json:"type,omitempty"`
	// updateTime property
	UpdateTime int64 `json:"updateTime,omitempty"`
	// workingType property
	WorkingType string `json:"workingType,omitempty"`
}

// OrderCancelResponse - Receive response from order.cancel
// Message name: Cancel Order (TRADE) Response
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


