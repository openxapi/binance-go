/*
Binance Ufutures API

OpenAPI specification for Binance cryptocurrency exchange - Ufutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ufutures

import (
	"encoding/json"
)

// checks if the UfuturesGetOrderAmendmentV1RespItemAmendment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetOrderAmendmentV1RespItemAmendment{}

// UfuturesGetOrderAmendmentV1RespItemAmendment struct for UfuturesGetOrderAmendmentV1RespItemAmendment
type UfuturesGetOrderAmendmentV1RespItemAmendment struct {
	Count *int32 `json:"count,omitempty"`
	OrigQty *UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty `json:"origQty,omitempty"`
	Price *UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty `json:"price,omitempty"`
}

// NewUfuturesGetOrderAmendmentV1RespItemAmendment instantiates a new UfuturesGetOrderAmendmentV1RespItemAmendment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetOrderAmendmentV1RespItemAmendment() *UfuturesGetOrderAmendmentV1RespItemAmendment {
	this := UfuturesGetOrderAmendmentV1RespItemAmendment{}
	return &this
}

// NewUfuturesGetOrderAmendmentV1RespItemAmendmentWithDefaults instantiates a new UfuturesGetOrderAmendmentV1RespItemAmendment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetOrderAmendmentV1RespItemAmendmentWithDefaults() *UfuturesGetOrderAmendmentV1RespItemAmendment {
	this := UfuturesGetOrderAmendmentV1RespItemAmendment{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) SetCount(v int32) {
	o.Count = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) GetOrigQty() UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty {
	if o == nil || IsNil(o.OrigQty) {
		var ret UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) GetOrigQtyOk() (*UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty and assigns it to the OrigQty field.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) SetOrigQty(v UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) {
	o.OrigQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) GetPrice() UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty {
	if o == nil || IsNil(o.Price) {
		var ret UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) GetPriceOk() (*UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty and assigns it to the Price field.
func (o *UfuturesGetOrderAmendmentV1RespItemAmendment) SetPrice(v UfuturesGetOrderAmendmentV1RespItemAmendmentOrigQty) {
	o.Price = &v
}

func (o UfuturesGetOrderAmendmentV1RespItemAmendment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetOrderAmendmentV1RespItemAmendment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableUfuturesGetOrderAmendmentV1RespItemAmendment struct {
	value *UfuturesGetOrderAmendmentV1RespItemAmendment
	isSet bool
}

func (v NullableUfuturesGetOrderAmendmentV1RespItemAmendment) Get() *UfuturesGetOrderAmendmentV1RespItemAmendment {
	return v.value
}

func (v *NullableUfuturesGetOrderAmendmentV1RespItemAmendment) Set(val *UfuturesGetOrderAmendmentV1RespItemAmendment) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetOrderAmendmentV1RespItemAmendment) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetOrderAmendmentV1RespItemAmendment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetOrderAmendmentV1RespItemAmendment(val *UfuturesGetOrderAmendmentV1RespItemAmendment) *NullableUfuturesGetOrderAmendmentV1RespItemAmendment {
	return &NullableUfuturesGetOrderAmendmentV1RespItemAmendment{value: val, isSet: true}
}

func (v NullableUfuturesGetOrderAmendmentV1RespItemAmendment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetOrderAmendmentV1RespItemAmendment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


