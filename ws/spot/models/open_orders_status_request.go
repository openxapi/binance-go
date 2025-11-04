package models

// OpenOrdersStatusRequest represents global message '#/components/messages/openOrdersStatusRequest'
type OpenOrdersStatusRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		RecvWindow string `json:"recvWindow,omitempty"` // The value cannot be greater than 60000.  Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
		Signature string `json:"signature"`
		Symbol string `json:"symbol,omitempty"` // If omitted, open orders for all symbols are returned
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


