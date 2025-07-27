package models

import (
	"encoding/json"
)

// MyPreventedMatchesRequestParams contains the parameters for MyPreventedMatchesRequest
type MyPreventedMatchesRequestParams struct {
	FromPreventedMatchId *int64 `json:"fromPreventedMatchId,omitempty"`
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	PreventedMatchId *int64 `json:"preventedMatchId,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// MyPreventedMatchesRequest - Send a myPreventedMatches request
// Message name: Account prevented matches (USER_DATA) Request
type MyPreventedMatchesRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *MyPreventedMatchesRequestParams `json:"params,omitempty"`
}

// NewMyPreventedMatchesRequest creates a new MyPreventedMatchesRequest instance
func NewMyPreventedMatchesRequest() *MyPreventedMatchesRequest {
	return &MyPreventedMatchesRequest{}
}

// String returns string representation of MyPreventedMatchesRequest
func (s MyPreventedMatchesRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *MyPreventedMatchesRequest) SetParams(value MyPreventedMatchesRequestParams) *MyPreventedMatchesRequest {
	r.Params = &value
	return r
}

// SetFromPreventedMatchId sets the fromPreventedMatchId parameter and returns the struct for method chaining
func (r *MyPreventedMatchesRequest) SetFromPreventedMatchId(value int64) *MyPreventedMatchesRequest {
	if r.Params == nil {
		r.Params = &MyPreventedMatchesRequestParams{}
	}
	r.Params.FromPreventedMatchId = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *MyPreventedMatchesRequest) SetLimit(value int64) *MyPreventedMatchesRequest {
	if r.Params == nil {
		r.Params = &MyPreventedMatchesRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *MyPreventedMatchesRequest) SetOrderId(value int64) *MyPreventedMatchesRequest {
	if r.Params == nil {
		r.Params = &MyPreventedMatchesRequestParams{}
	}
	r.Params.OrderId = &value
	return r
}

// SetPreventedMatchId sets the preventedMatchId parameter and returns the struct for method chaining
func (r *MyPreventedMatchesRequest) SetPreventedMatchId(value int64) *MyPreventedMatchesRequest {
	if r.Params == nil {
		r.Params = &MyPreventedMatchesRequestParams{}
	}
	r.Params.PreventedMatchId = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *MyPreventedMatchesRequest) SetRecvWindow(value int64) *MyPreventedMatchesRequest {
	if r.Params == nil {
		r.Params = &MyPreventedMatchesRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *MyPreventedMatchesRequest) SetSymbol(value string) *MyPreventedMatchesRequest {
	if r.Params == nil {
		r.Params = &MyPreventedMatchesRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *MyPreventedMatchesRequest) SetTimestamp(value int64) *MyPreventedMatchesRequest {
	if r.Params == nil {
		r.Params = &MyPreventedMatchesRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


