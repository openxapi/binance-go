package models

// SorOrderPlaceResponse represents global message '#/components/messages/sorOrderPlaceResponse'
type SorOrderPlaceResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		ClientOrderId string `json:"clientOrderId,omitempty"` // clientOrderId property
		CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"` // cummulativeQuoteQty property
		ExecutedQty string `json:"executedQty,omitempty"` // executedQty property
		Fills []struct {
			AllocId int64 `json:"allocId,omitempty"` // allocId property
			Commission string `json:"commission,omitempty"` // commission property
			CommissionAsset string `json:"commissionAsset,omitempty"` // commissionAsset property
			MatchType string `json:"matchType,omitempty"` // matchType property
			Price string `json:"price,omitempty"` // price property
			Qty string `json:"qty,omitempty"` // qty property
			TradeId int64 `json:"tradeId,omitempty"` // tradeId property
		} `json:"fills,omitempty"` // fills property
		OrderId int64 `json:"orderId,omitempty"` // orderId property
		OrderListId int64 `json:"orderListId,omitempty"` // orderListId property
		OrigQty string `json:"origQty,omitempty"` // origQty property
		OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"` // origQuoteOrderQty property
		Price string `json:"price,omitempty"` // price property
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // selfTradePreventionMode property
		Side string `json:"side,omitempty"` // side property
		Status string `json:"status,omitempty"` // status property
		Symbol string `json:"symbol,omitempty"` // symbol property
		TimeInForce string `json:"timeInForce,omitempty"` // timeInForce property
		TransactTime int64 `json:"transactTime,omitempty"` // transactTime property
		Type string `json:"type,omitempty"` // type property
		UsedSor bool `json:"usedSor,omitempty"` // usedSor property
		WorkingFloor string `json:"workingFloor,omitempty"` // workingFloor property
		WorkingTime int64 `json:"workingTime,omitempty"` // workingTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


