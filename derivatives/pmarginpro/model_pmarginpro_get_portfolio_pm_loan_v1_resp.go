/*
Binance Portfolio Margin Pro API

OpenAPI specification for Binance exchange - Pmarginpro API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmarginpro

import (
	"encoding/json"
)

// checks if the PmarginproGetPortfolioPmLoanV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginproGetPortfolioPmLoanV1Resp{}

// PmarginproGetPortfolioPmLoanV1Resp struct for PmarginproGetPortfolioPmLoanV1Resp
type PmarginproGetPortfolioPmLoanV1Resp struct {
	Amount *string `json:"amount,omitempty"`
	Asset *string `json:"asset,omitempty"`
}

// NewPmarginproGetPortfolioPmLoanV1Resp instantiates a new PmarginproGetPortfolioPmLoanV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginproGetPortfolioPmLoanV1Resp() *PmarginproGetPortfolioPmLoanV1Resp {
	this := PmarginproGetPortfolioPmLoanV1Resp{}
	return &this
}

// NewPmarginproGetPortfolioPmLoanV1RespWithDefaults instantiates a new PmarginproGetPortfolioPmLoanV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginproGetPortfolioPmLoanV1RespWithDefaults() *PmarginproGetPortfolioPmLoanV1Resp {
	this := PmarginproGetPortfolioPmLoanV1Resp{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PmarginproGetPortfolioPmLoanV1Resp) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproGetPortfolioPmLoanV1Resp) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PmarginproGetPortfolioPmLoanV1Resp) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *PmarginproGetPortfolioPmLoanV1Resp) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *PmarginproGetPortfolioPmLoanV1Resp) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproGetPortfolioPmLoanV1Resp) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *PmarginproGetPortfolioPmLoanV1Resp) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *PmarginproGetPortfolioPmLoanV1Resp) SetAsset(v string) {
	o.Asset = &v
}

func (o PmarginproGetPortfolioPmLoanV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginproGetPortfolioPmLoanV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	return toSerialize, nil
}

type NullablePmarginproGetPortfolioPmLoanV1Resp struct {
	value *PmarginproGetPortfolioPmLoanV1Resp
	isSet bool
}

func (v NullablePmarginproGetPortfolioPmLoanV1Resp) Get() *PmarginproGetPortfolioPmLoanV1Resp {
	return v.value
}

func (v *NullablePmarginproGetPortfolioPmLoanV1Resp) Set(val *PmarginproGetPortfolioPmLoanV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginproGetPortfolioPmLoanV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginproGetPortfolioPmLoanV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginproGetPortfolioPmLoanV1Resp(val *PmarginproGetPortfolioPmLoanV1Resp) *NullablePmarginproGetPortfolioPmLoanV1Resp {
	return &NullablePmarginproGetPortfolioPmLoanV1Resp{value: val, isSet: true}
}

func (v NullablePmarginproGetPortfolioPmLoanV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginproGetPortfolioPmLoanV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


