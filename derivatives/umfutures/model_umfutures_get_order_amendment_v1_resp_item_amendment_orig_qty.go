/*
Binance Umfutures API

OpenAPI specification for Binance cryptocurrency exchange - Umfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty{}

// UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty struct for UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty
type UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
}

// NewUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty instantiates a new UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty() *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty {
	this := UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty{}
	return &this
}

// NewUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQtyWithDefaults instantiates a new UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQtyWithDefaults() *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty {
	this := UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) GetAfter() string {
	if o == nil || IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) GetAfterOk() (*string, bool) {
	if o == nil || IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) HasAfter() bool {
	if o != nil && !IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) SetAfter(v string) {
	o.After = &v
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) GetBefore() string {
	if o == nil || IsNil(o.Before) {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) GetBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) HasBefore() bool {
	if o != nil && !IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) SetBefore(v string) {
	o.Before = &v
}

func (o UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.After) {
		toSerialize["after"] = o.After
	}
	if !IsNil(o.Before) {
		toSerialize["before"] = o.Before
	}
	return toSerialize, nil
}

type NullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty struct {
	value *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty
	isSet bool
}

func (v NullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) Get() *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty {
	return v.value
}

func (v *NullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) Set(val *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty(val *UmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) *NullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty {
	return &NullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty{value: val, isSet: true}
}

func (v NullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


