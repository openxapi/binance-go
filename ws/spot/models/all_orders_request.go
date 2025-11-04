package models

// AllOrdersRequest represents global message '#/components/messages/allOrdersRequest'
type AllOrdersRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		EndTime int64 `json:"endTime,omitempty"`
		Limit int `json:"limit,omitempty"` // Default: 500; Maximum: 1000
		OrderId int64 `json:"orderId,omitempty"` // Order ID to begin at
		RecvWindow string `json:"recvWindow,omitempty"` // The value cannot be greater than 60000.  Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
		Signature string `json:"signature"`
		StartTime int64 `json:"startTime,omitempty"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


