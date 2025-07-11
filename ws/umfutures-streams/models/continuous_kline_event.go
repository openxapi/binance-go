package models

import (
	"encoding/json"
)

// ContinuousKlineEvent represents ContinuousKlineEvent
type ContinuousKlineEvent struct {
	// Event type
	EventType string `json:"e,omitempty"`
	// Event time
	EventTime int64 `json:"E,omitempty"`
	// Pair
	Pair string `json:"ps,omitempty"`
	// Contract type
	ContractType string `json:"ct,omitempty"`
	// Kline data
	Kline interface{} `json:"k,omitempty"`
}

// String returns string representation of ContinuousKlineEvent
func (s ContinuousKlineEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


