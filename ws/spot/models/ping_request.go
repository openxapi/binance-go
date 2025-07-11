package models

import (
	"encoding/json"
)

// PingRequest - Send a ping request
// Message name: Test connectivity Request
type PingRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
}

// NewPingRequest creates a new PingRequest instance
func NewPingRequest() *PingRequest {
	return &PingRequest{}
}

// String returns string representation of PingRequest
func (s PingRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


