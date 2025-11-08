package models

// UiKlinesResponse represents global message '#/components/messages/uiKlinesResponse'
type UiKlinesResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result [][]interface{} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


