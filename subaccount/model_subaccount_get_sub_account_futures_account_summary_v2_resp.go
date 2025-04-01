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

// checks if the SubaccountGetSubAccountFuturesAccountSummaryV2Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountGetSubAccountFuturesAccountSummaryV2Resp{}

// SubaccountGetSubAccountFuturesAccountSummaryV2Resp struct for SubaccountGetSubAccountFuturesAccountSummaryV2Resp
type SubaccountGetSubAccountFuturesAccountSummaryV2Resp struct {
	FutureAccountSummaryResp *SubaccountGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp `json:"futureAccountSummaryResp,omitempty"`
}

// NewSubaccountGetSubAccountFuturesAccountSummaryV2Resp instantiates a new SubaccountGetSubAccountFuturesAccountSummaryV2Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountGetSubAccountFuturesAccountSummaryV2Resp() *SubaccountGetSubAccountFuturesAccountSummaryV2Resp {
	this := SubaccountGetSubAccountFuturesAccountSummaryV2Resp{}
	return &this
}

// NewSubaccountGetSubAccountFuturesAccountSummaryV2RespWithDefaults instantiates a new SubaccountGetSubAccountFuturesAccountSummaryV2Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountGetSubAccountFuturesAccountSummaryV2RespWithDefaults() *SubaccountGetSubAccountFuturesAccountSummaryV2Resp {
	this := SubaccountGetSubAccountFuturesAccountSummaryV2Resp{}
	return &this
}

// GetFutureAccountSummaryResp returns the FutureAccountSummaryResp field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountFuturesAccountSummaryV2Resp) GetFutureAccountSummaryResp() SubaccountGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp {
	if o == nil || IsNil(o.FutureAccountSummaryResp) {
		var ret SubaccountGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp
		return ret
	}
	return *o.FutureAccountSummaryResp
}

// GetFutureAccountSummaryRespOk returns a tuple with the FutureAccountSummaryResp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountFuturesAccountSummaryV2Resp) GetFutureAccountSummaryRespOk() (*SubaccountGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp, bool) {
	if o == nil || IsNil(o.FutureAccountSummaryResp) {
		return nil, false
	}
	return o.FutureAccountSummaryResp, true
}

// HasFutureAccountSummaryResp returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountFuturesAccountSummaryV2Resp) HasFutureAccountSummaryResp() bool {
	if o != nil && !IsNil(o.FutureAccountSummaryResp) {
		return true
	}

	return false
}

// SetFutureAccountSummaryResp gets a reference to the given SubaccountGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp and assigns it to the FutureAccountSummaryResp field.
func (o *SubaccountGetSubAccountFuturesAccountSummaryV2Resp) SetFutureAccountSummaryResp(v SubaccountGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) {
	o.FutureAccountSummaryResp = &v
}

func (o SubaccountGetSubAccountFuturesAccountSummaryV2Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountGetSubAccountFuturesAccountSummaryV2Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FutureAccountSummaryResp) {
		toSerialize["futureAccountSummaryResp"] = o.FutureAccountSummaryResp
	}
	return toSerialize, nil
}

type NullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp struct {
	value *SubaccountGetSubAccountFuturesAccountSummaryV2Resp
	isSet bool
}

func (v NullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp) Get() *SubaccountGetSubAccountFuturesAccountSummaryV2Resp {
	return v.value
}

func (v *NullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp) Set(val *SubaccountGetSubAccountFuturesAccountSummaryV2Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp(val *SubaccountGetSubAccountFuturesAccountSummaryV2Resp) *NullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp {
	return &NullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp{value: val, isSet: true}
}

func (v NullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountGetSubAccountFuturesAccountSummaryV2Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


