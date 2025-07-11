package models

import (
	"encoding/json"
)

// OrderPlaceRequestParams contains the parameters for OrderPlaceRequest
type OrderPlaceRequestParams struct {
	ApiKey string `json:"apiKey"`
	IcebergQty *string `json:"icebergQty,omitempty"`
	// Arbitrary unique ID among open orders. Automatically generated if not sent
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	// Select response format: ACK, RESULT, FULL.MARKET and LIMIT orders use FULL by default, other order types default to ACK.
	NewOrderRespType *string `json:"newOrderRespType,omitempty"`
	Price *string `json:"price,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	QuoteOrderQty *string `json:"quoteOrderQty,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	// The allowed enums is dependent on what is configured on the symbol. Supported values: STP Modes
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	// BUY or SELL
	Side string `json:"side"`
	Signature string `json:"signature"`
	StopPrice *string `json:"stopPrice,omitempty"`
	// Arbitrary numeric value identifying the order within an order strategy.
	StrategyId *int64 `json:"strategyId,omitempty"`
	// Arbitrary numeric value identifying the order strategy.Values smaller than 1000000 are reserved and cannot be used.
	StrategyType *int64 `json:"strategyType,omitempty"`
	Symbol string `json:"symbol"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Timestamp int64 `json:"timestamp"`
	// See Trailing Stop order FAQ
	TrailingDelta *int64 `json:"trailingDelta,omitempty"`
	Type string `json:"type"`
}

// OrderPlaceRequest - Send a order.place request
// Message name: Place new order (TRADE) Request
type OrderPlaceRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderPlaceRequestParams `json:"params,omitempty"`
}

// NewOrderPlaceRequest creates a new OrderPlaceRequest instance
func NewOrderPlaceRequest() *OrderPlaceRequest {
	return &OrderPlaceRequest{}
}

// String returns string representation of OrderPlaceRequest
func (s OrderPlaceRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderPlaceRequest) SetParams(value OrderPlaceRequestParams) *OrderPlaceRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetApiKey(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetIcebergQty sets the icebergQty parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetIcebergQty(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.IcebergQty = &value
	return r
}

// SetNewClientOrderId sets the newClientOrderId parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetNewClientOrderId(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.NewClientOrderId = &value
	return r
}

// SetNewOrderRespType sets the newOrderRespType parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetNewOrderRespType(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.NewOrderRespType = &value
	return r
}

// SetPrice sets the price parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetPrice(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.Price = &value
	return r
}

// SetQuantity sets the quantity parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetQuantity(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.Quantity = &value
	return r
}

// SetQuoteOrderQty sets the quoteOrderQty parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetQuoteOrderQty(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.QuoteOrderQty = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetRecvWindow(value int64) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSelfTradePreventionMode sets the selfTradePreventionMode parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetSelfTradePreventionMode(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.SelfTradePreventionMode = &value
	return r
}

// SetSide sets the side parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetSide(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.Side = value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetSignature(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetStopPrice sets the stopPrice parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetStopPrice(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.StopPrice = &value
	return r
}

// SetStrategyId sets the strategyId parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetStrategyId(value int64) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.StrategyId = &value
	return r
}

// SetStrategyType sets the strategyType parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetStrategyType(value int64) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.StrategyType = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetSymbol(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimeInForce sets the timeInForce parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetTimeInForce(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.TimeInForce = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetTimestamp(value int64) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}

// SetTrailingDelta sets the trailingDelta parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetTrailingDelta(value int64) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.TrailingDelta = &value
	return r
}

// SetType sets the type parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetType(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.Type = value
	return r
}


