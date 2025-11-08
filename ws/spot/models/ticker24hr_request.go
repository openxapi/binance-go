package models

// Ticker24hrRequest represents global message '#/components/messages/ticker24hrRequest'
type Ticker24hrRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Symbol string `json:"symbol,omitempty"` // Query ticker for a single symbol
		Symbols []string `json:"symbols,omitempty"` // Query ticker for multiple symbols
		Type string `json:"type,omitempty"` // Ticker type: FULL (default) or MINI
	} `json:"params,omitempty"` // params property
}


