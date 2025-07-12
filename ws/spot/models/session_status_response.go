package models

import (
	"encoding/json"
)

// SessionStatusResponseResult represents a nested object structure
type SessionStatusResponseResult struct {
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

// SessionStatusResponse - Receive response from session.status
// Message name: Query session status Response
type SessionStatusResponse struct {
	// id property
	Id string `json:"id,omitempty"`
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


