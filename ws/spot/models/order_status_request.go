package models

// OrderStatusRequest represents global message '#/components/messages/orderStatusRequest'
type OrderStatusRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		OrderId int64 `json:"orderId"` // Lookup order by orderId
		OrigClientOrderId string `json:"origClientOrderId,omitempty"` // Lookup order by clientOrderId
		RecvWindow string `json:"recvWindow,omitempty"` // The value cannot be greater than 60000. Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
		Signature string `json:"signature"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


