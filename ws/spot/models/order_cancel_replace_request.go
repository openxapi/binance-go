package models

// OrderCancelReplaceRequest represents global message '#/components/messages/orderCancelReplaceRequest'
type OrderCancelReplaceRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		CancelNewClientOrderId string `json:"cancelNewClientOrderId,omitempty"` // New ID for the canceled order. Automatically generated if not sent
		CancelOrderId int64 `json:"cancelOrderId"` // Cancel order by orderId
		CancelOrigClientOrderId string `json:"cancelOrigClientOrderId,omitempty"` // Cancel order by clientOrderId
		CancelReplaceMode string `json:"cancelReplaceMode"`
		CancelRestrictions string `json:"cancelRestrictions,omitempty"` // Supported values: ONLY_NEW - Cancel will succeed if the order status is NEW. ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED. For more information please refer to Regarding cancelRestrictions.
		IcebergQty string `json:"icebergQty,omitempty"`
		NewClientOrderId string `json:"newClientOrderId,omitempty"` // Arbitrary unique ID among open orders. Automatically generated if not sent
		NewOrderRespType string `json:"newOrderRespType,omitempty"` // Select response format: ACK, RESULT, FULL.                              MARKET and LIMIT orders produce FULL response by default,                 other order types default to ACK.
		OrderRateLimitExceededMode string `json:"orderRateLimitExceededMode,omitempty"` // Supported values:  DO_NOTHING (default)- will only attempt to cancel the order if account has not exceeded the unfilled order rate limit CANCEL_ONLY - will always cancel the order.
		Price string `json:"price,omitempty"`
		Quantity string `json:"quantity,omitempty"`
		QuoteOrderQty string `json:"quoteOrderQty,omitempty"`
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // The allowed enums is dependent on what is configured on the symbol.             Supported values: STP Modes.
		Side string `json:"side"` // BUY or SELL
		Signature string `json:"signature"`
		StopPrice string `json:"stopPrice,omitempty"`
		StrategyId int64 `json:"strategyId,omitempty"` // Arbitrary numeric value identifying the order within an order strategy.
		StrategyType int `json:"strategyType,omitempty"` // Arbitrary numeric value identifying the order strategy.             Values smaller than 1000000 are reserved and cannot be used.
		Symbol string `json:"symbol"`
		TimeInForce string `json:"timeInForce,omitempty"`
		Timestamp int64 `json:"timestamp"`
		TrailingDelta string `json:"trailingDelta,omitempty"` // See Trailing Stop order FAQ
		Type string `json:"type"`
	} `json:"params,omitempty"` // params property
}


