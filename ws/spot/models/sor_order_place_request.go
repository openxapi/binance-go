package models

import (
	"encoding/json"
)

// SorOrderPlaceRequestParams contains the parameters for SorOrderPlaceRequest
type SorOrderPlaceRequestParams struct {
	ApiKey string `json:"apiKey"`
	IcebergQty *string `json:"icebergQty,omitempty"`
	// Arbitrary unique ID among open orders. Automatically generated if not sent
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	// Select response format: ACK, RESULT, FULL.MARKET and LIMIT orders use FULL by default.
	NewOrderRespType *string `json:"newOrderRespType,omitempty"`
	// Applicable only to LIMIT order type
	Price *string `json:"price,omitempty"`
	Quantity string `json:"quantity"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	// The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	// BUY or SELL
	Side string `json:"side"`
	Signature string `json:"signature"`
	// Arbitrary numeric value identifying the order within an order strategy.
	StrategyId *int64 `json:"strategyId,omitempty"`
	// Arbitrary numeric value identifying the order strategy.Values smaller than 1000000 are reserved and cannot be used.
	StrategyType *int64 `json:"strategyType,omitempty"`
	Symbol string `json:"symbol"`
	// Applicable only to LIMIT order type
	TimeInForce *string `json:"timeInForce,omitempty"`
	Timestamp int64 `json:"timestamp"`
	Type string `json:"type"`
}

// SorOrderPlaceRequest - Send a sor.order.place request
// Message name: Place new order using SOR (TRADE) Request
type SorOrderPlaceRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *SorOrderPlaceRequestParams `json:"params,omitempty"`
}

// NewSorOrderPlaceRequest creates a new SorOrderPlaceRequest instance
func NewSorOrderPlaceRequest() *SorOrderPlaceRequest {
	return &SorOrderPlaceRequest{}
}

// String returns string representation of SorOrderPlaceRequest
func (s SorOrderPlaceRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetParams(value SorOrderPlaceRequestParams) *SorOrderPlaceRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetApiKey(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetIcebergQty sets the icebergQty parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetIcebergQty(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.IcebergQty = &value
	return r
}

// SetNewClientOrderId sets the newClientOrderId parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetNewClientOrderId(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.NewClientOrderId = &value
	return r
}

// SetNewOrderRespType sets the newOrderRespType parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetNewOrderRespType(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.NewOrderRespType = &value
	return r
}

// SetPrice sets the price parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetPrice(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.Price = &value
	return r
}

// SetQuantity sets the quantity parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetQuantity(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.Quantity = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetRecvWindow(value int64) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSelfTradePreventionMode sets the selfTradePreventionMode parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetSelfTradePreventionMode(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.SelfTradePreventionMode = &value
	return r
}

// SetSide sets the side parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetSide(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.Side = value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetSignature(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetStrategyId sets the strategyId parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetStrategyId(value int64) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.StrategyId = &value
	return r
}

// SetStrategyType sets the strategyType parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetStrategyType(value int64) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.StrategyType = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetSymbol(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimeInForce sets the timeInForce parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetTimeInForce(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.TimeInForce = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetTimestamp(value int64) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}

// SetType sets the type parameter and returns the struct for method chaining
func (r *SorOrderPlaceRequest) SetType(value string) *SorOrderPlaceRequest {
	if r.Params == nil {
		r.Params = &SorOrderPlaceRequestParams{}
	}
	r.Params.Type = value
	return r
}


