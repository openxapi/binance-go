package models

// OrderPlaceRequest represents global message '#/components/messages/orderPlaceRequest'
type OrderPlaceRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		ActivationPrice string `json:"activationPrice,omitempty"` // Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType)
		ApiKey string `json:"apiKey,omitempty"`
		CallbackRate string `json:"callbackRate,omitempty"` // Used with TRAILING_STOP_MARKET orders, min 0.1, max 10 where 1 for 1%
		ClosePosition string `json:"closePosition,omitempty"` // true, false；Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
		NewClientOrderId string `json:"newClientOrderId,omitempty"` // A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: ^[\.A-Z\:/a-z0-9_-]{1,36}$
		NewOrderRespType string `json:"newOrderRespType,omitempty"` // ACK,RESULT, default ACK
		PositionSide string `json:"positionSide,omitempty"` // Default BOTH for One-way Mode; LONG or SHORT for Hedge Mode.  It must be sent in Hedge Mode.
		Price string `json:"price,omitempty"`
		PriceMatch string `json:"priceMatch,omitempty"` // only available for LIMIT/STOP/TAKE_PROFIT order; can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20; Can't be passed together with price
		PriceProtect string `json:"priceProtect,omitempty"` // "TRUE" or "FALSE", default "FALSE". Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
		Quantity string `json:"quantity,omitempty"` // Quantity measured by contract number, Cannot be sent with closePosition=true
		RecvWindow int `json:"recvWindow,omitempty"`
		ReduceOnly string `json:"reduceOnly,omitempty"` // true or false. default false. Cannot be sent in Hedge Mode; cannot be sent with closePosition=true (Close-All)
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"` // NONE: No STP / EXPIRE_TAKER:expire taker order when STP triggers/ EXPIRE_MAKER:expire taker order when STP triggers/ EXPIRE_BOTH:expire both orders when STP triggers; default NONE
		Side string `json:"side"` // BUY or SELL
		Signature string `json:"signature,omitempty"`
		StopPrice string `json:"stopPrice,omitempty"` // Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
		Symbol string `json:"symbol"`
		TimeInForce string `json:"timeInForce,omitempty"`
		Timestamp int64 `json:"timestamp"`
		Type string `json:"type"` // LIMIT, MARKET, STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
		WorkingType string `json:"workingType,omitempty"` // stopPrice triggered by: "MARK_PRICE", "CONTRACT_PRICE". Default "CONTRACT_PRICE"
	} `json:"params,omitempty"` // params property
}


