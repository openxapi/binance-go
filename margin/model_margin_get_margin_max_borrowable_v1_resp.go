/*
Binance Margin API

OpenAPI specification for Binance cryptocurrency exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginGetMarginMaxBorrowableV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginMaxBorrowableV1Resp{}

// MarginGetMarginMaxBorrowableV1Resp struct for MarginGetMarginMaxBorrowableV1Resp
type MarginGetMarginMaxBorrowableV1Resp struct {
	Amount *string `json:"amount,omitempty"`
	BorrowLimit *string `json:"borrowLimit,omitempty"`
}

// NewMarginGetMarginMaxBorrowableV1Resp instantiates a new MarginGetMarginMaxBorrowableV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginMaxBorrowableV1Resp() *MarginGetMarginMaxBorrowableV1Resp {
	this := MarginGetMarginMaxBorrowableV1Resp{}
	return &this
}

// NewMarginGetMarginMaxBorrowableV1RespWithDefaults instantiates a new MarginGetMarginMaxBorrowableV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginMaxBorrowableV1RespWithDefaults() *MarginGetMarginMaxBorrowableV1Resp {
	this := MarginGetMarginMaxBorrowableV1Resp{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MarginGetMarginMaxBorrowableV1Resp) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginMaxBorrowableV1Resp) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MarginGetMarginMaxBorrowableV1Resp) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MarginGetMarginMaxBorrowableV1Resp) SetAmount(v string) {
	o.Amount = &v
}

// GetBorrowLimit returns the BorrowLimit field value if set, zero value otherwise.
func (o *MarginGetMarginMaxBorrowableV1Resp) GetBorrowLimit() string {
	if o == nil || IsNil(o.BorrowLimit) {
		var ret string
		return ret
	}
	return *o.BorrowLimit
}

// GetBorrowLimitOk returns a tuple with the BorrowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginMaxBorrowableV1Resp) GetBorrowLimitOk() (*string, bool) {
	if o == nil || IsNil(o.BorrowLimit) {
		return nil, false
	}
	return o.BorrowLimit, true
}

// HasBorrowLimit returns a boolean if a field has been set.
func (o *MarginGetMarginMaxBorrowableV1Resp) HasBorrowLimit() bool {
	if o != nil && !IsNil(o.BorrowLimit) {
		return true
	}

	return false
}

// SetBorrowLimit gets a reference to the given string and assigns it to the BorrowLimit field.
func (o *MarginGetMarginMaxBorrowableV1Resp) SetBorrowLimit(v string) {
	o.BorrowLimit = &v
}

func (o MarginGetMarginMaxBorrowableV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginMaxBorrowableV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BorrowLimit) {
		toSerialize["borrowLimit"] = o.BorrowLimit
	}
	return toSerialize, nil
}

type NullableMarginGetMarginMaxBorrowableV1Resp struct {
	value *MarginGetMarginMaxBorrowableV1Resp
	isSet bool
}

func (v NullableMarginGetMarginMaxBorrowableV1Resp) Get() *MarginGetMarginMaxBorrowableV1Resp {
	return v.value
}

func (v *NullableMarginGetMarginMaxBorrowableV1Resp) Set(val *MarginGetMarginMaxBorrowableV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginMaxBorrowableV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginMaxBorrowableV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginMaxBorrowableV1Resp(val *MarginGetMarginMaxBorrowableV1Resp) *NullableMarginGetMarginMaxBorrowableV1Resp {
	return &NullableMarginGetMarginMaxBorrowableV1Resp{value: val, isSet: true}
}

func (v NullableMarginGetMarginMaxBorrowableV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginMaxBorrowableV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


