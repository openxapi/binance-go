package models

// SessionStatusResponse represents global message '#/components/messages/sessionStatusResponse'
type SessionStatusResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // Current count
		Interval string `json:"interval,omitempty"` // Rate limit interval
		IntervalNum int `json:"intervalNum,omitempty"` // Interval number
		Limit int `json:"limit,omitempty"` // Rate limit
		RateLimitType string `json:"rateLimitType,omitempty"` // Rate limit type
	} `json:"rateLimits,omitempty"` // Rate limit information
	Result struct {
		ApiKey string `json:"apiKey,omitempty"` // apiKey property (shows current authenticated key or empty if not authenticated)
		AuthorizedSince int64 `json:"authorizedSince,omitempty"` // authorizedSince property (timestamp when authentication occurred)
		ConnectedSince int64 `json:"connectedSince,omitempty"` // connectedSince property (timestamp when WebSocket connection was established)
		ReturnRateLimits bool `json:"returnRateLimits,omitempty"` // returnRateLimits property
		ServerTime int64 `json:"serverTime,omitempty"` // serverTime property (current server timestamp)
		UserDataStream bool `json:"userDataStream,omitempty"` // userDataStream property (indicates if connection can receive user data stream events)
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


