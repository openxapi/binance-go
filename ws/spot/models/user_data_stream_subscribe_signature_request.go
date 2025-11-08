package models

// UserDataStreamSubscribeSignatureRequest represents global message '#/components/messages/userDataStreamSubscribeSignatureRequest'
type UserDataStreamSubscribeSignatureRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		Signature string `json:"signature"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


