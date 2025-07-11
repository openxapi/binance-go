package models

import (
	"encoding/json"
)

// ExchangeInfoResponseRateLimitsItem represents a nested object structure
type ExchangeInfoResponseRateLimitsItem struct {
	// count property (example: 20)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 6000)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// ExchangeInfoResponseResult represents a nested object structure
type ExchangeInfoResponseResult struct {
	// exchangeFilters property
	ExchangeFilters []interface{} `json:"exchangeFilters,omitempty"`
	// rateLimits property
	RateLimits []ExchangeInfoResponseResultRateLimitsItem `json:"rateLimits,omitempty"`
	// serverTime property
	ServerTime int64 `json:"serverTime,omitempty"`
	// sors property
	Sors []ExchangeInfoResponseResultSorsItem `json:"sors,omitempty"`
	// symbols property
	Symbols []ExchangeInfoResponseResultSymbolsItem `json:"symbols,omitempty"`
	// timezone property
	Timezone string `json:"timezone,omitempty"`
}

// ExchangeInfoResponseResultRateLimitsItem represents a nested object structure
type ExchangeInfoResponseResultRateLimitsItem struct {
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 6000)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// ExchangeInfoResponseResultSorsItem represents a nested object structure
type ExchangeInfoResponseResultSorsItem struct {
	// baseAsset property (example: "BTC")
	BaseAsset string `json:"baseAsset,omitempty"`
	// symbols property
	Symbols []string `json:"symbols,omitempty"`
}

// ExchangeInfoResponseResultSymbolsItem represents a nested object structure
type ExchangeInfoResponseResultSymbolsItem struct {
	// allowAmend property (example: false)
	AllowAmend bool `json:"allowAmend,omitempty"`
	// allowTrailingStop property (example: true)
	AllowTrailingStop bool `json:"allowTrailingStop,omitempty"`
	// allowedSelfTradePreventionModes property
	AllowedSelfTradePreventionModes []string `json:"allowedSelfTradePreventionModes,omitempty"`
	// baseAsset property (example: "BNB")
	BaseAsset string `json:"baseAsset,omitempty"`
	// baseAssetPrecision property (example: 8)
	BaseAssetPrecision int64 `json:"baseAssetPrecision,omitempty"`
	// baseCommissionPrecision property (example: 8)
	BaseCommissionPrecision int64 `json:"baseCommissionPrecision,omitempty"`
	// cancelReplaceAllowed property (example: true)
	CancelReplaceAllowed bool `json:"cancelReplaceAllowed,omitempty"`
	// defaultSelfTradePreventionMode property (example: "NONE")
	DefaultSelfTradePreventionMode string `json:"defaultSelfTradePreventionMode,omitempty"`
	// filters property
	Filters []ExchangeInfoResponseResultSymbolsItemFiltersItem `json:"filters,omitempty"`
	// icebergAllowed property (example: true)
	IcebergAllowed bool `json:"icebergAllowed,omitempty"`
	// isMarginTradingAllowed property (example: true)
	IsMarginTradingAllowed bool `json:"isMarginTradingAllowed,omitempty"`
	// isSpotTradingAllowed property (example: true)
	IsSpotTradingAllowed bool `json:"isSpotTradingAllowed,omitempty"`
	// ocoAllowed property (example: true)
	OcoAllowed bool `json:"ocoAllowed,omitempty"`
	// orderTypes property
	OrderTypes []string `json:"orderTypes,omitempty"`
	// otoAllowed property (example: true)
	OtoAllowed bool `json:"otoAllowed,omitempty"`
	// permissionSets property
	PermissionSets [][]string `json:"permissionSets,omitempty"`
	// permissions property
	Permissions []interface{} `json:"permissions,omitempty"`
	// quoteAsset property (example: "BTC")
	QuoteAsset string `json:"quoteAsset,omitempty"`
	// quoteAssetPrecision property (example: 8)
	QuoteAssetPrecision int64 `json:"quoteAssetPrecision,omitempty"`
	// quoteCommissionPrecision property (example: 8)
	QuoteCommissionPrecision int64 `json:"quoteCommissionPrecision,omitempty"`
	// quoteOrderQtyMarketAllowed property (example: true)
	QuoteOrderQtyMarketAllowed bool `json:"quoteOrderQtyMarketAllowed,omitempty"`
	// quotePrecision property (example: 8)
	QuotePrecision int64 `json:"quotePrecision,omitempty"`
	// status property (example: "TRADING")
	Status string `json:"status,omitempty"`
	// symbol property (example: "BNBBTC")
	Symbol string `json:"symbol,omitempty"`
}

// ExchangeInfoResponseResultSymbolsItemFiltersItem represents a nested object structure
type ExchangeInfoResponseResultSymbolsItemFiltersItem struct {
	// filterType property (example: "PRICE_FILTER")
	FilterType string `json:"filterType,omitempty"`
	// maxPrice property (example: "100000.00000000")
	MaxPrice string `json:"maxPrice,omitempty"`
	// minPrice property (example: "0.00000100")
	MinPrice string `json:"minPrice,omitempty"`
	// tickSize property (example: "0.00000100")
	TickSize string `json:"tickSize,omitempty"`
}

// ExchangeInfoResponse - Receive response from exchangeInfo
// Message name: Exchange information Response
type ExchangeInfoResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []ExchangeInfoResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result ExchangeInfoResponseResult `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of ExchangeInfoResponse
func (s ExchangeInfoResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


