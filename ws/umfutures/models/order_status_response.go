package models

import (
	"encoding/json"
)

// OrderStatusResponseResult represents a nested object structure
type OrderStatusResponseResult struct {
	// activatePrice property
	ActivatePrice string `json:"activatePrice,omitempty"`
	// avgPrice property
	AvgPrice string `json:"avgPrice,omitempty"`
	// clientOrderId property
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// closePosition property
	ClosePosition bool `json:"closePosition,omitempty"`
	// cumQuote property
	CumQuote string `json:"cumQuote,omitempty"`
	// executedQty property
	ExecutedQty string `json:"executedQty,omitempty"`
	// orderId property
	OrderId int64 `json:"orderId,omitempty"`
	// origQty property
	OrigQty string `json:"origQty,omitempty"`
	// origType property
	OrigType string `json:"origType,omitempty"`
	// positionSide property
	PositionSide string `json:"positionSide,omitempty"`
	// price property
	Price string `json:"price,omitempty"`
	// priceProtect property
	PriceProtect bool `json:"priceProtect,omitempty"`
	// priceRate property
	PriceRate string `json:"priceRate,omitempty"`
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
	// time property
	Time int64 `json:"time,omitempty"`
	// timeInForce property
	TimeInForce string `json:"timeInForce,omitempty"`
	// type property
	Type string `json:"type,omitempty"`
	// updateTime property
	UpdateTime int64 `json:"updateTime,omitempty"`
	// workingType property
	WorkingType string `json:"workingType,omitempty"`
}

// OrderStatusResponse - Receive response from order.status
// Message name: Query Order (USER_DATA) Response
type OrderStatusResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// result property
	Result OrderStatusResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderStatusResponse
func (s OrderStatusResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


