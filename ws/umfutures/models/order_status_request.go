package models

// OrderStatusRequest represents global message '#/components/messages/orderStatusRequest'
type OrderStatusRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		OrderId int64 `json:"orderId,omitempty"`
		OrigClientOrderId string `json:"origClientOrderId,omitempty"`
		RecvWindow int64 `json:"recvWindow,omitempty"`
		Signature string `json:"signature,omitempty"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


