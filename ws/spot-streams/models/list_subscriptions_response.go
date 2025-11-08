package models

// ListSubscriptionsResponse represents global message '#/components/messages/listSubscriptionsResponse'
type ListSubscriptionsResponse struct {
	Result []string `json:"result,omitempty"` // Array of active stream names
	Id MessageID `json:"id,omitempty"` // The id is used as an identifier to uniquely identify the messages going back and forth. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null 
}


