package models

import (
	"encoding/json"
)

// TimeRequest - Send a time request
// Message name: Check server time Request
type TimeRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
}

// NewTimeRequest creates a new TimeRequest instance
func NewTimeRequest() *TimeRequest {
	return &TimeRequest{}
}

// String returns string representation of TimeRequest
func (s TimeRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


