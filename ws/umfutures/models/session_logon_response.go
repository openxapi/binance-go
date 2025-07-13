package models

import (
	"encoding/json"
)

// SessionLogonResponseRateLimitsItem represents a nested object structure
type SessionLogonResponseRateLimitsItem struct {
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

// SessionLogonResponseResult represents a nested object structure
type SessionLogonResponseResult struct {
	// apiKey property
	ApiKey string `json:"apiKey,omitempty"`
	// authorizedSince property
	AuthorizedSince int64 `json:"authorizedSince,omitempty"`
	// connectedSince property
	ConnectedSince int64 `json:"connectedSince,omitempty"`
	// returnRateLimits property
	ReturnRateLimits bool `json:"returnRateLimits,omitempty"`
	// serverTime property
	ServerTime int64 `json:"serverTime,omitempty"`
	// userDataStream property
	UserDataStream bool `json:"userDataStream,omitempty"`
}

// SessionLogonResponse - Receive response from session.logon
// Message name: Log in with API key (SIGNED) Response
type SessionLogonResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// Rate limit information
	RateLimits []SessionLogonResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *SessionLogonResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of SessionLogonResponse
func (s SessionLogonResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


