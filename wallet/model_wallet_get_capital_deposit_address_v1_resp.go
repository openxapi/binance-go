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

// checks if the WalletGetCapitalDepositAddressV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletGetCapitalDepositAddressV1Resp{}

// WalletGetCapitalDepositAddressV1Resp struct for WalletGetCapitalDepositAddressV1Resp
type WalletGetCapitalDepositAddressV1Resp struct {
	Address *string `json:"address,omitempty"`
	Coin *string `json:"coin,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewWalletGetCapitalDepositAddressV1Resp instantiates a new WalletGetCapitalDepositAddressV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetCapitalDepositAddressV1Resp() *WalletGetCapitalDepositAddressV1Resp {
	this := WalletGetCapitalDepositAddressV1Resp{}
	return &this
}

// NewWalletGetCapitalDepositAddressV1RespWithDefaults instantiates a new WalletGetCapitalDepositAddressV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetCapitalDepositAddressV1RespWithDefaults() *WalletGetCapitalDepositAddressV1Resp {
	this := WalletGetCapitalDepositAddressV1Resp{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *WalletGetCapitalDepositAddressV1Resp) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetCapitalDepositAddressV1Resp) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *WalletGetCapitalDepositAddressV1Resp) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *WalletGetCapitalDepositAddressV1Resp) SetAddress(v string) {
	o.Address = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *WalletGetCapitalDepositAddressV1Resp) GetCoin() string {
	if o == nil || IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetCapitalDepositAddressV1Resp) GetCoinOk() (*string, bool) {
	if o == nil || IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *WalletGetCapitalDepositAddressV1Resp) HasCoin() bool {
	if o != nil && !IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *WalletGetCapitalDepositAddressV1Resp) SetCoin(v string) {
	o.Coin = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *WalletGetCapitalDepositAddressV1Resp) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetCapitalDepositAddressV1Resp) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *WalletGetCapitalDepositAddressV1Resp) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *WalletGetCapitalDepositAddressV1Resp) SetTag(v string) {
	o.Tag = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WalletGetCapitalDepositAddressV1Resp) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetCapitalDepositAddressV1Resp) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WalletGetCapitalDepositAddressV1Resp) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WalletGetCapitalDepositAddressV1Resp) SetUrl(v string) {
	o.Url = &v
}

func (o WalletGetCapitalDepositAddressV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletGetCapitalDepositAddressV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableWalletGetCapitalDepositAddressV1Resp struct {
	value *WalletGetCapitalDepositAddressV1Resp
	isSet bool
}

func (v NullableWalletGetCapitalDepositAddressV1Resp) Get() *WalletGetCapitalDepositAddressV1Resp {
	return v.value
}

func (v *NullableWalletGetCapitalDepositAddressV1Resp) Set(val *WalletGetCapitalDepositAddressV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetCapitalDepositAddressV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetCapitalDepositAddressV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetCapitalDepositAddressV1Resp(val *WalletGetCapitalDepositAddressV1Resp) *NullableWalletGetCapitalDepositAddressV1Resp {
	return &NullableWalletGetCapitalDepositAddressV1Resp{value: val, isSet: true}
}

func (v NullableWalletGetCapitalDepositAddressV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetCapitalDepositAddressV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


