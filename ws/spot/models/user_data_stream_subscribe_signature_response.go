package models

// UserDataStreamSubscribeSignatureResponse represents global message '#/components/messages/userDataStreamSubscribeSignatureResponse'
type UserDataStreamSubscribeSignatureResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	Result struct {
		SubscriptionId int64 `json:"subscriptionId,omitempty"` // subscriptionId property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


