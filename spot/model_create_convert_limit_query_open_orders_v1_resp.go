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

// checks if the CreateConvertLimitQueryOpenOrdersV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConvertLimitQueryOpenOrdersV1Resp{}

// CreateConvertLimitQueryOpenOrdersV1Resp struct for CreateConvertLimitQueryOpenOrdersV1Resp
type CreateConvertLimitQueryOpenOrdersV1Resp struct {
	List []CreateConvertLimitQueryOpenOrdersV1RespListInner `json:"list,omitempty"`
}

// NewCreateConvertLimitQueryOpenOrdersV1Resp instantiates a new CreateConvertLimitQueryOpenOrdersV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConvertLimitQueryOpenOrdersV1Resp() *CreateConvertLimitQueryOpenOrdersV1Resp {
	this := CreateConvertLimitQueryOpenOrdersV1Resp{}
	return &this
}

// NewCreateConvertLimitQueryOpenOrdersV1RespWithDefaults instantiates a new CreateConvertLimitQueryOpenOrdersV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConvertLimitQueryOpenOrdersV1RespWithDefaults() *CreateConvertLimitQueryOpenOrdersV1Resp {
	this := CreateConvertLimitQueryOpenOrdersV1Resp{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *CreateConvertLimitQueryOpenOrdersV1Resp) GetList() []CreateConvertLimitQueryOpenOrdersV1RespListInner {
	if o == nil || IsNil(o.List) {
		var ret []CreateConvertLimitQueryOpenOrdersV1RespListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConvertLimitQueryOpenOrdersV1Resp) GetListOk() ([]CreateConvertLimitQueryOpenOrdersV1RespListInner, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *CreateConvertLimitQueryOpenOrdersV1Resp) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []CreateConvertLimitQueryOpenOrdersV1RespListInner and assigns it to the List field.
func (o *CreateConvertLimitQueryOpenOrdersV1Resp) SetList(v []CreateConvertLimitQueryOpenOrdersV1RespListInner) {
	o.List = v
}

func (o CreateConvertLimitQueryOpenOrdersV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateConvertLimitQueryOpenOrdersV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableCreateConvertLimitQueryOpenOrdersV1Resp struct {
	value *CreateConvertLimitQueryOpenOrdersV1Resp
	isSet bool
}

func (v NullableCreateConvertLimitQueryOpenOrdersV1Resp) Get() *CreateConvertLimitQueryOpenOrdersV1Resp {
	return v.value
}

func (v *NullableCreateConvertLimitQueryOpenOrdersV1Resp) Set(val *CreateConvertLimitQueryOpenOrdersV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConvertLimitQueryOpenOrdersV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConvertLimitQueryOpenOrdersV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConvertLimitQueryOpenOrdersV1Resp(val *CreateConvertLimitQueryOpenOrdersV1Resp) *NullableCreateConvertLimitQueryOpenOrdersV1Resp {
	return &NullableCreateConvertLimitQueryOpenOrdersV1Resp{value: val, isSet: true}
}

func (v NullableCreateConvertLimitQueryOpenOrdersV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConvertLimitQueryOpenOrdersV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


