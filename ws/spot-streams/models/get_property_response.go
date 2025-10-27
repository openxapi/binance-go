package models

// GetPropertyResponse represents global message '#/components/messages/getPropertyResponse'
type GetPropertyResponse struct {
	Result interface{} `json:"result,omitempty"` // Property value (currently only 'combined' property supported)
	Id MessageID `json:"id,omitempty"` // The id is used as an identifier to uniquely identify the messages going back and forth. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null 
}


