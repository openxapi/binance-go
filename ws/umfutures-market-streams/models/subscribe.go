package models

import (
	"encoding/json"
)

// Subscribe represents Subscribe Request
type Subscribe struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Array of stream names to subscribe to
	ArrayOfStreamNamesToSubscribeTo []string `json:"params,omitempty"`
	// Request ID
	RequestId int `json:"id,omitempty"`
}

// String returns string representation of Subscribe
func (s Subscribe) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewSubscribe creates a new Subscribe instance
func NewSubscribe() *Subscribe {
	return &Subscribe{}
}


