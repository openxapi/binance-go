package models

import (
	"encoding/json"
)

// Unsubscribe represents Unsubscribe Request
type Unsubscribe struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Array of stream names to unsubscribe from
	ArrayOfStreamNamesToUnsubscribeFrom []string `json:"params,omitempty"`
	// Request ID
	RequestId string `json:"id,omitempty"`
}

// String returns string representation of Unsubscribe
func (s Unsubscribe) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewUnsubscribe creates a new Unsubscribe instance
func NewUnsubscribe() *Unsubscribe {
	return &Unsubscribe{}
}


