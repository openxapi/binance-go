package models

import (
	"fmt"
)

// ErrorMessage represents global message '#/components/messages/errorMessage'
type ErrorMessage struct {
	ErrorPayload struct {
		Code int `json:"code"` // Binance-specific error code.
		Msg string `json:"msg"` // Human-readable error message.
	} `json:"error"` // Details about the error that occurred.
	Id MessageID `json:"id"` // Unique identifier used to correlate requests and responses. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null
	RateLimits []struct {
		Count int `json:"count"` // Current usage count within the interval.
		Interval string `json:"interval"` // Time interval unit for the rate limit window.
		IntervalNum int `json:"intervalNum"` // Number of intervals the rate limit covers.
		Limit int `json:"limit"` // Maximum number of allowed operations for the interval.
		RateLimitType string `json:"rateLimitType"` // Rate limit type affected by the request.
	} `json:"rateLimits,omitempty"` // Rate limit counters associated with the request.
	Status int `json:"status"` // HTTP-style status code describing the error.
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
	case m.Status != 0:
		return fmt.Sprintf("status=%v", m.Status)
	default:
		return "unknown error"
	}
}

