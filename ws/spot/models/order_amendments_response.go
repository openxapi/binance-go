package models

// OrderAmendmentsResponse represents global message '#/components/messages/orderAmendmentsResponse'
type OrderAmendmentsResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		ExecutionId int64 `json:"executionId,omitempty"` // executionId property
		NewClientOrderId string `json:"newClientOrderId,omitempty"` // newClientOrderId property
		NewQty string `json:"newQty,omitempty"` // newQty property
		OrderId int64 `json:"orderId,omitempty"` // orderId property
		OrigClientOrderId string `json:"origClientOrderId,omitempty"` // origClientOrderId property
		OrigQty string `json:"origQty,omitempty"` // origQty property
		Symbol string `json:"symbol,omitempty"` // symbol property
		Time int64 `json:"time,omitempty"` // time property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


