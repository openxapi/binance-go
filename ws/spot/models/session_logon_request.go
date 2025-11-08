package models

// SessionLogonRequest represents global message '#/components/messages/sessionLogonRequest'
type SessionLogonRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000
		Signature string `json:"signature"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


