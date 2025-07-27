package models

import (
	"encoding/json"
)

// SessionLogoutResponseRateLimitsItem represents a nested object structure
type SessionLogoutResponseRateLimitsItem struct {
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

// SessionLogoutResponseResult represents a nested object structure
type SessionLogoutResponseResult struct {
	// apiKey property (empty after logout)
	ApiKey string `json:"apiKey,omitempty"`
	// authorizedSince property (empty after logout)
	AuthorizedSince string `json:"authorizedSince,omitempty"`
	// connectedSince property
	ConnectedSince int64 `json:"connectedSince,omitempty"`
	// returnRateLimits property
	ReturnRateLimits bool `json:"returnRateLimits,omitempty"`
	// serverTime property
	ServerTime int64 `json:"serverTime,omitempty"`
	// userDataStream property
	UserDataStream bool `json:"userDataStream,omitempty"`
}

// SessionLogoutResponse - Receive response from session.logout
// Message name: Log out of the session Response
type SessionLogoutResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// Rate limit information
	RateLimits []SessionLogoutResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result *SessionLogoutResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of SessionLogoutResponse
func (s SessionLogoutResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


