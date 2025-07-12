package models

import (
	"encoding/json"
)

// UserDataStreamPingResponseRateLimitsItem represents a nested object structure
type UserDataStreamPingResponseRateLimitsItem struct {
	// count property (example: 2)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 6000)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// UserDataStreamPingResponseResponse represents a nested object structure
type UserDataStreamPingResponseResponse struct {
}

// UserDataStreamPingResponse - Receive response from userDataStream.ping
// Message name: Ping user data stream (USER_STREAM) (Deprecated) Response
type UserDataStreamPingResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []UserDataStreamPingResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// response property
	Response *interface{} `json:"response,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of UserDataStreamPingResponse
func (s UserDataStreamPingResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


