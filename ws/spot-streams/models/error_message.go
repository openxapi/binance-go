package models

import (
	"fmt"
)

// ErrorMessage represents global message '#/components/messages/errorMessage'
type ErrorMessage struct {
	ErrorPayload struct {
		Code int `json:"code,omitempty"` // Error code
		Msg string `json:"msg,omitempty"` // Error message
	} `json:"error,omitempty"`
	Id MessageID `json:"id,omitempty"` // The id is used as an identifier to uniquely identify the messages going back and forth. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null 
}


// Error implements the error interface for ErrorMessage
func (m *ErrorMessage) Error() string {
	if m == nil {
		return ""
	}
	payload := m.ErrorPayload
	switch {
	case payload.Msg != "" && payload.Code != 0:
		return fmt.Sprintf("%v: %s", payload.Code, payload.Msg)
	case payload.Msg != "":
		return payload.Msg
	case payload.Code != 0:
		return fmt.Sprintf("%v", payload.Code)
	default:
		return "unknown error"
	}
}

