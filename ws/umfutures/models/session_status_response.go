package models

import (
	"encoding/json"
)

// SessionStatusResponseRateLimitsItem represents a nested object structure
type SessionStatusResponseRateLimitsItem struct {
	// Current count (example: 2)
	Count int64 `json:"count,omitempty"`
	// Rate limit interval (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// Interval number (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// Rate limit (example: 2400)
	Limit int64 `json:"limit,omitempty"`
	// Rate limit type (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// SessionStatusResponseResult represents a nested object structure
type SessionStatusResponseResult struct {
	// apiKey property (shows current authenticated key or empty if not authenticated)
	ApiKey string `json:"apiKey,omitempty"`
	// authorizedSince property (timestamp when authentication occurred)
	AuthorizedSince int64 `json:"authorizedSince,omitempty"`
	// connectedSince property (timestamp when WebSocket connection was established)
	ConnectedSince int64 `json:"connectedSince,omitempty"`
	// returnRateLimits property
	ReturnRateLimits bool `json:"returnRateLimits,omitempty"`
	// serverTime property (current server timestamp)
	ServerTime int64 `json:"serverTime,omitempty"`
	// userDataStream property (indicates if connection can receive user data stream events)
	UserDataStream bool `json:"userDataStream,omitempty"`
}

// SessionStatusResponse - Receive response from session.status
// Message name: Query session status Response
type SessionStatusResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// Rate limit information
	RateLimits []SessionStatusResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *SessionStatusResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of SessionStatusResponse
func (s SessionStatusResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


