package models

// SubscribeRequest represents global message '#/components/messages/subscribeRequest'
type SubscribeRequest struct {
	Method string `json:"method"` // Method name
	Params []string `json:"params"` // Array of stream names to subscribe to
	Id MessageID `json:"id"` // The id is used as an identifier to uniquely identify the messages going back and forth. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null 
}


