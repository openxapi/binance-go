package models

// OrderListPlaceOtocoRequest represents global message '#/components/messages/orderListPlaceOtocoRequest'
type OrderListPlaceOtocoRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		ListClientOrderId string `json:"listClientOrderId,omitempty"` // Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.  listClientOrderId is distinct from the workingClientOrderId, pendingAboveClientOrderId, and the pendingBelowClientOrderId.
		NewOrderRespType string `json:"newOrderRespType,omitempty"` // Format of the JSON response. Supported values: Order Response Type
		PendingAboveClientOrderId string `json:"pendingAboveClientOrderId,omitempty"` // Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
		PendingAboveIcebergQty string `json:"pendingAboveIcebergQty,omitempty"` // This can only be used if pendingAboveTimeInForce is GTC or if pendingAboveType is LIMIT_MAKER.
		PendingAbovePrice string `json:"pendingAbovePrice,omitempty"` // Can be used if pendingAboveType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
		PendingAboveStopPrice string `json:"pendingAboveStopPrice,omitempty"` // Can be used if pendingAboveType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT
		PendingAboveStrategyId int64 `json:"pendingAboveStrategyId,omitempty"` // Arbitrary numeric value identifying the pending above order within an order strategy.
		PendingAboveStrategyType int `json:"pendingAboveStrategyType,omitempty"` // Arbitrary numeric value identifying the pending above order strategy.  Values smaller than 1000000 are reserved and cannot be used.
		PendingAboveTimeInForce string `json:"pendingAboveTimeInForce,omitempty"`
		PendingAboveTrailingDelta string `json:"pendingAboveTrailingDelta,omitempty"` // See Trailing Stop FAQ
		PendingAboveType string `json:"pendingAboveType"` // Supported values: STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER, TAKE_PROFIT, TAKE_PROFIT_LIMIT
		PendingBelowClientOrderId string `json:"pendingBelowClientOrderId,omitempty"` // Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
		PendingBelowIcebergQty string `json:"pendingBelowIcebergQty,omitempty"` // This can only be used if pendingBelowTimeInForce is GTC, or if pendingBelowType is LIMIT_MAKER.
		PendingBelowPrice string `json:"pendingBelowPrice,omitempty"` // Can be used if pendingBelowType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT to specify the limit price.
		PendingBelowStopPrice string `json:"pendingBelowStopPrice,omitempty"` // Can be used if pendingBelowType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT. Either pendingBelowStopPrice or pendingBelowTrailingDelta or both, must be specified.
		PendingBelowStrategyId int64 `json:"pendingBelowStrategyId,omitempty"` // Arbitrary numeric value identifying the pending below order within an order strategy.
		PendingBelowStrategyType int `json:"pendingBelowStrategyType,omitempty"` // Arbitrary numeric value identifying the pending below order strategy.  Values smaller than 1000000 are reserved and cannot be used.
		PendingBelowTimeInForce string `json:"pendingBelowTimeInForce,omitempty"` // Supported values: Time In Force
		PendingBelowTrailingDelta string `json:"pendingBelowTrailingDelta,omitempty"`
		PendingBelowType string `json:"pendingBelowType,omitempty"` // Supported values: STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT,TAKE_PROFIT_LIMIT
		PendingQuantity string `json:"pendingQuantity"`
		PendingSide string `json:"pendingSide"` // Supported values: Order Side
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000.
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // The allowed values are dependent on what is configured on the symbol. Supported values: STP Modes
		Signature string `json:"signature"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
		WorkingClientOrderId string `json:"workingClientOrderId,omitempty"` // Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
		WorkingIcebergQty string `json:"workingIcebergQty,omitempty"` // This can only be used if workingTimeInForce is GTC.
		WorkingPrice string `json:"workingPrice"`
		WorkingQuantity string `json:"workingQuantity"`
		WorkingSide string `json:"workingSide"` // Supported values: Order Side
		WorkingStrategyId int64 `json:"workingStrategyId,omitempty"` // Arbitrary numeric value identifying the working order within an order strategy.
		WorkingStrategyType int `json:"workingStrategyType,omitempty"` // Arbitrary numeric value identifying the working order strategy.  Values smaller than 1000000 are reserved and cannot be used.
		WorkingTimeInForce string `json:"workingTimeInForce,omitempty"` // Supported values: Time In Force
		WorkingType string `json:"workingType"` // Supported values: LIMIT, LIMIT_MAKER
	} `json:"params,omitempty"` // params property
}


