/*
Binance Umfutures API

OpenAPI specification for Binance cryptocurrency exchange - Umfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the UmfuturesGetAdlQuantileV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmfuturesGetAdlQuantileV1RespItem{}

// UmfuturesGetAdlQuantileV1RespItem struct for UmfuturesGetAdlQuantileV1RespItem
type UmfuturesGetAdlQuantileV1RespItem struct {
	AdlQuantile *UmfuturesGetAdlQuantileV1RespItemAdlQuantile `json:"adlQuantile,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewUmfuturesGetAdlQuantileV1RespItem instantiates a new UmfuturesGetAdlQuantileV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmfuturesGetAdlQuantileV1RespItem() *UmfuturesGetAdlQuantileV1RespItem {
	this := UmfuturesGetAdlQuantileV1RespItem{}
	return &this
}

// NewUmfuturesGetAdlQuantileV1RespItemWithDefaults instantiates a new UmfuturesGetAdlQuantileV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmfuturesGetAdlQuantileV1RespItemWithDefaults() *UmfuturesGetAdlQuantileV1RespItem {
	this := UmfuturesGetAdlQuantileV1RespItem{}
	return &this
}

// GetAdlQuantile returns the AdlQuantile field value if set, zero value otherwise.
func (o *UmfuturesGetAdlQuantileV1RespItem) GetAdlQuantile() UmfuturesGetAdlQuantileV1RespItemAdlQuantile {
	if o == nil || IsNil(o.AdlQuantile) {
		var ret UmfuturesGetAdlQuantileV1RespItemAdlQuantile
		return ret
	}
	return *o.AdlQuantile
}

// GetAdlQuantileOk returns a tuple with the AdlQuantile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetAdlQuantileV1RespItem) GetAdlQuantileOk() (*UmfuturesGetAdlQuantileV1RespItemAdlQuantile, bool) {
	if o == nil || IsNil(o.AdlQuantile) {
		return nil, false
	}
	return o.AdlQuantile, true
}

// HasAdlQuantile returns a boolean if a field has been set.
func (o *UmfuturesGetAdlQuantileV1RespItem) HasAdlQuantile() bool {
	if o != nil && !IsNil(o.AdlQuantile) {
		return true
	}

	return false
}

// SetAdlQuantile gets a reference to the given UmfuturesGetAdlQuantileV1RespItemAdlQuantile and assigns it to the AdlQuantile field.
func (o *UmfuturesGetAdlQuantileV1RespItem) SetAdlQuantile(v UmfuturesGetAdlQuantileV1RespItemAdlQuantile) {
	o.AdlQuantile = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UmfuturesGetAdlQuantileV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetAdlQuantileV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UmfuturesGetAdlQuantileV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UmfuturesGetAdlQuantileV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

func (o UmfuturesGetAdlQuantileV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmfuturesGetAdlQuantileV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdlQuantile) {
		toSerialize["adlQuantile"] = o.AdlQuantile
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableUmfuturesGetAdlQuantileV1RespItem struct {
	value *UmfuturesGetAdlQuantileV1RespItem
	isSet bool
}

func (v NullableUmfuturesGetAdlQuantileV1RespItem) Get() *UmfuturesGetAdlQuantileV1RespItem {
	return v.value
}

func (v *NullableUmfuturesGetAdlQuantileV1RespItem) Set(val *UmfuturesGetAdlQuantileV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesGetAdlQuantileV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesGetAdlQuantileV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesGetAdlQuantileV1RespItem(val *UmfuturesGetAdlQuantileV1RespItem) *NullableUmfuturesGetAdlQuantileV1RespItem {
	return &NullableUmfuturesGetAdlQuantileV1RespItem{value: val, isSet: true}
}

func (v NullableUmfuturesGetAdlQuantileV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesGetAdlQuantileV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


