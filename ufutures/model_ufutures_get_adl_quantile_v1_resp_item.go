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

// checks if the UfuturesGetAdlQuantileV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetAdlQuantileV1RespItem{}

// UfuturesGetAdlQuantileV1RespItem struct for UfuturesGetAdlQuantileV1RespItem
type UfuturesGetAdlQuantileV1RespItem struct {
	AdlQuantile *UfuturesGetAdlQuantileV1RespItemAdlQuantile `json:"adlQuantile,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewUfuturesGetAdlQuantileV1RespItem instantiates a new UfuturesGetAdlQuantileV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetAdlQuantileV1RespItem() *UfuturesGetAdlQuantileV1RespItem {
	this := UfuturesGetAdlQuantileV1RespItem{}
	return &this
}

// NewUfuturesGetAdlQuantileV1RespItemWithDefaults instantiates a new UfuturesGetAdlQuantileV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetAdlQuantileV1RespItemWithDefaults() *UfuturesGetAdlQuantileV1RespItem {
	this := UfuturesGetAdlQuantileV1RespItem{}
	return &this
}

// GetAdlQuantile returns the AdlQuantile field value if set, zero value otherwise.
func (o *UfuturesGetAdlQuantileV1RespItem) GetAdlQuantile() UfuturesGetAdlQuantileV1RespItemAdlQuantile {
	if o == nil || IsNil(o.AdlQuantile) {
		var ret UfuturesGetAdlQuantileV1RespItemAdlQuantile
		return ret
	}
	return *o.AdlQuantile
}

// GetAdlQuantileOk returns a tuple with the AdlQuantile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAdlQuantileV1RespItem) GetAdlQuantileOk() (*UfuturesGetAdlQuantileV1RespItemAdlQuantile, bool) {
	if o == nil || IsNil(o.AdlQuantile) {
		return nil, false
	}
	return o.AdlQuantile, true
}

// HasAdlQuantile returns a boolean if a field has been set.
func (o *UfuturesGetAdlQuantileV1RespItem) HasAdlQuantile() bool {
	if o != nil && !IsNil(o.AdlQuantile) {
		return true
	}

	return false
}

// SetAdlQuantile gets a reference to the given UfuturesGetAdlQuantileV1RespItemAdlQuantile and assigns it to the AdlQuantile field.
func (o *UfuturesGetAdlQuantileV1RespItem) SetAdlQuantile(v UfuturesGetAdlQuantileV1RespItemAdlQuantile) {
	o.AdlQuantile = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UfuturesGetAdlQuantileV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAdlQuantileV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UfuturesGetAdlQuantileV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UfuturesGetAdlQuantileV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

func (o UfuturesGetAdlQuantileV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetAdlQuantileV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdlQuantile) {
		toSerialize["adlQuantile"] = o.AdlQuantile
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableUfuturesGetAdlQuantileV1RespItem struct {
	value *UfuturesGetAdlQuantileV1RespItem
	isSet bool
}

func (v NullableUfuturesGetAdlQuantileV1RespItem) Get() *UfuturesGetAdlQuantileV1RespItem {
	return v.value
}

func (v *NullableUfuturesGetAdlQuantileV1RespItem) Set(val *UfuturesGetAdlQuantileV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetAdlQuantileV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetAdlQuantileV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetAdlQuantileV1RespItem(val *UfuturesGetAdlQuantileV1RespItem) *NullableUfuturesGetAdlQuantileV1RespItem {
	return &NullableUfuturesGetAdlQuantileV1RespItem{value: val, isSet: true}
}

func (v NullableUfuturesGetAdlQuantileV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetAdlQuantileV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


