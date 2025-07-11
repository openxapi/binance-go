package models

import (
	"encoding/json"
)

// AccountCommissionRequestParams contains the parameters for AccountCommissionRequest
type AccountCommissionRequestParams struct {
	Symbol string `json:"symbol"`
}

// AccountCommissionRequest - Send a account.commission request
// Message name: Account Commission Rates (USER_DATA) Request
type AccountCommissionRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *AccountCommissionRequestParams `json:"params,omitempty"`
}

// NewAccountCommissionRequest creates a new AccountCommissionRequest instance
func NewAccountCommissionRequest() *AccountCommissionRequest {
	return &AccountCommissionRequest{}
}

// String returns string representation of AccountCommissionRequest
func (s AccountCommissionRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *AccountCommissionRequest) SetParams(value AccountCommissionRequestParams) *AccountCommissionRequest {
	r.Params = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *AccountCommissionRequest) SetSymbol(value string) *AccountCommissionRequest {
	if r.Params == nil {
		r.Params = &AccountCommissionRequestParams{}
	}
	r.Params.Symbol = value
	return r
}


