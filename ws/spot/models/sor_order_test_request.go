package models

import (
	"encoding/json"
)

// SorOrderTestRequestParams contains the parameters for SorOrderTestRequest
type SorOrderTestRequestParams struct {
	// apiKey property
	ApiKey *string `json:"apiKey,omitempty"`
	// price property
	Price *float64 `json:"price,omitempty"`
	// quantity property
	Quantity *float64 `json:"quantity,omitempty"`
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

// SorOrderTestRequest - Send a sor.order.test request
// Message name: Test new order using SOR (TRADE) Request
type SorOrderTestRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *SorOrderTestRequestParams `json:"params,omitempty"`
}

// NewSorOrderTestRequest creates a new SorOrderTestRequest instance
func NewSorOrderTestRequest() *SorOrderTestRequest {
	return &SorOrderTestRequest{}
}

// String returns string representation of SorOrderTestRequest
func (s SorOrderTestRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *SorOrderTestRequest) SetParams(value SorOrderTestRequestParams) *SorOrderTestRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *SorOrderTestRequest) SetApiKey(value string) *SorOrderTestRequest {
	if r.Params == nil {
		r.Params = &SorOrderTestRequestParams{}
	}
	r.Params.ApiKey = &value
	return r
}

// SetPrice sets the price parameter and returns the struct for method chaining
func (r *SorOrderTestRequest) SetPrice(value float64) *SorOrderTestRequest {
	if r.Params == nil {
		r.Params = &SorOrderTestRequestParams{}
	}
	r.Params.Price = &value
	return r
}

// SetQuantity sets the quantity parameter and returns the struct for method chaining
func (r *SorOrderTestRequest) SetQuantity(value float64) *SorOrderTestRequest {
	if r.Params == nil {
		r.Params = &SorOrderTestRequestParams{}
	}
	r.Params.Quantity = &value
	return r
}

// SetSide sets the side parameter and returns the struct for method chaining
func (r *SorOrderTestRequest) SetSide(value string) *SorOrderTestRequest {
	if r.Params == nil {
		r.Params = &SorOrderTestRequestParams{}
	}
	r.Params.Side = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *SorOrderTestRequest) SetSignature(value string) *SorOrderTestRequest {
	if r.Params == nil {
		r.Params = &SorOrderTestRequestParams{}
	}
	r.Params.Signature = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *SorOrderTestRequest) SetSymbol(value string) *SorOrderTestRequest {
	if r.Params == nil {
		r.Params = &SorOrderTestRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}

// SetTimeInForce sets the timeInForce parameter and returns the struct for method chaining
func (r *SorOrderTestRequest) SetTimeInForce(value string) *SorOrderTestRequest {
	if r.Params == nil {
		r.Params = &SorOrderTestRequestParams{}
	}
	r.Params.TimeInForce = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *SorOrderTestRequest) SetTimestamp(value int64) *SorOrderTestRequest {
	if r.Params == nil {
		r.Params = &SorOrderTestRequestParams{}
	}
	r.Params.Timestamp = &value
	return r
}

// SetType sets the type parameter and returns the struct for method chaining
func (r *SorOrderTestRequest) SetType(value string) *SorOrderTestRequest {
	if r.Params == nil {
		r.Params = &SorOrderTestRequestParams{}
	}
	r.Params.Type = &value
	return r
}


