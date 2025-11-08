package models

// MarkPriceEvent represents global message '#/components/messages/markPriceEvent'
type MarkPriceEvent struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	MarkPrice string `json:"p"` // Mark price
	EstimatedSettlePrice string `json:"P"` // Estimated settle price (only for quarterly contract)
	IndexPrice string `json:"i"` // Index price
	FundingRate string `json:"r"` // Funding rate (empty for delivery symbols)
	NextFundingTime int64 `json:"T"` // Next funding time (0 for delivery symbols)
}


