package models

// OrderListPlaceOtoRequest represents global message '#/components/messages/orderListPlaceOtoRequest'
type OrderListPlaceOtoRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey,omitempty"`
		ListClientOrderId string `json:"listClientOrderId,omitempty"` // Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.  listClientOrderId is distinct from the workingClientOrderId and the pendingClientOrderId.
		NewOrderRespType string `json:"newOrderRespType,omitempty"` // Format of the JSON response. Supported values: Order Response Type
		PendingClientOrderId string `json:"pendingClientOrderId,omitempty"` // Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
		PendingIcebergQty string `json:"pendingIcebergQty,omitempty"` // This can only be used if pendingTimeInForce is GTC, or if pendingType is LIMIT_MAKER.
		PendingPrice string `json:"pendingPrice,omitempty"`
		PendingQuantity string `json:"pendingQuantity"` // Sets the quantity for the pending order.
		PendingSide string `json:"pendingSide"` // Supported values: Order side
		PendingStopPrice string `json:"pendingStopPrice,omitempty"`
		PendingStrategyId int64 `json:"pendingStrategyId,omitempty"` // Arbitrary numeric value identifying the pending order within an order strategy.
		PendingStrategyType int `json:"pendingStrategyType,omitempty"` // Arbitrary numeric value identifying the pending order strategy.  Values smaller than 1000000 are reserved and cannot be used.
		PendingTimeInForce string `json:"pendingTimeInForce,omitempty"` // Supported values: Time In Force
		PendingTrailingDelta string `json:"pendingTrailingDelta,omitempty"`
		PendingType string `json:"pendingType"` // Supported values: Order types.  Note that MARKET orders using quoteOrderQty are not supported.
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000.
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // The allowed values are dependent on what is configured on the symbol. Supported values: STP Modes
		Signature string `json:"signature"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
		WorkingClientOrderId string `json:"workingClientOrderId,omitempty"` // Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
		WorkingIcebergQty string `json:"workingIcebergQty,omitempty"` // This can only be used if workingTimeInForce is GTC, or if workingType is LIMIT_MAKER.
		WorkingPrice string `json:"workingPrice"`
		WorkingQuantity string `json:"workingQuantity"` // Sets the quantity for the working order.
		WorkingSide string `json:"workingSide"` // Supported values: Order side
		WorkingStrategyId int64 `json:"workingStrategyId,omitempty"` // Arbitrary numeric value identifying the working order within an order strategy.
		WorkingStrategyType int `json:"workingStrategyType,omitempty"` // Arbitrary numeric value identifying the working order strategy.  Values smaller than 1000000 are reserved and cannot be used.
		WorkingTimeInForce string `json:"workingTimeInForce,omitempty"` // Supported values: Time In Force
		WorkingType string `json:"workingType"` // Supported values: LIMIT,LIMIT_MAKER
	} `json:"params,omitempty"` // params property
}


