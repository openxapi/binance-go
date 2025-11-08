package models

// OrderStatusResponse represents global message '#/components/messages/orderStatusResponse'
type OrderStatusResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		ClientOrderId string `json:"clientOrderId,omitempty"` // clientOrderId property
		CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"` // cummulativeQuoteQty property
		ExecutedQty string `json:"executedQty,omitempty"` // executedQty property
		IcebergQty string `json:"icebergQty,omitempty"` // icebergQty property
		IsWorking bool `json:"isWorking,omitempty"` // isWorking property
		OrderId int64 `json:"orderId,omitempty"` // orderId property
		OrderListId int64 `json:"orderListId,omitempty"` // orderListId property
		OrigQty string `json:"origQty,omitempty"` // origQty property
		OrigQuoteOrderQty string `json:"origQuoteOrderQty,omitempty"` // origQuoteOrderQty property
		PreventedMatchId int64 `json:"preventedMatchId,omitempty"` // preventedMatchId property
		PreventedQuantity string `json:"preventedQuantity,omitempty"` // preventedQuantity property
		Price string `json:"price,omitempty"` // price property
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // selfTradePreventionMode property
		Side string `json:"side,omitempty"` // side property
		Status string `json:"status,omitempty"` // status property
		StopPrice string `json:"stopPrice,omitempty"` // stopPrice property
		StrategyId int64 `json:"strategyId,omitempty"` // strategyId property
		StrategyType int `json:"strategyType,omitempty"` // strategyType property
		Symbol string `json:"symbol,omitempty"` // symbol property
		Time int64 `json:"time,omitempty"` // time property
		TimeInForce string `json:"timeInForce,omitempty"` // timeInForce property
		TrailingDelta int `json:"trailingDelta,omitempty"` // trailingDelta property
		TrailingTime int64 `json:"trailingTime,omitempty"` // trailingTime property
		Type string `json:"type,omitempty"` // type property
		UpdateTime int64 `json:"updateTime,omitempty"` // updateTime property
		WorkingTime int64 `json:"workingTime,omitempty"` // workingTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


