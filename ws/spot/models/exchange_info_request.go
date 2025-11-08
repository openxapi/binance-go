package models

// ExchangeInfoRequest represents global message '#/components/messages/exchangeInfoRequest'
type ExchangeInfoRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		Permissions []string `json:"permissions,omitempty"` // Filter symbols by permissions
		ShowPermissionSets bool `json:"showPermissionSets,omitempty"` // Controls whether the content of the permissionSets field is populated or not. Defaults to true.
		Symbol string `json:"symbol,omitempty"` // Describe a single symbol
		SymbolStatus string `json:"symbolStatus,omitempty"` // Filters symbols that have this tradingStatus. Valid values: TRADING, HALT, BREAK  Cannot be used in combination with symbol or symbols
		Symbols []string `json:"symbols,omitempty"` // Describe multiple symbols
	} `json:"params,omitempty"` // params property
}


