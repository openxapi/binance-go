package models

// V2AccountBalanceResponse represents global message '#/components/messages/v2AccountBalanceResponse'
type V2AccountBalanceResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		AccountAlias string `json:"accountAlias,omitempty"` // accountAlias property
		Asset string `json:"asset,omitempty"` // asset property
		AvailableBalance string `json:"availableBalance,omitempty"` // availableBalance property
		Balance string `json:"balance,omitempty"` // balance property
		CrossUnPnl string `json:"crossUnPnl,omitempty"` // crossUnPnl property
		CrossWalletBalance string `json:"crossWalletBalance,omitempty"` // crossWalletBalance property
		MarginAvailable bool `json:"marginAvailable,omitempty"` // marginAvailable property
		MaxWithdrawAmount string `json:"maxWithdrawAmount,omitempty"` // maxWithdrawAmount property
		UpdateTime int64 `json:"updateTime,omitempty"` // updateTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


