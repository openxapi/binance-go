/*
Binance Pmargin API

OpenAPI specification for Binance cryptocurrency exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginGetCmConditionalOpenOrdersV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetCmConditionalOpenOrdersV1RespItem{}

// PmarginGetCmConditionalOpenOrdersV1RespItem struct for PmarginGetCmConditionalOpenOrdersV1RespItem
type PmarginGetCmConditionalOpenOrdersV1RespItem struct {
	ActivatePrice *string `json:"activatePrice,omitempty"`
	BookTime *int64 `json:"bookTime,omitempty"`
	NewClientStrategyId *string `json:"newClientStrategyId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
	PositionSide *string `json:"positionSide,omitempty"`
	Price *string `json:"price,omitempty"`
	PriceRate *string `json:"priceRate,omitempty"`
	ReduceOnly *bool `json:"reduceOnly,omitempty"`
	Side *string `json:"side,omitempty"`
	StopPrice *string `json:"stopPrice,omitempty"`
	StrategyId *int64 `json:"strategyId,omitempty"`
	StrategyStatus *string `json:"strategyStatus,omitempty"`
	StrategyType *string `json:"strategyType,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewPmarginGetCmConditionalOpenOrdersV1RespItem instantiates a new PmarginGetCmConditionalOpenOrdersV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetCmConditionalOpenOrdersV1RespItem() *PmarginGetCmConditionalOpenOrdersV1RespItem {
	this := PmarginGetCmConditionalOpenOrdersV1RespItem{}
	return &this
}

// NewPmarginGetCmConditionalOpenOrdersV1RespItemWithDefaults instantiates a new PmarginGetCmConditionalOpenOrdersV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetCmConditionalOpenOrdersV1RespItemWithDefaults() *PmarginGetCmConditionalOpenOrdersV1RespItem {
	this := PmarginGetCmConditionalOpenOrdersV1RespItem{}
	return &this
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetActivatePrice() string {
	if o == nil || IsNil(o.ActivatePrice) {
		var ret string
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetActivatePriceOk() (*string, bool) {
	if o == nil || IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasActivatePrice() bool {
	if o != nil && !IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given string and assigns it to the ActivatePrice field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetActivatePrice(v string) {
	o.ActivatePrice = &v
}

// GetBookTime returns the BookTime field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetBookTime() int64 {
	if o == nil || IsNil(o.BookTime) {
		var ret int64
		return ret
	}
	return *o.BookTime
}

// GetBookTimeOk returns a tuple with the BookTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetBookTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.BookTime) {
		return nil, false
	}
	return o.BookTime, true
}

// HasBookTime returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasBookTime() bool {
	if o != nil && !IsNil(o.BookTime) {
		return true
	}

	return false
}

// SetBookTime gets a reference to the given int64 and assigns it to the BookTime field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetBookTime(v int64) {
	o.BookTime = &v
}

// GetNewClientStrategyId returns the NewClientStrategyId field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetNewClientStrategyId() string {
	if o == nil || IsNil(o.NewClientStrategyId) {
		var ret string
		return ret
	}
	return *o.NewClientStrategyId
}

// GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetNewClientStrategyIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewClientStrategyId) {
		return nil, false
	}
	return o.NewClientStrategyId, true
}

// HasNewClientStrategyId returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasNewClientStrategyId() bool {
	if o != nil && !IsNil(o.NewClientStrategyId) {
		return true
	}

	return false
}

// SetNewClientStrategyId gets a reference to the given string and assigns it to the NewClientStrategyId field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetNewClientStrategyId(v string) {
	o.NewClientStrategyId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetPrice(v string) {
	o.Price = &v
}

// GetPriceRate returns the PriceRate field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetPriceRate() string {
	if o == nil || IsNil(o.PriceRate) {
		var ret string
		return ret
	}
	return *o.PriceRate
}

// GetPriceRateOk returns a tuple with the PriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetPriceRateOk() (*string, bool) {
	if o == nil || IsNil(o.PriceRate) {
		return nil, false
	}
	return o.PriceRate, true
}

// HasPriceRate returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasPriceRate() bool {
	if o != nil && !IsNil(o.PriceRate) {
		return true
	}

	return false
}

// SetPriceRate gets a reference to the given string and assigns it to the PriceRate field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetPriceRate(v string) {
	o.PriceRate = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetReduceOnly() bool {
	if o == nil || IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasReduceOnly() bool {
	if o != nil && !IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetSide(v string) {
	o.Side = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetStrategyId returns the StrategyId field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetStrategyId() int64 {
	if o == nil || IsNil(o.StrategyId) {
		var ret int64
		return ret
	}
	return *o.StrategyId
}

// GetStrategyIdOk returns a tuple with the StrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetStrategyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.StrategyId) {
		return nil, false
	}
	return o.StrategyId, true
}

// HasStrategyId returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasStrategyId() bool {
	if o != nil && !IsNil(o.StrategyId) {
		return true
	}

	return false
}

// SetStrategyId gets a reference to the given int64 and assigns it to the StrategyId field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetStrategyId(v int64) {
	o.StrategyId = &v
}

// GetStrategyStatus returns the StrategyStatus field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetStrategyStatus() string {
	if o == nil || IsNil(o.StrategyStatus) {
		var ret string
		return ret
	}
	return *o.StrategyStatus
}

// GetStrategyStatusOk returns a tuple with the StrategyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetStrategyStatusOk() (*string, bool) {
	if o == nil || IsNil(o.StrategyStatus) {
		return nil, false
	}
	return o.StrategyStatus, true
}

// HasStrategyStatus returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasStrategyStatus() bool {
	if o != nil && !IsNil(o.StrategyStatus) {
		return true
	}

	return false
}

// SetStrategyStatus gets a reference to the given string and assigns it to the StrategyStatus field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetStrategyStatus(v string) {
	o.StrategyStatus = &v
}

// GetStrategyType returns the StrategyType field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetStrategyType() string {
	if o == nil || IsNil(o.StrategyType) {
		var ret string
		return ret
	}
	return *o.StrategyType
}

// GetStrategyTypeOk returns a tuple with the StrategyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetStrategyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StrategyType) {
		return nil, false
	}
	return o.StrategyType, true
}

// HasStrategyType returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasStrategyType() bool {
	if o != nil && !IsNil(o.StrategyType) {
		return true
	}

	return false
}

// SetStrategyType gets a reference to the given string and assigns it to the StrategyType field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetStrategyType(v string) {
	o.StrategyType = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *PmarginGetCmConditionalOpenOrdersV1RespItem) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o PmarginGetCmConditionalOpenOrdersV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetCmConditionalOpenOrdersV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivatePrice) {
		toSerialize["activatePrice"] = o.ActivatePrice
	}
	if !IsNil(o.BookTime) {
		toSerialize["bookTime"] = o.BookTime
	}
	if !IsNil(o.NewClientStrategyId) {
		toSerialize["newClientStrategyId"] = o.NewClientStrategyId
	}
	if !IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PriceRate) {
		toSerialize["priceRate"] = o.PriceRate
	}
	if !IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !IsNil(o.StrategyId) {
		toSerialize["strategyId"] = o.StrategyId
	}
	if !IsNil(o.StrategyStatus) {
		toSerialize["strategyStatus"] = o.StrategyStatus
	}
	if !IsNil(o.StrategyType) {
		toSerialize["strategyType"] = o.StrategyType
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullablePmarginGetCmConditionalOpenOrdersV1RespItem struct {
	value *PmarginGetCmConditionalOpenOrdersV1RespItem
	isSet bool
}

func (v NullablePmarginGetCmConditionalOpenOrdersV1RespItem) Get() *PmarginGetCmConditionalOpenOrdersV1RespItem {
	return v.value
}

func (v *NullablePmarginGetCmConditionalOpenOrdersV1RespItem) Set(val *PmarginGetCmConditionalOpenOrdersV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetCmConditionalOpenOrdersV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetCmConditionalOpenOrdersV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetCmConditionalOpenOrdersV1RespItem(val *PmarginGetCmConditionalOpenOrdersV1RespItem) *NullablePmarginGetCmConditionalOpenOrdersV1RespItem {
	return &NullablePmarginGetCmConditionalOpenOrdersV1RespItem{value: val, isSet: true}
}

func (v NullablePmarginGetCmConditionalOpenOrdersV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetCmConditionalOpenOrdersV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


