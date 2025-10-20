package models

// SubscribeResponse represents global message '#/components/messages/subscribeResponse'
type SubscribeResponse struct {
	Result struct {	} `json:"result,omitempty"` // Always null for successful subscription
	Id int64 `json:"id,omitempty"`
}


