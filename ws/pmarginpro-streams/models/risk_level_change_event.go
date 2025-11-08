package models

// RiskLevelChangeEvent represents global message '#/components/messages/riskLevelChangeEvent'
type RiskLevelChangeEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	UniMMRLevel string `json:"u,omitempty"` // UniMMR level
	RiskLevelStatus string `json:"s,omitempty"` // Risk level status
	AccountEquityInUSD string `json:"eq,omitempty"` // Account equity in USD
	ActualEquityWithoutCollateralRateInUSD string `json:"ae,omitempty"` // Actual equity without collateral rate in USD
	TotalMaintenanceMarginInUSD string `json:"m,omitempty"` // Total maintenance margin in USD
}


