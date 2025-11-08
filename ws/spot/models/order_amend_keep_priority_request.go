package models

// OrderAmendKeepPriorityRequest represents global message '#/components/messages/orderAmendKeepPriorityRequest'
type OrderAmendKeepPriorityRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		NewClientOrderId string `json:"newClientOrderId,omitempty"` // The new client order ID for the order after being amended.  If not sent, one will be randomly generated.  It is possible to reuse the current clientOrderId by sending it as the newClientOrderId.
		NewQty string `json:"newQty"` // newQty must be greater than 0 and less than the order's quantity.
		OrderId int64 `json:"orderId,omitempty"` // orderId or origClientOrderId must be sent
		OrigClientOrderId string `json:"origClientOrderId,omitempty"` // orderId or origClientOrderId must be sent
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000.
		Signature string `json:"signature,omitempty"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


