/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginCreateRepayFuturesSwitchV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginCreateRepayFuturesSwitchV1Resp{}

// PmarginCreateRepayFuturesSwitchV1Resp struct for PmarginCreateRepayFuturesSwitchV1Resp
type PmarginCreateRepayFuturesSwitchV1Resp struct {
	Msg *string `json:"msg,omitempty"`
}

// NewPmarginCreateRepayFuturesSwitchV1Resp instantiates a new PmarginCreateRepayFuturesSwitchV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginCreateRepayFuturesSwitchV1Resp() *PmarginCreateRepayFuturesSwitchV1Resp {
	this := PmarginCreateRepayFuturesSwitchV1Resp{}
	return &this
}

// NewPmarginCreateRepayFuturesSwitchV1RespWithDefaults instantiates a new PmarginCreateRepayFuturesSwitchV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginCreateRepayFuturesSwitchV1RespWithDefaults() *PmarginCreateRepayFuturesSwitchV1Resp {
	this := PmarginCreateRepayFuturesSwitchV1Resp{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *PmarginCreateRepayFuturesSwitchV1Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateRepayFuturesSwitchV1Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *PmarginCreateRepayFuturesSwitchV1Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *PmarginCreateRepayFuturesSwitchV1Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o PmarginCreateRepayFuturesSwitchV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginCreateRepayFuturesSwitchV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullablePmarginCreateRepayFuturesSwitchV1Resp struct {
	value *PmarginCreateRepayFuturesSwitchV1Resp
	isSet bool
}

func (v NullablePmarginCreateRepayFuturesSwitchV1Resp) Get() *PmarginCreateRepayFuturesSwitchV1Resp {
	return v.value
}

func (v *NullablePmarginCreateRepayFuturesSwitchV1Resp) Set(val *PmarginCreateRepayFuturesSwitchV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginCreateRepayFuturesSwitchV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginCreateRepayFuturesSwitchV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginCreateRepayFuturesSwitchV1Resp(val *PmarginCreateRepayFuturesSwitchV1Resp) *NullablePmarginCreateRepayFuturesSwitchV1Resp {
	return &NullablePmarginCreateRepayFuturesSwitchV1Resp{value: val, isSet: true}
}

func (v NullablePmarginCreateRepayFuturesSwitchV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginCreateRepayFuturesSwitchV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


