package models

// SessionLogoutResponse represents global message '#/components/messages/sessionLogoutResponse'
type SessionLogoutResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // Current count
		Interval string `json:"interval,omitempty"` // Rate limit interval
		IntervalNum int `json:"intervalNum,omitempty"` // Interval number
		Limit int `json:"limit,omitempty"` // Rate limit
		RateLimitType string `json:"rateLimitType,omitempty"` // Rate limit type
	} `json:"rateLimits,omitempty"` // Rate limit information
	Result struct {
		ApiKey string `json:"apiKey,omitempty"` // apiKey property (empty after logout)
		AuthorizedSince string `json:"authorizedSince,omitempty"` // authorizedSince property (empty after logout)
		ConnectedSince int64 `json:"connectedSince,omitempty"` // connectedSince property
		ReturnRateLimits bool `json:"returnRateLimits,omitempty"` // returnRateLimits property
		ServerTime int64 `json:"serverTime,omitempty"` // serverTime property
		UserDataStream bool `json:"userDataStream,omitempty"` // userDataStream property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


