package models

// OrderTradeUpdateEvent represents global message '#/components/messages/orderTradeUpdateEvent'
type OrderTradeUpdateEvent struct {
	Event struct {
		EventType string `json:"e,omitempty"` // Event Type
		EventTime int64 `json:"E,omitempty"` // Event Time (milliseconds)
		TransactionTime int64 `json:"T,omitempty"` // Transaction Time (milliseconds)
		OrderDetails struct {
			Symbol string `json:"s,omitempty"` // Symbol
			ClientOrderID string `json:"c,omitempty"` // Client Order ID
			Side string `json:"S,omitempty"` // Side
			OrderType string `json:"o,omitempty"` // Order Type
			TimeInForce string `json:"f,omitempty"` // Time in Force
			OriginalQuantity string `json:"q,omitempty"` // Original Quantity
			OriginalPrice string `json:"p,omitempty"` // Original Price
			AveragePrice string `json:"ap,omitempty"` // Average Price
			StopPrice string `json:"sp,omitempty"` // Stop Price
			ExecutionType string `json:"x,omitempty"` // Execution Type
			OrderStatus string `json:"X,omitempty"` // Order Status
			OrderID int64 `json:"i,omitempty"` // Order ID
			OrderLastFilledQuantity string `json:"l,omitempty"` // Order Last Filled Quantity
			OrderFilledAccumulatedQuantity string `json:"z,omitempty"` // Order Filled Accumulated Quantity
			LastFilledPrice string `json:"L,omitempty"` // Last Filled Price
			CommissionAsset string `json:"N,omitempty"` // Commission Asset
			Commission string `json:"n,omitempty"` // Commission
			OrderTradeTime int64 `json:"T,omitempty"` // Order Trade Time (milliseconds)
			TradeID int64 `json:"t,omitempty"` // Trade ID
			BidsNotional string `json:"b,omitempty"` // Bids Notional
			AskNotional string `json:"a,omitempty"` // Ask Notional
			MakerSide bool `json:"m,omitempty"` // Maker side
			IsThisReduceOnly bool `json:"R,omitempty"` // Is this reduce only
			StopPriceWorkingType string `json:"wt,omitempty"` // Stop Price Working Type
			OriginalOrderType string `json:"ot,omitempty"` // Original Order Type
			PositionSide string `json:"ps,omitempty"` // Position Side
			IfCloseAll bool `json:"cp,omitempty"` // If Close-All
			ActivationPrice string `json:"AP,omitempty"` // Activation Price (only for TRAILING_STOP_MARKET)
			CallbackRate string `json:"cr,omitempty"` // Callback Rate (only for TRAILING_STOP_MARKET)
			ProtectPosition bool `json:"pP,omitempty"` // Protect position
			Ignore int64 `json:"si,omitempty"` // Ignore
			Ignore2 int64 `json:"ss,omitempty"` // Ignore
			RealizedProfit string `json:"rp,omitempty"` // Realized Profit
			STPMode string `json:"V,omitempty"` // STP mode
			PriceMatchMode string `json:"pm,omitempty"` // Price match mode
			TIFGTDOrderAutoCancelTime int64 `json:"gtd,omitempty"` // TIF GTD order auto cancel time
		} `json:"o,omitempty"` // Order details
	} `json:"event,omitempty"`
}


