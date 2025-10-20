package models

// UnsubscribeResponse represents global message '#/components/messages/unsubscribeResponse'
type UnsubscribeResponse struct {
	Result struct {	} `json:"result,omitempty"` // Always null for successful unsubscription
	Id int64 `json:"id,omitempty"`
}


