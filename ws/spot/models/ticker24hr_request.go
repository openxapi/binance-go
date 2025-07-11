package models

import (
	"encoding/json"
)

// Ticker24hrRequestParams contains the parameters for Ticker24hrRequest
type Ticker24hrRequestParams struct {
	// Query ticker for a single symbol
	Symbol *string `json:"symbol,omitempty"`
	// Query ticker for multiple symbols
	Symbols *[]string `json:"symbols,omitempty"`
	// Ticker type: FULL (default) or MINI
	Type *string `json:"type,omitempty"`
}

// Ticker24hrRequest - Send a ticker.24hr request
// Message name: 24hr ticker price change statistics Request
type Ticker24hrRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *Ticker24hrRequestParams `json:"params,omitempty"`
}

// NewTicker24hrRequest creates a new Ticker24hrRequest instance
func NewTicker24hrRequest() *Ticker24hrRequest {
	return &Ticker24hrRequest{}
}

// String returns string representation of Ticker24hrRequest
func (s Ticker24hrRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *Ticker24hrRequest) SetParams(value Ticker24hrRequestParams) *Ticker24hrRequest {
	r.Params = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *Ticker24hrRequest) SetSymbol(value string) *Ticker24hrRequest {
	if r.Params == nil {
		r.Params = &Ticker24hrRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}

// SetSymbols sets the symbols parameter and returns the struct for method chaining
func (r *Ticker24hrRequest) SetSymbols(value []string) *Ticker24hrRequest {
	if r.Params == nil {
		r.Params = &Ticker24hrRequestParams{}
	}
	r.Params.Symbols = &value
	return r
}

// SetType sets the type parameter and returns the struct for method chaining
func (r *Ticker24hrRequest) SetType(value string) *Ticker24hrRequest {
	if r.Params == nil {
		r.Params = &Ticker24hrRequestParams{}
	}
	r.Params.Type = &value
	return r
}


