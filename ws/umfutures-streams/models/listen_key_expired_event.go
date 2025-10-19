package models

// ListenKeyExpiredEvent represents global message '#/components/messages/listenKeyExpiredEvent'
type ListenKeyExpiredEvent struct {
	Event struct {
		EventType string `json:"e,omitempty"` // Event type
		EventTime int64 `json:"E,omitempty"` // Event time (timestamp)
		TheExpiredListenKey string `json:"listenKey,omitempty"` // The expired listen key
	} `json:"event,omitempty"`
}


