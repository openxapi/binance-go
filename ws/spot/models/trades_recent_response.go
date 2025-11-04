package models

// TradesRecentResponse represents global message '#/components/messages/tradesRecentResponse'
type TradesRecentResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		Id MessageID `json:"id,omitempty"` // id property
		IsBestMatch bool `json:"isBestMatch,omitempty"` // isBestMatch property
		IsBuyerMaker bool `json:"isBuyerMaker,omitempty"` // isBuyerMaker property
		Price string `json:"price,omitempty"` // price property
		Qty string `json:"qty,omitempty"` // qty property
		QuoteQty string `json:"quoteQty,omitempty"` // quoteQty property
		Time int64 `json:"time,omitempty"` // time property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


