/*
Binance Sub Account API

OpenAPI specification for Binance exchange - Subaccount API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package subaccount

import (
	"encoding/json"
)

// checks if the SubaccountGetSubAccountFuturesAccountV2Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountGetSubAccountFuturesAccountV2Resp{}

// SubaccountGetSubAccountFuturesAccountV2Resp struct for SubaccountGetSubAccountFuturesAccountV2Resp
type SubaccountGetSubAccountFuturesAccountV2Resp struct {
	FutureAccountResp *SubaccountGetSubAccountFuturesAccountV2RespFutureAccountResp `json:"futureAccountResp,omitempty"`
}

// NewSubaccountGetSubAccountFuturesAccountV2Resp instantiates a new SubaccountGetSubAccountFuturesAccountV2Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountGetSubAccountFuturesAccountV2Resp() *SubaccountGetSubAccountFuturesAccountV2Resp {
	this := SubaccountGetSubAccountFuturesAccountV2Resp{}
	return &this
}

// NewSubaccountGetSubAccountFuturesAccountV2RespWithDefaults instantiates a new SubaccountGetSubAccountFuturesAccountV2Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountGetSubAccountFuturesAccountV2RespWithDefaults() *SubaccountGetSubAccountFuturesAccountV2Resp {
	this := SubaccountGetSubAccountFuturesAccountV2Resp{}
	return &this
}

// GetFutureAccountResp returns the FutureAccountResp field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountFuturesAccountV2Resp) GetFutureAccountResp() SubaccountGetSubAccountFuturesAccountV2RespFutureAccountResp {
	if o == nil || IsNil(o.FutureAccountResp) {
		var ret SubaccountGetSubAccountFuturesAccountV2RespFutureAccountResp
		return ret
	}
	return *o.FutureAccountResp
}

// GetFutureAccountRespOk returns a tuple with the FutureAccountResp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountFuturesAccountV2Resp) GetFutureAccountRespOk() (*SubaccountGetSubAccountFuturesAccountV2RespFutureAccountResp, bool) {
	if o == nil || IsNil(o.FutureAccountResp) {
		return nil, false
	}
	return o.FutureAccountResp, true
}

// HasFutureAccountResp returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountFuturesAccountV2Resp) HasFutureAccountResp() bool {
	if o != nil && !IsNil(o.FutureAccountResp) {
		return true
	}

	return false
}

// SetFutureAccountResp gets a reference to the given SubaccountGetSubAccountFuturesAccountV2RespFutureAccountResp and assigns it to the FutureAccountResp field.
func (o *SubaccountGetSubAccountFuturesAccountV2Resp) SetFutureAccountResp(v SubaccountGetSubAccountFuturesAccountV2RespFutureAccountResp) {
	o.FutureAccountResp = &v
}

func (o SubaccountGetSubAccountFuturesAccountV2Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountGetSubAccountFuturesAccountV2Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FutureAccountResp) {
		toSerialize["futureAccountResp"] = o.FutureAccountResp
	}
	return toSerialize, nil
}

type NullableSubaccountGetSubAccountFuturesAccountV2Resp struct {
	value *SubaccountGetSubAccountFuturesAccountV2Resp
	isSet bool
}

func (v NullableSubaccountGetSubAccountFuturesAccountV2Resp) Get() *SubaccountGetSubAccountFuturesAccountV2Resp {
	return v.value
}

func (v *NullableSubaccountGetSubAccountFuturesAccountV2Resp) Set(val *SubaccountGetSubAccountFuturesAccountV2Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountGetSubAccountFuturesAccountV2Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountGetSubAccountFuturesAccountV2Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountGetSubAccountFuturesAccountV2Resp(val *SubaccountGetSubAccountFuturesAccountV2Resp) *NullableSubaccountGetSubAccountFuturesAccountV2Resp {
	return &NullableSubaccountGetSubAccountFuturesAccountV2Resp{value: val, isSet: true}
}

func (v NullableSubaccountGetSubAccountFuturesAccountV2Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountGetSubAccountFuturesAccountV2Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


