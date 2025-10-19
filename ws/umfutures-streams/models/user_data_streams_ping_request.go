package models

// UserDataStreamsPingRequest represents global message '#/components/messages/userDataStreamsPingRequest'
type UserDataStreamsPingRequest struct {
	Id int64 `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"` // apiKey property
	} `json:"params,omitempty"` // params property
}


