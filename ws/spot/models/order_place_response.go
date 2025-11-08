package models

// OrderPlaceResponse represents global message '#/components/messages/orderPlaceResponse'
type OrderPlaceResponse struct {
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
		OrderId int64 `json:"orderId,omitempty"` // orderId property
		OrderListId int64 `json:"orderListId,omitempty"` // orderListId property
		Symbol string `json:"symbol,omitempty"` // symbol property
		TransactTime int64 `json:"transactTime,omitempty"` // transactTime property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


