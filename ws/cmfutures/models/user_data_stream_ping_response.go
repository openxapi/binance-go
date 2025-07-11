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
	// limit property (example: 2400)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// UserDataStreamPingResponseResult represents a nested object structure
type UserDataStreamPingResponseResult struct {
	// listenKey property
	ListenKey string `json:"listenKey,omitempty"`
}

// UserDataStreamPingResponse - Receive response from userDataStream.ping
// Message name: Keepalive User Data Stream (USER_STREAM) Response
type UserDataStreamPingResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []UserDataStreamPingResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result UserDataStreamPingResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of UserDataStreamPingResponse
func (s UserDataStreamPingResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


