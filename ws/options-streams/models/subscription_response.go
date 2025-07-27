package models

import (
	"encoding/json"
)

// SubscriptionResponse represents SubscriptionResponse
type SubscriptionResponse struct {
	// Always null for successful subscription
	AlwaysNullForSuccessfulSubscription interface{} `json:"result,omitempty"`
	Id int `json:"id,omitempty"`
}

// String returns string representation of SubscriptionResponse
func (s SubscriptionResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


