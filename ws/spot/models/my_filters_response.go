package models

// MyFiltersResponse represents global message '#/components/messages/myFiltersResponse'
type MyFiltersResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	Result struct {
		AssetFilters []struct {
			Asset string `json:"asset,omitempty"` // asset property
			FilterType string `json:"filterType,omitempty"` // filterType property
			Limit string `json:"limit,omitempty"` // limit property
		} `json:"assetFilters,omitempty"` // assetFilters property
		ExchangeFilters []struct {
			FilterType string `json:"filterType,omitempty"` // filterType property
			MaxNumOrders int `json:"maxNumOrders,omitempty"` // maxNumOrders property
		} `json:"exchangeFilters,omitempty"` // exchangeFilters property
		SymbolFilters []struct {
			FilterType string `json:"filterType,omitempty"` // filterType property
			MaxNumOrderLists int `json:"maxNumOrderLists,omitempty"` // maxNumOrderLists property
		} `json:"symbolFilters,omitempty"` // symbolFilters property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


