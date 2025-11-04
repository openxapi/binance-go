package models

// OrderCancelRequest represents global message '#/components/messages/orderCancelRequest'
type OrderCancelRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		CancelRestrictions string `json:"cancelRestrictions,omitempty"` // Supported values: ONLY_NEW - Cancel will succeed if the order status is NEW. ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED.
		NewClientOrderId string `json:"newClientOrderId,omitempty"` // New ID for the canceled order. Automatically generated if not sent
		OrderId int64 `json:"orderId"` // Cancel order by orderId
		OrigClientOrderId string `json:"origClientOrderId,omitempty"` // Cancel order by clientOrderId
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000
		Signature string `json:"signature"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


