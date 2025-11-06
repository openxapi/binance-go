package models

// TickerBookRequest represents global message '#/components/messages/tickerBookRequest'
type TickerBookRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Symbol string `json:"symbol,omitempty"`
	} `json:"params,omitempty"` // params property
}


