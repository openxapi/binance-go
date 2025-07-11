package models

import (
	"encoding/json"
)

// TickerTradingDayRequestParams contains the parameters for TickerTradingDayRequest
type TickerTradingDayRequestParams struct {
	// Query ticker of a single symbol
	Symbol *string `json:"symbol,omitempty"`
	// Query ticker for multiple symbols
	Symbols *[]string `json:"symbols,omitempty"`
	// Default: 0 (UTC)
	TimeZone *string `json:"timeZone,omitempty"`
	// Supported values: FULL or MINI. If none provided, the default is FULL
	Type *string `json:"type,omitempty"`
}

// TickerTradingDayRequest - Send a ticker.tradingDay request
// Message name: Trading Day Ticker Request
type TickerTradingDayRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *TickerTradingDayRequestParams `json:"params,omitempty"`
}

// NewTickerTradingDayRequest creates a new TickerTradingDayRequest instance
func NewTickerTradingDayRequest() *TickerTradingDayRequest {
	return &TickerTradingDayRequest{}
}

// String returns string representation of TickerTradingDayRequest
func (s TickerTradingDayRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *TickerTradingDayRequest) SetParams(value TickerTradingDayRequestParams) *TickerTradingDayRequest {
	r.Params = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *TickerTradingDayRequest) SetSymbol(value string) *TickerTradingDayRequest {
	if r.Params == nil {
		r.Params = &TickerTradingDayRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}

// SetSymbols sets the symbols parameter and returns the struct for method chaining
func (r *TickerTradingDayRequest) SetSymbols(value []string) *TickerTradingDayRequest {
	if r.Params == nil {
		r.Params = &TickerTradingDayRequestParams{}
	}
	r.Params.Symbols = &value
	return r
}

// SetTimeZone sets the timeZone parameter and returns the struct for method chaining
func (r *TickerTradingDayRequest) SetTimeZone(value string) *TickerTradingDayRequest {
	if r.Params == nil {
		r.Params = &TickerTradingDayRequestParams{}
	}
	r.Params.TimeZone = &value
	return r
}

// SetType sets the type parameter and returns the struct for method chaining
func (r *TickerTradingDayRequest) SetType(value string) *TickerTradingDayRequest {
	if r.Params == nil {
		r.Params = &TickerTradingDayRequestParams{}
	}
	r.Params.Type = &value
	return r
}


