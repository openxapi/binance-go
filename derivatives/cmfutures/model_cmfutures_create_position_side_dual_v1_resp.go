/*
Binance COIN-M Futures API

OpenAPI specification for Binance exchange - Cmfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
)

// checks if the CmfuturesCreatePositionSideDualV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CmfuturesCreatePositionSideDualV1Resp{}

// CmfuturesCreatePositionSideDualV1Resp struct for CmfuturesCreatePositionSideDualV1Resp
type CmfuturesCreatePositionSideDualV1Resp struct {
	Code *int32 `json:"code,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewCmfuturesCreatePositionSideDualV1Resp instantiates a new CmfuturesCreatePositionSideDualV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmfuturesCreatePositionSideDualV1Resp() *CmfuturesCreatePositionSideDualV1Resp {
	this := CmfuturesCreatePositionSideDualV1Resp{}
	return &this
}

// NewCmfuturesCreatePositionSideDualV1RespWithDefaults instantiates a new CmfuturesCreatePositionSideDualV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmfuturesCreatePositionSideDualV1RespWithDefaults() *CmfuturesCreatePositionSideDualV1Resp {
	this := CmfuturesCreatePositionSideDualV1Resp{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CmfuturesCreatePositionSideDualV1Resp) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreatePositionSideDualV1Resp) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CmfuturesCreatePositionSideDualV1Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *CmfuturesCreatePositionSideDualV1Resp) SetCode(v int32) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CmfuturesCreatePositionSideDualV1Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreatePositionSideDualV1Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CmfuturesCreatePositionSideDualV1Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CmfuturesCreatePositionSideDualV1Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o CmfuturesCreatePositionSideDualV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmfuturesCreatePositionSideDualV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableCmfuturesCreatePositionSideDualV1Resp struct {
	value *CmfuturesCreatePositionSideDualV1Resp
	isSet bool
}

func (v NullableCmfuturesCreatePositionSideDualV1Resp) Get() *CmfuturesCreatePositionSideDualV1Resp {
	return v.value
}

func (v *NullableCmfuturesCreatePositionSideDualV1Resp) Set(val *CmfuturesCreatePositionSideDualV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCmfuturesCreatePositionSideDualV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCmfuturesCreatePositionSideDualV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmfuturesCreatePositionSideDualV1Resp(val *CmfuturesCreatePositionSideDualV1Resp) *NullableCmfuturesCreatePositionSideDualV1Resp {
	return &NullableCmfuturesCreatePositionSideDualV1Resp{value: val, isSet: true}
}

func (v NullableCmfuturesCreatePositionSideDualV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmfuturesCreatePositionSideDualV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


