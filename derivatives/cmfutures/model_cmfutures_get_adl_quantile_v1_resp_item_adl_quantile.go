/*
Binance COIN-M Futures API

OpenAPI specification for Binance exchange - Cmfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
)

// checks if the CmfuturesGetAdlQuantileV1RespItemAdlQuantile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CmfuturesGetAdlQuantileV1RespItemAdlQuantile{}

// CmfuturesGetAdlQuantileV1RespItemAdlQuantile struct for CmfuturesGetAdlQuantileV1RespItemAdlQuantile
type CmfuturesGetAdlQuantileV1RespItemAdlQuantile struct {
	HEDGE *int32 `json:"HEDGE,omitempty"`
	LONG *int32 `json:"LONG,omitempty"`
	SHORT *int32 `json:"SHORT,omitempty"`
}

// NewCmfuturesGetAdlQuantileV1RespItemAdlQuantile instantiates a new CmfuturesGetAdlQuantileV1RespItemAdlQuantile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmfuturesGetAdlQuantileV1RespItemAdlQuantile() *CmfuturesGetAdlQuantileV1RespItemAdlQuantile {
	this := CmfuturesGetAdlQuantileV1RespItemAdlQuantile{}
	return &this
}

// NewCmfuturesGetAdlQuantileV1RespItemAdlQuantileWithDefaults instantiates a new CmfuturesGetAdlQuantileV1RespItemAdlQuantile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmfuturesGetAdlQuantileV1RespItemAdlQuantileWithDefaults() *CmfuturesGetAdlQuantileV1RespItemAdlQuantile {
	this := CmfuturesGetAdlQuantileV1RespItemAdlQuantile{}
	return &this
}

// GetHEDGE returns the HEDGE field value if set, zero value otherwise.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) GetHEDGE() int32 {
	if o == nil || IsNil(o.HEDGE) {
		var ret int32
		return ret
	}
	return *o.HEDGE
}

// GetHEDGEOk returns a tuple with the HEDGE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) GetHEDGEOk() (*int32, bool) {
	if o == nil || IsNil(o.HEDGE) {
		return nil, false
	}
	return o.HEDGE, true
}

// HasHEDGE returns a boolean if a field has been set.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) HasHEDGE() bool {
	if o != nil && !IsNil(o.HEDGE) {
		return true
	}

	return false
}

// SetHEDGE gets a reference to the given int32 and assigns it to the HEDGE field.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) SetHEDGE(v int32) {
	o.HEDGE = &v
}

// GetLONG returns the LONG field value if set, zero value otherwise.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) GetLONG() int32 {
	if o == nil || IsNil(o.LONG) {
		var ret int32
		return ret
	}
	return *o.LONG
}

// GetLONGOk returns a tuple with the LONG field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) GetLONGOk() (*int32, bool) {
	if o == nil || IsNil(o.LONG) {
		return nil, false
	}
	return o.LONG, true
}

// HasLONG returns a boolean if a field has been set.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) HasLONG() bool {
	if o != nil && !IsNil(o.LONG) {
		return true
	}

	return false
}

// SetLONG gets a reference to the given int32 and assigns it to the LONG field.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) SetLONG(v int32) {
	o.LONG = &v
}

// GetSHORT returns the SHORT field value if set, zero value otherwise.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) GetSHORT() int32 {
	if o == nil || IsNil(o.SHORT) {
		var ret int32
		return ret
	}
	return *o.SHORT
}

// GetSHORTOk returns a tuple with the SHORT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) GetSHORTOk() (*int32, bool) {
	if o == nil || IsNil(o.SHORT) {
		return nil, false
	}
	return o.SHORT, true
}

// HasSHORT returns a boolean if a field has been set.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) HasSHORT() bool {
	if o != nil && !IsNil(o.SHORT) {
		return true
	}

	return false
}

// SetSHORT gets a reference to the given int32 and assigns it to the SHORT field.
func (o *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) SetSHORT(v int32) {
	o.SHORT = &v
}

func (o CmfuturesGetAdlQuantileV1RespItemAdlQuantile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmfuturesGetAdlQuantileV1RespItemAdlQuantile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HEDGE) {
		toSerialize["HEDGE"] = o.HEDGE
	}
	if !IsNil(o.LONG) {
		toSerialize["LONG"] = o.LONG
	}
	if !IsNil(o.SHORT) {
		toSerialize["SHORT"] = o.SHORT
	}
	return toSerialize, nil
}

type NullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile struct {
	value *CmfuturesGetAdlQuantileV1RespItemAdlQuantile
	isSet bool
}

func (v NullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile) Get() *CmfuturesGetAdlQuantileV1RespItemAdlQuantile {
	return v.value
}

func (v *NullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile) Set(val *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) {
	v.value = val
	v.isSet = true
}

func (v NullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile) IsSet() bool {
	return v.isSet
}

func (v *NullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile(val *CmfuturesGetAdlQuantileV1RespItemAdlQuantile) *NullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile {
	return &NullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile{value: val, isSet: true}
}

func (v NullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmfuturesGetAdlQuantileV1RespItemAdlQuantile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


