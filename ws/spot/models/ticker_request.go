package models

// TickerRequest represents global message '#/components/messages/tickerRequest'
type TickerRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Symbol string `json:"symbol,omitempty"` // Query ticker of a single symbol
		Symbols []string `json:"symbols,omitempty"` // Query ticker for multiple symbols
		Type string `json:"type,omitempty"` // Ticker type: FULL (default) or MINI
		WindowSize string `json:"windowSize,omitempty"` // Default 1d
	} `json:"params,omitempty"` // params property
}


