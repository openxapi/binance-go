package models

// ListStatusEvent represents global message '#/components/messages/listStatusEvent'
type ListStatusEvent struct {
	Event struct {
		ListClientOrderID string `json:"C,omitempty"` // List Client Order ID
		EventTime int64 `json:"E"` // Event Time (milliseconds)
		ListOrderStatus string `json:"L,omitempty"` // List Order Status
		ArrayOfOrdersInTheList []struct {
			ClientOrderId string `json:"c,omitempty"` // ClientOrderId
			OrderId int64 `json:"i,omitempty"` // OrderId
			Symbol string `json:"s,omitempty"` // Symbol
		} `json:"O,omitempty"` // Array of orders in the list
		TransactionTime int64 `json:"T"` // Transaction Time (milliseconds)
		ContingencyType string `json:"c,omitempty"` // Contingency Type
		EventType string `json:"e"` // Event Type
		OrderListId int64 `json:"g,omitempty"` // OrderListId
		ListStatusType string `json:"l,omitempty"` // List Status Type
		ListRejectReason string `json:"r,omitempty"` // List Reject Reason
		Symbol string `json:"s,omitempty"` // Symbol
	} `json:"event,omitempty"`
	SubscriptionId int64 `json:"subscriptionId,omitempty"` // subscriptionId
}


