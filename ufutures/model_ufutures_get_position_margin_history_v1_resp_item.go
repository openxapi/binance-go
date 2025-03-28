/*
Binance Ufutures API

OpenAPI specification for Binance cryptocurrency exchange - Ufutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ufutures

import (
	"encoding/json"
)

// checks if the UfuturesGetPositionMarginHistoryV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetPositionMarginHistoryV1RespItem{}

// UfuturesGetPositionMarginHistoryV1RespItem struct for UfuturesGetPositionMarginHistoryV1RespItem
type UfuturesGetPositionMarginHistoryV1RespItem struct {
	Amount *string `json:"amount,omitempty"`
	Asset *string `json:"asset,omitempty"`
	DeltaType *string `json:"deltaType,omitempty"`
	PositionSide *string `json:"positionSide,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Time *int64 `json:"time,omitempty"`
	Type *int32 `json:"type,omitempty"`
}

// NewUfuturesGetPositionMarginHistoryV1RespItem instantiates a new UfuturesGetPositionMarginHistoryV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetPositionMarginHistoryV1RespItem() *UfuturesGetPositionMarginHistoryV1RespItem {
	this := UfuturesGetPositionMarginHistoryV1RespItem{}
	return &this
}

// NewUfuturesGetPositionMarginHistoryV1RespItemWithDefaults instantiates a new UfuturesGetPositionMarginHistoryV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetPositionMarginHistoryV1RespItemWithDefaults() *UfuturesGetPositionMarginHistoryV1RespItem {
	this := UfuturesGetPositionMarginHistoryV1RespItem{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) SetAsset(v string) {
	o.Asset = &v
}

// GetDeltaType returns the DeltaType field value if set, zero value otherwise.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetDeltaType() string {
	if o == nil || IsNil(o.DeltaType) {
		var ret string
		return ret
	}
	return *o.DeltaType
}

// GetDeltaTypeOk returns a tuple with the DeltaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetDeltaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeltaType) {
		return nil, false
	}
	return o.DeltaType, true
}

// HasDeltaType returns a boolean if a field has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) HasDeltaType() bool {
	if o != nil && !IsNil(o.DeltaType) {
		return true
	}

	return false
}

// SetDeltaType gets a reference to the given string and assigns it to the DeltaType field.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) SetDeltaType(v string) {
	o.DeltaType = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) SetTime(v int64) {
	o.Time = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UfuturesGetPositionMarginHistoryV1RespItem) SetType(v int32) {
	o.Type = &v
}

func (o UfuturesGetPositionMarginHistoryV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetPositionMarginHistoryV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.DeltaType) {
		toSerialize["deltaType"] = o.DeltaType
	}
	if !IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUfuturesGetPositionMarginHistoryV1RespItem struct {
	value *UfuturesGetPositionMarginHistoryV1RespItem
	isSet bool
}

func (v NullableUfuturesGetPositionMarginHistoryV1RespItem) Get() *UfuturesGetPositionMarginHistoryV1RespItem {
	return v.value
}

func (v *NullableUfuturesGetPositionMarginHistoryV1RespItem) Set(val *UfuturesGetPositionMarginHistoryV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetPositionMarginHistoryV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetPositionMarginHistoryV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetPositionMarginHistoryV1RespItem(val *UfuturesGetPositionMarginHistoryV1RespItem) *NullableUfuturesGetPositionMarginHistoryV1RespItem {
	return &NullableUfuturesGetPositionMarginHistoryV1RespItem{value: val, isSet: true}
}

func (v NullableUfuturesGetPositionMarginHistoryV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetPositionMarginHistoryV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


