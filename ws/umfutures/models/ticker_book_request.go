package models

import (
	"encoding/json"
)

// TickerBookRequestParams contains the parameters for TickerBookRequest
type TickerBookRequestParams struct {
	Symbol *string `json:"symbol,omitempty"`
}

// TickerBookRequest - Send a ticker.book request
// Message name: Symbol Order Book Ticker Request
type TickerBookRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *TickerBookRequestParams `json:"params,omitempty"`
}

// NewTickerBookRequest creates a new TickerBookRequest instance
func NewTickerBookRequest() *TickerBookRequest {
	return &TickerBookRequest{}
}

// String returns string representation of TickerBookRequest
func (s TickerBookRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *TickerBookRequest) SetParams(value TickerBookRequestParams) *TickerBookRequest {
	r.Params = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *TickerBookRequest) SetSymbol(value string) *TickerBookRequest {
	if r.Params == nil {
		r.Params = &TickerBookRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}


