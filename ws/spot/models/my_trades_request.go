package models

import (
	"encoding/json"
)

// MyTradesRequestParams contains the parameters for MyTradesRequest
type MyTradesRequestParams struct {
	ApiKey string `json:"apiKey"`
	EndTime *int64 `json:"endTime,omitempty"`
	// First trade ID to query
	FromId *int64 `json:"fromId,omitempty"`
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	StartTime *int64 `json:"startTime,omitempty"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// MyTradesRequest - Send a myTrades request
// Message name: Account trade history (USER_DATA) Request
type MyTradesRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *MyTradesRequestParams `json:"params,omitempty"`
}

// NewMyTradesRequest creates a new MyTradesRequest instance
func NewMyTradesRequest() *MyTradesRequest {
	return &MyTradesRequest{}
}

// String returns string representation of MyTradesRequest
func (s MyTradesRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *MyTradesRequest) SetParams(value MyTradesRequestParams) *MyTradesRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetApiKey(value string) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetEndTime sets the endTime parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetEndTime(value int64) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.EndTime = &value
	return r
}

// SetFromId sets the fromId parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetFromId(value int64) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.FromId = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetLimit(value int64) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetOrderId(value int64) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.OrderId = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetRecvWindow(value int64) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetSignature(value string) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetStartTime sets the startTime parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetStartTime(value int64) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.StartTime = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetSymbol(value string) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *MyTradesRequest) SetTimestamp(value int64) *MyTradesRequest {
	if r.Params == nil {
		r.Params = &MyTradesRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


