package models

// SetPropertyResponse represents global message '#/components/messages/setPropertyResponse'
type SetPropertyResponse struct {
	Result interface{} `json:"result,omitempty"` // Always null for successful property setting
	Id MessageID `json:"id,omitempty"` // The id is used as an identifier to uniquely identify the messages going back and forth. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null 
}


