package models

// SessionLogonRequest represents global message '#/components/messages/sessionLogonRequest'
type SessionLogonRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"` // apiKey property
		Signature string `json:"signature"` // signature property
		Timestamp int64 `json:"timestamp"` // timestamp property
	} `json:"params,omitempty"` // params property
}


