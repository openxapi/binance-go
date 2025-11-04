package models

// OrderAmendKeepPriorityResponse represents global message '#/components/messages/orderAmendKeepPriorityResponse'
type OrderAmendKeepPriorityResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		AmendedOrder struct {
			ClientOrderId string `json:"clientOrderId,omitempty"` // clientOrderId property
			CumulativeQuoteQty string `json:"cumulativeQuoteQty,omitempty"` // cumulativeQuoteQty property
			ExecutedQty string `json:"executedQty,omitempty"` // executedQty property
			OrderId int64 `json:"orderId,omitempty"` // orderId property
			OrderListId int64 `json:"orderListId,omitempty"` // orderListId property
			OrigClientOrderId string `json:"origClientOrderId,omitempty"` // origClientOrderId property
			PreventedQty string `json:"preventedQty,omitempty"` // preventedQty property
			Price string `json:"price,omitempty"` // price property
			Qty string `json:"qty,omitempty"` // qty property
			QuoteOrderQty string `json:"quoteOrderQty,omitempty"` // quoteOrderQty property
			SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // selfTradePreventionMode property
			Side string `json:"side,omitempty"` // side property
			Status string `json:"status,omitempty"` // status property
			Symbol string `json:"symbol,omitempty"` // symbol property
			TimeInForce string `json:"timeInForce,omitempty"` // timeInForce property
			Type string `json:"type,omitempty"` // type property
			WorkingTime int64 `json:"workingTime,omitempty"` // workingTime property
		} `json:"amendedOrder,omitempty"` // amendedOrder property
		ExecutionId int64 `json:"executionId,omitempty"` // executionId property
		TransactTime int64 `json:"transactTime,omitempty"` // transactTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


