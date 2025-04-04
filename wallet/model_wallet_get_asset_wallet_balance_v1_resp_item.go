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

// checks if the WalletGetAssetWalletBalanceV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletGetAssetWalletBalanceV1RespItem{}

// WalletGetAssetWalletBalanceV1RespItem struct for WalletGetAssetWalletBalanceV1RespItem
type WalletGetAssetWalletBalanceV1RespItem struct {
	Activate *bool `json:"activate,omitempty"`
	Balance *string `json:"balance,omitempty"`
	WalletName *string `json:"walletName,omitempty"`
}

// NewWalletGetAssetWalletBalanceV1RespItem instantiates a new WalletGetAssetWalletBalanceV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetAssetWalletBalanceV1RespItem() *WalletGetAssetWalletBalanceV1RespItem {
	this := WalletGetAssetWalletBalanceV1RespItem{}
	return &this
}

// NewWalletGetAssetWalletBalanceV1RespItemWithDefaults instantiates a new WalletGetAssetWalletBalanceV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetAssetWalletBalanceV1RespItemWithDefaults() *WalletGetAssetWalletBalanceV1RespItem {
	this := WalletGetAssetWalletBalanceV1RespItem{}
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *WalletGetAssetWalletBalanceV1RespItem) GetActivate() bool {
	if o == nil || IsNil(o.Activate) {
		var ret bool
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetWalletBalanceV1RespItem) GetActivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *WalletGetAssetWalletBalanceV1RespItem) HasActivate() bool {
	if o != nil && !IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given bool and assigns it to the Activate field.
func (o *WalletGetAssetWalletBalanceV1RespItem) SetActivate(v bool) {
	o.Activate = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *WalletGetAssetWalletBalanceV1RespItem) GetBalance() string {
	if o == nil || IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetWalletBalanceV1RespItem) GetBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *WalletGetAssetWalletBalanceV1RespItem) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *WalletGetAssetWalletBalanceV1RespItem) SetBalance(v string) {
	o.Balance = &v
}

// GetWalletName returns the WalletName field value if set, zero value otherwise.
func (o *WalletGetAssetWalletBalanceV1RespItem) GetWalletName() string {
	if o == nil || IsNil(o.WalletName) {
		var ret string
		return ret
	}
	return *o.WalletName
}

// GetWalletNameOk returns a tuple with the WalletName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetWalletBalanceV1RespItem) GetWalletNameOk() (*string, bool) {
	if o == nil || IsNil(o.WalletName) {
		return nil, false
	}
	return o.WalletName, true
}

// HasWalletName returns a boolean if a field has been set.
func (o *WalletGetAssetWalletBalanceV1RespItem) HasWalletName() bool {
	if o != nil && !IsNil(o.WalletName) {
		return true
	}

	return false
}

// SetWalletName gets a reference to the given string and assigns it to the WalletName field.
func (o *WalletGetAssetWalletBalanceV1RespItem) SetWalletName(v string) {
	o.WalletName = &v
}

func (o WalletGetAssetWalletBalanceV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletGetAssetWalletBalanceV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.WalletName) {
		toSerialize["walletName"] = o.WalletName
	}
	return toSerialize, nil
}

type NullableWalletGetAssetWalletBalanceV1RespItem struct {
	value *WalletGetAssetWalletBalanceV1RespItem
	isSet bool
}

func (v NullableWalletGetAssetWalletBalanceV1RespItem) Get() *WalletGetAssetWalletBalanceV1RespItem {
	return v.value
}

func (v *NullableWalletGetAssetWalletBalanceV1RespItem) Set(val *WalletGetAssetWalletBalanceV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetAssetWalletBalanceV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetAssetWalletBalanceV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetAssetWalletBalanceV1RespItem(val *WalletGetAssetWalletBalanceV1RespItem) *NullableWalletGetAssetWalletBalanceV1RespItem {
	return &NullableWalletGetAssetWalletBalanceV1RespItem{value: val, isSet: true}
}

func (v NullableWalletGetAssetWalletBalanceV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetAssetWalletBalanceV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


