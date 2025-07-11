package models

import (
	"encoding/json"
)

// PartialDepthEvent represents PartialDepthEvent
type PartialDepthEvent struct {
	// Last update ID
	LastUpdateId int64 `json:"lastUpdateId,omitempty"`
	// Bids
	Bids [][]string `json:"bids,omitempty"`
	// Asks
	Asks [][]string `json:"asks,omitempty"`
}

// String returns string representation of PartialDepthEvent
func (s PartialDepthEvent) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


