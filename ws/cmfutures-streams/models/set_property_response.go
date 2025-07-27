package models

import (
	"encoding/json"
)

// SetPropertyResponse represents SetPropertyResponse
type SetPropertyResponse struct {
	// Always null for successful property setting
	AlwaysNullForSuccessfulPropertySetting interface{} `json:"result,omitempty"`
	Id string `json:"id,omitempty"`
}

// String returns string representation of SetPropertyResponse
func (s SetPropertyResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


