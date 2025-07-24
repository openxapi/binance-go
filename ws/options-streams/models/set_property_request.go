package models

import (
	"encoding/json"
)

// SetPropertyRequest represents SetPropertyRequest
type SetPropertyRequest struct {
	// Method name
	MethodName string `json:"method,omitempty"`
	// Array containing property name and value
	ArrayContainingPropertyNameAndValue []interface{} `json:"params,omitempty"`
	// Request ID
	RequestId int `json:"id,omitempty"`
}

// String returns string representation of SetPropertyRequest
func (s SetPropertyRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


