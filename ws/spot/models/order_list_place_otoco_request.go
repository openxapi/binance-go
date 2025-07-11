package models

import (
	"encoding/json"
)

// OrderListPlaceOtocoRequestParams contains the parameters for OrderListPlaceOtocoRequest
type OrderListPlaceOtocoRequestParams struct {
	// Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.  listClientOrderId is distinct from the workingClientOrderId, pendingAboveClientOrderId, and the pendingBelowClientOrderId.
	ListClientOrderId *string `json:"listClientOrderId,omitempty"`
	// Format of the JSON response. Supported values: Order Response Type
	NewOrderRespType *string `json:"newOrderRespType,omitempty"`
	// Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
	PendingAboveClientOrderId *string `json:"pendingAboveClientOrderId,omitempty"`
	// This can only be used if pendingAboveTimeInForce is GTC or if pendingAboveType is LIMIT_MAKER.
	PendingAboveIcebergQty *string `json:"pendingAboveIcebergQty,omitempty"`
	// Can be used if pendingAboveType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
	PendingAbovePrice *string `json:"pendingAbovePrice,omitempty"`
	// Can be used if pendingAboveType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT
	PendingAboveStopPrice *string `json:"pendingAboveStopPrice,omitempty"`
	// Arbitrary numeric value identifying the pending above order within an order strategy.
	PendingAboveStrategyId *int64 `json:"pendingAboveStrategyId,omitempty"`
	// Arbitrary numeric value identifying the pending above order strategy.  Values smaller than 1000000 are reserved and cannot be used.
	PendingAboveStrategyType *int64 `json:"pendingAboveStrategyType,omitempty"`
	PendingAboveTimeInForce *string `json:"pendingAboveTimeInForce,omitempty"`
	// See Trailing Stop FAQ
	PendingAboveTrailingDelta *string `json:"pendingAboveTrailingDelta,omitempty"`
	// Supported values: STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER, TAKE_PROFIT, TAKE_PROFIT_LIMIT
	PendingAboveType string `json:"pendingAboveType"`
	// Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
	PendingBelowClientOrderId *string `json:"pendingBelowClientOrderId,omitempty"`
	// This can only be used if pendingBelowTimeInForce is GTC, or if pendingBelowType is LIMIT_MAKER.
	PendingBelowIcebergQty *string `json:"pendingBelowIcebergQty,omitempty"`
	// Can be used if pendingBelowType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT to specify the limit price.
	PendingBelowPrice *string `json:"pendingBelowPrice,omitempty"`
	// Can be used if pendingBelowType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT. Either pendingBelowStopPrice or pendingBelowTrailingDelta or both, must be specified.
	PendingBelowStopPrice *string `json:"pendingBelowStopPrice,omitempty"`
	// Arbitrary numeric value identifying the pending below order within an order strategy.
	PendingBelowStrategyId *int64 `json:"pendingBelowStrategyId,omitempty"`
	// Arbitrary numeric value identifying the pending below order strategy.  Values smaller than 1000000 are reserved and cannot be used.
	PendingBelowStrategyType *int64 `json:"pendingBelowStrategyType,omitempty"`
	// Supported values: Time In Force
	PendingBelowTimeInForce *string `json:"pendingBelowTimeInForce,omitempty"`
	PendingBelowTrailingDelta *string `json:"pendingBelowTrailingDelta,omitempty"`
	// Supported values: STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT,TAKE_PROFIT_LIMIT
	PendingBelowType *string `json:"pendingBelowType,omitempty"`
	PendingQuantity string `json:"pendingQuantity"`
	// Supported values: Order Side
	PendingSide string `json:"pendingSide"`
	// The value cannot be greater than 60000.
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	// The allowed values are dependent on what is configured on the symbol. Supported values: STP Modes
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Signature string `json:"signature"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
	// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
	WorkingClientOrderId *string `json:"workingClientOrderId,omitempty"`
	// This can only be used if workingTimeInForce is GTC.
	WorkingIcebergQty *string `json:"workingIcebergQty,omitempty"`
	WorkingPrice string `json:"workingPrice"`
	WorkingQuantity string `json:"workingQuantity"`
	// Supported values: Order Side
	WorkingSide string `json:"workingSide"`
	// Arbitrary numeric value identifying the working order within an order strategy.
	WorkingStrategyId *int64 `json:"workingStrategyId,omitempty"`
	// Arbitrary numeric value identifying the working order strategy.  Values smaller than 1000000 are reserved and cannot be used.
	WorkingStrategyType *int64 `json:"workingStrategyType,omitempty"`
	// Supported values: Time In Force
	WorkingTimeInForce *string `json:"workingTimeInForce,omitempty"`
	// Supported values: LIMIT, LIMIT_MAKER
	WorkingType string `json:"workingType"`
}

// OrderListPlaceOtocoRequest - Send a orderList.place.otoco request
// Message name: Place new Order list - OTOCO (TRADE) Request
type OrderListPlaceOtocoRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderListPlaceOtocoRequestParams `json:"params,omitempty"`
}

// NewOrderListPlaceOtocoRequest creates a new OrderListPlaceOtocoRequest instance
func NewOrderListPlaceOtocoRequest() *OrderListPlaceOtocoRequest {
	return &OrderListPlaceOtocoRequest{}
}

// String returns string representation of OrderListPlaceOtocoRequest
func (s OrderListPlaceOtocoRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetParams(value OrderListPlaceOtocoRequestParams) *OrderListPlaceOtocoRequest {
	r.Params = &value
	return r
}

// SetListClientOrderId sets the listClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetListClientOrderId(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.ListClientOrderId = &value
	return r
}

// SetNewOrderRespType sets the newOrderRespType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetNewOrderRespType(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.NewOrderRespType = &value
	return r
}

// SetPendingAboveClientOrderId sets the pendingAboveClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingAboveClientOrderId(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingAboveClientOrderId = &value
	return r
}

// SetPendingAboveIcebergQty sets the pendingAboveIcebergQty parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingAboveIcebergQty(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingAboveIcebergQty = &value
	return r
}

// SetPendingAbovePrice sets the pendingAbovePrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingAbovePrice(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingAbovePrice = &value
	return r
}

// SetPendingAboveStopPrice sets the pendingAboveStopPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingAboveStopPrice(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingAboveStopPrice = &value
	return r
}

// SetPendingAboveStrategyId sets the pendingAboveStrategyId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingAboveStrategyId(value int64) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingAboveStrategyId = &value
	return r
}

// SetPendingAboveStrategyType sets the pendingAboveStrategyType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingAboveStrategyType(value int64) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingAboveStrategyType = &value
	return r
}

// SetPendingAboveTimeInForce sets the pendingAboveTimeInForce parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingAboveTimeInForce(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingAboveTimeInForce = &value
	return r
}

// SetPendingAboveTrailingDelta sets the pendingAboveTrailingDelta parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingAboveTrailingDelta(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingAboveTrailingDelta = &value
	return r
}

// SetPendingAboveType sets the pendingAboveType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingAboveType(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingAboveType = value
	return r
}

// SetPendingBelowClientOrderId sets the pendingBelowClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingBelowClientOrderId(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingBelowClientOrderId = &value
	return r
}

// SetPendingBelowIcebergQty sets the pendingBelowIcebergQty parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingBelowIcebergQty(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingBelowIcebergQty = &value
	return r
}

// SetPendingBelowPrice sets the pendingBelowPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingBelowPrice(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingBelowPrice = &value
	return r
}

// SetPendingBelowStopPrice sets the pendingBelowStopPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingBelowStopPrice(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingBelowStopPrice = &value
	return r
}

// SetPendingBelowStrategyId sets the pendingBelowStrategyId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingBelowStrategyId(value int64) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingBelowStrategyId = &value
	return r
}

// SetPendingBelowStrategyType sets the pendingBelowStrategyType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingBelowStrategyType(value int64) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingBelowStrategyType = &value
	return r
}

// SetPendingBelowTimeInForce sets the pendingBelowTimeInForce parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingBelowTimeInForce(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingBelowTimeInForce = &value
	return r
}

// SetPendingBelowTrailingDelta sets the pendingBelowTrailingDelta parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingBelowTrailingDelta(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingBelowTrailingDelta = &value
	return r
}

// SetPendingBelowType sets the pendingBelowType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingBelowType(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingBelowType = &value
	return r
}

// SetPendingQuantity sets the pendingQuantity parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingQuantity(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingQuantity = value
	return r
}

// SetPendingSide sets the pendingSide parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetPendingSide(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.PendingSide = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetRecvWindow(value int64) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSelfTradePreventionMode sets the selfTradePreventionMode parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetSelfTradePreventionMode(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.SelfTradePreventionMode = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetSignature(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetSymbol(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetTimestamp(value int64) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}

// SetWorkingClientOrderId sets the workingClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetWorkingClientOrderId(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.WorkingClientOrderId = &value
	return r
}

// SetWorkingIcebergQty sets the workingIcebergQty parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetWorkingIcebergQty(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.WorkingIcebergQty = &value
	return r
}

// SetWorkingPrice sets the workingPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetWorkingPrice(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.WorkingPrice = value
	return r
}

// SetWorkingQuantity sets the workingQuantity parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetWorkingQuantity(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.WorkingQuantity = value
	return r
}

// SetWorkingSide sets the workingSide parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetWorkingSide(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.WorkingSide = value
	return r
}

// SetWorkingStrategyId sets the workingStrategyId parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetWorkingStrategyId(value int64) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.WorkingStrategyId = &value
	return r
}

// SetWorkingStrategyType sets the workingStrategyType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetWorkingStrategyType(value int64) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.WorkingStrategyType = &value
	return r
}

// SetWorkingTimeInForce sets the workingTimeInForce parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetWorkingTimeInForce(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.WorkingTimeInForce = &value
	return r
}

// SetWorkingType sets the workingType parameter and returns the struct for method chaining
func (r *OrderListPlaceOtocoRequest) SetWorkingType(value string) *OrderListPlaceOtocoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOtocoRequestParams{}
	}
	r.Params.WorkingType = value
	return r
}


