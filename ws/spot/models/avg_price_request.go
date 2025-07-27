package models

import (
	"encoding/json"
)

// AvgPriceRequestParams contains the parameters for AvgPriceRequest
type AvgPriceRequestParams struct {
	Symbol string `json:"symbol"`
}

// AvgPriceRequest - Send a avgPrice request
// Message name: Current average price Request
type AvgPriceRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *AvgPriceRequestParams `json:"params,omitempty"`
}

// NewAvgPriceRequest creates a new AvgPriceRequest instance
func NewAvgPriceRequest() *AvgPriceRequest {
	return &AvgPriceRequest{}
}

// String returns string representation of AvgPriceRequest
func (s AvgPriceRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *AvgPriceRequest) SetParams(value AvgPriceRequestParams) *AvgPriceRequest {
	r.Params = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *AvgPriceRequest) SetSymbol(value string) *AvgPriceRequest {
	if r.Params == nil {
		r.Params = &AvgPriceRequestParams{}
	}
	r.Params.Symbol = value
	return r
}


