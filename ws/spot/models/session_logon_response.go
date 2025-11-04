package models

// SessionLogonResponse represents global message '#/components/messages/sessionLogonResponse'
type SessionLogonResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	Result struct {
		ApiKey string `json:"apiKey,omitempty"` // apiKey property
		AuthorizedSince int `json:"authorizedSince,omitempty"` // authorizedSince property
		ConnectedSince int `json:"connectedSince,omitempty"` // connectedSince property
		ReturnRateLimits bool `json:"returnRateLimits,omitempty"` // returnRateLimits property
		ServerTime int64 `json:"serverTime,omitempty"` // serverTime property
		UserDataStream bool `json:"userDataStream,omitempty"` // userDataStream property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


