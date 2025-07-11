package models

import (
	"encoding/json"
)

// MyAllocationsRequestParams contains the parameters for MyAllocationsRequest
type MyAllocationsRequestParams struct {
	EndTime *int64 `json:"endTime,omitempty"`
	FromAllocationId *int64 `json:"fromAllocationId,omitempty"`
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	Symbol string `json:"symbol"`
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// MyAllocationsRequest - Send a myAllocations request
// Message name: Account allocations (USER_DATA) Request
type MyAllocationsRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *MyAllocationsRequestParams `json:"params,omitempty"`
}

// NewMyAllocationsRequest creates a new MyAllocationsRequest instance
func NewMyAllocationsRequest() *MyAllocationsRequest {
	return &MyAllocationsRequest{}
}

// String returns string representation of MyAllocationsRequest
func (s MyAllocationsRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *MyAllocationsRequest) SetParams(value MyAllocationsRequestParams) *MyAllocationsRequest {
	r.Params = &value
	return r
}

// SetEndTime sets the endTime parameter and returns the struct for method chaining
func (r *MyAllocationsRequest) SetEndTime(value int64) *MyAllocationsRequest {
	if r.Params == nil {
		r.Params = &MyAllocationsRequestParams{}
	}
	r.Params.EndTime = &value
	return r
}

// SetFromAllocationId sets the fromAllocationId parameter and returns the struct for method chaining
func (r *MyAllocationsRequest) SetFromAllocationId(value int64) *MyAllocationsRequest {
	if r.Params == nil {
		r.Params = &MyAllocationsRequestParams{}
	}
	r.Params.FromAllocationId = &value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *MyAllocationsRequest) SetLimit(value int64) *MyAllocationsRequest {
	if r.Params == nil {
		r.Params = &MyAllocationsRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *MyAllocationsRequest) SetOrderId(value int64) *MyAllocationsRequest {
	if r.Params == nil {
		r.Params = &MyAllocationsRequestParams{}
	}
	r.Params.OrderId = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *MyAllocationsRequest) SetRecvWindow(value int64) *MyAllocationsRequest {
	if r.Params == nil {
		r.Params = &MyAllocationsRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetStartTime sets the startTime parameter and returns the struct for method chaining
func (r *MyAllocationsRequest) SetStartTime(value int64) *MyAllocationsRequest {
	if r.Params == nil {
		r.Params = &MyAllocationsRequestParams{}
	}
	r.Params.StartTime = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *MyAllocationsRequest) SetSymbol(value string) *MyAllocationsRequest {
	if r.Params == nil {
		r.Params = &MyAllocationsRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *MyAllocationsRequest) SetTimestamp(value int64) *MyAllocationsRequest {
	if r.Params == nil {
		r.Params = &MyAllocationsRequestParams{}
	}
	r.Params.Timestamp = &value
	return r
}


