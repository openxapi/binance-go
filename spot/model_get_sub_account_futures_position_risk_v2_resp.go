/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the GetSubAccountFuturesPositionRiskV2Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubAccountFuturesPositionRiskV2Resp{}

// GetSubAccountFuturesPositionRiskV2Resp struct for GetSubAccountFuturesPositionRiskV2Resp
type GetSubAccountFuturesPositionRiskV2Resp struct {
	FuturePositionRiskVos []GetSubAccountFuturesPositionRiskV2RespFuturePositionRiskVosInner `json:"futurePositionRiskVos,omitempty"`
}

// NewGetSubAccountFuturesPositionRiskV2Resp instantiates a new GetSubAccountFuturesPositionRiskV2Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountFuturesPositionRiskV2Resp() *GetSubAccountFuturesPositionRiskV2Resp {
	this := GetSubAccountFuturesPositionRiskV2Resp{}
	return &this
}

// NewGetSubAccountFuturesPositionRiskV2RespWithDefaults instantiates a new GetSubAccountFuturesPositionRiskV2Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountFuturesPositionRiskV2RespWithDefaults() *GetSubAccountFuturesPositionRiskV2Resp {
	this := GetSubAccountFuturesPositionRiskV2Resp{}
	return &this
}

// GetFuturePositionRiskVos returns the FuturePositionRiskVos field value if set, zero value otherwise.
func (o *GetSubAccountFuturesPositionRiskV2Resp) GetFuturePositionRiskVos() []GetSubAccountFuturesPositionRiskV2RespFuturePositionRiskVosInner {
	if o == nil || IsNil(o.FuturePositionRiskVos) {
		var ret []GetSubAccountFuturesPositionRiskV2RespFuturePositionRiskVosInner
		return ret
	}
	return o.FuturePositionRiskVos
}

// GetFuturePositionRiskVosOk returns a tuple with the FuturePositionRiskVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesPositionRiskV2Resp) GetFuturePositionRiskVosOk() ([]GetSubAccountFuturesPositionRiskV2RespFuturePositionRiskVosInner, bool) {
	if o == nil || IsNil(o.FuturePositionRiskVos) {
		return nil, false
	}
	return o.FuturePositionRiskVos, true
}

// HasFuturePositionRiskVos returns a boolean if a field has been set.
func (o *GetSubAccountFuturesPositionRiskV2Resp) HasFuturePositionRiskVos() bool {
	if o != nil && !IsNil(o.FuturePositionRiskVos) {
		return true
	}

	return false
}

// SetFuturePositionRiskVos gets a reference to the given []GetSubAccountFuturesPositionRiskV2RespFuturePositionRiskVosInner and assigns it to the FuturePositionRiskVos field.
func (o *GetSubAccountFuturesPositionRiskV2Resp) SetFuturePositionRiskVos(v []GetSubAccountFuturesPositionRiskV2RespFuturePositionRiskVosInner) {
	o.FuturePositionRiskVos = v
}

func (o GetSubAccountFuturesPositionRiskV2Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountFuturesPositionRiskV2Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FuturePositionRiskVos) {
		toSerialize["futurePositionRiskVos"] = o.FuturePositionRiskVos
	}
	return toSerialize, nil
}

type NullableGetSubAccountFuturesPositionRiskV2Resp struct {
	value *GetSubAccountFuturesPositionRiskV2Resp
	isSet bool
}

func (v NullableGetSubAccountFuturesPositionRiskV2Resp) Get() *GetSubAccountFuturesPositionRiskV2Resp {
	return v.value
}

func (v *NullableGetSubAccountFuturesPositionRiskV2Resp) Set(val *GetSubAccountFuturesPositionRiskV2Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountFuturesPositionRiskV2Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountFuturesPositionRiskV2Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountFuturesPositionRiskV2Resp(val *GetSubAccountFuturesPositionRiskV2Resp) *NullableGetSubAccountFuturesPositionRiskV2Resp {
	return &NullableGetSubAccountFuturesPositionRiskV2Resp{value: val, isSet: true}
}

func (v NullableGetSubAccountFuturesPositionRiskV2Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountFuturesPositionRiskV2Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


