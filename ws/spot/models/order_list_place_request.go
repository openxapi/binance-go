package models

// OrderListPlaceRequest represents global message '#/components/messages/orderListPlaceRequest'
type OrderListPlaceRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ApiKey string `json:"apiKey"`
		LimitClientOrderId string `json:"limitClientOrderId,omitempty"` // Arbitrary unique ID among open orders for the limit order. Automatically generated if not sent
		LimitIcebergQty string `json:"limitIcebergQty,omitempty"`
		LimitStrategyId int64 `json:"limitStrategyId,omitempty"` // Arbitrary numeric value identifying the limit order within an order strategy.
		LimitStrategyType int `json:"limitStrategyType,omitempty"` // Arbitrary numeric value identifying the limit order strategy.Values smaller than 1000000 are reserved and cannot be used.
		ListClientOrderId string `json:"listClientOrderId,omitempty"` // Arbitrary unique ID among open order lists. Automatically generated if not sent
		NewOrderRespType string `json:"newOrderRespType,omitempty"` // Select response format: ACK, RESULT, FULL (default)
		Price string `json:"price"` // Price for the limit order
		Quantity string `json:"quantity"`
		RecvWindow int64 `json:"recvWindow,omitempty"` // The value cannot be greater than 60000
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes
		Side string `json:"side"` // BUY or SELL
		Signature string `json:"signature"`
		StopClientOrderId string `json:"stopClientOrderId,omitempty"` // Arbitrary unique ID among open orders for the stop order. Automatically generated if not sent
		StopIcebergQty string `json:"stopIcebergQty,omitempty"`
		StopLimitPrice string `json:"stopLimitPrice,omitempty"`
		StopLimitTimeInForce string `json:"stopLimitTimeInForce,omitempty"` // See order.place for available options
		StopPrice string `json:"stopPrice"` // Either stopPrice or trailingDelta, or both must be specified
		StopStrategyId int64 `json:"stopStrategyId,omitempty"` // Arbitrary numeric value identifying the stop order within an order strategy.
		StopStrategyType int `json:"stopStrategyType,omitempty"` // Arbitrary numeric value identifying the stop order strategy.Values smaller than 1000000 are reserved and cannot be used.
		Symbol string `json:"symbol"`
		Timestamp int64 `json:"timestamp"`
		TrailingDelta int `json:"trailingDelta"` // See Trailing Stop order FAQ
	} `json:"params,omitempty"` // params property
}


