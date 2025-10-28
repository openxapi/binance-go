package models

// ListenKeyExpiredEvent represents global message '#/components/messages/listenKeyExpiredEvent'
type ListenKeyExpiredEvent struct {
	EventType string `json:"e,omitempty"` // Event type
	EventTime int64 `json:"E,omitempty"` // Event time (timestamp)
	ListenKey string `json:"listenKey,omitempty"` // listen key
}


