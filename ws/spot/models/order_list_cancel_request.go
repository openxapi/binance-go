package models

// OrderListCancelRequest represents global message '#/components/messages/orderListCancelRequest'
type OrderListCancelRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		ListClientOrderId string `json:"listClientOrderId,omitempty"` // Cancel order list by listClientId
		NewClientOrderId string `json:"newClientOrderId,omitempty"` // New ID for the canceled order list. Automatically generated if not sent
		OrderListId int64 `json:"orderListId"` // Cancel order list by orderListId
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000
		Signature string `json:"signature"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


