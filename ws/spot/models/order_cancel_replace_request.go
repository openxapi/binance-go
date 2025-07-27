package models

import (
	"encoding/json"
)

// OrderCancelReplaceRequestParams contains the parameters for OrderCancelReplaceRequest
type OrderCancelReplaceRequestParams struct {
	ApiKey string `json:"apiKey"`
	// New ID for the canceled order. Automatically generated if not sent
	CancelNewClientOrderId *string `json:"cancelNewClientOrderId,omitempty"`
	// Cancel order by orderId
	CancelOrderId int64 `json:"cancelOrderId"`
	// Cancel order by clientOrderId
	CancelOrigClientOrderId *string `json:"cancelOrigClientOrderId,omitempty"`
	CancelReplaceMode string `json:"cancelReplaceMode"`
	// Supported values: ONLY_NEW - Cancel will succeed if the order status is NEW. ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED. For more information please refer to Regarding cancelRestrictions.
	CancelRestrictions *string `json:"cancelRestrictions,omitempty"`
	IcebergQty *string `json:"icebergQty,omitempty"`
	// Arbitrary unique ID among open orders. Automatically generated if not sent
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	// Select response format: ACK, RESULT, FULL.
	// 
	// MARKET and LIMIT orders produce FULL response by default,
	// other order types default to ACK.
	NewOrderRespType *string `json:"newOrderRespType,omitempty"`
	// Supported values:  DO_NOTHING (default)- will only attempt to cancel the order if account has not exceeded the unfilled order rate limit CANCEL_ONLY - will always cancel the order.
	OrderRateLimitExceededMode *string `json:"orderRateLimitExceededMode,omitempty"`
	Price *string `json:"price,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	QuoteOrderQty *string `json:"quoteOrderQty,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	// The allowed enums is dependent on what is configured on the symbol.
	// Supported values: STP Modes.
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	// BUY or SELL
	Side string `json:"side"`
	Signature string `json:"signature"`
	StopPrice *string `json:"stopPrice,omitempty"`
	// Arbitrary numeric value identifying the order within an order strategy.
	StrategyId *int64 `json:"strategyId,omitempty"`
	// Arbitrary numeric value identifying the order strategy.
	// Values smaller than 1000000 are reserved and cannot be used.
	StrategyType *int64 `json:"strategyType,omitempty"`
	Symbol string `json:"symbol"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Timestamp int64 `json:"timestamp"`
	// See Trailing Stop order FAQ
	TrailingDelta *string `json:"trailingDelta,omitempty"`
	Type string `json:"type"`
}

// OrderCancelReplaceRequest - Send a order.cancelReplace request
// Message name: Cancel and replace order (TRADE) Request
type OrderCancelReplaceRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderCancelReplaceRequestParams `json:"params,omitempty"`
}

// NewOrderCancelReplaceRequest creates a new OrderCancelReplaceRequest instance
func NewOrderCancelReplaceRequest() *OrderCancelReplaceRequest {
	return &OrderCancelReplaceRequest{}
}

// String returns string representation of OrderCancelReplaceRequest
func (s OrderCancelReplaceRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetParams(value OrderCancelReplaceRequestParams) *OrderCancelReplaceRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetApiKey(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetCancelNewClientOrderId sets the cancelNewClientOrderId parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetCancelNewClientOrderId(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.CancelNewClientOrderId = &value
	return r
}

// SetCancelOrderId sets the cancelOrderId parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetCancelOrderId(value int64) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.CancelOrderId = value
	return r
}

// SetCancelOrigClientOrderId sets the cancelOrigClientOrderId parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetCancelOrigClientOrderId(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.CancelOrigClientOrderId = &value
	return r
}

// SetCancelReplaceMode sets the cancelReplaceMode parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetCancelReplaceMode(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.CancelReplaceMode = value
	return r
}

// SetCancelRestrictions sets the cancelRestrictions parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetCancelRestrictions(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.CancelRestrictions = &value
	return r
}

// SetIcebergQty sets the icebergQty parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetIcebergQty(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.IcebergQty = &value
	return r
}

// SetNewClientOrderId sets the newClientOrderId parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetNewClientOrderId(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.NewClientOrderId = &value
	return r
}

// SetNewOrderRespType sets the newOrderRespType parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetNewOrderRespType(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.NewOrderRespType = &value
	return r
}

// SetOrderRateLimitExceededMode sets the orderRateLimitExceededMode parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetOrderRateLimitExceededMode(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.OrderRateLimitExceededMode = &value
	return r
}

// SetPrice sets the price parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetPrice(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.Price = &value
	return r
}

// SetQuantity sets the quantity parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetQuantity(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.Quantity = &value
	return r
}

// SetQuoteOrderQty sets the quoteOrderQty parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetQuoteOrderQty(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.QuoteOrderQty = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetRecvWindow(value int64) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSelfTradePreventionMode sets the selfTradePreventionMode parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetSelfTradePreventionMode(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.SelfTradePreventionMode = &value
	return r
}

// SetSide sets the side parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetSide(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.Side = value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetSignature(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetStopPrice sets the stopPrice parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetStopPrice(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.StopPrice = &value
	return r
}

// SetStrategyId sets the strategyId parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetStrategyId(value int64) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.StrategyId = &value
	return r
}

// SetStrategyType sets the strategyType parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetStrategyType(value int64) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.StrategyType = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetSymbol(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimeInForce sets the timeInForce parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetTimeInForce(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.TimeInForce = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetTimestamp(value int64) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}

// SetTrailingDelta sets the trailingDelta parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetTrailingDelta(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.TrailingDelta = &value
	return r
}

// SetType sets the type parameter and returns the struct for method chaining
func (r *OrderCancelReplaceRequest) SetType(value string) *OrderCancelReplaceRequest {
	if r.Params == nil {
		r.Params = &OrderCancelReplaceRequestParams{}
	}
	r.Params.Type = value
	return r
}


