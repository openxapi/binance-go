package models

// RiskLevelChangeEvent represents global message '#/components/messages/riskLevelChangeEvent'
type RiskLevelChangeEvent struct {
	EventType string `json:"e"` // Event Type
	EventTime int64 `json:"E"` // Event Time (milliseconds)
	MarginBalance string `json:"mb,omitempty"` // Margin Balance
	MaintenanceMargin string `json:"mm,omitempty"` // Maintenance Margin
	RiskLevel string `json:"s"` // Risk Level
}


