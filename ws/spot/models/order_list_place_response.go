package models

// OrderListPlaceResponse represents global message '#/components/messages/orderListPlaceResponse'
type OrderListPlaceResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		ContingencyType string `json:"contingencyType,omitempty"` // contingencyType property
		ListClientOrderId string `json:"listClientOrderId,omitempty"` // listClientOrderId property
		ListOrderStatus string `json:"listOrderStatus,omitempty"` // listOrderStatus property
		ListStatusType string `json:"listStatusType,omitempty"` // listStatusType property
		OrderListId int64 `json:"orderListId,omitempty"` // orderListId property
		OrderReports []struct {
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
			StopPrice string `json:"stopPrice,omitempty"` // stopPrice property
			Symbol string `json:"symbol,omitempty"` // symbol property
			TimeInForce string `json:"timeInForce,omitempty"` // timeInForce property
			TransactTime int64 `json:"transactTime,omitempty"` // transactTime property
			Type string `json:"type,omitempty"` // type property
			WorkingTime int64 `json:"workingTime,omitempty"` // workingTime property
		} `json:"orderReports,omitempty"` // orderReports property
		Orders []struct {
			ClientOrderId string `json:"clientOrderId,omitempty"` // clientOrderId property
			OrderId int64 `json:"orderId,omitempty"` // orderId property
			Symbol string `json:"symbol,omitempty"` // symbol property
		} `json:"orders,omitempty"` // orders property
		Symbol string `json:"symbol,omitempty"` // symbol property
		TransactionTime int64 `json:"transactionTime,omitempty"` // transactionTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


