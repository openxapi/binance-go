package models

// AllMarkPricesEventItem is the item type for AllMarkPricesEvent
type AllMarkPricesEventItem struct {
	EventType string `json:"e"` // Event type
	EventTime int64 `json:"E"` // Event time
	Symbol string `json:"s"` // Symbol
	MarkPrice string `json:"p"` // Mark price
	EstimatedSettlePrice string `json:"P"` // Estimated settle price (only for quarterly contract)
	IndexPrice string `json:"i"` // Index price
	FundingRate string `json:"r"` // Funding rate
	NextFundingTime int64 `json:"T"` // Next funding time
}

// AllMarkPricesEvent is an array of AllMarkPricesEventItem
type AllMarkPricesEvent []AllMarkPricesEventItem

