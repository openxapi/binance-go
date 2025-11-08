package models

// AvgPriceResponse represents global message '#/components/messages/avgPriceResponse'
type AvgPriceResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		CloseTime int64 `json:"closeTime,omitempty"` // closeTime property
		Mins int `json:"mins,omitempty"` // mins property
		Price string `json:"price,omitempty"` // price property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


