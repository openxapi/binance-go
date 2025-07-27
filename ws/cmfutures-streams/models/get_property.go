package models

import (
	"encoding/json"
)

// GetProperty represents Get Property Request
type GetProperty struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Array containing property name to retrieve
	ArrayContainingPropertyNameToRetrieve []string `json:"params,omitempty"`
	// Request ID
	RequestId string `json:"id,omitempty"`
}

// String returns string representation of GetProperty
func (s GetProperty) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewGetProperty creates a new GetProperty instance
func NewGetProperty() *GetProperty {
	return &GetProperty{}
}


