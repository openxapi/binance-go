package models

import (
	"encoding/json"
)

// GetPropertyRequest represents GetPropertyRequest
type GetPropertyRequest struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Array containing property name to retrieve
	ArrayContainingPropertyNameToRetrieve []string `json:"params,omitempty"`
	// Request ID
	RequestId int `json:"id,omitempty"`
}

// String returns string representation of GetPropertyRequest
func (s GetPropertyRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


