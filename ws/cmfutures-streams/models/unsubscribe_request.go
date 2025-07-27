package models

import (
	"encoding/json"
)

// UnsubscribeRequest represents UnsubscribeRequest
type UnsubscribeRequest struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Array of stream names to unsubscribe from
	ArrayOfStreamNamesToUnsubscribeFrom []string `json:"params,omitempty"`
	// Request ID
	RequestId string `json:"id,omitempty"`
}

// String returns string representation of UnsubscribeRequest
func (s UnsubscribeRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


