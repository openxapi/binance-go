package models

import (
	"encoding/json"
)

// OrderModifyResponseRateLimitsItem represents a nested object structure
type OrderModifyResponseRateLimitsItem struct {
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

// OrderModifyResponseResult represents a nested object structure
type OrderModifyResponseResult struct {
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

// OrderModifyResponse - Receive response from order.modify
// Message name: Modify Order (TRADE) Response
type OrderModifyResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderModifyResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result OrderModifyResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderModifyResponse
func (s OrderModifyResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


