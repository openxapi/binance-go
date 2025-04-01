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

// checks if the UmfuturesDeleteAllOpenOrdersV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmfuturesDeleteAllOpenOrdersV1Resp{}

// UmfuturesDeleteAllOpenOrdersV1Resp struct for UmfuturesDeleteAllOpenOrdersV1Resp
type UmfuturesDeleteAllOpenOrdersV1Resp struct {
	Code *int32 `json:"code,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewUmfuturesDeleteAllOpenOrdersV1Resp instantiates a new UmfuturesDeleteAllOpenOrdersV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmfuturesDeleteAllOpenOrdersV1Resp() *UmfuturesDeleteAllOpenOrdersV1Resp {
	this := UmfuturesDeleteAllOpenOrdersV1Resp{}
	return &this
}

// NewUmfuturesDeleteAllOpenOrdersV1RespWithDefaults instantiates a new UmfuturesDeleteAllOpenOrdersV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmfuturesDeleteAllOpenOrdersV1RespWithDefaults() *UmfuturesDeleteAllOpenOrdersV1Resp {
	this := UmfuturesDeleteAllOpenOrdersV1Resp{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UmfuturesDeleteAllOpenOrdersV1Resp) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesDeleteAllOpenOrdersV1Resp) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UmfuturesDeleteAllOpenOrdersV1Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *UmfuturesDeleteAllOpenOrdersV1Resp) SetCode(v int32) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *UmfuturesDeleteAllOpenOrdersV1Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesDeleteAllOpenOrdersV1Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *UmfuturesDeleteAllOpenOrdersV1Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *UmfuturesDeleteAllOpenOrdersV1Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o UmfuturesDeleteAllOpenOrdersV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmfuturesDeleteAllOpenOrdersV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableUmfuturesDeleteAllOpenOrdersV1Resp struct {
	value *UmfuturesDeleteAllOpenOrdersV1Resp
	isSet bool
}

func (v NullableUmfuturesDeleteAllOpenOrdersV1Resp) Get() *UmfuturesDeleteAllOpenOrdersV1Resp {
	return v.value
}

func (v *NullableUmfuturesDeleteAllOpenOrdersV1Resp) Set(val *UmfuturesDeleteAllOpenOrdersV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesDeleteAllOpenOrdersV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesDeleteAllOpenOrdersV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesDeleteAllOpenOrdersV1Resp(val *UmfuturesDeleteAllOpenOrdersV1Resp) *NullableUmfuturesDeleteAllOpenOrdersV1Resp {
	return &NullableUmfuturesDeleteAllOpenOrdersV1Resp{value: val, isSet: true}
}

func (v NullableUmfuturesDeleteAllOpenOrdersV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesDeleteAllOpenOrdersV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


