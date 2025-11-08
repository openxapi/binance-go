package models

// TickerPriceRequest represents global message '#/components/messages/tickerPriceRequest'
type TickerPriceRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Symbol string `json:"symbol,omitempty"` // Query price for a single symbol
		Symbols []string `json:"symbols,omitempty"` // Query price for multiple symbols
	} `json:"params,omitempty"` // params property
}


