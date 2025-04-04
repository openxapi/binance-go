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

// checks if the SpotGetOrderV3Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotGetOrderV3Resp{}

// SpotGetOrderV3Resp struct for SpotGetOrderV3Resp
type SpotGetOrderV3Resp struct {
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	CummulativeQuoteQty *string `json:"cummulativeQuoteQty,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	IcebergQty *string `json:"icebergQty,omitempty"`
	IsWorking *bool `json:"isWorking,omitempty"`
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
	Time *int64 `json:"time,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
	WorkingTime *int64 `json:"workingTime,omitempty"`
}

// NewSpotGetOrderV3Resp instantiates a new SpotGetOrderV3Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotGetOrderV3Resp() *SpotGetOrderV3Resp {
	this := SpotGetOrderV3Resp{}
	return &this
}

// NewSpotGetOrderV3RespWithDefaults instantiates a new SpotGetOrderV3Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotGetOrderV3RespWithDefaults() *SpotGetOrderV3Resp {
	this := SpotGetOrderV3Resp{}
	return &this
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *SpotGetOrderV3Resp) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetCummulativeQuoteQty() string {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasCummulativeQuoteQty() bool {
	if o != nil && !IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *SpotGetOrderV3Resp) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *SpotGetOrderV3Resp) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetIcebergQty returns the IcebergQty field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetIcebergQty() string {
	if o == nil || IsNil(o.IcebergQty) {
		var ret string
		return ret
	}
	return *o.IcebergQty
}

// GetIcebergQtyOk returns a tuple with the IcebergQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetIcebergQtyOk() (*string, bool) {
	if o == nil || IsNil(o.IcebergQty) {
		return nil, false
	}
	return o.IcebergQty, true
}

// HasIcebergQty returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasIcebergQty() bool {
	if o != nil && !IsNil(o.IcebergQty) {
		return true
	}

	return false
}

// SetIcebergQty gets a reference to the given string and assigns it to the IcebergQty field.
func (o *SpotGetOrderV3Resp) SetIcebergQty(v string) {
	o.IcebergQty = &v
}

// GetIsWorking returns the IsWorking field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetIsWorking() bool {
	if o == nil || IsNil(o.IsWorking) {
		var ret bool
		return ret
	}
	return *o.IsWorking
}

// GetIsWorkingOk returns a tuple with the IsWorking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetIsWorkingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWorking) {
		return nil, false
	}
	return o.IsWorking, true
}

// HasIsWorking returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasIsWorking() bool {
	if o != nil && !IsNil(o.IsWorking) {
		return true
	}

	return false
}

// SetIsWorking gets a reference to the given bool and assigns it to the IsWorking field.
func (o *SpotGetOrderV3Resp) SetIsWorking(v bool) {
	o.IsWorking = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *SpotGetOrderV3Resp) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetOrderListId() int64 {
	if o == nil || IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetOrderListIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasOrderListId() bool {
	if o != nil && !IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *SpotGetOrderV3Resp) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *SpotGetOrderV3Resp) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetOrigQuoteOrderQty() string {
	if o == nil || IsNil(o.OrigQuoteOrderQty) {
		var ret string
		return ret
	}
	return *o.OrigQuoteOrderQty
}

// GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetOrigQuoteOrderQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQuoteOrderQty) {
		return nil, false
	}
	return o.OrigQuoteOrderQty, true
}

// HasOrigQuoteOrderQty returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasOrigQuoteOrderQty() bool {
	if o != nil && !IsNil(o.OrigQuoteOrderQty) {
		return true
	}

	return false
}

// SetOrigQuoteOrderQty gets a reference to the given string and assigns it to the OrigQuoteOrderQty field.
func (o *SpotGetOrderV3Resp) SetOrigQuoteOrderQty(v string) {
	o.OrigQuoteOrderQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SpotGetOrderV3Resp) SetPrice(v string) {
	o.Price = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetSelfTradePreventionMode() string {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasSelfTradePreventionMode() bool {
	if o != nil && !IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *SpotGetOrderV3Resp) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *SpotGetOrderV3Resp) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SpotGetOrderV3Resp) SetStatus(v string) {
	o.Status = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *SpotGetOrderV3Resp) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SpotGetOrderV3Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SpotGetOrderV3Resp) SetTime(v int64) {
	o.Time = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *SpotGetOrderV3Resp) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SpotGetOrderV3Resp) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *SpotGetOrderV3Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetWorkingTime returns the WorkingTime field value if set, zero value otherwise.
func (o *SpotGetOrderV3Resp) GetWorkingTime() int64 {
	if o == nil || IsNil(o.WorkingTime) {
		var ret int64
		return ret
	}
	return *o.WorkingTime
}

// GetWorkingTimeOk returns a tuple with the WorkingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetOrderV3Resp) GetWorkingTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.WorkingTime) {
		return nil, false
	}
	return o.WorkingTime, true
}

// HasWorkingTime returns a boolean if a field has been set.
func (o *SpotGetOrderV3Resp) HasWorkingTime() bool {
	if o != nil && !IsNil(o.WorkingTime) {
		return true
	}

	return false
}

// SetWorkingTime gets a reference to the given int64 and assigns it to the WorkingTime field.
func (o *SpotGetOrderV3Resp) SetWorkingTime(v int64) {
	o.WorkingTime = &v
}

func (o SpotGetOrderV3Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotGetOrderV3Resp) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsWorking) {
		toSerialize["isWorking"] = o.IsWorking
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
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.WorkingTime) {
		toSerialize["workingTime"] = o.WorkingTime
	}
	return toSerialize, nil
}

type NullableSpotGetOrderV3Resp struct {
	value *SpotGetOrderV3Resp
	isSet bool
}

func (v NullableSpotGetOrderV3Resp) Get() *SpotGetOrderV3Resp {
	return v.value
}

func (v *NullableSpotGetOrderV3Resp) Set(val *SpotGetOrderV3Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotGetOrderV3Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotGetOrderV3Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotGetOrderV3Resp(val *SpotGetOrderV3Resp) *NullableSpotGetOrderV3Resp {
	return &NullableSpotGetOrderV3Resp{value: val, isSet: true}
}

func (v NullableSpotGetOrderV3Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotGetOrderV3Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


