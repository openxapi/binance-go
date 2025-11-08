package models

// AvgPriceRequest represents global message '#/components/messages/avgPriceRequest'
type AvgPriceRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Symbol string `json:"symbol"`
	} `json:"params,omitempty"` // params property
}


