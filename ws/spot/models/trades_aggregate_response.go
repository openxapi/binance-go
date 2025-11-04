package models

// TradesAggregateResponse represents global message '#/components/messages/tradesAggregateResponse'
type TradesAggregateResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		M bool `json:"M,omitempty"` // M property
		T int `json:"T,omitempty"` // T property
		A int `json:"a,omitempty"` // a property
		F int `json:"f,omitempty"` // f property
		L int `json:"l,omitempty"` // l property
		M2 bool `json:"m,omitempty"` // m property
		P string `json:"p,omitempty"` // p property
		Q string `json:"q,omitempty"` // q property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


