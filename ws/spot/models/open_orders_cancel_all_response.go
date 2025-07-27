package models

import (
	"encoding/json"
)

// OpenOrdersCancelAllResponseRateLimitsItem represents a nested object structure
type OpenOrdersCancelAllResponseRateLimitsItem struct {
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

// OpenOrdersCancelAllResponseResultItem represents a nested object structure
type OpenOrdersCancelAllResponseResultItem struct {
	// clientOrderId property (example: "91fe37ce9e69c90d6358c0")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property (example: "0.23416100")
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property (example: "0.00001000")
	ExecutedQty string `json:"executedQty,omitempty"`
	// icebergQty property (example: "0.00000000")
	IcebergQty string `json:"icebergQty,omitempty"`
	// orderId property (example: 12569099453)
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property (example: -1)
	OrderListId int64 `json:"orderListId,omitempty"`
	// origClientOrderId property (example: "4d96324ff9d44481926157")
	OrigClientOrderId string `json:"origClientOrderId,omitempty"`
	// origQty property (example: "0.00847000")
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property (example: "0.000000")
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// price property (example: "23416.10000000")
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property (example: "NONE")
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// side property (example: "SELL")
	Side string `json:"side,omitempty"`
	// status property (example: "CANCELED")
	Status string `json:"status,omitempty"`
	// stopPrice property (example: "0.00000000")
	StopPrice string `json:"stopPrice,omitempty"`
	// strategyId property (example: 37463720)
	StrategyId int64 `json:"strategyId,omitempty"`
	// strategyType property (example: 1000000)
	StrategyType int64 `json:"strategyType,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property (example: "GTC")
	TimeInForce string `json:"timeInForce,omitempty"`
	// trailingDelta property (example: 0)
	TrailingDelta int64 `json:"trailingDelta,omitempty"`
	// trailingTime property (example: -1)
	TrailingTime int64 `json:"trailingTime,omitempty"`
	// transactTime property (example: 1684804350068)
	TransactTime int64 `json:"transactTime,omitempty"`
	// type property (example: "LIMIT")
	Type string `json:"type,omitempty"`
}

// OpenOrdersCancelAllResponse - Receive response from openOrders.cancelAll
// Message name: Cancel open orders (TRADE) Response
type OpenOrdersCancelAllResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OpenOrdersCancelAllResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []OpenOrdersCancelAllResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OpenOrdersCancelAllResponse
func (s OpenOrdersCancelAllResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


