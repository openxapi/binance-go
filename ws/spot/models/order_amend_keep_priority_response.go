package models

import (
	"encoding/json"
)

// OrderAmendKeepPriorityResponseRateLimitsItem represents a nested object structure
type OrderAmendKeepPriorityResponseRateLimitsItem struct {
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

// OrderAmendKeepPriorityResponseResult represents a nested object structure
type OrderAmendKeepPriorityResponseResult struct {
	// amendedOrder property
	AmendedOrder *OrderAmendKeepPriorityResponseResultAmendedOrder `json:"amendedOrder,omitempty"`
	// executionId property
	ExecutionId int64 `json:"executionId,omitempty"`
	// transactTime property
	TransactTime int64 `json:"transactTime,omitempty"`
}

// OrderAmendKeepPriorityResponseResultAmendedOrder represents a nested object structure
type OrderAmendKeepPriorityResponseResultAmendedOrder struct {
	// clientOrderId property
	ClientOrderId string `json:"clientOrderId,omitempty"`
	// cumulativeQuoteQty property
	CumulativeQuoteQty string `json:"cumulativeQuoteQty,omitempty"`
	// executedQty property
	ExecutedQty string `json:"executedQty,omitempty"`
	// orderId property
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property
	OrderListId int64 `json:"orderListId,omitempty"`
	// origClientOrderId property
	OrigClientOrderId string `json:"origClientOrderId,omitempty"`
	// preventedQty property
	PreventedQty string `json:"preventedQty,omitempty"`
	// price property
	Price string `json:"price,omitempty"`
	// qty property
	Qty string `json:"qty,omitempty"`
	// quoteOrderQty property
	QuoteOrderQty string `json:"quoteOrderQty,omitempty"`
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
	// type property
	Type string `json:"type,omitempty"`
	// workingTime property
	WorkingTime int64 `json:"workingTime,omitempty"`
}

// OrderAmendKeepPriorityResponse - Receive response from order.amend.keepPriority
// Message name: Order Amend Keep Priority (TRADE) Response
type OrderAmendKeepPriorityResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []OrderAmendKeepPriorityResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *OrderAmendKeepPriorityResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of OrderAmendKeepPriorityResponse
func (s OrderAmendKeepPriorityResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


