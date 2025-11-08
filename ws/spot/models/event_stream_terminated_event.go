package models

// EventStreamTerminatedEvent represents global message '#/components/messages/eventStreamTerminatedEvent'
type EventStreamTerminatedEvent struct {
	Event struct {
		EventTime int64 `json:"E"` // Event Time (milliseconds)
		EventType string `json:"e"` // Event Type
	} `json:"event,omitempty"`
	SubscriptionId int64 `json:"subscriptionId,omitempty"` // subscriptionId
}


