package models

// OrderCancelReplaceResponse represents global message '#/components/messages/orderCancelReplaceResponse'
type OrderCancelReplaceResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		CancelResponse struct {
			ClientOrderId string `json:"clientOrderId,omitempty"` // clientOrderId property
			CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"` // cummulativeQuoteQty property
			ExecutedQty string `json:"executedQty,omitempty"` // executedQty property
			OrderId int64 `json:"orderId,omitempty"` // orderId property
			OrderListId int64 `json:"orderListId,omitempty"` // orderListId property
			OrigClientOrderId string `json:"origClientOrderId,omitempty"` // origClientOrderId property
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
		} `json:"cancelResponse,omitempty"` // cancelResponse property
		CancelResult string `json:"cancelResult,omitempty"` // cancelResult property
		NewOrderResponse struct {
			ClientOrderId string `json:"clientOrderId,omitempty"` // clientOrderId property
			CummulativeQuoteQty string `json:"cummulativeQuoteQty,omitempty"` // cummulativeQuoteQty property
			ExecutedQty string `json:"executedQty,omitempty"` // executedQty property
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
		} `json:"newOrderResponse,omitempty"` // newOrderResponse property
		NewOrderResult string `json:"newOrderResult,omitempty"` // newOrderResult property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


