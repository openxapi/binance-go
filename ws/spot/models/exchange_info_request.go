package models

import (
	"encoding/json"
)

// ExchangeInfoRequestParams contains the parameters for ExchangeInfoRequest
type ExchangeInfoRequestParams struct {
	// Filter symbols by permissions
	Permissions *[]string `json:"permissions,omitempty"`
	// Controls whether the content of the permissionSets field is populated or not. Defaults to true.
	ShowPermissionSets *bool `json:"showPermissionSets,omitempty"`
	// Describe a single symbol
	Symbol *string `json:"symbol,omitempty"`
	// Filters symbols that have this tradingStatus. Valid values: TRADING, HALT, BREAK  Cannot be used in combination with symbol or symbols
	SymbolStatus *string `json:"symbolStatus,omitempty"`
	// Describe multiple symbols
	Symbols *[]string `json:"symbols,omitempty"`
}

// ExchangeInfoRequest - Send a exchangeInfo request
// Message name: Exchange information Request
type ExchangeInfoRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *ExchangeInfoRequestParams `json:"params,omitempty"`
}

// NewExchangeInfoRequest creates a new ExchangeInfoRequest instance
func NewExchangeInfoRequest() *ExchangeInfoRequest {
	return &ExchangeInfoRequest{}
}

// String returns string representation of ExchangeInfoRequest
func (s ExchangeInfoRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *ExchangeInfoRequest) SetParams(value ExchangeInfoRequestParams) *ExchangeInfoRequest {
	r.Params = &value
	return r
}

// SetPermissions sets the permissions parameter and returns the struct for method chaining
func (r *ExchangeInfoRequest) SetPermissions(value []string) *ExchangeInfoRequest {
	if r.Params == nil {
		r.Params = &ExchangeInfoRequestParams{}
	}
	r.Params.Permissions = &value
	return r
}

// SetShowPermissionSets sets the showPermissionSets parameter and returns the struct for method chaining
func (r *ExchangeInfoRequest) SetShowPermissionSets(value bool) *ExchangeInfoRequest {
	if r.Params == nil {
		r.Params = &ExchangeInfoRequestParams{}
	}
	r.Params.ShowPermissionSets = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *ExchangeInfoRequest) SetSymbol(value string) *ExchangeInfoRequest {
	if r.Params == nil {
		r.Params = &ExchangeInfoRequestParams{}
	}
	r.Params.Symbol = &value
	return r
}

// SetSymbolStatus sets the symbolStatus parameter and returns the struct for method chaining
func (r *ExchangeInfoRequest) SetSymbolStatus(value string) *ExchangeInfoRequest {
	if r.Params == nil {
		r.Params = &ExchangeInfoRequestParams{}
	}
	r.Params.SymbolStatus = &value
	return r
}

// SetSymbols sets the symbols parameter and returns the struct for method chaining
func (r *ExchangeInfoRequest) SetSymbols(value []string) *ExchangeInfoRequest {
	if r.Params == nil {
		r.Params = &ExchangeInfoRequestParams{}
	}
	r.Params.Symbols = &value
	return r
}


