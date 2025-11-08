package models

// AllOrderListsRequest represents global message '#/components/messages/allOrderListsRequest'
type AllOrderListsRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		EndTime int64 `json:"endTime,omitempty"`
		FromId int64 `json:"fromId,omitempty"` // Order list ID to begin at
		Limit int `json:"limit,omitempty"` // Default: 500; Maximum: 1000
		RecvWindow string `json:"recvWindow,omitempty"` // The value cannot be greater than 60000.  Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
		Signature string `json:"signature"`
		StartTime int64 `json:"startTime,omitempty"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


