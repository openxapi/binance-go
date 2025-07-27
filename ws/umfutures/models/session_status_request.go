package models

import (
	"encoding/json"
)

// SessionStatusRequest - Send a session.status request
// Message name: Query session status Request
type SessionStatusRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
}

// NewSessionStatusRequest creates a new SessionStatusRequest instance
func NewSessionStatusRequest() *SessionStatusRequest {
	return &SessionStatusRequest{}
}

// String returns string representation of SessionStatusRequest
func (s SessionStatusRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


