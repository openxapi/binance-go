package models

// ContractInfoEvent represents global message '#/components/messages/contractInfoEvent'
type ContractInfoEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	Pair string `json:"ps"` // Pair
	ContractType string `json:"ct"` // Contract type
	DeliveryDateTime int64 `json:"dt"` // Delivery date time
	OnboardDateTime int64 `json:"ot"` // onboard date time
	ContractStatus string `json:"cs"` // Contract status
	BracketInformation []struct {
		NotionalBracket int `json:"bs"` // Notional bracket
		FloorNotionalOfThisBracket int64 `json:"bnf"` // Floor notional of this bracket
		CapNotionalOfThisBracket int64 `json:"bnc"` // Cap notional of this bracket
		MaintenanceRatioForThisBracket float64 `json:"mmr"` // Maintenance ratio for this bracket
		AuxiliaryCalculationNumber int64 `json:"cf"` // Auxiliary calculation number
		MinimumLeverage int `json:"mi"` // Minimum leverage
		MaximumLeverage int `json:"ma"` // Maximum leverage
	} `json:"bks"` // Bracket information
}


