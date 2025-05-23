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

// checks if the GetSubAccountUniversalTransferV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubAccountUniversalTransferV1Resp{}

// GetSubAccountUniversalTransferV1Resp struct for GetSubAccountUniversalTransferV1Resp
type GetSubAccountUniversalTransferV1Resp struct {
	Result []GetSubAccountUniversalTransferV1RespResultInner `json:"result,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewGetSubAccountUniversalTransferV1Resp instantiates a new GetSubAccountUniversalTransferV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountUniversalTransferV1Resp() *GetSubAccountUniversalTransferV1Resp {
	this := GetSubAccountUniversalTransferV1Resp{}
	return &this
}

// NewGetSubAccountUniversalTransferV1RespWithDefaults instantiates a new GetSubAccountUniversalTransferV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountUniversalTransferV1RespWithDefaults() *GetSubAccountUniversalTransferV1Resp {
	this := GetSubAccountUniversalTransferV1Resp{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetSubAccountUniversalTransferV1Resp) GetResult() []GetSubAccountUniversalTransferV1RespResultInner {
	if o == nil || IsNil(o.Result) {
		var ret []GetSubAccountUniversalTransferV1RespResultInner
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountUniversalTransferV1Resp) GetResultOk() ([]GetSubAccountUniversalTransferV1RespResultInner, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetSubAccountUniversalTransferV1Resp) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []GetSubAccountUniversalTransferV1RespResultInner and assigns it to the Result field.
func (o *GetSubAccountUniversalTransferV1Resp) SetResult(v []GetSubAccountUniversalTransferV1RespResultInner) {
	o.Result = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *GetSubAccountUniversalTransferV1Resp) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountUniversalTransferV1Resp) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *GetSubAccountUniversalTransferV1Resp) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *GetSubAccountUniversalTransferV1Resp) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o GetSubAccountUniversalTransferV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountUniversalTransferV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableGetSubAccountUniversalTransferV1Resp struct {
	value *GetSubAccountUniversalTransferV1Resp
	isSet bool
}

func (v NullableGetSubAccountUniversalTransferV1Resp) Get() *GetSubAccountUniversalTransferV1Resp {
	return v.value
}

func (v *NullableGetSubAccountUniversalTransferV1Resp) Set(val *GetSubAccountUniversalTransferV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountUniversalTransferV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountUniversalTransferV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountUniversalTransferV1Resp(val *GetSubAccountUniversalTransferV1Resp) *NullableGetSubAccountUniversalTransferV1Resp {
	return &NullableGetSubAccountUniversalTransferV1Resp{value: val, isSet: true}
}

func (v NullableGetSubAccountUniversalTransferV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountUniversalTransferV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


