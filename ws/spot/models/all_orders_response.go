package models

import (
	"encoding/json"
)

// AllOrdersResponseRateLimitsItem represents a nested object structure
type AllOrdersResponseRateLimitsItem struct {
	// count property (example: 20)
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

// AllOrdersResponseResultItem represents a nested object structure
type AllOrdersResponseResultItem struct {
	// clientOrderId property (example: "4d96324ff9d44481926157")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property (example: "198.33521500")
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property (example: "0.00847000")
	ExecutedQty string `json:"executedQty,omitempty"`
	// icebergQty property (example: "0.00000000")
	IcebergQty string `json:"icebergQty,omitempty"`
	// isWorking property (example: true)
	IsWorking bool `json:"isWorking,omitempty"`
	// orderId property (example: 12569099453)
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property (example: -1)
	OrderListId int64 `json:"orderListId,omitempty"`
	// origQty property (example: "0.00847000")
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property (example: "0.00000000")
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// preventedMatchId property (example: 0)
	PreventedMatchId int64 `json:"preventedMatchId,omitempty"`
	// preventedQuantity property (example: "1.200000")
	PreventedQuantity string `json:"preventedQuantity,omitempty"`
	// price property (example: "23416.10000000")
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property (example: "NONE")
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// side property (example: "SELL")
	Side string `json:"side,omitempty"`
	// status property (example: "FILLED")
	Status string `json:"status,omitempty"`
	// stopPrice property (example: "0.00000000")
	StopPrice string `json:"stopPrice,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// time property (example: 1660801715639)
	Time int64 `json:"time,omitempty"`
	// timeInForce property (example: "GTC")
	TimeInForce string `json:"timeInForce,omitempty"`
	// type property (example: "LIMIT")
	Type string `json:"type,omitempty"`
	// updateTime property (example: 1660801717945)
	UpdateTime int64 `json:"updateTime,omitempty"`
	// workingTime property (example: 1660801715639)
	WorkingTime int64 `json:"workingTime,omitempty"`
}

// AllOrdersResponse - Receive response from allOrders
// Message name: Account order history (USER_DATA) Response
type AllOrdersResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []AllOrdersResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []AllOrdersResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of AllOrdersResponse
func (s AllOrdersResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


