package models

import (
	"encoding/json"
)

// SubscribeResponse represents Subscription Response
type SubscribeResponse struct {
	// Always null for successful subscription
	AlwaysNullForSuccessfulSubscription interface{} `json:"result,omitempty"`
	Id string `json:"id,omitempty"`
}

// String returns string representation of SubscribeResponse
func (s SubscribeResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// NewSubscribeResponse creates a new SubscribeResponse instance
func NewSubscribeResponse() *SubscribeResponse {
	return &SubscribeResponse{}
}


