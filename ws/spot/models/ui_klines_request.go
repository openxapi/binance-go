package models

import (
	"encoding/json"
)

// UiKlinesRequestParams contains the parameters for UiKlinesRequest
type UiKlinesRequestParams struct {
	EndTime *int64 `json:"endTime,omitempty"`
	// See klines
	Interval string `json:"interval"`
	// Default: 500; Maximum: 1000
	Limit *int64 `json:"limit,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	Symbol string `json:"symbol"`
	// Default: 0 (UTC)
	TimeZone *string `json:"timeZone,omitempty"`
}

// UiKlinesRequest - Send a uiKlines request
// Message name: UI Klines Request
type UiKlinesRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *UiKlinesRequestParams `json:"params,omitempty"`
}

// NewUiKlinesRequest creates a new UiKlinesRequest instance
func NewUiKlinesRequest() *UiKlinesRequest {
	return &UiKlinesRequest{}
}

// String returns string representation of UiKlinesRequest
func (s UiKlinesRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *UiKlinesRequest) SetParams(value UiKlinesRequestParams) *UiKlinesRequest {
	r.Params = &value
	return r
}

// SetEndTime sets the endTime parameter and returns the struct for method chaining
func (r *UiKlinesRequest) SetEndTime(value int64) *UiKlinesRequest {
	if r.Params == nil {
		r.Params = &UiKlinesRequestParams{}
	}
	r.Params.EndTime = &value
	return r
}

// SetInterval sets the interval parameter and returns the struct for method chaining
func (r *UiKlinesRequest) SetInterval(value string) *UiKlinesRequest {
	if r.Params == nil {
		r.Params = &UiKlinesRequestParams{}
	}
	r.Params.Interval = value
	return r
}

// SetLimit sets the limit parameter and returns the struct for method chaining
func (r *UiKlinesRequest) SetLimit(value int64) *UiKlinesRequest {
	if r.Params == nil {
		r.Params = &UiKlinesRequestParams{}
	}
	r.Params.Limit = &value
	return r
}

// SetStartTime sets the startTime parameter and returns the struct for method chaining
func (r *UiKlinesRequest) SetStartTime(value int64) *UiKlinesRequest {
	if r.Params == nil {
		r.Params = &UiKlinesRequestParams{}
	}
	r.Params.StartTime = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *UiKlinesRequest) SetSymbol(value string) *UiKlinesRequest {
	if r.Params == nil {
		r.Params = &UiKlinesRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimeZone sets the timeZone parameter and returns the struct for method chaining
func (r *UiKlinesRequest) SetTimeZone(value string) *UiKlinesRequest {
	if r.Params == nil {
		r.Params = &UiKlinesRequestParams{}
	}
	r.Params.TimeZone = &value
	return r
}


