package models

// DepthRequest represents global message '#/components/messages/depthRequest'
type DepthRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Limit int `json:"limit,omitempty"` // Default: 100; Maximum: 5000
		Symbol string `json:"symbol"`
	} `json:"params,omitempty"` // params property
}


