package models

// TradesRecentRequest represents global message '#/components/messages/tradesRecentRequest'
type TradesRecentRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Limit int `json:"limit,omitempty"` // Default: 500; Maximum: 1000
		Symbol string `json:"symbol"`
	} `json:"params,omitempty"` // params property
}


