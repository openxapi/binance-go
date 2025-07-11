package models

import (
	"encoding/json"
)

// OrderListPlaceOtoRequestParams contains the parameters for OrderListPlaceOtoRequest
type OrderListPlaceOtoRequestParams struct {
	// Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.  listClientOrderId is distinct from the workingClientOrderId and the pendingClientOrderId.
	ListClientOrderId *string `json:"listClientOrderId,omitempty"`
	// Format of the JSON response. Supported values: Order Response Type
	NewOrderRespType *string `json:"newOrderRespType,omitempty"`
	// Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
	PendingClientOrderId *string `json:"pendingClientOrderId,omitempty"`
	// This can only be used if pendingTimeInForce is GTC, or if pendingType is LIMIT_MAKER.
	PendingIcebergQty *string `json:"pendingIcebergQty,omitempty"`
	PendingPrice *string `json:"pendingPrice,omitempty"`
	// Sets the quantity for the pending order.
	PendingQuantity string `json:"pendingQuantity"`
	// Supported values: Order side
	PendingSide string `json:"pendingSide"`
	PendingStopPrice *string `json:"pendingStopPrice,omitempty"`
	// Arbitrary numeric value identifying the pending order within an order strategy.
	PendingStrategyId *int64 `json:"pendingStrategyId,omitempty"`
	// Arbitrary numeric value identifying the pending order strategy.  Values smaller than 1000000 are reserved and cannot be used.
	PendingStrategyType *int64 `json:"pendingStrategyType,omitempty"`
	// Supported values: Time In Force
	PendingTimeInForce *string `json:"pendingTimeInForce,omitempty"`
	PendingTrailingDelta *string `json:"pendingTrailingDelta,omitempty"`
	// Supported values: Order types.  Note that MARKET orders using quoteOrderQty are not supported.
	PendingType string `json:"pendingType"`
	// The value cannot be greater than 60000.
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	// The allowed values are dependent on what is configured on the symbol. Supported values: STP Modes
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Signature string `json:"signature"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
	// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
	WorkingClientOrderId *string `json:"workingClientOrderId,omitempty"`
	// This can only be used if workingTimeInForce is GTC, or if workingType is LIMIT_MAKER.
	WorkingIcebergQty *string `json:"workingIcebergQty,omitempty"`
	WorkingPrice string `json:"workingPrice"`
	// Sets the quantity for the working order.
	WorkingQuantity string `json:"workingQuantity"`
	// Supported values: Order side
	WorkingSide string `json:"workingSide"`
	// Arbitrary numeric value identifying the working order within an order strategy.
	WorkingStrategyId *int64 `json:"workingStrategyId,omitempty"`
	// Arbitrary numeric value identifying the working order strategy.  Values smaller than 1000000 are reserved and cannot be used.
	WorkingStrategyType *int64 `json:"workingStrategyType,omitempty"`
	// Supported values: Time In Force
	WorkingTimeInForce *string `json:"workingTimeInForce,omitempty"`
	// Supported values: LIMIT,LIMIT_MAKER
	WorkingType string `json:"workingType"`
}

// OrderListPlaceOtoRequest - Send a orderList.place.oto request
// Message name: Place new Order list - OTO (TRADE) Request
type OrderListPlaceOtoRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderListPlaceOtoRequestParams `json:"params,omitempty"`
}

// NewOrderListPlaceOtoRequest creates a new OrderListPlaceOtoRequest instance
func NewOrderListPlaceOtoRequest() *OrderListPlaceOtoRequest {
	return &OrderListPlaceOtoRequest{}
}

// String returns string representation of OrderListPlaceOtoRequest
func (s OrderListPlaceOtoRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetParams(value OrderListPlaceOtoRequestParams) *OrderListPlaceOtoRequest {
	r.Params = &value
	return r
}

// SetListClientOrderId sets the listClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetListClientOrderId(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.ListClientOrderId = &value
	return r
}

// SetNewOrderRespType sets the newOrderRespType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetNewOrderRespType(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.NewOrderRespType = &value
	return r
}

// SetPendingClientOrderId sets the pendingClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingClientOrderId(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingClientOrderId = &value
	return r
}

// SetPendingIcebergQty sets the pendingIcebergQty parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingIcebergQty(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingIcebergQty = &value
	return r
}

// SetPendingPrice sets the pendingPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingPrice(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingPrice = &value
	return r
}

// SetPendingQuantity sets the pendingQuantity parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingQuantity(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingQuantity = value
	return r
}

// SetPendingSide sets the pendingSide parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingSide(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingSide = value
	return r
}

// SetPendingStopPrice sets the pendingStopPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingStopPrice(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingStopPrice = &value
	return r
}

// SetPendingStrategyId sets the pendingStrategyId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingStrategyId(value int64) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingStrategyId = &value
	return r
}

// SetPendingStrategyType sets the pendingStrategyType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingStrategyType(value int64) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingStrategyType = &value
	return r
}

// SetPendingTimeInForce sets the pendingTimeInForce parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingTimeInForce(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingTimeInForce = &value
	return r
}

// SetPendingTrailingDelta sets the pendingTrailingDelta parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingTrailingDelta(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingTrailingDelta = &value
	return r
}

// SetPendingType sets the pendingType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetPendingType(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.PendingType = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetRecvWindow(value int64) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSelfTradePreventionMode sets the selfTradePreventionMode parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetSelfTradePreventionMode(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.SelfTradePreventionMode = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetSignature(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetSymbol(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetTimestamp(value int64) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}

// SetWorkingClientOrderId sets the workingClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetWorkingClientOrderId(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.WorkingClientOrderId = &value
	return r
}

// SetWorkingIcebergQty sets the workingIcebergQty parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetWorkingIcebergQty(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.WorkingIcebergQty = &value
	return r
}

// SetWorkingPrice sets the workingPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetWorkingPrice(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.WorkingPrice = value
	return r
}

// SetWorkingQuantity sets the workingQuantity parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetWorkingQuantity(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.WorkingQuantity = value
	return r
}

// SetWorkingSide sets the workingSide parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetWorkingSide(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.WorkingSide = value
	return r
}

// SetWorkingStrategyId sets the workingStrategyId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetWorkingStrategyId(value int64) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.WorkingStrategyId = &value
	return r
}

// SetWorkingStrategyType sets the workingStrategyType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetWorkingStrategyType(value int64) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.WorkingStrategyType = &value
	return r
}

// SetWorkingTimeInForce sets the workingTimeInForce parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetWorkingTimeInForce(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.WorkingTimeInForce = &value
	return r
}

// SetWorkingType sets the workingType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtoRequest) SetWorkingType(value string) *OrderListPlaceOtoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtoRequestParams{}
	}
	r.Params.WorkingType = value
	return r
}


