/*
Binance Pmarginpro API

OpenAPI specification for Binance cryptocurrency exchange - Pmarginpro API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmarginpro

import (
	"encoding/json"
)

// checks if the PmarginproCreatePortfolioBnbTransferV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginproCreatePortfolioBnbTransferV1Resp{}

// PmarginproCreatePortfolioBnbTransferV1Resp struct for PmarginproCreatePortfolioBnbTransferV1Resp
type PmarginproCreatePortfolioBnbTransferV1Resp struct {
	TranId *int64 `json:"tranId,omitempty"`
}

// NewPmarginproCreatePortfolioBnbTransferV1Resp instantiates a new PmarginproCreatePortfolioBnbTransferV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginproCreatePortfolioBnbTransferV1Resp() *PmarginproCreatePortfolioBnbTransferV1Resp {
	this := PmarginproCreatePortfolioBnbTransferV1Resp{}
	return &this
}

// NewPmarginproCreatePortfolioBnbTransferV1RespWithDefaults instantiates a new PmarginproCreatePortfolioBnbTransferV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginproCreatePortfolioBnbTransferV1RespWithDefaults() *PmarginproCreatePortfolioBnbTransferV1Resp {
	this := PmarginproCreatePortfolioBnbTransferV1Resp{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *PmarginproCreatePortfolioBnbTransferV1Resp) GetTranId() int64 {
	if o == nil || IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproCreatePortfolioBnbTransferV1Resp) GetTranIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *PmarginproCreatePortfolioBnbTransferV1Resp) HasTranId() bool {
	if o != nil && !IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *PmarginproCreatePortfolioBnbTransferV1Resp) SetTranId(v int64) {
	o.TranId = &v
}

func (o PmarginproCreatePortfolioBnbTransferV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginproCreatePortfolioBnbTransferV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	return toSerialize, nil
}

type NullablePmarginproCreatePortfolioBnbTransferV1Resp struct {
	value *PmarginproCreatePortfolioBnbTransferV1Resp
	isSet bool
}

func (v NullablePmarginproCreatePortfolioBnbTransferV1Resp) Get() *PmarginproCreatePortfolioBnbTransferV1Resp {
	return v.value
}

func (v *NullablePmarginproCreatePortfolioBnbTransferV1Resp) Set(val *PmarginproCreatePortfolioBnbTransferV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginproCreatePortfolioBnbTransferV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginproCreatePortfolioBnbTransferV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginproCreatePortfolioBnbTransferV1Resp(val *PmarginproCreatePortfolioBnbTransferV1Resp) *NullablePmarginproCreatePortfolioBnbTransferV1Resp {
	return &NullablePmarginproCreatePortfolioBnbTransferV1Resp{value: val, isSet: true}
}

func (v NullablePmarginproCreatePortfolioBnbTransferV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginproCreatePortfolioBnbTransferV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


