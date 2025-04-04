/*
Binance Wallet API

OpenAPI specification for Binance exchange - Wallet API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wallet

import (
	"encoding/json"
)

// checks if the WalletGetAccountApiTradingStatusV1RespDataTriggerCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletGetAccountApiTradingStatusV1RespDataTriggerCondition{}

// WalletGetAccountApiTradingStatusV1RespDataTriggerCondition struct for WalletGetAccountApiTradingStatusV1RespDataTriggerCondition
type WalletGetAccountApiTradingStatusV1RespDataTriggerCondition struct {
	GCR *int32 `json:"GCR,omitempty"`
	IFER *int32 `json:"IFER,omitempty"`
	UFR *int32 `json:"UFR,omitempty"`
}

// NewWalletGetAccountApiTradingStatusV1RespDataTriggerCondition instantiates a new WalletGetAccountApiTradingStatusV1RespDataTriggerCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetAccountApiTradingStatusV1RespDataTriggerCondition() *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition {
	this := WalletGetAccountApiTradingStatusV1RespDataTriggerCondition{}
	return &this
}

// NewWalletGetAccountApiTradingStatusV1RespDataTriggerConditionWithDefaults instantiates a new WalletGetAccountApiTradingStatusV1RespDataTriggerCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetAccountApiTradingStatusV1RespDataTriggerConditionWithDefaults() *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition {
	this := WalletGetAccountApiTradingStatusV1RespDataTriggerCondition{}
	return &this
}

// GetGCR returns the GCR field value if set, zero value otherwise.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) GetGCR() int32 {
	if o == nil || IsNil(o.GCR) {
		var ret int32
		return ret
	}
	return *o.GCR
}

// GetGCROk returns a tuple with the GCR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) GetGCROk() (*int32, bool) {
	if o == nil || IsNil(o.GCR) {
		return nil, false
	}
	return o.GCR, true
}

// HasGCR returns a boolean if a field has been set.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) HasGCR() bool {
	if o != nil && !IsNil(o.GCR) {
		return true
	}

	return false
}

// SetGCR gets a reference to the given int32 and assigns it to the GCR field.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) SetGCR(v int32) {
	o.GCR = &v
}

// GetIFER returns the IFER field value if set, zero value otherwise.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) GetIFER() int32 {
	if o == nil || IsNil(o.IFER) {
		var ret int32
		return ret
	}
	return *o.IFER
}

// GetIFEROk returns a tuple with the IFER field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) GetIFEROk() (*int32, bool) {
	if o == nil || IsNil(o.IFER) {
		return nil, false
	}
	return o.IFER, true
}

// HasIFER returns a boolean if a field has been set.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) HasIFER() bool {
	if o != nil && !IsNil(o.IFER) {
		return true
	}

	return false
}

// SetIFER gets a reference to the given int32 and assigns it to the IFER field.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) SetIFER(v int32) {
	o.IFER = &v
}

// GetUFR returns the UFR field value if set, zero value otherwise.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) GetUFR() int32 {
	if o == nil || IsNil(o.UFR) {
		var ret int32
		return ret
	}
	return *o.UFR
}

// GetUFROk returns a tuple with the UFR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) GetUFROk() (*int32, bool) {
	if o == nil || IsNil(o.UFR) {
		return nil, false
	}
	return o.UFR, true
}

// HasUFR returns a boolean if a field has been set.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) HasUFR() bool {
	if o != nil && !IsNil(o.UFR) {
		return true
	}

	return false
}

// SetUFR gets a reference to the given int32 and assigns it to the UFR field.
func (o *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) SetUFR(v int32) {
	o.UFR = &v
}

func (o WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GCR) {
		toSerialize["GCR"] = o.GCR
	}
	if !IsNil(o.IFER) {
		toSerialize["IFER"] = o.IFER
	}
	if !IsNil(o.UFR) {
		toSerialize["UFR"] = o.UFR
	}
	return toSerialize, nil
}

type NullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition struct {
	value *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition
	isSet bool
}

func (v NullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition) Get() *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition {
	return v.value
}

func (v *NullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition) Set(val *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition(val *WalletGetAccountApiTradingStatusV1RespDataTriggerCondition) *NullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition {
	return &NullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition{value: val, isSet: true}
}

func (v NullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetAccountApiTradingStatusV1RespDataTriggerCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


