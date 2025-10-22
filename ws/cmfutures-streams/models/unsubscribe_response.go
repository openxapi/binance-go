package models

// UnsubscribeResponse represents global message '#/components/messages/unsubscribeResponse'
type UnsubscribeResponse struct {
	Result interface{} `json:"result,omitempty"` // Always null for successful subscription
	Id int64 `json:"id,omitempty"`
}


