package models

// ContractInfoEvent represents global message '#/components/messages/contractInfoEvent'
type ContractInfoEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	Pair string `json:"ps"` // Pair
	ContractType string `json:"ct"` // Contract type
	DeliveryDateTime int64 `json:"dt"` // Delivery date time
	OnboardDateTime int64 `json:"ot"` // Onboard date time
	ContractStatus string `json:"cs"` // Contract status
	BracketLevels []struct {
		NotionalBracket int `json:"bs"` // Notional bracket
		FloorNotionalOfThisBracket int `json:"bnf"` // Floor notional of this bracket
		CapNotionalOfThisBracket int `json:"bnc"` // Cap notional of this bracket
		MaintenanceMarginRate float64 `json:"mmr"` // Maintenance margin rate
		AuxiliaryNumberForQuickCalculation int `json:"cf"` // Auxiliary number for quick calculation
		MinLeverageForThisBracket int `json:"mi"` // Min leverage for this bracket
		MaxLeverageForThisBracket int `json:"ma"` // Max leverage for this bracket
	} `json:"bks"` // Bracket levels (only present when bracket updated)
}


