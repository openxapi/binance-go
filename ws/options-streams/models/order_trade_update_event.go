package models

// OrderTradeUpdateEvent represents global message '#/components/messages/orderTradeUpdateEvent'
type OrderTradeUpdateEvent struct {
	EventType string `json:"e"` // Event Type
	EventTime int64 `json:"E"` // Event Time (milliseconds)
	OrderDetailsArray []struct {
		Symbol string `json:"s"` // Symbol
		OrderID string `json:"oid"` // Order ID
		ClientOrderID string `json:"c,omitempty"` // Client Order ID
		OrderPrice string `json:"p,omitempty"` // Order Price
		OrderQuantity string `json:"q,omitempty"` // Order Quantity (+ve for BUY, -ve for SELL)
		CompletedTradeVolume string `json:"e,omitempty"` // Completed Trade Volume
		CompletedTradeAmount string `json:"ec,omitempty"` // Completed Trade Amount
		Fee string `json:"f,omitempty"` // Fee
		OrderCreateTime int64 `json:"T,omitempty"` // Order Create Time (milliseconds)
		OrderUpdateTime int64 `json:"t,omitempty"` // Order Update Time (milliseconds)
		OrderType string `json:"oty,omitempty"` // Order Type
		OrderStatus string `json:"S,omitempty"` // Order Status
		TimeInForce string `json:"tif,omitempty"` // Time in Force
		PostOnlyFlag bool `json:"po,omitempty"` // Post Only flag
		ReduceOnlyFlag bool `json:"r,omitempty"` // Reduce Only flag
		NotUsedCurrently int `json:"stp,omitempty"` // Not used currently
		FillsArray []struct {
			TradeID string `json:"t,omitempty"` // Trade ID
			TradePrice string `json:"p,omitempty"` // Trade Price
			TradeQuantity string `json:"q,omitempty"` // Trade Quantity
			CommissionRebate string `json:"f,omitempty"` // Commission/Rebate
			TradeTime int64 `json:"T,omitempty"` // Trade Time (milliseconds)
			TakerMaker string `json:"m,omitempty"` // Taker/Maker
		} `json:"fi,omitempty"` // Fills Array
	} `json:"o"` // Order details array
}


