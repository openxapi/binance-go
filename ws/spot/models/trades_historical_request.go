package models

// TradesHistoricalRequest represents global message '#/components/messages/tradesHistoricalRequest'
type TradesHistoricalRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		FromId int64 `json:"fromId,omitempty"` // Trade ID to begin at
		Limit int `json:"limit,omitempty"` // Default: 500; Maximum: 1000
		Symbol string `json:"symbol"`
	} `json:"params,omitempty"` // params property
}


