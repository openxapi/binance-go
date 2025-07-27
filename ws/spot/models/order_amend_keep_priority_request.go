package models

import (
	"encoding/json"
)

// OrderAmendKeepPriorityRequestParams contains the parameters for OrderAmendKeepPriorityRequest
type OrderAmendKeepPriorityRequestParams struct {
	// The new client order ID for the order after being amended.  If not sent, one will be randomly generated.  It is possible to reuse the current clientOrderId by sending it as the newClientOrderId.
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	// newQty must be greater than 0 and less than the order's quantity.
	NewQty string `json:"newQty"`
	// orderId or origClientOrderId must be sent
	OrderId *int64 `json:"orderId,omitempty"`
	// orderId or origClientOrderId must be sent
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"`
	// The value cannot be greater than 60000.
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Symbol string `json:"symbol"`
	Timestamp int64 `json:"timestamp"`
}

// OrderAmendKeepPriorityRequest - Send a order.amend.keepPriority request
// Message name: Order Amend Keep Priority (TRADE) Request
type OrderAmendKeepPriorityRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *OrderAmendKeepPriorityRequestParams `json:"params,omitempty"`
}

// NewOrderAmendKeepPriorityRequest creates a new OrderAmendKeepPriorityRequest instance
func NewOrderAmendKeepPriorityRequest() *OrderAmendKeepPriorityRequest {
	return &OrderAmendKeepPriorityRequest{}
}

// String returns string representation of OrderAmendKeepPriorityRequest
func (s OrderAmendKeepPriorityRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *OrderAmendKeepPriorityRequest) SetParams(value OrderAmendKeepPriorityRequestParams) *OrderAmendKeepPriorityRequest {
	r.Params = &value
	return r
}

// SetNewClientOrderId sets the newClientOrderId parameter and returns the struct for method chaining
func (r *OrderAmendKeepPriorityRequest) SetNewClientOrderId(value string) *OrderAmendKeepPriorityRequest {
	if r.Params == nil {
		r.Params = &OrderAmendKeepPriorityRequestParams{}
	}
	r.Params.NewClientOrderId = &value
	return r
}

// SetNewQty sets the newQty parameter and returns the struct for method chaining
func (r *OrderAmendKeepPriorityRequest) SetNewQty(value string) *OrderAmendKeepPriorityRequest {
	if r.Params == nil {
		r.Params = &OrderAmendKeepPriorityRequestParams{}
	}
	r.Params.NewQty = value
	return r
}

// SetOrderId sets the orderId parameter and returns the struct for method chaining
func (r *OrderAmendKeepPriorityRequest) SetOrderId(value int64) *OrderAmendKeepPriorityRequest {
	if r.Params == nil {
		r.Params = &OrderAmendKeepPriorityRequestParams{}
	}
	r.Params.OrderId = &value
	return r
}

// SetOrigClientOrderId sets the origClientOrderId parameter and returns the struct for method chaining
func (r *OrderAmendKeepPriorityRequest) SetOrigClientOrderId(value string) *OrderAmendKeepPriorityRequest {
	if r.Params == nil {
		r.Params = &OrderAmendKeepPriorityRequestParams{}
	}
	r.Params.OrigClientOrderId = &value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *OrderAmendKeepPriorityRequest) SetRecvWindow(value int64) *OrderAmendKeepPriorityRequest {
	if r.Params == nil {
		r.Params = &OrderAmendKeepPriorityRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSymbol sets the symbol parameter and returns the struct for method chaining
func (r *OrderAmendKeepPriorityRequest) SetSymbol(value string) *OrderAmendKeepPriorityRequest {
	if r.Params == nil {
		r.Params = &OrderAmendKeepPriorityRequestParams{}
	}
	r.Params.Symbol = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *OrderAmendKeepPriorityRequest) SetTimestamp(value int64) *OrderAmendKeepPriorityRequest {
	if r.Params == nil {
		r.Params = &OrderAmendKeepPriorityRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


