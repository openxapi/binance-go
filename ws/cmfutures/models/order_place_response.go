package models

// OrderPlaceResponse represents global message '#/components/messages/orderPlaceResponse'
type OrderPlaceResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		AvgPrice string `json:"avgPrice,omitempty"` // avgPrice property
		ClientOrderId string `json:"clientOrderId,omitempty"` // clientOrderId property
		ClosePosition bool `json:"closePosition,omitempty"` // closePosition property
		CumBase string `json:"cumBase,omitempty"` // cumBase property
		CumQty string `json:"cumQty,omitempty"` // cumQty property
		ExecutedQty string `json:"executedQty,omitempty"` // executedQty property
		OrderId int64 `json:"orderId,omitempty"` // orderId property
		OrigQty string `json:"origQty,omitempty"` // origQty property
		OrigType string `json:"origType,omitempty"` // origType property
		Pair string `json:"pair,omitempty"` // pair property
		PositionSide string `json:"positionSide,omitempty"` // positionSide property
		Price string `json:"price,omitempty"` // price property
		PriceProtect bool `json:"priceProtect,omitempty"` // priceProtect property
		ReduceOnly bool `json:"reduceOnly,omitempty"` // reduceOnly property
		Side string `json:"side,omitempty"` // side property
		Status string `json:"status,omitempty"` // status property
		StopPrice string `json:"stopPrice,omitempty"` // stopPrice property
		Symbol string `json:"symbol,omitempty"` // symbol property
		TimeInForce string `json:"timeInForce,omitempty"` // timeInForce property
		Type string `json:"type,omitempty"` // type property
		UpdateTime int64 `json:"updateTime,omitempty"` // updateTime property
		WorkingType string `json:"workingType,omitempty"` // workingType property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


