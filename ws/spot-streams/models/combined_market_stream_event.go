package models

// CombinedMarketStreamEvent represents global message '#/components/messages/combinedMarketStreamEvent'
type CombinedMarketStreamEvent struct {
	Stream string `json:"stream"` // Original stream name that generated this event
	Data interface{} `json:"data"` // Original stream event data (unwrapped). This is exactly the same data that would be received when connecting to the individual stream directly. The event type can be determined by examining the 'stream' field.  For all-market streams (!ticker@arr, !miniTicker@arr), the data will be an array of events instead of a single event object. 
}


