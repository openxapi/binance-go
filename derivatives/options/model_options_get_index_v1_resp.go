/*
Binance Options API

OpenAPI specification for Binance exchange - Options API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
)

// checks if the OptionsGetIndexV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsGetIndexV1Resp{}

// OptionsGetIndexV1Resp struct for OptionsGetIndexV1Resp
type OptionsGetIndexV1Resp struct {
	IndexPrice *string `json:"indexPrice,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewOptionsGetIndexV1Resp instantiates a new OptionsGetIndexV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsGetIndexV1Resp() *OptionsGetIndexV1Resp {
	this := OptionsGetIndexV1Resp{}
	return &this
}

// NewOptionsGetIndexV1RespWithDefaults instantiates a new OptionsGetIndexV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsGetIndexV1RespWithDefaults() *OptionsGetIndexV1Resp {
	this := OptionsGetIndexV1Resp{}
	return &this
}

// GetIndexPrice returns the IndexPrice field value if set, zero value otherwise.
func (o *OptionsGetIndexV1Resp) GetIndexPrice() string {
	if o == nil || IsNil(o.IndexPrice) {
		var ret string
		return ret
	}
	return *o.IndexPrice
}

// GetIndexPriceOk returns a tuple with the IndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetIndexV1Resp) GetIndexPriceOk() (*string, bool) {
	if o == nil || IsNil(o.IndexPrice) {
		return nil, false
	}
	return o.IndexPrice, true
}

// HasIndexPrice returns a boolean if a field has been set.
func (o *OptionsGetIndexV1Resp) HasIndexPrice() bool {
	if o != nil && !IsNil(o.IndexPrice) {
		return true
	}

	return false
}

// SetIndexPrice gets a reference to the given string and assigns it to the IndexPrice field.
func (o *OptionsGetIndexV1Resp) SetIndexPrice(v string) {
	o.IndexPrice = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OptionsGetIndexV1Resp) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetIndexV1Resp) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OptionsGetIndexV1Resp) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OptionsGetIndexV1Resp) SetTime(v int64) {
	o.Time = &v
}

func (o OptionsGetIndexV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionsGetIndexV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IndexPrice) {
		toSerialize["indexPrice"] = o.IndexPrice
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableOptionsGetIndexV1Resp struct {
	value *OptionsGetIndexV1Resp
	isSet bool
}

func (v NullableOptionsGetIndexV1Resp) Get() *OptionsGetIndexV1Resp {
	return v.value
}

func (v *NullableOptionsGetIndexV1Resp) Set(val *OptionsGetIndexV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsGetIndexV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsGetIndexV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsGetIndexV1Resp(val *OptionsGetIndexV1Resp) *NullableOptionsGetIndexV1Resp {
	return &NullableOptionsGetIndexV1Resp{value: val, isSet: true}
}

func (v NullableOptionsGetIndexV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsGetIndexV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


