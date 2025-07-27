package models

import (
	"encoding/json"
)

// ErrorResponse represents ErrorResponse
type ErrorResponse struct {
	Error *ErrorResponseError `json:"error,omitempty"`
	Id string `json:"id,omitempty"`
}

// ErrorResponseError represents the error details
type ErrorResponseError struct {
	// Error code
	ErrorCode int `json:"code,omitempty"`
	// Error message
	ErrorMessage string `json:"msg,omitempty"`
}

// String returns string representation of ErrorResponse
func (s ErrorResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


