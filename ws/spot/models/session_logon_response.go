package models

import (
	"encoding/json"
)

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


