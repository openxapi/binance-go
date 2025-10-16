package models

// SubscriptionResponse represents global message '#/components/messages/subscriptionResponse'
type SubscriptionResponse struct {
	Result struct {	} `json:"result,omitempty"` // Always null for successful subscription
	Id int64 `json:"id,omitempty"`
}


