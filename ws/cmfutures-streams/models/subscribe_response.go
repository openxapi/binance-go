package models

// SubscribeResponse represents global message '#/components/messages/subscribeResponse'
type SubscribeResponse struct {
	Result interface{} `json:"result,omitempty"` // Always null for successful subscription
	Id MessageID `json:"id,omitempty"`
}


