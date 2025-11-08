package models

// ListSubscriptionsRequest represents global message '#/components/messages/listSubscriptionsRequest'
type ListSubscriptionsRequest struct {
	Method string `json:"method"` // Method name
	Id MessageID `json:"id"` // The id is used as an identifier to uniquely identify the messages going back and forth. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null 
}


