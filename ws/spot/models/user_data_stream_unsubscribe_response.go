package models

import (
	"encoding/json"
)

// UserDataStreamUnsubscribeResponseResult represents a nested object structure
type UserDataStreamUnsubscribeResponseResult struct {
}

// UserDataStreamUnsubscribeResponse - Receive response from userDataStream.unsubscribe
// Message name: Unsubscribe from User Data Stream (USER_STREAM) Response
type UserDataStreamUnsubscribeResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// result property
	Result interface{} `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of UserDataStreamUnsubscribeResponse
func (s UserDataStreamUnsubscribeResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


