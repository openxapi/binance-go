package models

import (
	"encoding/json"
)

// TickerRequestParams contains the parameters for TickerRequest
type TickerRequestParams struct {
	// Query ticker of a single symbol
	Symbol *string `json:"symbol,omitempty"`
	// Query ticker for multiple symbols
	Symbols *[]string `json:"symbols,omitempty"`
	// Ticker type: FULL (default) or MINI
	Type *string `json:"type,omitempty"`
	// Default 1d
	WindowSize *string `json:"windowSize,omitempty"`
}

// TickerRequest - Send a ticker request
// Message name: Rolling window price change statistics Request
type TickerRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *TickerRequestParams `json:"params,omitempty"`
}

// NewTickerRequest creates a new TickerRequest instance
func NewTickerRequest() *TickerRequest {
	return &TickerRequest{}
}

// String returns string representation of TickerRequest
func (s TickerRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *TickerRequest) SetParams(value TickerRequestParams) *TickerRequest {
	r.Params = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *TickerRequest) SetSymbol(value string) *TickerRequest {
	if r.Params == nil {
		r.Params = &TickerRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}

// SetSymbols sets the symbols parameter and returns the struct for method chaining
func (r *TickerRequest) SetSymbols(value []string) *TickerRequest {
	if r.Params == nil {
		r.Params = &TickerRequestParams{}
	}
	r.Params.Symbols = &value
	return r
}

// SetType sets the type parameter and returns the struct for method chaining
func (r *TickerRequest) SetType(value string) *TickerRequest {
	if r.Params == nil {
		r.Params = &TickerRequestParams{}
	}
	r.Params.Type = &value
	return r
}

// SetWindowSize sets the windowSize parameter and returns the struct for method chaining
func (r *TickerRequest) SetWindowSize(value string) *TickerRequest {
	if r.Params == nil {
		r.Params = &TickerRequestParams{}
	}
	r.Params.WindowSize = &value
	return r
}


