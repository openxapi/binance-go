package models

// SessionLogonRequest represents global message '#/components/messages/sessionLogonRequest'
type SessionLogonRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"` // API Key (Only Ed25519 keys are supported)
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000
		Signature string `json:"signature"` // Ed25519 signature
		Timestamp int64 `json:"timestamp"` // Timestamp in milliseconds
	} `json:"params,omitempty"` // params property
}


