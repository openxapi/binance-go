package models

import (
	"encoding/json"
)

// ErrorMessageError represents a nested object structure
type ErrorMessageError struct {
	// Error code
	Code int64 `json:"code,omitempty"`
	// Error message
	Msg string `json:"msg,omitempty"`
}

// ErrorMessage - Message structure
// Message name: Error Message
type ErrorMessage struct {
	Error *ErrorMessageError `json:"error,omitempty"`
	// Request ID echo
	Id int64 `json:"id,omitempty"`
}

// String returns string representation of ErrorMessage
func (s ErrorMessage) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


