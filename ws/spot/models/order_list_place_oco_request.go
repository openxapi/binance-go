package models

import (
	"encoding/json"
)

// OrderListPlaceOcoRequestParams contains the parameters for OrderListPlaceOcoRequest
type OrderListPlaceOcoRequestParams struct {
	// Arbitrary unique ID among open orders for the above order. Automatically generated if not sent
	AboveClientOrderId *string `json:"aboveClientOrderId,omitempty"`
	// Note that this can only be used if aboveTimeInForce is GTC.
	AboveIcebergQty *int64 `json:"aboveIcebergQty,omitempty"`
	// Can be used if aboveType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
	AbovePrice *string `json:"abovePrice,omitempty"`
	// Can be used if aboveType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT. Either aboveStopPrice or aboveTrailingDelta or both, must be specified.
	AboveStopPrice *string `json:"aboveStopPrice,omitempty"`
	// Arbitrary numeric value identifying the above order within an order strategy.
	AboveStrategyId *int64 `json:"aboveStrategyId,omitempty"`
	// Arbitrary numeric value identifying the above order strategy. Values smaller than 1000000 are reserved and cannot be used.
	AboveStrategyType *int64 `json:"aboveStrategyType,omitempty"`
	// Required if aboveType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT.
	AboveTimeInForce *string `json:"aboveTimeInForce,omitempty"`
	// See Trailing Stop order FAQ.
	AboveTrailingDelta *int64 `json:"aboveTrailingDelta,omitempty"`
	// Supported values: STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER, TAKE_PROFIT, TAKE_PROFIT_LIMIT
	AboveType string `json:"aboveType"`
	ApiKey string `json:"apiKey"`
	BelowClientOrderId *string `json:"belowClientOrderId,omitempty"`
	// Note that this can only be used if belowTimeInForce is GTC.
	BelowIcebergQty *int64 `json:"belowIcebergQty,omitempty"`
	// Can be used if belowType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
	BelowPrice *string `json:"belowPrice,omitempty"`
	// Can be used if belowType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT. Either belowStopPrice or belowTrailingDelta or both, must be specified.
	BelowStopPrice *string `json:"belowStopPrice,omitempty"`
	// Arbitrary numeric value identifying the below order within an order strategy.
	BelowStrategyId *int64 `json:"belowStrategyId,omitempty"`
	// Arbitrary numeric value identifying the below order strategy. Values smaller than 1000000 are reserved and cannot be used.
	BelowStrategyType *int64 `json:"belowStrategyType,omitempty"`
	// Required if belowType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT
	BelowTimeInForce *string `json:"belowTimeInForce,omitempty"`
	// See Trailing Stop order FAQ.
	BelowTrailingDelta *int64 `json:"belowTrailingDelta,omitempty"`
	// Supported values: STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT,TAKE_PROFIT_LIMIT
	BelowType string `json:"belowType"`
	// Arbitrary unique ID among open order lists. Automatically generated if not sent.  A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.  listClientOrderId is distinct from the aboveClientOrderId and the belowCLientOrderId.
	ListClientOrderId *string `json:"listClientOrderId,omitempty"`
	// Select response format: ACK, RESULT, FULL
	NewOrderRespType *string `json:"newOrderRespType,omitempty"`
	// Quantity for both orders of the order list.
	Quantity string `json:"quantity"`
	// The value cannot be greater than 60000.
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	// The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	// BUY or SELL
	Side string `json:"side"`
	Signature string `json:"signature"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// OrderListPlaceOcoRequest - Send a orderList.place.oco request
// Message name: Place new Order list - OCO (TRADE) Request
type OrderListPlaceOcoRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderListPlaceOcoRequestParams `json:"params,omitempty"`
}

// NewOrderListPlaceOcoRequest creates a new OrderListPlaceOcoRequest instance
func NewOrderListPlaceOcoRequest() *OrderListPlaceOcoRequest {
	return &OrderListPlaceOcoRequest{}
}

// String returns string representation of OrderListPlaceOcoRequest
func (s OrderListPlaceOcoRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetParams(value OrderListPlaceOcoRequestParams) *OrderListPlaceOcoRequest {
	r.Params = &value
	return r
}

// SetAboveClientOrderId sets the aboveClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetAboveClientOrderId(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.AboveClientOrderId = &value
	return r
}

// SetAboveIcebergQty sets the aboveIcebergQty parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetAboveIcebergQty(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.AboveIcebergQty = &value
	return r
}

// SetAbovePrice sets the abovePrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetAbovePrice(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.AbovePrice = &value
	return r
}

// SetAboveStopPrice sets the aboveStopPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetAboveStopPrice(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.AboveStopPrice = &value
	return r
}

// SetAboveStrategyId sets the aboveStrategyId parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetAboveStrategyId(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.AboveStrategyId = &value
	return r
}

// SetAboveStrategyType sets the aboveStrategyType parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetAboveStrategyType(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.AboveStrategyType = &value
	return r
}

// SetAboveTimeInForce sets the aboveTimeInForce parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetAboveTimeInForce(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.AboveTimeInForce = &value
	return r
}

// SetAboveTrailingDelta sets the aboveTrailingDelta parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetAboveTrailingDelta(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.AboveTrailingDelta = &value
	return r
}

// SetAboveType sets the aboveType parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetAboveType(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.AboveType = value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetApiKey(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetBelowClientOrderId sets the belowClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetBelowClientOrderId(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.BelowClientOrderId = &value
	return r
}

// SetBelowIcebergQty sets the belowIcebergQty parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetBelowIcebergQty(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.BelowIcebergQty = &value
	return r
}

// SetBelowPrice sets the belowPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetBelowPrice(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.BelowPrice = &value
	return r
}

// SetBelowStopPrice sets the belowStopPrice parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetBelowStopPrice(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.BelowStopPrice = &value
	return r
}

// SetBelowStrategyId sets the belowStrategyId parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetBelowStrategyId(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.BelowStrategyId = &value
	return r
}

// SetBelowStrategyType sets the belowStrategyType parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetBelowStrategyType(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.BelowStrategyType = &value
	return r
}

// SetBelowTimeInForce sets the belowTimeInForce parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetBelowTimeInForce(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.BelowTimeInForce = &value
	return r
}

// SetBelowTrailingDelta sets the belowTrailingDelta parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetBelowTrailingDelta(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.BelowTrailingDelta = &value
	return r
}

// SetBelowType sets the belowType parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetBelowType(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.BelowType = value
	return r
}

// SetListClientOrderId sets the listClientOrderId parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetListClientOrderId(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.ListClientOrderId = &value
	return r
}

// SetNewOrderRespType sets the newOrderRespType parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetNewOrderRespType(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.NewOrderRespType = &value
	return r
}

// SetQuantity sets the quantity parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetQuantity(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.Quantity = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetRecvWindow(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSelfTradePreventionMode sets the selfTradePreventionMode parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetSelfTradePreventionMode(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.SelfTradePreventionMode = &value
	return r
}

// SetSide sets the side parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetSide(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.Side = value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetSignature(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetSymbol(value string) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderListPlaceOcoRequest) SetTimestamp(value int64) *OrderListPlaceOcoRequest {
	if r.Params == nil {
		r.Params = &OrderListPlaceOcoRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


