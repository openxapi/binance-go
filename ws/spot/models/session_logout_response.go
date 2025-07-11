package models

import (
	"encoding/json"
)

// SessionLogoutResponseResult represents a nested object structure
type SessionLogoutResponseResult struct {
	// apiKey property
	ApiKey string `json:"apiKey,omitempty"`
	// authorizedSince property
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
	// result property
	Result SessionLogoutResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of SessionLogoutResponse
func (s SessionLogoutResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


