package models

// UserDataStreamStartRequest represents global message '#/components/messages/userDataStreamStartRequest'
type UserDataStreamStartRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"` // apiKey property
	} `json:"params,omitempty"` // params property
}


