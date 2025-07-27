package models

import (
	"encoding/json"
)

// GetPropertyResponse represents GetPropertyResponse
type GetPropertyResponse struct {
	// Property value (currently only 'combined' property supported)
	PropertyValue interface{} `json:"result,omitempty"`
	Id string `json:"id,omitempty"`
}

// String returns string representation of GetPropertyResponse
func (s GetPropertyResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


