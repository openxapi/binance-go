package models

// AccountPositionResponse represents global message '#/components/messages/accountPositionResponse'
type AccountPositionResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		BreakEvenPrice string `json:"breakEvenPrice,omitempty"` // breakEvenPrice property
		EntryPrice string `json:"entryPrice,omitempty"` // entryPrice property
		IsAutoAddMargin string `json:"isAutoAddMargin,omitempty"` // isAutoAddMargin property
		IsolatedMargin string `json:"isolatedMargin,omitempty"` // isolatedMargin property
		IsolatedWallet string `json:"isolatedWallet,omitempty"` // isolatedWallet property
		Leverage string `json:"leverage,omitempty"` // leverage property
		LiquidationPrice string `json:"liquidationPrice,omitempty"` // liquidationPrice property
		MarginType string `json:"marginType,omitempty"` // marginType property
		MarkPrice string `json:"markPrice,omitempty"` // markPrice property
		MaxQty string `json:"maxQty,omitempty"` // maxQty property
		NotionalValue string `json:"notionalValue,omitempty"` // notionalValue property
		PositionAmt string `json:"positionAmt,omitempty"` // positionAmt property
		PositionSide string `json:"positionSide,omitempty"` // positionSide property
		Symbol string `json:"symbol,omitempty"` // symbol property
		UnRealizedProfit string `json:"unRealizedProfit,omitempty"` // unRealizedProfit property
		UpdateTime int64 `json:"updateTime,omitempty"` // updateTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


