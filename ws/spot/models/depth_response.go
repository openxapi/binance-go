package models

// DepthResponse represents global message '#/components/messages/depthResponse'
type DepthResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		Asks [][]string `json:"asks,omitempty"` // asks property
		Bids [][]string `json:"bids,omitempty"` // bids property
		LastUpdateId int64 `json:"lastUpdateId,omitempty"` // lastUpdateId property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


