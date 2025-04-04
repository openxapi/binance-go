/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginUpdateCmOrderV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginUpdateCmOrderV1Resp{}

// PmarginUpdateCmOrderV1Resp struct for PmarginUpdateCmOrderV1Resp
type PmarginUpdateCmOrderV1Resp struct {
	AvgPrice *string `json:"avgPrice,omitempty"`
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	CumBase *string `json:"cumBase,omitempty"`
	CumQty *string `json:"cumQty,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
	OrigType *string `json:"origType,omitempty"`
	Pair *string `json:"pair,omitempty"`
	PositionSide *string `json:"positionSide,omitempty"`
	Price *string `json:"price,omitempty"`
	ReduceOnly *bool `json:"reduceOnly,omitempty"`
	Side *string `json:"side,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewPmarginUpdateCmOrderV1Resp instantiates a new PmarginUpdateCmOrderV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginUpdateCmOrderV1Resp() *PmarginUpdateCmOrderV1Resp {
	this := PmarginUpdateCmOrderV1Resp{}
	return &this
}

// NewPmarginUpdateCmOrderV1RespWithDefaults instantiates a new PmarginUpdateCmOrderV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginUpdateCmOrderV1RespWithDefaults() *PmarginUpdateCmOrderV1Resp {
	this := PmarginUpdateCmOrderV1Resp{}
	return &this
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetAvgPrice() string {
	if o == nil || IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetAvgPriceOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasAvgPrice() bool {
	if o != nil && !IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *PmarginUpdateCmOrderV1Resp) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *PmarginUpdateCmOrderV1Resp) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCumBase returns the CumBase field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetCumBase() string {
	if o == nil || IsNil(o.CumBase) {
		var ret string
		return ret
	}
	return *o.CumBase
}

// GetCumBaseOk returns a tuple with the CumBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetCumBaseOk() (*string, bool) {
	if o == nil || IsNil(o.CumBase) {
		return nil, false
	}
	return o.CumBase, true
}

// HasCumBase returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasCumBase() bool {
	if o != nil && !IsNil(o.CumBase) {
		return true
	}

	return false
}

// SetCumBase gets a reference to the given string and assigns it to the CumBase field.
func (o *PmarginUpdateCmOrderV1Resp) SetCumBase(v string) {
	o.CumBase = &v
}

// GetCumQty returns the CumQty field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetCumQty() string {
	if o == nil || IsNil(o.CumQty) {
		var ret string
		return ret
	}
	return *o.CumQty
}

// GetCumQtyOk returns a tuple with the CumQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetCumQtyOk() (*string, bool) {
	if o == nil || IsNil(o.CumQty) {
		return nil, false
	}
	return o.CumQty, true
}

// HasCumQty returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasCumQty() bool {
	if o != nil && !IsNil(o.CumQty) {
		return true
	}

	return false
}

// SetCumQty gets a reference to the given string and assigns it to the CumQty field.
func (o *PmarginUpdateCmOrderV1Resp) SetCumQty(v string) {
	o.CumQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *PmarginUpdateCmOrderV1Resp) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *PmarginUpdateCmOrderV1Resp) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *PmarginUpdateCmOrderV1Resp) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetOrigType returns the OrigType field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetOrigType() string {
	if o == nil || IsNil(o.OrigType) {
		var ret string
		return ret
	}
	return *o.OrigType
}

// GetOrigTypeOk returns a tuple with the OrigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetOrigTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrigType) {
		return nil, false
	}
	return o.OrigType, true
}

// HasOrigType returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasOrigType() bool {
	if o != nil && !IsNil(o.OrigType) {
		return true
	}

	return false
}

// SetOrigType gets a reference to the given string and assigns it to the OrigType field.
func (o *PmarginUpdateCmOrderV1Resp) SetOrigType(v string) {
	o.OrigType = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetPair() string {
	if o == nil || IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetPairOk() (*string, bool) {
	if o == nil || IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasPair() bool {
	if o != nil && !IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *PmarginUpdateCmOrderV1Resp) SetPair(v string) {
	o.Pair = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *PmarginUpdateCmOrderV1Resp) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *PmarginUpdateCmOrderV1Resp) SetPrice(v string) {
	o.Price = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetReduceOnly() bool {
	if o == nil || IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasReduceOnly() bool {
	if o != nil && !IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *PmarginUpdateCmOrderV1Resp) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *PmarginUpdateCmOrderV1Resp) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PmarginUpdateCmOrderV1Resp) SetStatus(v string) {
	o.Status = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PmarginUpdateCmOrderV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *PmarginUpdateCmOrderV1Resp) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PmarginUpdateCmOrderV1Resp) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PmarginUpdateCmOrderV1Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginUpdateCmOrderV1Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PmarginUpdateCmOrderV1Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *PmarginUpdateCmOrderV1Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o PmarginUpdateCmOrderV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginUpdateCmOrderV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !IsNil(o.CumBase) {
		toSerialize["cumBase"] = o.CumBase
	}
	if !IsNil(o.CumQty) {
		toSerialize["cumQty"] = o.CumQty
	}
	if !IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !IsNil(o.OrigType) {
		toSerialize["origType"] = o.OrigType
	}
	if !IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullablePmarginUpdateCmOrderV1Resp struct {
	value *PmarginUpdateCmOrderV1Resp
	isSet bool
}

func (v NullablePmarginUpdateCmOrderV1Resp) Get() *PmarginUpdateCmOrderV1Resp {
	return v.value
}

func (v *NullablePmarginUpdateCmOrderV1Resp) Set(val *PmarginUpdateCmOrderV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginUpdateCmOrderV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginUpdateCmOrderV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginUpdateCmOrderV1Resp(val *PmarginUpdateCmOrderV1Resp) *NullablePmarginUpdateCmOrderV1Resp {
	return &NullablePmarginUpdateCmOrderV1Resp{value: val, isSet: true}
}

func (v NullablePmarginUpdateCmOrderV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginUpdateCmOrderV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


