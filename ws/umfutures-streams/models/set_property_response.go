package models

// SetPropertyResponse represents global message '#/components/messages/setPropertyResponse'
type SetPropertyResponse struct {
	Result interface{} `json:"result,omitempty"` // Always null for successful property setting
	Id int64 `json:"id,omitempty"`
}


