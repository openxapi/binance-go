package models

import (
	"encoding/json"
)

// SessionLogoutRequest - Send a session.logout request
// Message name: SIGNED request example (Ed25519) Request
type SessionLogoutRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
}

// NewSessionLogoutRequest creates a new SessionLogoutRequest instance
func NewSessionLogoutRequest() *SessionLogoutRequest {
	return &SessionLogoutRequest{}
}

// String returns string representation of SessionLogoutRequest
func (s SessionLogoutRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


