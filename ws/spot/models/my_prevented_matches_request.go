package models

// MyPreventedMatchesRequest represents global message '#/components/messages/myPreventedMatchesRequest'
type MyPreventedMatchesRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		FromPreventedMatchId int64 `json:"fromPreventedMatchId,omitempty"`
		Limit int `json:"limit,omitempty"` // Default: 500; Maximum: 1000
		OrderId int64 `json:"orderId,omitempty"`
		PreventedMatchId int64 `json:"preventedMatchId,omitempty"`
		RecvWindow string `json:"recvWindow,omitempty"` // The value cannot be greater than 60000.  Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
		Signature string `json:"signature,omitempty"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


