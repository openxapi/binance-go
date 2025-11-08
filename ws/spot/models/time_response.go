package models

// TimeResponse represents global message '#/components/messages/timeResponse'
type TimeResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		ServerTime int64 `json:"serverTime,omitempty"` // serverTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


