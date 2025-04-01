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

// checks if the SubaccountGetSubAccountAssetsV4Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountGetSubAccountAssetsV4Resp{}

// SubaccountGetSubAccountAssetsV4Resp struct for SubaccountGetSubAccountAssetsV4Resp
type SubaccountGetSubAccountAssetsV4Resp struct {
	Balances []SubaccountGetSubAccountAssetsV4RespBalancesInner `json:"balances,omitempty"`
}

// NewSubaccountGetSubAccountAssetsV4Resp instantiates a new SubaccountGetSubAccountAssetsV4Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountGetSubAccountAssetsV4Resp() *SubaccountGetSubAccountAssetsV4Resp {
	this := SubaccountGetSubAccountAssetsV4Resp{}
	return &this
}

// NewSubaccountGetSubAccountAssetsV4RespWithDefaults instantiates a new SubaccountGetSubAccountAssetsV4Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountGetSubAccountAssetsV4RespWithDefaults() *SubaccountGetSubAccountAssetsV4Resp {
	this := SubaccountGetSubAccountAssetsV4Resp{}
	return &this
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountAssetsV4Resp) GetBalances() []SubaccountGetSubAccountAssetsV4RespBalancesInner {
	if o == nil || IsNil(o.Balances) {
		var ret []SubaccountGetSubAccountAssetsV4RespBalancesInner
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountAssetsV4Resp) GetBalancesOk() ([]SubaccountGetSubAccountAssetsV4RespBalancesInner, bool) {
	if o == nil || IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountAssetsV4Resp) HasBalances() bool {
	if o != nil && !IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []SubaccountGetSubAccountAssetsV4RespBalancesInner and assigns it to the Balances field.
func (o *SubaccountGetSubAccountAssetsV4Resp) SetBalances(v []SubaccountGetSubAccountAssetsV4RespBalancesInner) {
	o.Balances = v
}

func (o SubaccountGetSubAccountAssetsV4Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountGetSubAccountAssetsV4Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	return toSerialize, nil
}

type NullableSubaccountGetSubAccountAssetsV4Resp struct {
	value *SubaccountGetSubAccountAssetsV4Resp
	isSet bool
}

func (v NullableSubaccountGetSubAccountAssetsV4Resp) Get() *SubaccountGetSubAccountAssetsV4Resp {
	return v.value
}

func (v *NullableSubaccountGetSubAccountAssetsV4Resp) Set(val *SubaccountGetSubAccountAssetsV4Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountGetSubAccountAssetsV4Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountGetSubAccountAssetsV4Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountGetSubAccountAssetsV4Resp(val *SubaccountGetSubAccountAssetsV4Resp) *NullableSubaccountGetSubAccountAssetsV4Resp {
	return &NullableSubaccountGetSubAccountAssetsV4Resp{value: val, isSet: true}
}

func (v NullableSubaccountGetSubAccountAssetsV4Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountGetSubAccountAssetsV4Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


