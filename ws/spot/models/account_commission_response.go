package models

// AccountCommissionResponse represents global message '#/components/messages/accountCommissionResponse'
type AccountCommissionResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		Discount struct {
			Discount string `json:"discount,omitempty"` // discount property
			DiscountAsset string `json:"discountAsset,omitempty"` // discountAsset property
			EnabledForAccount bool `json:"enabledForAccount,omitempty"` // enabledForAccount property
			EnabledForSymbol bool `json:"enabledForSymbol,omitempty"` // enabledForSymbol property
		} `json:"discount,omitempty"` // discount property
		SpecialCommission struct {
			Buyer string `json:"buyer,omitempty"` // buyer property
			Maker string `json:"maker,omitempty"` // maker property
			Seller string `json:"seller,omitempty"` // seller property
			Taker string `json:"taker,omitempty"` // taker property
		} `json:"specialCommission,omitempty"` // specialCommission property
		StandardCommission struct {
			Buyer string `json:"buyer,omitempty"` // buyer property
			Maker string `json:"maker,omitempty"` // maker property
			Seller string `json:"seller,omitempty"` // seller property
			Taker string `json:"taker,omitempty"` // taker property
		} `json:"standardCommission,omitempty"` // standardCommission property
		Symbol string `json:"symbol,omitempty"` // symbol property
		TaxCommission struct {
			Buyer string `json:"buyer,omitempty"` // buyer property
			Maker string `json:"maker,omitempty"` // maker property
			Seller string `json:"seller,omitempty"` // seller property
			Taker string `json:"taker,omitempty"` // taker property
		} `json:"taxCommission,omitempty"` // taxCommission property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


