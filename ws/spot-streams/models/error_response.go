package models

import (
	"encoding/json"
)

// ErrorResponse represents ErrorResponse
type ErrorResponse struct {
	Error interface{} `json:"error,omitempty"`
	// Request ID echo
	RequestIdEcho int `json:"id,omitempty"`
}

// String returns string representation of ErrorResponse
func (s ErrorResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


