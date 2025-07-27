package models

import (
	"encoding/json"
)

// ListSubscriptionsResponse represents ListSubscriptionsResponse
type ListSubscriptionsResponse struct {
	// Array of active stream names
	ArrayOfActiveStreamNames []string `json:"result,omitempty"`
	Id string `json:"id,omitempty"`
}

// String returns string representation of ListSubscriptionsResponse
func (s ListSubscriptionsResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


