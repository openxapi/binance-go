package models

import (
	"encoding/json"
)

// UserDataStreamStartResponseRateLimitsItem represents a nested object structure
type UserDataStreamStartResponseRateLimitsItem struct {
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

// UserDataStreamStartResponseResult represents a nested object structure
type UserDataStreamStartResponseResult struct {
	// listenKey property
	ListenKey string `json:"listenKey,omitempty"`
}

// UserDataStreamStartResponse - Receive response from userDataStream.start
// Message name: Start User Data Stream (USER_STREAM) Response
type UserDataStreamStartResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []UserDataStreamStartResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result UserDataStreamStartResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of UserDataStreamStartResponse
func (s UserDataStreamStartResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


