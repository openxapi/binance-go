package models

import (
	"encoding/json"
)

// SorOrderPlaceResponseRateLimitsItem represents a nested object structure
type SorOrderPlaceResponseRateLimitsItem struct {
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

// SorOrderPlaceResponseResultItem represents a nested object structure
type SorOrderPlaceResponseResultItem struct {
	// clientOrderId property (example: "sBI1KM6nNtOfj5tccZSKly")
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cummulativeQuoteQty property (example: "14000.00000000")
	CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"`
	// executedQty property (example: "0.50000000")
	ExecutedQty string `json:"executedQty,omitempty"`
	// fills property
	Fills []SorOrderPlaceResponseResultItemFillsItem `json:"fills,omitempty"`
	// orderId property (example: 2)
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property (example: -1)
	OrderListId int64 `json:"orderListId,omitempty"`
	// origQty property (example: "0.50000000")
	OrigQty string `json:"origQty,omitempty"`
	// origQuoteOrderQty property (example: "0.000000")
	OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"`
	// price property (example: "31000.00000000")
	Price string `json:"price,omitempty"`
	// selfTradePreventionMode property (example: "NONE")
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	// side property (example: "BUY")
	Side string `json:"side,omitempty"`
	// status property (example: "FILLED")
	Status string `json:"status,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// timeInForce property (example: "GTC")
	TimeInForce string `json:"timeInForce,omitempty"`
	// transactTime property (example: 1689149087774)
	TransactTime int64 `json:"transactTime,omitempty"`
	// type property (example: "LIMIT")
	Type string `json:"type,omitempty"`
	// usedSor property (example: true)
	UsedSor bool `json:"usedSor,omitempty"`
	// workingFloor property (example: "SOR")
	WorkingFloor string `json:"workingFloor,omitempty"`
	// workingTime property (example: 1689149087774)
	WorkingTime int64 `json:"workingTime,omitempty"`
}

// SorOrderPlaceResponseResultItemFillsItem represents a nested object structure
type SorOrderPlaceResponseResultItemFillsItem struct {
	// allocId property (example: 0)
	AllocId int64 `json:"allocId,omitempty"`
	// commission property (example: "0.00000000")
	Commission string `json:"commission,omitempty"`
	// commissionAsset property (example: "BTC")
	CommissionAsset string `json:"commissionAsset,omitempty"`
	// matchType property (example: "ONE_PARTY_TRADE_REPORT")
	MatchType string `json:"matchType,omitempty"`
	// price property (example: "28000.00000000")
	Price string `json:"price,omitempty"`
	// qty property (example: "0.50000000")
	Qty string `json:"qty,omitempty"`
	// tradeId property (example: -1)
	TradeId int64 `json:"tradeId,omitempty"`
}

// SorOrderPlaceResponse - Receive response from sor.order.place
// Message name: Place new order using SOR (TRADE) Response
type SorOrderPlaceResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []SorOrderPlaceResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []SorOrderPlaceResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of SorOrderPlaceResponse
func (s SorOrderPlaceResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


