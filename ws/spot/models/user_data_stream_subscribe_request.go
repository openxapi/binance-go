package models

import (
	"encoding/json"
)

// UserDataStreamSubscribeRequest - Send a userDataStream.subscribe request
// Message name: Subscribe to User Data Stream (USER_STREAM) Request
type UserDataStreamSubscribeRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
}

// NewUserDataStreamSubscribeRequest creates a new UserDataStreamSubscribeRequest instance
func NewUserDataStreamSubscribeRequest() *UserDataStreamSubscribeRequest {
	return &UserDataStreamSubscribeRequest{}
}

// String returns string representation of UserDataStreamSubscribeRequest
func (s UserDataStreamSubscribeRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


