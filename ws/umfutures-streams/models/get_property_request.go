package models

// GetPropertyRequest represents global message '#/components/messages/getPropertyRequest'
type GetPropertyRequest struct {
	Method string `json:"method"` // Method name
	Params []string `json:"params"` // Array containing property name to retrieve
	Id int64 `json:"id"` // Request ID
}


