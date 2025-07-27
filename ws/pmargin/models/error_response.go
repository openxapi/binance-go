package models

import (
	"encoding/json"
)

// ErrorResponseError represents a nested object structure
type ErrorResponseError struct {
	// Error code
	Code int64 `json:"code,omitempty"`
	// Error message
	Msg string `json:"msg,omitempty"`
}

// ErrorResponse represents ErrorResponse
type ErrorResponse struct {
	Error ErrorResponseError `json:"error,omitempty"`
	Id int64 `json:"id,omitempty"`
}

// String returns string representation of ErrorResponse
func (s ErrorResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


