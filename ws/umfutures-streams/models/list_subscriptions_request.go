package models

import (
	"encoding/json"
)

// ListSubscriptionsRequest represents ListSubscriptionsRequest
type ListSubscriptionsRequest struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Request ID
	RequestId string `json:"id,omitempty"`
}

// String returns string representation of ListSubscriptionsRequest
func (s ListSubscriptionsRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


