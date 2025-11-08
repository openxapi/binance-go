package models

// OrderAmendmentsRequest represents global message '#/components/messages/orderAmendmentsRequest'
type OrderAmendmentsRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		FromExecutionId int64 `json:"fromExecutionId,omitempty"`
		Limit int `json:"limit,omitempty"` // Default:500; Maximum: 1000
		OrderId int64 `json:"orderId"`
		RecvWindow string `json:"recvWindow,omitempty"` // The value cannot be greater than 60000.  Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
		Signature string `json:"signature,omitempty"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


