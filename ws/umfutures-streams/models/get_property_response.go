package models

// GetPropertyResponse represents global message '#/components/messages/getPropertyResponse'
type GetPropertyResponse struct {
	Result interface{} `json:"result,omitempty"` // Property value (currently only 'combined' property supported)
	Id int64 `json:"id,omitempty"`
}


