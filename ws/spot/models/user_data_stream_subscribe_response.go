package models

import (
	"encoding/json"
)

// UserDataStreamSubscribeResponseResult represents a nested object structure
type UserDataStreamSubscribeResponseResult struct {
}

// UserDataStreamSubscribeResponse - Receive response from userDataStream.subscribe
// Message name: Subscribe to User Data Stream (USER_STREAM) Response
type UserDataStreamSubscribeResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// result property
	Result *interface{} `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of UserDataStreamSubscribeResponse
func (s UserDataStreamSubscribeResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


