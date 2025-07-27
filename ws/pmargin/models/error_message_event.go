package models

import (
	"encoding/json"
)

// ErrorMessageEventError represents a nested object structure
type ErrorMessageEventError struct {
	// Error code
	Code int64 `json:"code,omitempty"`
	// Error message
	Msg string `json:"msg,omitempty"`
}

// ErrorMessageEvent - Message structure
// Message name: Error Message
type ErrorMessageEvent struct {
	Error *ErrorMessageEventError `json:"error,omitempty"`
	Id int64 `json:"id,omitempty"`
}

// String returns string representation of ErrorMessageEvent
func (s ErrorMessageEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


