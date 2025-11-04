package models

// SorOrderTestRequest represents global message '#/components/messages/sorOrderTestRequest'
type SorOrderTestRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"` // apiKey property
		Price float64 `json:"price,omitempty"` // price property
		Quantity float64 `json:"quantity,omitempty"` // quantity property
		Side string `json:"side,omitempty"` // side property
		Signature string `json:"signature,omitempty"` // signature property
		Symbol string `json:"symbol,omitempty"` // symbol property
		TimeInForce string `json:"timeInForce,omitempty"` // timeInForce property
		Timestamp int64 `json:"timestamp,omitempty"` // timestamp property
		Type string `json:"type,omitempty"` // type property
	} `json:"params,omitempty"` // params property
}


