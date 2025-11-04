package models

// MyPreventedMatchesResponse represents global message '#/components/messages/myPreventedMatchesResponse'
type MyPreventedMatchesResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		MakerOrderId int64 `json:"makerOrderId,omitempty"` // makerOrderId property
		MakerPreventedQuantity string `json:"makerPreventedQuantity,omitempty"` // makerPreventedQuantity property
		MakerSymbol string `json:"makerSymbol,omitempty"` // makerSymbol property
		PreventedMatchId int64 `json:"preventedMatchId,omitempty"` // preventedMatchId property
		Price string `json:"price,omitempty"` // price property
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // selfTradePreventionMode property
		Symbol string `json:"symbol,omitempty"` // symbol property
		TakerOrderId int64 `json:"takerOrderId,omitempty"` // takerOrderId property
		TradeGroupId int64 `json:"tradeGroupId,omitempty"` // tradeGroupId property
		TransactTime int64 `json:"transactTime,omitempty"` // transactTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


