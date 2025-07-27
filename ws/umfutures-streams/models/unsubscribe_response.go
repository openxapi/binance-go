package models

import (
	"encoding/json"
)

// UnsubscribeResponse represents Subscription Response
type UnsubscribeResponse struct {
	// Always null for successful subscription
	AlwaysNullForSuccessfulSubscription interface{} `json:"result,omitempty"`
	Id string `json:"id,omitempty"`
}

// String returns string representation of UnsubscribeResponse
func (s UnsubscribeResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewUnsubscribeResponse creates a new UnsubscribeResponse instance
func NewUnsubscribeResponse() *UnsubscribeResponse {
	return &UnsubscribeResponse{}
}


