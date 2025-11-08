package models

// TickerBookResponse represents global message '#/components/messages/tickerBookResponse'
type TickerBookResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		AskPrice string `json:"askPrice,omitempty"` // askPrice property
		AskQty string `json:"askQty,omitempty"` // askQty property
		BidPrice string `json:"bidPrice,omitempty"` // bidPrice property
		BidQty string `json:"bidQty,omitempty"` // bidQty property
		Symbol string `json:"symbol,omitempty"` // symbol property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


