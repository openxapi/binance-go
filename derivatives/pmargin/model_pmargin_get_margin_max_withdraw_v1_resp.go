/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginGetMarginMaxWithdrawV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetMarginMaxWithdrawV1Resp{}

// PmarginGetMarginMaxWithdrawV1Resp struct for PmarginGetMarginMaxWithdrawV1Resp
type PmarginGetMarginMaxWithdrawV1Resp struct {
	Amount *string `json:"amount,omitempty"`
}

// NewPmarginGetMarginMaxWithdrawV1Resp instantiates a new PmarginGetMarginMaxWithdrawV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetMarginMaxWithdrawV1Resp() *PmarginGetMarginMaxWithdrawV1Resp {
	this := PmarginGetMarginMaxWithdrawV1Resp{}
	return &this
}

// NewPmarginGetMarginMaxWithdrawV1RespWithDefaults instantiates a new PmarginGetMarginMaxWithdrawV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetMarginMaxWithdrawV1RespWithDefaults() *PmarginGetMarginMaxWithdrawV1Resp {
	this := PmarginGetMarginMaxWithdrawV1Resp{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PmarginGetMarginMaxWithdrawV1Resp) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginMaxWithdrawV1Resp) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PmarginGetMarginMaxWithdrawV1Resp) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *PmarginGetMarginMaxWithdrawV1Resp) SetAmount(v string) {
	o.Amount = &v
}

func (o PmarginGetMarginMaxWithdrawV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetMarginMaxWithdrawV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullablePmarginGetMarginMaxWithdrawV1Resp struct {
	value *PmarginGetMarginMaxWithdrawV1Resp
	isSet bool
}

func (v NullablePmarginGetMarginMaxWithdrawV1Resp) Get() *PmarginGetMarginMaxWithdrawV1Resp {
	return v.value
}

func (v *NullablePmarginGetMarginMaxWithdrawV1Resp) Set(val *PmarginGetMarginMaxWithdrawV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetMarginMaxWithdrawV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetMarginMaxWithdrawV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetMarginMaxWithdrawV1Resp(val *PmarginGetMarginMaxWithdrawV1Resp) *NullablePmarginGetMarginMaxWithdrawV1Resp {
	return &NullablePmarginGetMarginMaxWithdrawV1Resp{value: val, isSet: true}
}

func (v NullablePmarginGetMarginMaxWithdrawV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetMarginMaxWithdrawV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


