package models

import (
	"encoding/json"
)

// OrderPlaceRequestParams contains the parameters for OrderPlaceRequest
type OrderPlaceRequestParams struct {
	// Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType)
	ActivationPrice *string `json:"activationPrice,omitempty"`
	// Used with TRAILING_STOP_MARKET orders, min 0.1, max 10 where 1 for 1%
	CallbackRate *string `json:"callbackRate,omitempty"`
	// true, false；Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
	ClosePosition *string `json:"closePosition,omitempty"`
	// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: ^[\.A-Z\:/a-z0-9_-]{1,36}$
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	// ACK,RESULT, default ACK
	NewOrderRespType *string `json:"newOrderRespType,omitempty"`
	// Default BOTH for One-way Mode; LONG or SHORT for Hedge Mode.  It must be sent in Hedge Mode.
	PositionSide *string `json:"positionSide,omitempty"`
	Price *string `json:"price,omitempty"`
	// only available for LIMIT/STOP/TAKE_PROFIT order; can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20; Can't be passed together with price
	PriceMatch *string `json:"priceMatch,omitempty"`
	// "TRUE" or "FALSE", default "FALSE". Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	PriceProtect *string `json:"priceProtect,omitempty"`
	// Quantity measured by contract number, Cannot be sent with closePosition=true
	Quantity *string `json:"quantity,omitempty"`
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	// true or false. default false. Cannot be sent in Hedge Mode; cannot be sent with closePosition=true (Close-All)
	ReduceOnly *string `json:"reduceOnly,omitempty"`
	// NONE: No STP / EXPIRE_TAKER:expire taker order when STP triggers/ EXPIRE_MAKER:expire taker order when STP triggers/ EXPIRE_BOTH:expire both orders when STP triggers; default NONE
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	// BUY or SELL
	Side string `json:"side"`
	// Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	StopPrice *string `json:"stopPrice,omitempty"`
	Symbol string `json:"symbol"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Timestamp int64 `json:"timestamp"`
	// LIMIT, MARKET, STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
	Type string `json:"type"`
	// stopPrice triggered by: "MARK_PRICE", "CONTRACT_PRICE". Default "CONTRACT_PRICE"
	WorkingType *string `json:"workingType,omitempty"`
}

// OrderPlaceRequest - Send a order.place request
// Message name: New Order(TRADE) Request
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

// SetActivationPrice sets the activationPrice parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetActivationPrice(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.ActivationPrice = &value
	return r
}

// SetCallbackRate sets the callbackRate parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetCallbackRate(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.CallbackRate = &value
	return r
}

// SetClosePosition sets the closePosition parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetClosePosition(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.ClosePosition = &value
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

// SetPositionSide sets the positionSide parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetPositionSide(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.PositionSide = &value
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

// SetPriceMatch sets the priceMatch parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetPriceMatch(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.PriceMatch = &value
	return r
}

// SetPriceProtect sets the priceProtect parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetPriceProtect(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.PriceProtect = &value
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

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetRecvWindow(value int64) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetReduceOnly sets the reduceOnly parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetReduceOnly(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.ReduceOnly = &value
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

// SetStopPrice sets the stopPrice parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetStopPrice(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.StopPrice = &value
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

// SetType sets the type parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetType(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.Type = value
	return r
}

// SetWorkingType sets the workingType parameter and returns the struct for method chaining
func (r *OrderPlaceRequest) SetWorkingType(value string) *OrderPlaceRequest {
	if r.Params == nil {
		r.Params = &OrderPlaceRequestParams{}
	}
	r.Params.WorkingType = &value
	return r
}


