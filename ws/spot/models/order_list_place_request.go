package models

import (
	"encoding/json"
)

// OrderListPlaceRequestParams contains the parameters for OrderListPlaceRequest
type OrderListPlaceRequestParams struct {
	ApiKey string `json:"apiKey"`
	// Arbitrary unique ID among open orders for the limit order. Automatically generated if not sent
	LimitClientOrderId *string `json:"limitClientOrderId,omitempty"`
	LimitIcebergQty *string `json:"limitIcebergQty,omitempty"`
	// Arbitrary numeric value identifying the limit order within an order strategy.
	LimitStrategyId *int64 `json:"limitStrategyId,omitempty"`
	// Arbitrary numeric value identifying the limit order strategy.Values smaller than 1000000 are reserved and cannot be used.
	LimitStrategyType *int64 `json:"limitStrategyType,omitempty"`
	// Arbitrary unique ID among open order lists. Automatically generated if not sent
	ListClientOrderId *string `json:"listClientOrderId,omitempty"`
	// Select response format: ACK, RESULT, FULL (default)
	NewOrderRespType *string `json:"newOrderRespType,omitempty"`
	// Price for the limit order
	Price string `json:"price"`
	Quantity string `json:"quantity"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	// The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	// BUY or SELL
	Side string `json:"side"`
	Signature string `json:"signature"`
	// Arbitrary unique ID among open orders for the stop order. Automatically generated if not sent
	StopClientOrderId *string `json:"stopClientOrderId,omitempty"`
	StopIcebergQty *string `json:"stopIcebergQty,omitempty"`
	StopLimitPrice *string `json:"stopLimitPrice,omitempty"`
	// See order.place for available options
	StopLimitTimeInForce *string `json:"stopLimitTimeInForce,omitempty"`
	// Either stopPrice or trailingDelta, or both must be specified
	StopPrice string `json:"stopPrice"`
	// Arbitrary numeric value identifying the stop order within an order strategy.
	StopStrategyId *int64 `json:"stopStrategyId,omitempty"`
	// Arbitrary numeric value identifying the stop order strategy.Values smaller than 1000000 are reserved and cannot be used.
	StopStrategyType *int64 `json:"stopStrategyType,omitempty"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
	// See Trailing Stop order FAQ
	TrailingDelta int64 `json:"trailingDelta"`
}

// OrderListPlaceRequest - Send a orderList.place request
// Message name: Place new OCO - Deprecated (TRADE) Request
type OrderListPlaceRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderListPlaceRequestParams `json:"params,omitempty"`
}

// NewOrderListPlaceRequest creates a new OrderListPlaceRequest instance
func NewOrderListPlaceRequest() *OrderListPlaceRequest {
	return &OrderListPlaceRequest{}
}

// String returns string representation of OrderListPlaceRequest
func (s OrderListPlaceRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetParams(value OrderListPlaceRequestParams) *OrderListPlaceRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetApiKey(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetLimitClientOrderId sets the limitClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetLimitClientOrderId(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.LimitClientOrderId = &value
	return r
}

// SetLimitIcebergQty sets the limitIcebergQty parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetLimitIcebergQty(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.LimitIcebergQty = &value
	return r
}

// SetLimitStrategyId sets the limitStrategyId parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetLimitStrategyId(value int64) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.LimitStrategyId = &value
	return r
}

// SetLimitStrategyType sets the limitStrategyType parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetLimitStrategyType(value int64) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.LimitStrategyType = &value
	return r
}

// SetListClientOrderId sets the listClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetListClientOrderId(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.ListClientOrderId = &value
	return r
}

// SetNewOrderRespType sets the newOrderRespType parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetNewOrderRespType(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.NewOrderRespType = &value
	return r
}

// SetPrice sets the price parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetPrice(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.Price = value
	return r
}

// SetQuantity sets the quantity parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetQuantity(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.Quantity = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetRecvWindow(value int64) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSelfTradePreventionMode sets the selfTradePreventionMode parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetSelfTradePreventionMode(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.SelfTradePreventionMode = &value
	return r
}

// SetSide sets the side parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetSide(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.Side = value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetSignature(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetStopClientOrderId sets the stopClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetStopClientOrderId(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.StopClientOrderId = &value
	return r
}

// SetStopIcebergQty sets the stopIcebergQty parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetStopIcebergQty(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.StopIcebergQty = &value
	return r
}

// SetStopLimitPrice sets the stopLimitPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetStopLimitPrice(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.StopLimitPrice = &value
	return r
}

// SetStopLimitTimeInForce sets the stopLimitTimeInForce parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetStopLimitTimeInForce(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.StopLimitTimeInForce = &value
	return r
}

// SetStopPrice sets the stopPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetStopPrice(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.StopPrice = value
	return r
}

// SetStopStrategyId sets the stopStrategyId parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetStopStrategyId(value int64) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.StopStrategyId = &value
	return r
}

// SetStopStrategyType sets the stopStrategyType parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetStopStrategyType(value int64) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.StopStrategyType = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetSymbol(value string) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetTimestamp(value int64) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}

// SetTrailingDelta sets the trailingDelta parameter and returns the struct for method chaining
func (r *OrderListPlaceRequest) SetTrailingDelta(value int64) *OrderListPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceRequestParams{}
	}
	r.Params.TrailingDelta = value
	return r
}


