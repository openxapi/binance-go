package models

import (
	"encoding/json"
)

// ListSubscriptions represents List Subscriptions Request
type ListSubscriptions struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Request ID
	RequestId int `json:"id,omitempty"`
}

// String returns string representation of ListSubscriptions
func (s ListSubscriptions) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewListSubscriptions creates a new ListSubscriptions instance
func NewListSubscriptions() *ListSubscriptions {
	return &ListSubscriptions{}
}


