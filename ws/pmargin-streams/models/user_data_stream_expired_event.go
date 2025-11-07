package models

// UserDataStreamExpiredEvent represents global message '#/components/messages/userDataStreamExpiredEvent'
type UserDataStreamExpiredEvent struct {
	EventType string `json:"e,omitempty"` // Event Type
	EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
}


