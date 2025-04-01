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

// checks if the WalletGetAssetAssetDetailV1RespValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletGetAssetAssetDetailV1RespValue{}

// WalletGetAssetAssetDetailV1RespValue struct for WalletGetAssetAssetDetailV1RespValue
type WalletGetAssetAssetDetailV1RespValue struct {
	DepositStatus *bool `json:"depositStatus,omitempty"`
	DepositTip *string `json:"depositTip,omitempty"`
	MinWithdrawAmount *string `json:"minWithdrawAmount,omitempty"`
	WithdrawFee *int32 `json:"withdrawFee,omitempty"`
	WithdrawStatus *bool `json:"withdrawStatus,omitempty"`
}

// NewWalletGetAssetAssetDetailV1RespValue instantiates a new WalletGetAssetAssetDetailV1RespValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetAssetAssetDetailV1RespValue() *WalletGetAssetAssetDetailV1RespValue {
	this := WalletGetAssetAssetDetailV1RespValue{}
	return &this
}

// NewWalletGetAssetAssetDetailV1RespValueWithDefaults instantiates a new WalletGetAssetAssetDetailV1RespValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetAssetAssetDetailV1RespValueWithDefaults() *WalletGetAssetAssetDetailV1RespValue {
	this := WalletGetAssetAssetDetailV1RespValue{}
	return &this
}

// GetDepositStatus returns the DepositStatus field value if set, zero value otherwise.
func (o *WalletGetAssetAssetDetailV1RespValue) GetDepositStatus() bool {
	if o == nil || IsNil(o.DepositStatus) {
		var ret bool
		return ret
	}
	return *o.DepositStatus
}

// GetDepositStatusOk returns a tuple with the DepositStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) GetDepositStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.DepositStatus) {
		return nil, false
	}
	return o.DepositStatus, true
}

// HasDepositStatus returns a boolean if a field has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) HasDepositStatus() bool {
	if o != nil && !IsNil(o.DepositStatus) {
		return true
	}

	return false
}

// SetDepositStatus gets a reference to the given bool and assigns it to the DepositStatus field.
func (o *WalletGetAssetAssetDetailV1RespValue) SetDepositStatus(v bool) {
	o.DepositStatus = &v
}

// GetDepositTip returns the DepositTip field value if set, zero value otherwise.
func (o *WalletGetAssetAssetDetailV1RespValue) GetDepositTip() string {
	if o == nil || IsNil(o.DepositTip) {
		var ret string
		return ret
	}
	return *o.DepositTip
}

// GetDepositTipOk returns a tuple with the DepositTip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) GetDepositTipOk() (*string, bool) {
	if o == nil || IsNil(o.DepositTip) {
		return nil, false
	}
	return o.DepositTip, true
}

// HasDepositTip returns a boolean if a field has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) HasDepositTip() bool {
	if o != nil && !IsNil(o.DepositTip) {
		return true
	}

	return false
}

// SetDepositTip gets a reference to the given string and assigns it to the DepositTip field.
func (o *WalletGetAssetAssetDetailV1RespValue) SetDepositTip(v string) {
	o.DepositTip = &v
}

// GetMinWithdrawAmount returns the MinWithdrawAmount field value if set, zero value otherwise.
func (o *WalletGetAssetAssetDetailV1RespValue) GetMinWithdrawAmount() string {
	if o == nil || IsNil(o.MinWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MinWithdrawAmount
}

// GetMinWithdrawAmountOk returns a tuple with the MinWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) GetMinWithdrawAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MinWithdrawAmount) {
		return nil, false
	}
	return o.MinWithdrawAmount, true
}

// HasMinWithdrawAmount returns a boolean if a field has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) HasMinWithdrawAmount() bool {
	if o != nil && !IsNil(o.MinWithdrawAmount) {
		return true
	}

	return false
}

