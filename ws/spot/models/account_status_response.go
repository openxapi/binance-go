package models

// AccountStatusResponse represents global message '#/components/messages/accountStatusResponse'
type AccountStatusResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		AccountType string `json:"accountType,omitempty"` // accountType property
		Balances []struct {
			Asset string `json:"asset,omitempty"` // asset property
			Free string `json:"free,omitempty"` // free property
			Locked string `json:"locked,omitempty"` // locked property
		} `json:"balances,omitempty"` // balances property
		Brokered bool `json:"brokered,omitempty"` // brokered property
		BuyerCommission int `json:"buyerCommission,omitempty"` // buyerCommission property
		CanDeposit bool `json:"canDeposit,omitempty"` // canDeposit property
		CanTrade bool `json:"canTrade,omitempty"` // canTrade property
		CanWithdraw bool `json:"canWithdraw,omitempty"` // canWithdraw property
		CommissionRates struct {
			Buyer string `json:"buyer,omitempty"` // buyer property
			Maker string `json:"maker,omitempty"` // maker property
			Seller string `json:"seller,omitempty"` // seller property
			Taker string `json:"taker,omitempty"` // taker property
		} `json:"commissionRates,omitempty"` // commissionRates property
		MakerCommission int `json:"makerCommission,omitempty"` // makerCommission property
		Permissions []string `json:"permissions,omitempty"` // permissions property
		PreventSor bool `json:"preventSor,omitempty"` // preventSor property
		RequireSelfTradePrevention bool `json:"requireSelfTradePrevention,omitempty"` // requireSelfTradePrevention property
		SellerCommission int `json:"sellerCommission,omitempty"` // sellerCommission property
		TakerCommission int `json:"takerCommission,omitempty"` // takerCommission property
		Uid int `json:"uid,omitempty"` // uid property
		UpdateTime int64 `json:"updateTime,omitempty"` // updateTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


