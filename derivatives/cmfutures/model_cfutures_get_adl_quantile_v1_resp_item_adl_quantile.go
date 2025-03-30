/*
Binance Cfutures API

OpenAPI specification for Binance cryptocurrency exchange - Cfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
)

// checks if the CfuturesGetAdlQuantileV1RespItemAdlQuantile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CfuturesGetAdlQuantileV1RespItemAdlQuantile{}

// CfuturesGetAdlQuantileV1RespItemAdlQuantile struct for CfuturesGetAdlQuantileV1RespItemAdlQuantile
type CfuturesGetAdlQuantileV1RespItemAdlQuantile struct {
	HEDGE *int32 `json:"HEDGE,omitempty"`
	LONG *int32 `json:"LONG,omitempty"`
	SHORT *int32 `json:"SHORT,omitempty"`
}

// NewCfuturesGetAdlQuantileV1RespItemAdlQuantile instantiates a new CfuturesGetAdlQuantileV1RespItemAdlQuantile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCfuturesGetAdlQuantileV1RespItemAdlQuantile() *CfuturesGetAdlQuantileV1RespItemAdlQuantile {
	this := CfuturesGetAdlQuantileV1RespItemAdlQuantile{}
	return &this
}

// NewCfuturesGetAdlQuantileV1RespItemAdlQuantileWithDefaults instantiates a new CfuturesGetAdlQuantileV1RespItemAdlQuantile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCfuturesGetAdlQuantileV1RespItemAdlQuantileWithDefaults() *CfuturesGetAdlQuantileV1RespItemAdlQuantile {
	this := CfuturesGetAdlQuantileV1RespItemAdlQuantile{}
	return &this
}

// GetHEDGE returns the HEDGE field value if set, zero value otherwise.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) GetHEDGE() int32 {
	if o == nil || IsNil(o.HEDGE) {
		var ret int32
		return ret
	}
	return *o.HEDGE
}

// GetHEDGEOk returns a tuple with the HEDGE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) GetHEDGEOk() (*int32, bool) {
	if o == nil || IsNil(o.HEDGE) {
		return nil, false
	}
	return o.HEDGE, true
}

// HasHEDGE returns a boolean if a field has been set.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) HasHEDGE() bool {
	if o != nil && !IsNil(o.HEDGE) {
		return true
	}

	return false
}

// SetHEDGE gets a reference to the given int32 and assigns it to the HEDGE field.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) SetHEDGE(v int32) {
	o.HEDGE = &v
}

// GetLONG returns the LONG field value if set, zero value otherwise.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) GetLONG() int32 {
	if o == nil || IsNil(o.LONG) {
		var ret int32
		return ret
	}
	return *o.LONG
}

// GetLONGOk returns a tuple with the LONG field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) GetLONGOk() (*int32, bool) {
	if o == nil || IsNil(o.LONG) {
		return nil, false
	}
	return o.LONG, true
}

// HasLONG returns a boolean if a field has been set.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) HasLONG() bool {
	if o != nil && !IsNil(o.LONG) {
		return true
	}

	return false
}

// SetLONG gets a reference to the given int32 and assigns it to the LONG field.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) SetLONG(v int32) {
	o.LONG = &v
}

// GetSHORT returns the SHORT field value if set, zero value otherwise.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) GetSHORT() int32 {
	if o == nil || IsNil(o.SHORT) {
		var ret int32
		return ret
	}
	return *o.SHORT
}

// GetSHORTOk returns a tuple with the SHORT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) GetSHORTOk() (*int32, bool) {
	if o == nil || IsNil(o.SHORT) {
		return nil, false
	}
	return o.SHORT, true
}

// HasSHORT returns a boolean if a field has been set.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) HasSHORT() bool {
	if o != nil && !IsNil(o.SHORT) {
		return true
	}

	return false
}

// SetSHORT gets a reference to the given int32 and assigns it to the SHORT field.
func (o *CfuturesGetAdlQuantileV1RespItemAdlQuantile) SetSHORT(v int32) {
	o.SHORT = &v
}

func (o CfuturesGetAdlQuantileV1RespItemAdlQuantile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CfuturesGetAdlQuantileV1RespItemAdlQuantile) ToMap() (map[string]interface{}, error) {
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

type NullableCfuturesGetAdlQuantileV1RespItemAdlQuantile struct {
	value *CfuturesGetAdlQuantileV1RespItemAdlQuantile
	isSet bool
}

func (v NullableCfuturesGetAdlQuantileV1RespItemAdlQuantile) Get() *CfuturesGetAdlQuantileV1RespItemAdlQuantile {
	return v.value
}

func (v *NullableCfuturesGetAdlQuantileV1RespItemAdlQuantile) Set(val *CfuturesGetAdlQuantileV1RespItemAdlQuantile) {
	v.value = val
	v.isSet = true
}

func (v NullableCfuturesGetAdlQuantileV1RespItemAdlQuantile) IsSet() bool {
	return v.isSet
}

func (v *NullableCfuturesGetAdlQuantileV1RespItemAdlQuantile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCfuturesGetAdlQuantileV1RespItemAdlQuantile(val *CfuturesGetAdlQuantileV1RespItemAdlQuantile) *NullableCfuturesGetAdlQuantileV1RespItemAdlQuantile {
	return &NullableCfuturesGetAdlQuantileV1RespItemAdlQuantile{value: val, isSet: true}
}

func (v NullableCfuturesGetAdlQuantileV1RespItemAdlQuantile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCfuturesGetAdlQuantileV1RespItemAdlQuantile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


