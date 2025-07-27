package models

import (
	"encoding/json"
)

// UserDataStreamUnsubscribeRequest - Send a userDataStream.unsubscribe request
// Message name: Unsubscribe from User Data Stream (USER_STREAM) Request
type UserDataStreamUnsubscribeRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
}

// NewUserDataStreamUnsubscribeRequest creates a new UserDataStreamUnsubscribeRequest instance
func NewUserDataStreamUnsubscribeRequest() *UserDataStreamUnsubscribeRequest {
	return &UserDataStreamUnsubscribeRequest{}
}

// String returns string representation of UserDataStreamUnsubscribeRequest
func (s UserDataStreamUnsubscribeRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


