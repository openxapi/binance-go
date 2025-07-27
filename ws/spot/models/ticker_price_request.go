package models

import (
	"encoding/json"
)

// TickerPriceRequestParams contains the parameters for TickerPriceRequest
type TickerPriceRequestParams struct {
	// Query price for a single symbol
	Symbol *string `json:"symbol,omitempty"`
	// Query price for multiple symbols
	Symbols *[]string `json:"symbols,omitempty"`
}

// TickerPriceRequest - Send a ticker.price request
// Message name: Symbol price ticker Request
type TickerPriceRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *TickerPriceRequestParams `json:"params,omitempty"`
}

// NewTickerPriceRequest creates a new TickerPriceRequest instance
func NewTickerPriceRequest() *TickerPriceRequest {
	return &TickerPriceRequest{}
}

// String returns string representation of TickerPriceRequest
func (s TickerPriceRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *TickerPriceRequest) SetParams(value TickerPriceRequestParams) *TickerPriceRequest {
	r.Params = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *TickerPriceRequest) SetSymbol(value string) *TickerPriceRequest {
	if r.Params == nil {
		r.Params = &TickerPriceRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}

// SetSymbols sets the symbols parameter and returns the struct for method chaining
func (r *TickerPriceRequest) SetSymbols(value []string) *TickerPriceRequest {
	if r.Params == nil {
		r.Params = &TickerPriceRequestParams{}
	}
	r.Params.Symbols = &value
	return r
}


