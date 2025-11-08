package models

// PingRequest represents global message '#/components/messages/pingRequest'
type PingRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
}


