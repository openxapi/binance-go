package models

// UserDataStreamPingRequest represents global message '#/components/messages/userDataStreamPingRequest'
type UserDataStreamPingRequest struct {
	Id int64 `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"` // Required if session is not authenticated via session.logon
	} `json:"params,omitempty"` // params property
}


