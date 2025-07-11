package models

import (
	"encoding/json"
)

// KlinesRequestParams contains the parameters for KlinesRequest
type KlinesRequestParams struct {
	EndTime *int64 `json:"endTime,omitempty"`
	Interval string `json:"interval"`
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	Symbol string `json:"symbol"`
	// Default: 0 (UTC)
	TimeZone *string `json:"timeZone,omitempty"`
}

// KlinesRequest - Send a klines request
// Message name: Klines Request
type KlinesRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *KlinesRequestParams `json:"params,omitempty"`
}

// NewKlinesRequest creates a new KlinesRequest instance
func NewKlinesRequest() *KlinesRequest {
	return &KlinesRequest{}
}

// String returns string representation of KlinesRequest
func (s KlinesRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *KlinesRequest) SetParams(value KlinesRequestParams) *KlinesRequest {
	r.Params = &value
	return r
}

// SetEndTime sets the endTime parameter and returns the struct for method chaining
func (r *KlinesRequest) SetEndTime(value int64) *KlinesRequest {
	if r.Params == nil {
		r.Params = &KlinesRequestParams{}
	}
	r.Params.EndTime = &value
	return r
}

// SetInterval sets the interval parameter and returns the struct for method chaining
func (r *KlinesRequest) SetInterval(value string) *KlinesRequest {
	if r.Params == nil {
		r.Params = &KlinesRequestParams{}
	}
	r.Params.Interval = value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *KlinesRequest) SetLimit(value int64) *KlinesRequest {
	if r.Params == nil {
		r.Params = &KlinesRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetStartTime sets the startTime parameter and returns the struct for method chaining
func (r *KlinesRequest) SetStartTime(value int64) *KlinesRequest {
	if r.Params == nil {
		r.Params = &KlinesRequestParams{}
	}
	r.Params.StartTime = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *KlinesRequest) SetSymbol(value string) *KlinesRequest {
	if r.Params == nil {
		r.Params = &KlinesRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimeZone sets the timeZone parameter and returns the struct for method chaining
func (r *KlinesRequest) SetTimeZone(value string) *KlinesRequest {
	if r.Params == nil {
		r.Params = &KlinesRequestParams{}
	}
	r.Params.TimeZone = &value
	return r
}


