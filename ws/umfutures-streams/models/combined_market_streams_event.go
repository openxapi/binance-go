package models

// CombinedMarketStreamsEvent represents global message '#/components/messages/combinedMarketStreamsEvent'
type CombinedMarketStreamsEvent struct {
	Stream string `json:"stream"` // Original stream name that generated this event
	Data interface{} `json:"data"` // Original stream event data (unwrapped). This is exactly the same data that would be received when connecting to the individual stream directly. The event type can be determined by examining the 'stream' field. 
}


