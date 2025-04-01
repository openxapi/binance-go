/*
Binance Algo API

OpenAPI specification for Binance cryptocurrency exchange - Algo API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package algo

import (
	"encoding/json"
)

// checks if the AlgoCreateAlgoFuturesNewOrderTwapV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlgoCreateAlgoFuturesNewOrderTwapV1Resp{}

// AlgoCreateAlgoFuturesNewOrderTwapV1Resp struct for AlgoCreateAlgoFuturesNewOrderTwapV1Resp
type AlgoCreateAlgoFuturesNewOrderTwapV1Resp struct {
	ClientAlgoId *string `json:"clientAlgoId,omitempty"`
	Code *int32 `json:"code,omitempty"`
	Msg *string `json:"msg,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewAlgoCreateAlgoFuturesNewOrderTwapV1Resp instantiates a new AlgoCreateAlgoFuturesNewOrderTwapV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlgoCreateAlgoFuturesNewOrderTwapV1Resp() *AlgoCreateAlgoFuturesNewOrderTwapV1Resp {
	this := AlgoCreateAlgoFuturesNewOrderTwapV1Resp{}
	return &this
}

// NewAlgoCreateAlgoFuturesNewOrderTwapV1RespWithDefaults instantiates a new AlgoCreateAlgoFuturesNewOrderTwapV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlgoCreateAlgoFuturesNewOrderTwapV1RespWithDefaults() *AlgoCreateAlgoFuturesNewOrderTwapV1Resp {
	this := AlgoCreateAlgoFuturesNewOrderTwapV1Resp{}
	return &this
}

// GetClientAlgoId returns the ClientAlgoId field value if set, zero value otherwise.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) GetClientAlgoId() string {
	if o == nil || IsNil(o.ClientAlgoId) {
		var ret string
		return ret
	}
	return *o.ClientAlgoId
}

// GetClientAlgoIdOk returns a tuple with the ClientAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) GetClientAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientAlgoId) {
		return nil, false
	}
	return o.ClientAlgoId, true
}

// HasClientAlgoId returns a boolean if a field has been set.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) HasClientAlgoId() bool {
	if o != nil && !IsNil(o.ClientAlgoId) {
		return true
	}

	return false
}

// SetClientAlgoId gets a reference to the given string and assigns it to the ClientAlgoId field.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) SetClientAlgoId(v string) {
	o.ClientAlgoId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) SetCode(v int32) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) SetMsg(v string) {
	o.Msg = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) SetSuccess(v bool) {
	o.Success = &v
}

func (o AlgoCreateAlgoFuturesNewOrderTwapV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlgoCreateAlgoFuturesNewOrderTwapV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientAlgoId) {
		toSerialize["clientAlgoId"] = o.ClientAlgoId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp struct {
	value *AlgoCreateAlgoFuturesNewOrderTwapV1Resp
	isSet bool
}

func (v NullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp) Get() *AlgoCreateAlgoFuturesNewOrderTwapV1Resp {
	return v.value
}

func (v *NullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp) Set(val *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp(val *AlgoCreateAlgoFuturesNewOrderTwapV1Resp) *NullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp {
	return &NullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp{value: val, isSet: true}
}

func (v NullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgoCreateAlgoFuturesNewOrderTwapV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


