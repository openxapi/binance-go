package models

// OrderModifyRequest represents global message '#/components/messages/orderModifyRequest'
type OrderModifyRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		OrderId int64 `json:"orderId,omitempty"`
		OrigClientOrderId string `json:"origClientOrderId,omitempty"`
		Price string `json:"price"`
		PriceMatch string `json:"priceMatch,omitempty"` // only avaliable for LIMIT/STOP/TAKE_PROFIT order; can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20; Can't be passed together with price
		Quantity string `json:"quantity"` // Order quantity, cannot be sent with closePosition=true
		RecvWindow int64 `json:"recvWindow,omitempty"`
		Side string `json:"side"` // SELL, BUY
		Signature string `json:"signature,omitempty"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


