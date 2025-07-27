package models

import (
	"encoding/json"
)

// SetProperty represents Set Property Request
type SetProperty struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Array containing property name and value
	ArrayContainingPropertyNameAndValue []interface{} `json:"params,omitempty"`
	// Request ID
	RequestId string `json:"id,omitempty"`
}

// String returns string representation of SetProperty
func (s SetProperty) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewSetProperty creates a new SetProperty instance
func NewSetProperty() *SetProperty {
	return &SetProperty{}
}


