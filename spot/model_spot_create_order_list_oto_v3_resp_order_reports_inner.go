/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotCreateOrderListOtoV3RespOrderReportsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotCreateOrderListOtoV3RespOrderReportsInner{}

// SpotCreateOrderListOtoV3RespOrderReportsInner struct for SpotCreateOrderListOtoV3RespOrderReportsInner
type SpotCreateOrderListOtoV3RespOrderReportsInner struct {
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	CummulativeQuoteQty *string `json:"cummulativeQuoteQty,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	OrderListId *int64 `json:"orderListId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
	OrigQuoteOrderQty *string `json:"origQuoteOrderQty,omitempty"`
	Price *string `json:"price,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Side *string `json:"side,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	TransactTime *int64 `json:"transactTime,omitempty"`
	Type *string `json:"type,omitempty"`
	WorkingTime *int64 `json:"workingTime,omitempty"`
}

// NewSpotCreateOrderListOtoV3RespOrderReportsInner instantiates a new SpotCreateOrderListOtoV3RespOrderReportsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotCreateOrderListOtoV3RespOrderReportsInner() *SpotCreateOrderListOtoV3RespOrderReportsInner {
	this := SpotCreateOrderListOtoV3RespOrderReportsInner{}
	return &this
}

// NewSpotCreateOrderListOtoV3RespOrderReportsInnerWithDefaults instantiates a new SpotCreateOrderListOtoV3RespOrderReportsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotCreateOrderListOtoV3RespOrderReportsInnerWithDefaults() *SpotCreateOrderListOtoV3RespOrderReportsInner {
	this := SpotCreateOrderListOtoV3RespOrderReportsInner{}
	return &this
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetCummulativeQuoteQty() string {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasCummulativeQuoteQty() bool {
	if o != nil && !IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetOrderListId() int64 {
	if o == nil || IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasOrderListId() bool {
	if o != nil && !IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetOrigQuoteOrderQty() string {
	if o == nil || IsNil(o.OrigQuoteOrderQty) {
		var ret string
		return ret
	}
	return *o.OrigQuoteOrderQty
}

// GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetOrigQuoteOrderQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQuoteOrderQty) {
		return nil, false
	}
	return o.OrigQuoteOrderQty, true
}

// HasOrigQuoteOrderQty returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasOrigQuoteOrderQty() bool {
	if o != nil && !IsNil(o.OrigQuoteOrderQty) {
		return true
	}

	return false
}

// SetOrigQuoteOrderQty gets a reference to the given string and assigns it to the OrigQuoteOrderQty field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetOrigQuoteOrderQty(v string) {
	o.OrigQuoteOrderQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetPrice(v string) {
	o.Price = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetSelfTradePreventionMode() string {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasSelfTradePreventionMode() bool {
	if o != nil && !IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetStatus(v string) {
	o.Status = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetTransactTime() int64 {
	if o == nil || IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetTransactTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasTransactTime() bool {
	if o != nil && !IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetTransactTime(v int64) {
	o.TransactTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetType(v string) {
	o.Type = &v
}

// GetWorkingTime returns the WorkingTime field value if set, zero value otherwise.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetWorkingTime() int64 {
	if o == nil || IsNil(o.WorkingTime) {
		var ret int64
		return ret
	}
	return *o.WorkingTime
}

// GetWorkingTimeOk returns a tuple with the WorkingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) GetWorkingTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.WorkingTime) {
		return nil, false
	}
	return o.WorkingTime, true
}

// HasWorkingTime returns a boolean if a field has been set.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) HasWorkingTime() bool {
	if o != nil && !IsNil(o.WorkingTime) {
		return true
	}

	return false
}

// SetWorkingTime gets a reference to the given int64 and assigns it to the WorkingTime field.
func (o *SpotCreateOrderListOtoV3RespOrderReportsInner) SetWorkingTime(v int64) {
	o.WorkingTime = &v
}

func (o SpotCreateOrderListOtoV3RespOrderReportsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotCreateOrderListOtoV3RespOrderReportsInner) ToMap() (map[string]interface{}, error) {
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

type NullableSpotCreateOrderListOtoV3RespOrderReportsInner struct {
	value *SpotCreateOrderListOtoV3RespOrderReportsInner
	isSet bool
}

func (v NullableSpotCreateOrderListOtoV3RespOrderReportsInner) Get() *SpotCreateOrderListOtoV3RespOrderReportsInner {
	return v.value
}

func (v *NullableSpotCreateOrderListOtoV3RespOrderReportsInner) Set(val *SpotCreateOrderListOtoV3RespOrderReportsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotCreateOrderListOtoV3RespOrderReportsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotCreateOrderListOtoV3RespOrderReportsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotCreateOrderListOtoV3RespOrderReportsInner(val *SpotCreateOrderListOtoV3RespOrderReportsInner) *NullableSpotCreateOrderListOtoV3RespOrderReportsInner {
	return &NullableSpotCreateOrderListOtoV3RespOrderReportsInner{value: val, isSet: true}
}

func (v NullableSpotCreateOrderListOtoV3RespOrderReportsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotCreateOrderListOtoV3RespOrderReportsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


