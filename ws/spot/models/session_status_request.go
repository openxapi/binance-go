package models

// SessionStatusRequest represents global message '#/components/messages/sessionStatusRequest'
type SessionStatusRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
}


