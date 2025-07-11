package models

import (
	"encoding/json"
)

// OrderModifyRequestParams contains the parameters for OrderModifyRequest
type OrderModifyRequestParams struct {
	OrderId *int64 `json:"orderId,omitempty"`
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"`
	Price string `json:"price"`
	// only avaliable for LIMIT/STOP/TAKE_PROFIT order; can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20; Can't be passed together with price
	PriceMatch *string `json:"priceMatch,omitempty"`
	// Order quantity, cannot be sent with closePosition=true
	Quantity string `json:"quantity"`
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	// SELL, BUY
	Side string `json:"side"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// OrderModifyRequest - Send a order.modify request
// Message name: Modify Order (TRADE) Request
type OrderModifyRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderModifyRequestParams `json:"params,omitempty"`
}

// NewOrderModifyRequest creates a new OrderModifyRequest instance
func NewOrderModifyRequest() *OrderModifyRequest {
	return &OrderModifyRequest{}
}

// String returns string representation of OrderModifyRequest
func (s OrderModifyRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderModifyRequest) SetParams(value OrderModifyRequestParams) *OrderModifyRequest {
	r.Params = &value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *OrderModifyRequest) SetOrderId(value int64) *OrderModifyRequest {
	if r.Params == nil {
		r.Params = &OrderModifyRequestParams{}
	}
	r.Params.OrderId = &value
	return r
}

// SetOrigClientOrderId sets the origClientOrderId parameter and returns the struct for method chaining
func (r *OrderModifyRequest) SetOrigClientOrderId(value string) *OrderModifyRequest {
	if r.Params == nil {
		r.Params = &OrderModifyRequestParams{}
	}
	r.Params.OrigClientOrderId = &value
	return r
}

// SetPrice sets the price parameter and returns the struct for method chaining
func (r *OrderModifyRequest) SetPrice(value string) *OrderModifyRequest {
	if r.Params == nil {
		r.Params = &OrderModifyRequestParams{}
	}
	r.Params.Price = value
	return r
}

// SetPriceMatch sets the priceMatch parameter and returns the struct for method chaining
func (r *OrderModifyRequest) SetPriceMatch(value string) *OrderModifyRequest {
	if r.Params == nil {
		r.Params = &OrderModifyRequestParams{}
	}
	r.Params.PriceMatch = &value
	return r
}

// SetQuantity sets the quantity parameter and returns the struct for method chaining
func (r *OrderModifyRequest) SetQuantity(value string) *OrderModifyRequest {
	if r.Params == nil {
		r.Params = &OrderModifyRequestParams{}
	}
	r.Params.Quantity = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderModifyRequest) SetRecvWindow(value int64) *OrderModifyRequest {
	if r.Params == nil {
		r.Params = &OrderModifyRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSide sets the side parameter and returns the struct for method chaining
func (r *OrderModifyRequest) SetSide(value string) *OrderModifyRequest {
	if r.Params == nil {
		r.Params = &OrderModifyRequestParams{}
	}
	r.Params.Side = value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderModifyRequest) SetSymbol(value string) *OrderModifyRequest {
	if r.Params == nil {
		r.Params = &OrderModifyRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderModifyRequest) SetTimestamp(value int64) *OrderModifyRequest {
	if r.Params == nil {
		r.Params = &OrderModifyRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


