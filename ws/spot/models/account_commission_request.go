package models

// AccountCommissionRequest represents global message '#/components/messages/accountCommissionRequest'
type AccountCommissionRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		Signature string `json:"signature,omitempty"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp,omitempty"`
	} `json:"params,omitempty"` // params property
}


