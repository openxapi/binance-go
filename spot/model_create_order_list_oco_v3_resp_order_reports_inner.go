/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the CreateOrderListOcoV3RespOrderReportsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrderListOcoV3RespOrderReportsInner{}

// CreateOrderListOcoV3RespOrderReportsInner struct for CreateOrderListOcoV3RespOrderReportsInner
type CreateOrderListOcoV3RespOrderReportsInner struct {
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	CummulativeQuoteQty *string `json:"cummulativeQuoteQty,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	IcebergQty *string `json:"icebergQty,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	OrderListId *int64 `json:"orderListId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
	OrigQuoteOrderQty *string `json:"origQuoteOrderQty,omitempty"`
	Price *string `json:"price,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Side *string `json:"side,omitempty"`
	Status *string `json:"status,omitempty"`
	StopPrice *string `json:"stopPrice,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	TransactTime *int64 `json:"transactTime,omitempty"`
	Type *string `json:"type,omitempty"`
	WorkingTime *int64 `json:"workingTime,omitempty"`
}

// NewCreateOrderListOcoV3RespOrderReportsInner instantiates a new CreateOrderListOcoV3RespOrderReportsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrderListOcoV3RespOrderReportsInner() *CreateOrderListOcoV3RespOrderReportsInner {
	this := CreateOrderListOcoV3RespOrderReportsInner{}
	return &this
}

// NewCreateOrderListOcoV3RespOrderReportsInnerWithDefaults instantiates a new CreateOrderListOcoV3RespOrderReportsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderListOcoV3RespOrderReportsInnerWithDefaults() *CreateOrderListOcoV3RespOrderReportsInner {
	this := CreateOrderListOcoV3RespOrderReportsInner{}
	return &this
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetCummulativeQuoteQty() string {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasCummulativeQuoteQty() bool {
	if o != nil && !IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetIcebergQty returns the IcebergQty field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetIcebergQty() string {
	if o == nil || IsNil(o.IcebergQty) {
		var ret string
		return ret
	}
	return *o.IcebergQty
}

// GetIcebergQtyOk returns a tuple with the IcebergQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetIcebergQtyOk() (*string, bool) {
	if o == nil || IsNil(o.IcebergQty) {
		return nil, false
	}
	return o.IcebergQty, true
}

// HasIcebergQty returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasIcebergQty() bool {
	if o != nil && !IsNil(o.IcebergQty) {
		return true
	}

	return false
}

// SetIcebergQty gets a reference to the given string and assigns it to the IcebergQty field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetIcebergQty(v string) {
	o.IcebergQty = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetOrderListId() int64 {
	if o == nil || IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasOrderListId() bool {
	if o != nil && !IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetOrigQuoteOrderQty() string {
	if o == nil || IsNil(o.OrigQuoteOrderQty) {
		var ret string
		return ret
	}
	return *o.OrigQuoteOrderQty
}

// GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetOrigQuoteOrderQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQuoteOrderQty) {
		return nil, false
	}
	return o.OrigQuoteOrderQty, true
}

// HasOrigQuoteOrderQty returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasOrigQuoteOrderQty() bool {
	if o != nil && !IsNil(o.OrigQuoteOrderQty) {
		return true
	}

	return false
}

// SetOrigQuoteOrderQty gets a reference to the given string and assigns it to the OrigQuoteOrderQty field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetOrigQuoteOrderQty(v string) {
	o.OrigQuoteOrderQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetPrice(v string) {
	o.Price = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetSelfTradePreventionMode() string {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasSelfTradePreventionMode() bool {
	if o != nil && !IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetStatus(v string) {
	o.Status = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetTransactTime() int64 {
	if o == nil || IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetTransactTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasTransactTime() bool {
	if o != nil && !IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetTransactTime(v int64) {
	o.TransactTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetType(v string) {
	o.Type = &v
}

// GetWorkingTime returns the WorkingTime field value if set, zero value otherwise.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetWorkingTime() int64 {
	if o == nil || IsNil(o.WorkingTime) {
		var ret int64
		return ret
	}
	return *o.WorkingTime
}

// GetWorkingTimeOk returns a tuple with the WorkingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) GetWorkingTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.WorkingTime) {
		return nil, false
	}
	return o.WorkingTime, true
}

// HasWorkingTime returns a boolean if a field has been set.
func (o *CreateOrderListOcoV3RespOrderReportsInner) HasWorkingTime() bool {
	if o != nil && !IsNil(o.WorkingTime) {
		return true
	}

	return false
}

// SetWorkingTime gets a reference to the given int64 and assigns it to the WorkingTime field.
func (o *CreateOrderListOcoV3RespOrderReportsInner) SetWorkingTime(v int64) {
	o.WorkingTime = &v
}

func (o CreateOrderListOcoV3RespOrderReportsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrderListOcoV3RespOrderReportsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !IsNil(o.CummulativeQuoteQty) {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if !IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !IsNil(o.IcebergQty) {
		toSerialize["icebergQty"] = o.IcebergQty
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !IsNil(o.OrigQuoteOrderQty) {
		toSerialize["origQuoteOrderQty"] = o.OrigQuoteOrderQty
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !IsNil(o.TransactTime) {
		toSerialize["transactTime"] = o.TransactTime
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.WorkingTime) {
		toSerialize["workingTime"] = o.WorkingTime
	}
	return toSerialize, nil
}

type NullableCreateOrderListOcoV3RespOrderReportsInner struct {
	value *CreateOrderListOcoV3RespOrderReportsInner
	isSet bool
}

func (v NullableCreateOrderListOcoV3RespOrderReportsInner) Get() *CreateOrderListOcoV3RespOrderReportsInner {
	return v.value
}

func (v *NullableCreateOrderListOcoV3RespOrderReportsInner) Set(val *CreateOrderListOcoV3RespOrderReportsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderListOcoV3RespOrderReportsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderListOcoV3RespOrderReportsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderListOcoV3RespOrderReportsInner(val *CreateOrderListOcoV3RespOrderReportsInner) *NullableCreateOrderListOcoV3RespOrderReportsInner {
	return &NullableCreateOrderListOcoV3RespOrderReportsInner{value: val, isSet: true}
}

func (v NullableCreateOrderListOcoV3RespOrderReportsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderListOcoV3RespOrderReportsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


