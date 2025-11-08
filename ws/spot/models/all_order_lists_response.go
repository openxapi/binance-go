package models

// AllOrderListsResponse represents global message '#/components/messages/allOrderListsResponse'
type AllOrderListsResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		ContingencyType string `json:"contingencyType,omitempty"` // contingencyType property
		ListClientOrderId string `json:"listClientOrderId,omitempty"` // listClientOrderId property
		ListOrderStatus string `json:"listOrderStatus,omitempty"` // listOrderStatus property
		ListStatusType string `json:"listStatusType,omitempty"` // listStatusType property
		OrderListId int64 `json:"orderListId,omitempty"` // orderListId property
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


