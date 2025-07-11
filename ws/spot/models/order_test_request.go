package models

import (
	"encoding/json"
)

// OrderTestRequestParams contains the parameters for OrderTestRequest
type OrderTestRequestParams struct {
	// apiKey property
	ApiKey *string `json:"apiKey,omitempty"`
	// price property
	Price *string `json:"price,omitempty"`
	// quantity property
	Quantity *string `json:"quantity,omitempty"`
	// side property
	Side *string `json:"side,omitempty"`
	// signature property
	Signature *string `json:"signature,omitempty"`
	// symbol property
	Symbol *string `json:"symbol,omitempty"`
	// timeInForce property
	TimeInForce *string `json:"timeInForce,omitempty"`
	// timestamp property
	Timestamp *int64 `json:"timestamp,omitempty"`
	// type property
	Type *string `json:"type,omitempty"`
}

// OrderTestRequest - Send a order.test request
// Message name: Test new order (TRADE) Request
type OrderTestRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderTestRequestParams `json:"params,omitempty"`
}

// NewOrderTestRequest creates a new OrderTestRequest instance
func NewOrderTestRequest() *OrderTestRequest {
	return &OrderTestRequest{}
}

// String returns string representation of OrderTestRequest
func (s OrderTestRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderTestRequest) SetParams(value OrderTestRequestParams) *OrderTestRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *OrderTestRequest) SetApiKey(value string) *OrderTestRequest {
	if r.Params == nil {
		r.Params = &OrderTestRequestParams{}
	}
	r.Params.ApiKey = &value
	return r
}

// SetPrice sets the price parameter and returns the struct for method chaining
func (r *OrderTestRequest) SetPrice(value string) *OrderTestRequest {
	if r.Params == nil {
		r.Params = &OrderTestRequestParams{}
	}
	r.Params.Price = &value
	return r
}

// SetQuantity sets the quantity parameter and returns the struct for method chaining
func (r *OrderTestRequest) SetQuantity(value string) *OrderTestRequest {
	if r.Params == nil {
		r.Params = &OrderTestRequestParams{}
	}
	r.Params.Quantity = &value
	return r
}

// SetSide sets the side parameter and returns the struct for method chaining
func (r *OrderTestRequest) SetSide(value string) *OrderTestRequest {
	if r.Params == nil {
		r.Params = &OrderTestRequestParams{}
	}
	r.Params.Side = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *OrderTestRequest) SetSignature(value string) *OrderTestRequest {
	if r.Params == nil {
		r.Params = &OrderTestRequestParams{}
	}
	r.Params.Signature = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderTestRequest) SetSymbol(value string) *OrderTestRequest {
	if r.Params == nil {
		r.Params = &OrderTestRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}

// SetTimeInForce sets the timeInForce parameter and returns the struct for method chaining
func (r *OrderTestRequest) SetTimeInForce(value string) *OrderTestRequest {
	if r.Params == nil {
		r.Params = &OrderTestRequestParams{}
	}
	r.Params.TimeInForce = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderTestRequest) SetTimestamp(value int64) *OrderTestRequest {
	if r.Params == nil {
		r.Params = &OrderTestRequestParams{}
	}
	r.Params.Timestamp = &value
	return r
}

// SetType sets the type parameter and returns the struct for method chaining
func (r *OrderTestRequest) SetType(value string) *OrderTestRequest {
	if r.Params == nil {
		r.Params = &OrderTestRequestParams{}
	}
	r.Params.Type = &value
	return r
}


