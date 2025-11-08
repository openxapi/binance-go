package models

// OrderListPlaceOcoRequest represents global message '#/components/messages/orderListPlaceOcoRequest'
type OrderListPlaceOcoRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		AboveClientOrderId string `json:"aboveClientOrderId,omitempty"` // Arbitrary unique ID among open orders for the above order. Automatically generated if not sent
		AboveIcebergQty int64 `json:"aboveIcebergQty,omitempty"` // Note that this can only be used if aboveTimeInForce is GTC.
		AbovePrice string `json:"abovePrice,omitempty"` // Can be used if aboveType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
		AboveStopPrice string `json:"aboveStopPrice,omitempty"` // Can be used if aboveType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT. Either aboveStopPrice or aboveTrailingDelta or both, must be specified.
		AboveStrategyId int64 `json:"aboveStrategyId,omitempty"` // Arbitrary numeric value identifying the above order within an order strategy.
		AboveStrategyType int `json:"aboveStrategyType,omitempty"` // Arbitrary numeric value identifying the above order strategy. Values smaller than 1000000 are reserved and cannot be used.
		AboveTimeInForce string `json:"aboveTimeInForce,omitempty"` // Required if aboveType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT.
		AboveTrailingDelta int64 `json:"aboveTrailingDelta,omitempty"` // See Trailing Stop order FAQ.
		AboveType string `json:"aboveType"` // Supported values: STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER, TAKE_PROFIT, TAKE_PROFIT_LIMIT
		ApiKey string `json:"apiKey"`
		BelowClientOrderId string `json:"belowClientOrderId,omitempty"`
		BelowIcebergQty int64 `json:"belowIcebergQty,omitempty"` // Note that this can only be used if belowTimeInForce is GTC.
		BelowPrice string `json:"belowPrice,omitempty"` // Can be used if belowType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
		BelowStopPrice string `json:"belowStopPrice,omitempty"` // Can be used if belowType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT. Either belowStopPrice or belowTrailingDelta or both, must be specified.
		BelowStrategyId int64 `json:"belowStrategyId,omitempty"` // Arbitrary numeric value identifying the below order within an order strategy.
		BelowStrategyType int `json:"belowStrategyType,omitempty"` // Arbitrary numeric value identifying the below order strategy. Values smaller than 1000000 are reserved and cannot be used.
		BelowTimeInForce string `json:"belowTimeInForce,omitempty"` // Required if belowType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT
		BelowTrailingDelta int64 `json:"belowTrailingDelta,omitempty"` // See Trailing Stop order FAQ.
		BelowType string `json:"belowType"` // Supported values: STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT,TAKE_PROFIT_LIMIT
		ListClientOrderId string `json:"listClientOrderId,omitempty"` // Arbitrary unique ID among open order lists. Automatically generated if not sent.  A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.  listClientOrderId is distinct from the aboveClientOrderId and the belowCLientOrderId.
		NewOrderRespType string `json:"newOrderRespType,omitempty"` // Select response format: ACK, RESULT, FULL
		Quantity string `json:"quantity"` // Quantity for both orders of the order list.
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000.
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
		Side string `json:"side"` // BUY or SELL
		Signature string `json:"signature"`
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
	} `json:"params,omitempty"` // params property
}


