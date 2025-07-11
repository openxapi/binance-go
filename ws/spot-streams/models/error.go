package models

import (
	"encoding/json"
)

// Error represents Error Message
type Error struct {
	Error interface{} `json:"error,omitempty"`
	// Request ID echo
	RequestIdEcho int `json:"id,omitempty"`
}

// String returns string representation of Error
func (s Error) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewError creates a new Error instance
func NewError() *Error {
	return &Error{}
}


