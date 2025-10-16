package models

// UnsubscriptionResponse represents global message '#/components/messages/unsubscriptionResponse'
type UnsubscriptionResponse struct {
	Result struct {	} `json:"result,omitempty"` // Always null for successful unsubscription
	Id int64 `json:"id,omitempty"`
}


