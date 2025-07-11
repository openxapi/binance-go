package models

import (
	"encoding/json"
)

// AllOrdersRequestParams contains the parameters for AllOrdersRequest
type AllOrdersRequestParams struct {
	ApiKey string `json:"apiKey"`
	EndTime *int64 `json:"endTime,omitempty"`
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	// Order ID to begin at
	OrderId *int64 `json:"orderId,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	StartTime *int64 `json:"startTime,omitempty"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// AllOrdersRequest - Send a allOrders request
// Message name: Account order history (USER_DATA) Request
type AllOrdersRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *AllOrdersRequestParams `json:"params,omitempty"`
}

// NewAllOrdersRequest creates a new AllOrdersRequest instance
func NewAllOrdersRequest() *AllOrdersRequest {
	return &AllOrdersRequest{}
}

// String returns string representation of AllOrdersRequest
func (s AllOrdersRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *AllOrdersRequest) SetParams(value AllOrdersRequestParams) *AllOrdersRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *AllOrdersRequest) SetApiKey(value string) *AllOrdersRequest {
	if r.Params == nil {
		r.Params = &AllOrdersRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetEndTime sets the endTime parameter and returns the struct for method chaining
func (r *AllOrdersRequest) SetEndTime(value int64) *AllOrdersRequest {
	if r.Params == nil {
		r.Params = &AllOrdersRequestParams{}
	}
	r.Params.EndTime = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *AllOrdersRequest) SetLimit(value int64) *AllOrdersRequest {
	if r.Params == nil {
		r.Params = &AllOrdersRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *AllOrdersRequest) SetOrderId(value int64) *AllOrdersRequest {
	if r.Params == nil {
		r.Params = &AllOrdersRequestParams{}
	}
	r.Params.OrderId = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *AllOrdersRequest) SetRecvWindow(value int64) *AllOrdersRequest {
	if r.Params == nil {
		r.Params = &AllOrdersRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *AllOrdersRequest) SetSignature(value string) *AllOrdersRequest {
	if r.Params == nil {
		r.Params = &AllOrdersRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetStartTime sets the startTime parameter and returns the struct for method chaining
func (r *AllOrdersRequest) SetStartTime(value int64) *AllOrdersRequest {
	if r.Params == nil {
		r.Params = &AllOrdersRequestParams{}
	}
	r.Params.StartTime = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *AllOrdersRequest) SetSymbol(value string) *AllOrdersRequest {
	if r.Params == nil {
		r.Params = &AllOrdersRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *AllOrdersRequest) SetTimestamp(value int64) *AllOrdersRequest {
	if r.Params == nil {
		r.Params = &AllOrdersRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


