package models

// OrderListStatusRequest represents global message '#/components/messages/orderListStatusRequest'
type OrderListStatusRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		OrderListId int64 `json:"orderListId,omitempty"` // Query order list by orderListId.orderListId or origClientOrderId must be provided.
		OrigClientOrderId string `json:"origClientOrderId,omitempty"` // Query order list by listClientOrderId.orderListId or origClientOrderId must be provided.
		RecvWindow string `json:"recvWindow,omitempty"` // The value cannot be greater than 60000. Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
		Signature string `json:"signature,omitempty"`
		Timestamp int64 `json:"timestamp,omitempty"`
	} `json:"params,omitempty"` // params property
}


