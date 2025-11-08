package models

// SetPropertyRequest represents global message '#/components/messages/setPropertyRequest'
type SetPropertyRequest struct {
	Method string `json:"method"` // Method name
	Params []interface{} `json:"params"` // Array containing property name and value
	Id MessageID `json:"id"` // The id is used as an identifier to uniquely identify the messages going back and forth. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null 
}