// SetMinWithdrawAmount gets a reference to the given string and assigns it to the MinWithdrawAmount field.
func (o *WalletGetAssetAssetDetailV1RespValue) SetMinWithdrawAmount(v string) {
	o.MinWithdrawAmount = &v
}

// GetWithdrawFee returns the WithdrawFee field value if set, zero value otherwise.
func (o *WalletGetAssetAssetDetailV1RespValue) GetWithdrawFee() int32 {
	if o == nil || IsNil(o.WithdrawFee) {
		var ret int32
		return ret
	}
	return *o.WithdrawFee
}

// GetWithdrawFeeOk returns a tuple with the WithdrawFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) GetWithdrawFeeOk() (*int32, bool) {
	if o == nil || IsNil(o.WithdrawFee) {
		return nil, false
	}
	return o.WithdrawFee, true
}

// HasWithdrawFee returns a boolean if a field has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) HasWithdrawFee() bool {
	if o != nil && !IsNil(o.WithdrawFee) {
		return true
	}

	return false
}

// SetWithdrawFee gets a reference to the given int32 and assigns it to the WithdrawFee field.
func (o *WalletGetAssetAssetDetailV1RespValue) SetWithdrawFee(v int32) {
	o.WithdrawFee = &v
}

// GetWithdrawStatus returns the WithdrawStatus field value if set, zero value otherwise.
func (o *WalletGetAssetAssetDetailV1RespValue) GetWithdrawStatus() bool {
	if o == nil || IsNil(o.WithdrawStatus) {
		var ret bool
		return ret
	}
	return *o.WithdrawStatus
}

// GetWithdrawStatusOk returns a tuple with the WithdrawStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) GetWithdrawStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.WithdrawStatus) {
		return nil, false
	}
	return o.WithdrawStatus, true
}

// HasWithdrawStatus returns a boolean if a field has been set.
func (o *WalletGetAssetAssetDetailV1RespValue) HasWithdrawStatus() bool {
	if o != nil && !IsNil(o.WithdrawStatus) {
		return true
	}

	return false
}

// SetWithdrawStatus gets a reference to the given bool and assigns it to the WithdrawStatus field.
func (o *WalletGetAssetAssetDetailV1RespValue) SetWithdrawStatus(v bool) {
	o.WithdrawStatus = &v
}

func (o WalletGetAssetAssetDetailV1RespValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletGetAssetAssetDetailV1RespValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DepositStatus) {
		toSerialize["depositStatus"] = o.DepositStatus
	}
	if !IsNil(o.DepositTip) {
		toSerialize["depositTip"] = o.DepositTip
	}
	if !IsNil(o.MinWithdrawAmount) {
		toSerialize["minWithdrawAmount"] = o.MinWithdrawAmount
	}
	if !IsNil(o.WithdrawFee) {
		toSerialize["withdrawFee"] = o.WithdrawFee
	}
	if !IsNil(o.WithdrawStatus) {
		toSerialize["withdrawStatus"] = o.WithdrawStatus
	}
	return toSerialize, nil
}

type NullableWalletGetAssetAssetDetailV1RespValue struct {
	value *WalletGetAssetAssetDetailV1RespValue
	isSet bool
}

func (v NullableWalletGetAssetAssetDetailV1RespValue) Get() *WalletGetAssetAssetDetailV1RespValue {
	return v.value
}

func (v *NullableWalletGetAssetAssetDetailV1RespValue) Set(val *WalletGetAssetAssetDetailV1RespValue) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetAssetAssetDetailV1RespValue) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetAssetAssetDetailV1RespValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetAssetAssetDetailV1RespValue(val *WalletGetAssetAssetDetailV1RespValue) *NullableWalletGetAssetAssetDetailV1RespValue {
	return &NullableWalletGetAssetAssetDetailV1RespValue{value: val, isSet: true}
}

func (v NullableWalletGetAssetAssetDetailV1RespValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetAssetAssetDetailV1RespValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


