package models

// GetPropertyRequest represents global message '#/components/messages/getPropertyRequest'
type GetPropertyRequest struct {
	Method string `json:"method"` // Method name
	Params []string `json:"params"` // Array containing property name to retrieve
	Id MessageID `json:"id"` // The id is used as an identifier to uniquely identify the messages going back and forth. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null 
}


