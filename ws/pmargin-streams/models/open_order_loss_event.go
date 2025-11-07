package models

// OpenOrderLossEvent represents global message '#/components/messages/openOrderLossEvent'
type OpenOrderLossEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
	ArrayOfLossUpdates []struct {
		Asset string `json:"a,omitempty"` // Asset
		LossAmount string `json:"o,omitempty"` // Loss Amount (negative number)
	} `json:"O,omitempty"` // Array of loss updates
}


