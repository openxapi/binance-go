package models

// ExchangeInfoResponse represents global message '#/components/messages/exchangeInfoResponse'
type ExchangeInfoResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		ExchangeFilters []struct {		} `json:"exchangeFilters,omitempty"` // exchangeFilters property
		RateLimits []struct {
			Interval string `json:"interval,omitempty"` // interval property
			IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
			Limit int `json:"limit,omitempty"` // limit property
			RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
		} `json:"rateLimits,omitempty"` // rateLimits property
		ServerTime int64 `json:"serverTime,omitempty"` // serverTime property
		Sors []struct {
			BaseAsset string `json:"baseAsset,omitempty"` // baseAsset property
			Symbols []string `json:"symbols,omitempty"` // symbols property
		} `json:"sors,omitempty"` // sors property
		Symbols []struct {
			AllowAmend bool `json:"allowAmend,omitempty"` // allowAmend property
			AllowTrailingStop bool `json:"allowTrailingStop,omitempty"` // allowTrailingStop property
			AllowedSelfTradePreventionModes []string `json:"allowedSelfTradePreventionModes,omitempty"` // allowedSelfTradePreventionModes property
			BaseAsset string `json:"baseAsset,omitempty"` // baseAsset property
			BaseAssetPrecision int `json:"baseAssetPrecision,omitempty"` // baseAssetPrecision property
			BaseCommissionPrecision int `json:"baseCommissionPrecision,omitempty"` // baseCommissionPrecision property
			CancelReplaceAllowed bool `json:"cancelReplaceAllowed,omitempty"` // cancelReplaceAllowed property
			DefaultSelfTradePreventionMode string `json:"defaultSelfTradePreventionMode,omitempty"` // defaultSelfTradePreventionMode property
			Filters []struct {
				FilterType string `json:"filterType,omitempty"` // filterType property
				MaxPrice string `json:"maxPrice,omitempty"` // maxPrice property
				MinPrice string `json:"minPrice,omitempty"` // minPrice property
				TickSize string `json:"tickSize,omitempty"` // tickSize property
			} `json:"filters,omitempty"` // filters property
			IcebergAllowed bool `json:"icebergAllowed,omitempty"` // icebergAllowed property
			IsMarginTradingAllowed bool `json:"isMarginTradingAllowed,omitempty"` // isMarginTradingAllowed property
			IsSpotTradingAllowed bool `json:"isSpotTradingAllowed,omitempty"` // isSpotTradingAllowed property
			OcoAllowed bool `json:"ocoAllowed,omitempty"` // ocoAllowed property
			OrderTypes []string `json:"orderTypes,omitempty"` // orderTypes property
			OtoAllowed bool `json:"otoAllowed,omitempty"` // otoAllowed property
			PermissionSets [][]string `json:"permissionSets,omitempty"` // permissionSets property
			Permissions []struct {			} `json:"permissions,omitempty"` // permissions property
			QuoteAsset string `json:"quoteAsset,omitempty"` // quoteAsset property
			QuoteAssetPrecision int `json:"quoteAssetPrecision,omitempty"` // quoteAssetPrecision property
			QuoteCommissionPrecision int `json:"quoteCommissionPrecision,omitempty"` // quoteCommissionPrecision property
			QuoteOrderQtyMarketAllowed bool `json:"quoteOrderQtyMarketAllowed,omitempty"` // quoteOrderQtyMarketAllowed property
			QuotePrecision int `json:"quotePrecision,omitempty"` // quotePrecision property
			Status string `json:"status,omitempty"` // status property
			Symbol string `json:"symbol,omitempty"` // symbol property
		} `json:"symbols,omitempty"` // symbols property
		Timezone string `json:"timezone,omitempty"` // timezone property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


