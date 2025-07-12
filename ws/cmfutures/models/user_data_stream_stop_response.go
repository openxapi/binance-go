package models

import (
	"encoding/json"
)

// UserDataStreamStopResponseRateLimitsItem represents a nested object structure
type UserDataStreamStopResponseRateLimitsItem struct {
	// count property (example: 2)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 2400)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// UserDataStreamStopResponseResult represents a nested object structure
type UserDataStreamStopResponseResult struct {
}

// UserDataStreamStopResponse - Receive response from userDataStream.stop
// Message name: Close User Data Stream (USER_STREAM) Response
type UserDataStreamStopResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []UserDataStreamStopResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *interface{} `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of UserDataStreamStopResponse
func (s UserDataStreamStopResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


