package models

// TradesAggregateRequest represents global message '#/components/messages/tradesAggregateRequest'
type TradesAggregateRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		EndTime int64 `json:"endTime,omitempty"`
		FromId int64 `json:"fromId,omitempty"` // Aggregate trade ID to begin at
		Limit int `json:"limit,omitempty"` // Default: 500; Maximum: 1000
		StartTime int64 `json:"startTime,omitempty"`
		Symbol string `json:"symbol"`
	} `json:"params,omitempty"` // params property
}


