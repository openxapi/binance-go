package models

// TickerBookRequest represents global message '#/components/messages/tickerBookRequest'
type TickerBookRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Symbol string `json:"symbol,omitempty"` // Query ticker for a single symbol
		Symbols []string `json:"symbols,omitempty"` // Query ticker for multiple symbols
	} `json:"params,omitempty"` // params property
}


