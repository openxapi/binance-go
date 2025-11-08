package models

// UserDataStreamUnsubscribeRequest represents global message '#/components/messages/userDataStreamUnsubscribeRequest'
type UserDataStreamUnsubscribeRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		SubscriptionId int64 `json:"subscriptionId,omitempty"` // When called with no parameter, this will close all subscriptions. When called with the subscriptionId parameter, this will attempt to close the subscription with that subscription id, if it exists.
	} `json:"params,omitempty"` // params property
}


