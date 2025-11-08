package models

// OpenOrdersCancelAllRequest represents global message '#/components/messages/openOrdersCancelAllRequest'
type OpenOrdersCancelAllRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000
		Signature string `json:"signature"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


