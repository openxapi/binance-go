package models

// SetPropertyRequest represents global message '#/components/messages/setPropertyRequest'
type SetPropertyRequest struct {
	Method string `json:"method"` // Method name
	Params []interface{} `json:"params"` // Array containing property name and value
	Id MessageID `json:"id"` // Request ID
}


