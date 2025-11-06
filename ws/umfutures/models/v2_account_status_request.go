package models

// V2AccountStatusRequest represents global message '#/components/messages/v2AccountStatusRequest'
type V2AccountStatusRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		RecvWindow int64 `json:"recvWindow,omitempty"`
		Signature string `json:"signature,omitempty"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


