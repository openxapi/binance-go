package models

// GetPropertyRequest represents global message '#/components/messages/getPropertyRequest'
type GetPropertyRequest struct {
	Method string `json:"method"` // Method name
	Params []string `json:"params"` // Array containing property name to retrieve
	Id MessageID `json:"id"` // Request ID
}


