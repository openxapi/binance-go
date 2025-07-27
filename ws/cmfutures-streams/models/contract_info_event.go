package models

import (
	"encoding/json"
)

// ContractInfoEvent represents ContractInfoEvent
type ContractInfoEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Symbol
	Symbol string `json:"s,omitempty"`
	// Pair
	Pair string `json:"ps,omitempty"`
	// Contract type
	ContractType string `json:"ct,omitempty"`
	// Delivery date time
	DeliveryDateTime int64 `json:"dt,omitempty"`
	// Onboard date time
	OnboardDateTime int64 `json:"ot,omitempty"`
	// Contract status
	ContractStatus string `json:"cs,omitempty"`
	// Bracket information (optional)
	BracketInformation []ContractInfoEventBracketInformationItem `json:"bks,omitempty"`
}

// ContractInfoEventBracketInformationItem represents the bracketinformation item details
type ContractInfoEventBracketInformationItem struct {
	// Notional bracket
	NotionalBracket int `json:"bs,omitempty"`
	// Floor notional
	FloorNotional int `json:"bnf,omitempty"`
	// Cap notional
	CapNotional int `json:"bnc,omitempty"`
	// Maintenance ratio
	MaintenanceRatio float64 `json:"mmr,omitempty"`
	// Auxiliary calculation number
	AuxiliaryCalculationNumber int `json:"cf,omitempty"`
	// Minimum leverage
	MinimumLeverage int `json:"mi,omitempty"`
	// Maximum leverage
	MaximumLeverage int `json:"ma,omitempty"`
}

// String returns string representation of ContractInfoEvent
func (s ContractInfoEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


