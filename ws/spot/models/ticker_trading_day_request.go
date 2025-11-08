package models

// TickerTradingDayRequest represents global message '#/components/messages/tickerTradingDayRequest'
type TickerTradingDayRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Symbol string `json:"symbol,omitempty"` // Query ticker of a single symbol
		Symbols []string `json:"symbols,omitempty"` // Query ticker for multiple symbols
		TimeZone string `json:"timeZone,omitempty"` // Default: 0 (UTC)
		Type string `json:"type,omitempty"` // Supported values: FULL or MINI. If none provided, the default is FULL
	} `json:"params,omitempty"` // params property
}


