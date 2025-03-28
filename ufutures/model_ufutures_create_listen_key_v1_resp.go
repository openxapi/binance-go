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

// checks if the UfuturesCreateListenKeyV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesCreateListenKeyV1Resp{}

// UfuturesCreateListenKeyV1Resp struct for UfuturesCreateListenKeyV1Resp
type UfuturesCreateListenKeyV1Resp struct {
	ListenKey *string `json:"listenKey,omitempty"`
}

// NewUfuturesCreateListenKeyV1Resp instantiates a new UfuturesCreateListenKeyV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesCreateListenKeyV1Resp() *UfuturesCreateListenKeyV1Resp {
	this := UfuturesCreateListenKeyV1Resp{}
	return &this
}

// NewUfuturesCreateListenKeyV1RespWithDefaults instantiates a new UfuturesCreateListenKeyV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesCreateListenKeyV1RespWithDefaults() *UfuturesCreateListenKeyV1Resp {
	this := UfuturesCreateListenKeyV1Resp{}
	return &this
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *UfuturesCreateListenKeyV1Resp) GetListenKey() string {
	if o == nil || IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesCreateListenKeyV1Resp) GetListenKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *UfuturesCreateListenKeyV1Resp) HasListenKey() bool {
	if o != nil && !IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *UfuturesCreateListenKeyV1Resp) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o UfuturesCreateListenKeyV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesCreateListenKeyV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}
	return toSerialize, nil
}

type NullableUfuturesCreateListenKeyV1Resp struct {
	value *UfuturesCreateListenKeyV1Resp
	isSet bool
}

func (v NullableUfuturesCreateListenKeyV1Resp) Get() *UfuturesCreateListenKeyV1Resp {
	return v.value
}

func (v *NullableUfuturesCreateListenKeyV1Resp) Set(val *UfuturesCreateListenKeyV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesCreateListenKeyV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesCreateListenKeyV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesCreateListenKeyV1Resp(val *UfuturesCreateListenKeyV1Resp) *NullableUfuturesCreateListenKeyV1Resp {
	return &NullableUfuturesCreateListenKeyV1Resp{value: val, isSet: true}
}

func (v NullableUfuturesCreateListenKeyV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesCreateListenKeyV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


