package models

import (
	"encoding/json"
)

// AllOrderListsRequestParams contains the parameters for AllOrderListsRequest
type AllOrderListsRequestParams struct {
	ApiKey string `json:"apiKey"`
	EndTime *int64 `json:"endTime,omitempty"`
	// Order list ID to begin at
	FromId *int64 `json:"fromId,omitempty"`
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	StartTime *int64 `json:"startTime,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

// AllOrderListsRequest - Send a allOrderLists request
// Message name: Account order list history (USER_DATA) Request
type AllOrderListsRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *AllOrderListsRequestParams `json:"params,omitempty"`
}

// NewAllOrderListsRequest creates a new AllOrderListsRequest instance
func NewAllOrderListsRequest() *AllOrderListsRequest {
	return &AllOrderListsRequest{}
}

// String returns string representation of AllOrderListsRequest
func (s AllOrderListsRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *AllOrderListsRequest) SetParams(value AllOrderListsRequestParams) *AllOrderListsRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *AllOrderListsRequest) SetApiKey(value string) *AllOrderListsRequest {
	if r.Params == nil {
		r.Params = &AllOrderListsRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetEndTime sets the endTime parameter and returns the struct for method chaining
func (r *AllOrderListsRequest) SetEndTime(value int64) *AllOrderListsRequest {
	if r.Params == nil {
		r.Params = &AllOrderListsRequestParams{}
	}
	r.Params.EndTime = &value
	return r
}

// SetFromId sets the fromId parameter and returns the struct for method chaining
func (r *AllOrderListsRequest) SetFromId(value int64) *AllOrderListsRequest {
	if r.Params == nil {
		r.Params = &AllOrderListsRequestParams{}
	}
	r.Params.FromId = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *AllOrderListsRequest) SetLimit(value int64) *AllOrderListsRequest {
	if r.Params == nil {
		r.Params = &AllOrderListsRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *AllOrderListsRequest) SetRecvWindow(value int64) *AllOrderListsRequest {
	if r.Params == nil {
		r.Params = &AllOrderListsRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *AllOrderListsRequest) SetSignature(value string) *AllOrderListsRequest {
	if r.Params == nil {
		r.Params = &AllOrderListsRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetStartTime sets the startTime parameter and returns the struct for method chaining
func (r *AllOrderListsRequest) SetStartTime(value int64) *AllOrderListsRequest {
	if r.Params == nil {
		r.Params = &AllOrderListsRequestParams{}
	}
	r.Params.StartTime = &value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *AllOrderListsRequest) SetTimestamp(value int64) *AllOrderListsRequest {
	if r.Params == nil {
		r.Params = &AllOrderListsRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


