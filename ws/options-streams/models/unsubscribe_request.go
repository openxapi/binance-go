package models

// UnsubscribeRequest represents global message '#/components/messages/unsubscribeRequest'
type UnsubscribeRequest struct {
	Method string `json:"method"` // Method name
	Params []string `json:"params"` // Array of stream names to unsubscribe from
	Id int64 `json:"id"` // Request ID
}


