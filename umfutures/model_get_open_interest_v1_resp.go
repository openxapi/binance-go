/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the GetOpenInterestV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOpenInterestV1Resp{}

// GetOpenInterestV1Resp struct for GetOpenInterestV1Resp
type GetOpenInterestV1Resp struct {
	OpenInterest *string `json:"openInterest,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewGetOpenInterestV1Resp instantiates a new GetOpenInterestV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOpenInterestV1Resp() *GetOpenInterestV1Resp {
	this := GetOpenInterestV1Resp{}
	return &this
}

// NewGetOpenInterestV1RespWithDefaults instantiates a new GetOpenInterestV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOpenInterestV1RespWithDefaults() *GetOpenInterestV1Resp {
	this := GetOpenInterestV1Resp{}
	return &this
}

// GetOpenInterest returns the OpenInterest field value if set, zero value otherwise.
func (o *GetOpenInterestV1Resp) GetOpenInterest() string {
	if o == nil || IsNil(o.OpenInterest) {
		var ret string
		return ret
	}
	return *o.OpenInterest
}

// GetOpenInterestOk returns a tuple with the OpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenInterestV1Resp) GetOpenInterestOk() (*string, bool) {
	if o == nil || IsNil(o.OpenInterest) {
		return nil, false
	}
	return o.OpenInterest, true
}

// HasOpenInterest returns a boolean if a field has been set.
func (o *GetOpenInterestV1Resp) HasOpenInterest() bool {
	if o != nil && !IsNil(o.OpenInterest) {
		return true
	}

	return false
}

// SetOpenInterest gets a reference to the given string and assigns it to the OpenInterest field.
func (o *GetOpenInterestV1Resp) SetOpenInterest(v string) {
	o.OpenInterest = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetOpenInterestV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenInterestV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetOpenInterestV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetOpenInterestV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetOpenInterestV1Resp) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenInterestV1Resp) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetOpenInterestV1Resp) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetOpenInterestV1Resp) SetTime(v int64) {
	o.Time = &v
}

func (o GetOpenInterestV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOpenInterestV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpenInterest) {
		toSerialize["openInterest"] = o.OpenInterest
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableGetOpenInterestV1Resp struct {
	value *GetOpenInterestV1Resp
	isSet bool
}

func (v NullableGetOpenInterestV1Resp) Get() *GetOpenInterestV1Resp {
	return v.value
}

func (v *NullableGetOpenInterestV1Resp) Set(val *GetOpenInterestV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOpenInterestV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOpenInterestV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOpenInterestV1Resp(val *GetOpenInterestV1Resp) *NullableGetOpenInterestV1Resp {
	return &NullableGetOpenInterestV1Resp{value: val, isSet: true}
}

func (v NullableGetOpenInterestV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOpenInterestV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


