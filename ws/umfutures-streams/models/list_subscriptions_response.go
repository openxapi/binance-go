package models

// ListSubscriptionsResponse represents global message '#/components/messages/listSubscriptionsResponse'
type ListSubscriptionsResponse struct {
	Result []string `json:"result,omitempty"` // Array of active stream names
	Id int64 `json:"id,omitempty"`
}


