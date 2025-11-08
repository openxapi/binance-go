package models

// V2AccountPositionResponse represents global message '#/components/messages/v2AccountPositionResponse'
type V2AccountPositionResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		Adl int `json:"adl,omitempty"` // adl property
		AskNotional string `json:"askNotional,omitempty"` // askNotional property
		BidNotional string `json:"bidNotional,omitempty"` // bidNotional property
		BreakEvenPrice string `json:"breakEvenPrice,omitempty"` // breakEvenPrice property
		EntryPrice string `json:"entryPrice,omitempty"` // entryPrice property
		InitialMargin string `json:"initialMargin,omitempty"` // initialMargin property
		IsolatedMargin string `json:"isolatedMargin,omitempty"` // isolatedMargin property
		IsolatedWallet string `json:"isolatedWallet,omitempty"` // isolatedWallet property
		LiquidationPrice string `json:"liquidationPrice,omitempty"` // liquidationPrice property
		MaintMargin string `json:"maintMargin,omitempty"` // maintMargin property
		MarginAsset string `json:"marginAsset,omitempty"` // marginAsset property
		MarkPrice string `json:"markPrice,omitempty"` // markPrice property
		Notional string `json:"notional,omitempty"` // notional property
		OpenOrderInitialMargin string `json:"openOrderInitialMargin,omitempty"` // openOrderInitialMargin property
		PositionAmt string `json:"positionAmt,omitempty"` // positionAmt property
		PositionInitialMargin string `json:"positionInitialMargin,omitempty"` // positionInitialMargin property
		PositionSide string `json:"positionSide,omitempty"` // positionSide property
		Symbol string `json:"symbol,omitempty"` // symbol property
		UnrealizedProfit string `json:"unrealizedProfit,omitempty"` // unrealizedProfit property
		UpdateTime int64 `json:"updateTime,omitempty"` // updateTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


