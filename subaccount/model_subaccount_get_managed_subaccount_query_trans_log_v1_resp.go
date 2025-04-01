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

// checks if the SubaccountGetManagedSubaccountQueryTransLogV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountGetManagedSubaccountQueryTransLogV1Resp{}

// SubaccountGetManagedSubaccountQueryTransLogV1Resp struct for SubaccountGetManagedSubaccountQueryTransLogV1Resp
type SubaccountGetManagedSubaccountQueryTransLogV1Resp struct {
	Count *int32 `json:"count,omitempty"`
	ManagerSubTransferHistoryVos []SubaccountGetManagedSubaccountQueryTransLogForInvestorV1RespManagerSubTransferHistoryVosInner `json:"managerSubTransferHistoryVos,omitempty"`
}

// NewSubaccountGetManagedSubaccountQueryTransLogV1Resp instantiates a new SubaccountGetManagedSubaccountQueryTransLogV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountGetManagedSubaccountQueryTransLogV1Resp() *SubaccountGetManagedSubaccountQueryTransLogV1Resp {
	this := SubaccountGetManagedSubaccountQueryTransLogV1Resp{}
	return &this
}

// NewSubaccountGetManagedSubaccountQueryTransLogV1RespWithDefaults instantiates a new SubaccountGetManagedSubaccountQueryTransLogV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountGetManagedSubaccountQueryTransLogV1RespWithDefaults() *SubaccountGetManagedSubaccountQueryTransLogV1Resp {
	this := SubaccountGetManagedSubaccountQueryTransLogV1Resp{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SubaccountGetManagedSubaccountQueryTransLogV1Resp) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetManagedSubaccountQueryTransLogV1Resp) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SubaccountGetManagedSubaccountQueryTransLogV1Resp) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SubaccountGetManagedSubaccountQueryTransLogV1Resp) SetCount(v int32) {
	o.Count = &v
}

// GetManagerSubTransferHistoryVos returns the ManagerSubTransferHistoryVos field value if set, zero value otherwise.
func (o *SubaccountGetManagedSubaccountQueryTransLogV1Resp) GetManagerSubTransferHistoryVos() []SubaccountGetManagedSubaccountQueryTransLogForInvestorV1RespManagerSubTransferHistoryVosInner {
	if o == nil || IsNil(o.ManagerSubTransferHistoryVos) {
		var ret []SubaccountGetManagedSubaccountQueryTransLogForInvestorV1RespManagerSubTransferHistoryVosInner
		return ret
	}
	return o.ManagerSubTransferHistoryVos
}

// GetManagerSubTransferHistoryVosOk returns a tuple with the ManagerSubTransferHistoryVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetManagedSubaccountQueryTransLogV1Resp) GetManagerSubTransferHistoryVosOk() ([]SubaccountGetManagedSubaccountQueryTransLogForInvestorV1RespManagerSubTransferHistoryVosInner, bool) {
	if o == nil || IsNil(o.ManagerSubTransferHistoryVos) {
		return nil, false
	}
	return o.ManagerSubTransferHistoryVos, true
}

// HasManagerSubTransferHistoryVos returns a boolean if a field has been set.
func (o *SubaccountGetManagedSubaccountQueryTransLogV1Resp) HasManagerSubTransferHistoryVos() bool {
	if o != nil && !IsNil(o.ManagerSubTransferHistoryVos) {
		return true
	}

	return false
}

// SetManagerSubTransferHistoryVos gets a reference to the given []SubaccountGetManagedSubaccountQueryTransLogForInvestorV1RespManagerSubTransferHistoryVosInner and assigns it to the ManagerSubTransferHistoryVos field.
func (o *SubaccountGetManagedSubaccountQueryTransLogV1Resp) SetManagerSubTransferHistoryVos(v []SubaccountGetManagedSubaccountQueryTransLogForInvestorV1RespManagerSubTransferHistoryVosInner) {
	o.ManagerSubTransferHistoryVos = v
}

func (o SubaccountGetManagedSubaccountQueryTransLogV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountGetManagedSubaccountQueryTransLogV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.ManagerSubTransferHistoryVos) {
		toSerialize["managerSubTransferHistoryVos"] = o.ManagerSubTransferHistoryVos
	}
	return toSerialize, nil
}

type NullableSubaccountGetManagedSubaccountQueryTransLogV1Resp struct {
	value *SubaccountGetManagedSubaccountQueryTransLogV1Resp
	isSet bool
}

func (v NullableSubaccountGetManagedSubaccountQueryTransLogV1Resp) Get() *SubaccountGetManagedSubaccountQueryTransLogV1Resp {
	return v.value
}

func (v *NullableSubaccountGetManagedSubaccountQueryTransLogV1Resp) Set(val *SubaccountGetManagedSubaccountQueryTransLogV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountGetManagedSubaccountQueryTransLogV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountGetManagedSubaccountQueryTransLogV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountGetManagedSubaccountQueryTransLogV1Resp(val *SubaccountGetManagedSubaccountQueryTransLogV1Resp) *NullableSubaccountGetManagedSubaccountQueryTransLogV1Resp {
	return &NullableSubaccountGetManagedSubaccountQueryTransLogV1Resp{value: val, isSet: true}
}

func (v NullableSubaccountGetManagedSubaccountQueryTransLogV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountGetManagedSubaccountQueryTransLogV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


