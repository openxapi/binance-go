/*
Binance Options API

OpenAPI specification for Binance exchange - Options API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
)

// checks if the GetTimeV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeV1Resp{}

// GetTimeV1Resp struct for GetTimeV1Resp
type GetTimeV1Resp struct {
	ServerTime *int64 `json:"serverTime,omitempty"`
}

// NewGetTimeV1Resp instantiates a new GetTimeV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeV1Resp() *GetTimeV1Resp {
	this := GetTimeV1Resp{}
	return &this
}

// NewGetTimeV1RespWithDefaults instantiates a new GetTimeV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeV1RespWithDefaults() *GetTimeV1Resp {
	this := GetTimeV1Resp{}
	return &this
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *GetTimeV1Resp) GetServerTime() int64 {
	if o == nil || IsNil(o.ServerTime) {
		var ret int64
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeV1Resp) GetServerTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ServerTime) {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *GetTimeV1Resp) HasServerTime() bool {
	if o != nil && !IsNil(o.ServerTime) {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int64 and assigns it to the ServerTime field.
func (o *GetTimeV1Resp) SetServerTime(v int64) {
	o.ServerTime = &v
}

func (o GetTimeV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerTime) {
		toSerialize["serverTime"] = o.ServerTime
	}
	return toSerialize, nil
}

type NullableGetTimeV1Resp struct {
	value *GetTimeV1Resp
	isSet bool
}

func (v NullableGetTimeV1Resp) Get() *GetTimeV1Resp {
	return v.value
}

func (v *NullableGetTimeV1Resp) Set(val *GetTimeV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeV1Resp(val *GetTimeV1Resp) *NullableGetTimeV1Resp {
	return &NullableGetTimeV1Resp{value: val, isSet: true}
}

func (v NullableGetTimeV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


