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

// checks if the WalletGetAssetTradeFeeV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletGetAssetTradeFeeV1RespItem{}

// WalletGetAssetTradeFeeV1RespItem struct for WalletGetAssetTradeFeeV1RespItem
type WalletGetAssetTradeFeeV1RespItem struct {
	MakerCommission *string `json:"makerCommission,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TakerCommission *string `json:"takerCommission,omitempty"`
}

// NewWalletGetAssetTradeFeeV1RespItem instantiates a new WalletGetAssetTradeFeeV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetAssetTradeFeeV1RespItem() *WalletGetAssetTradeFeeV1RespItem {
	this := WalletGetAssetTradeFeeV1RespItem{}
	return &this
}

// NewWalletGetAssetTradeFeeV1RespItemWithDefaults instantiates a new WalletGetAssetTradeFeeV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetAssetTradeFeeV1RespItemWithDefaults() *WalletGetAssetTradeFeeV1RespItem {
	this := WalletGetAssetTradeFeeV1RespItem{}
	return &this
}

// GetMakerCommission returns the MakerCommission field value if set, zero value otherwise.
func (o *WalletGetAssetTradeFeeV1RespItem) GetMakerCommission() string {
	if o == nil || IsNil(o.MakerCommission) {
		var ret string
		return ret
	}
	return *o.MakerCommission
}

// GetMakerCommissionOk returns a tuple with the MakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTradeFeeV1RespItem) GetMakerCommissionOk() (*string, bool) {
	if o == nil || IsNil(o.MakerCommission) {
		return nil, false
	}
	return o.MakerCommission, true
}

// HasMakerCommission returns a boolean if a field has been set.
func (o *WalletGetAssetTradeFeeV1RespItem) HasMakerCommission() bool {
	if o != nil && !IsNil(o.MakerCommission) {
		return true
	}

	return false
}

// SetMakerCommission gets a reference to the given string and assigns it to the MakerCommission field.
func (o *WalletGetAssetTradeFeeV1RespItem) SetMakerCommission(v string) {
	o.MakerCommission = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *WalletGetAssetTradeFeeV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTradeFeeV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *WalletGetAssetTradeFeeV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *WalletGetAssetTradeFeeV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTakerCommission returns the TakerCommission field value if set, zero value otherwise.
func (o *WalletGetAssetTradeFeeV1RespItem) GetTakerCommission() string {
	if o == nil || IsNil(o.TakerCommission) {
		var ret string
		return ret
	}
	return *o.TakerCommission
}

// GetTakerCommissionOk returns a tuple with the TakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTradeFeeV1RespItem) GetTakerCommissionOk() (*string, bool) {
	if o == nil || IsNil(o.TakerCommission) {
		return nil, false
	}
	return o.TakerCommission, true
}

// HasTakerCommission returns a boolean if a field has been set.
func (o *WalletGetAssetTradeFeeV1RespItem) HasTakerCommission() bool {
	if o != nil && !IsNil(o.TakerCommission) {
		return true
	}

	return false
}

// SetTakerCommission gets a reference to the given string and assigns it to the TakerCommission field.
func (o *WalletGetAssetTradeFeeV1RespItem) SetTakerCommission(v string) {
	o.TakerCommission = &v
}

func (o WalletGetAssetTradeFeeV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletGetAssetTradeFeeV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MakerCommission) {
		toSerialize["makerCommission"] = o.MakerCommission
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.TakerCommission) {
		toSerialize["takerCommission"] = o.TakerCommission
	}
	return toSerialize, nil
}

type NullableWalletGetAssetTradeFeeV1RespItem struct {
	value *WalletGetAssetTradeFeeV1RespItem
	isSet bool
}

func (v NullableWalletGetAssetTradeFeeV1RespItem) Get() *WalletGetAssetTradeFeeV1RespItem {
	return v.value
}

func (v *NullableWalletGetAssetTradeFeeV1RespItem) Set(val *WalletGetAssetTradeFeeV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetAssetTradeFeeV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetAssetTradeFeeV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetAssetTradeFeeV1RespItem(val *WalletGetAssetTradeFeeV1RespItem) *NullableWalletGetAssetTradeFeeV1RespItem {
	return &NullableWalletGetAssetTradeFeeV1RespItem{value: val, isSet: true}
}

func (v NullableWalletGetAssetTradeFeeV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetAssetTradeFeeV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


