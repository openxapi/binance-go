/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the CreateCmPositionSideDualV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCmPositionSideDualV1Resp{}

// CreateCmPositionSideDualV1Resp struct for CreateCmPositionSideDualV1Resp
type CreateCmPositionSideDualV1Resp struct {
	Code *int32 `json:"code,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewCreateCmPositionSideDualV1Resp instantiates a new CreateCmPositionSideDualV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCmPositionSideDualV1Resp() *CreateCmPositionSideDualV1Resp {
	this := CreateCmPositionSideDualV1Resp{}
	return &this
}

// NewCreateCmPositionSideDualV1RespWithDefaults instantiates a new CreateCmPositionSideDualV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCmPositionSideDualV1RespWithDefaults() *CreateCmPositionSideDualV1Resp {
	this := CreateCmPositionSideDualV1Resp{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateCmPositionSideDualV1Resp) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCmPositionSideDualV1Resp) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateCmPositionSideDualV1Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *CreateCmPositionSideDualV1Resp) SetCode(v int32) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateCmPositionSideDualV1Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCmPositionSideDualV1Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CreateCmPositionSideDualV1Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateCmPositionSideDualV1Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o CreateCmPositionSideDualV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCmPositionSideDualV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableCreateCmPositionSideDualV1Resp struct {
	value *CreateCmPositionSideDualV1Resp
	isSet bool
}

func (v NullableCreateCmPositionSideDualV1Resp) Get() *CreateCmPositionSideDualV1Resp {
	return v.value
}

func (v *NullableCreateCmPositionSideDualV1Resp) Set(val *CreateCmPositionSideDualV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCmPositionSideDualV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCmPositionSideDualV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCmPositionSideDualV1Resp(val *CreateCmPositionSideDualV1Resp) *NullableCreateCmPositionSideDualV1Resp {
	return &NullableCreateCmPositionSideDualV1Resp{value: val, isSet: true}
}

func (v NullableCreateCmPositionSideDualV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCmPositionSideDualV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


