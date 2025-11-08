package models

// UnsubscribeResponse represents global message '#/components/messages/unsubscribeResponse'
type UnsubscribeResponse struct {
	Result interface{} `json:"result,omitempty"` // Always null for successful subscription
	Id MessageID `json:"id,omitempty"`
}


