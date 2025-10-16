package models

// SubscribeRequest represents global message '#/components/messages/subscribeRequest'
type SubscribeRequest struct {
	Method string `json:"method"` // Method name
	Params []string `json:"params"` // Array of stream names to subscribe to
	Id int64 `json:"id"` // Request ID
}


