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

// checks if the MarginCreateMarginBorrowRepayV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginCreateMarginBorrowRepayV1Resp{}

// MarginCreateMarginBorrowRepayV1Resp struct for MarginCreateMarginBorrowRepayV1Resp
type MarginCreateMarginBorrowRepayV1Resp struct {
	TranId *int64 `json:"tranId,omitempty"`
}

// NewMarginCreateMarginBorrowRepayV1Resp instantiates a new MarginCreateMarginBorrowRepayV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCreateMarginBorrowRepayV1Resp() *MarginCreateMarginBorrowRepayV1Resp {
	this := MarginCreateMarginBorrowRepayV1Resp{}
	return &this
}

// NewMarginCreateMarginBorrowRepayV1RespWithDefaults instantiates a new MarginCreateMarginBorrowRepayV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCreateMarginBorrowRepayV1RespWithDefaults() *MarginCreateMarginBorrowRepayV1Resp {
	this := MarginCreateMarginBorrowRepayV1Resp{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *MarginCreateMarginBorrowRepayV1Resp) GetTranId() int64 {
	if o == nil || IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginBorrowRepayV1Resp) GetTranIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *MarginCreateMarginBorrowRepayV1Resp) HasTranId() bool {
	if o != nil && !IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *MarginCreateMarginBorrowRepayV1Resp) SetTranId(v int64) {
	o.TranId = &v
}

func (o MarginCreateMarginBorrowRepayV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginCreateMarginBorrowRepayV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	return toSerialize, nil
}

type NullableMarginCreateMarginBorrowRepayV1Resp struct {
	value *MarginCreateMarginBorrowRepayV1Resp
	isSet bool
}

func (v NullableMarginCreateMarginBorrowRepayV1Resp) Get() *MarginCreateMarginBorrowRepayV1Resp {
	return v.value
}

func (v *NullableMarginCreateMarginBorrowRepayV1Resp) Set(val *MarginCreateMarginBorrowRepayV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCreateMarginBorrowRepayV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCreateMarginBorrowRepayV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCreateMarginBorrowRepayV1Resp(val *MarginCreateMarginBorrowRepayV1Resp) *NullableMarginCreateMarginBorrowRepayV1Resp {
	return &NullableMarginCreateMarginBorrowRepayV1Resp{value: val, isSet: true}
}

func (v NullableMarginCreateMarginBorrowRepayV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCreateMarginBorrowRepayV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


