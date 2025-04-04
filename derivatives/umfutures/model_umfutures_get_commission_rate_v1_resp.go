/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the UmfuturesGetCommissionRateV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmfuturesGetCommissionRateV1Resp{}

// UmfuturesGetCommissionRateV1Resp struct for UmfuturesGetCommissionRateV1Resp
type UmfuturesGetCommissionRateV1Resp struct {
	MakerCommissionRate *string `json:"makerCommissionRate,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TakerCommissionRate *string `json:"takerCommissionRate,omitempty"`
}

// NewUmfuturesGetCommissionRateV1Resp instantiates a new UmfuturesGetCommissionRateV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmfuturesGetCommissionRateV1Resp() *UmfuturesGetCommissionRateV1Resp {
	this := UmfuturesGetCommissionRateV1Resp{}
	return &this
}

// NewUmfuturesGetCommissionRateV1RespWithDefaults instantiates a new UmfuturesGetCommissionRateV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmfuturesGetCommissionRateV1RespWithDefaults() *UmfuturesGetCommissionRateV1Resp {
	this := UmfuturesGetCommissionRateV1Resp{}
	return &this
}

// GetMakerCommissionRate returns the MakerCommissionRate field value if set, zero value otherwise.
func (o *UmfuturesGetCommissionRateV1Resp) GetMakerCommissionRate() string {
	if o == nil || IsNil(o.MakerCommissionRate) {
		var ret string
		return ret
	}
	return *o.MakerCommissionRate
}

// GetMakerCommissionRateOk returns a tuple with the MakerCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetCommissionRateV1Resp) GetMakerCommissionRateOk() (*string, bool) {
	if o == nil || IsNil(o.MakerCommissionRate) {
		return nil, false
	}
	return o.MakerCommissionRate, true
}

// HasMakerCommissionRate returns a boolean if a field has been set.
func (o *UmfuturesGetCommissionRateV1Resp) HasMakerCommissionRate() bool {
	if o != nil && !IsNil(o.MakerCommissionRate) {
		return true
	}

	return false
}

// SetMakerCommissionRate gets a reference to the given string and assigns it to the MakerCommissionRate field.
func (o *UmfuturesGetCommissionRateV1Resp) SetMakerCommissionRate(v string) {
	o.MakerCommissionRate = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UmfuturesGetCommissionRateV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetCommissionRateV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UmfuturesGetCommissionRateV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UmfuturesGetCommissionRateV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTakerCommissionRate returns the TakerCommissionRate field value if set, zero value otherwise.
func (o *UmfuturesGetCommissionRateV1Resp) GetTakerCommissionRate() string {
	if o == nil || IsNil(o.TakerCommissionRate) {
		var ret string
		return ret
	}
	return *o.TakerCommissionRate
}

// GetTakerCommissionRateOk returns a tuple with the TakerCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetCommissionRateV1Resp) GetTakerCommissionRateOk() (*string, bool) {
	if o == nil || IsNil(o.TakerCommissionRate) {
		return nil, false
	}
	return o.TakerCommissionRate, true
}

// HasTakerCommissionRate returns a boolean if a field has been set.
func (o *UmfuturesGetCommissionRateV1Resp) HasTakerCommissionRate() bool {
	if o != nil && !IsNil(o.TakerCommissionRate) {
		return true
	}

	return false
}

// SetTakerCommissionRate gets a reference to the given string and assigns it to the TakerCommissionRate field.
func (o *UmfuturesGetCommissionRateV1Resp) SetTakerCommissionRate(v string) {
	o.TakerCommissionRate = &v
}

func (o UmfuturesGetCommissionRateV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmfuturesGetCommissionRateV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MakerCommissionRate) {
		toSerialize["makerCommissionRate"] = o.MakerCommissionRate
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.TakerCommissionRate) {
		toSerialize["takerCommissionRate"] = o.TakerCommissionRate
	}
	return toSerialize, nil
}

type NullableUmfuturesGetCommissionRateV1Resp struct {
	value *UmfuturesGetCommissionRateV1Resp
	isSet bool
}

func (v NullableUmfuturesGetCommissionRateV1Resp) Get() *UmfuturesGetCommissionRateV1Resp {
	return v.value
}

func (v *NullableUmfuturesGetCommissionRateV1Resp) Set(val *UmfuturesGetCommissionRateV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesGetCommissionRateV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesGetCommissionRateV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesGetCommissionRateV1Resp(val *UmfuturesGetCommissionRateV1Resp) *NullableUmfuturesGetCommissionRateV1Resp {
	return &NullableUmfuturesGetCommissionRateV1Resp{value: val, isSet: true}
}

func (v NullableUmfuturesGetCommissionRateV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesGetCommissionRateV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


