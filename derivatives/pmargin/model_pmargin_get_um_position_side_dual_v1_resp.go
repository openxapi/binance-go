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

// checks if the PmarginGetUmPositionSideDualV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetUmPositionSideDualV1Resp{}

// PmarginGetUmPositionSideDualV1Resp struct for PmarginGetUmPositionSideDualV1Resp
type PmarginGetUmPositionSideDualV1Resp struct {
	DualSidePosition *bool `json:"dualSidePosition,omitempty"`
}

// NewPmarginGetUmPositionSideDualV1Resp instantiates a new PmarginGetUmPositionSideDualV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetUmPositionSideDualV1Resp() *PmarginGetUmPositionSideDualV1Resp {
	this := PmarginGetUmPositionSideDualV1Resp{}
	return &this
}

// NewPmarginGetUmPositionSideDualV1RespWithDefaults instantiates a new PmarginGetUmPositionSideDualV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetUmPositionSideDualV1RespWithDefaults() *PmarginGetUmPositionSideDualV1Resp {
	this := PmarginGetUmPositionSideDualV1Resp{}
	return &this
}

// GetDualSidePosition returns the DualSidePosition field value if set, zero value otherwise.
func (o *PmarginGetUmPositionSideDualV1Resp) GetDualSidePosition() bool {
	if o == nil || IsNil(o.DualSidePosition) {
		var ret bool
		return ret
	}
	return *o.DualSidePosition
}

// GetDualSidePositionOk returns a tuple with the DualSidePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetUmPositionSideDualV1Resp) GetDualSidePositionOk() (*bool, bool) {
	if o == nil || IsNil(o.DualSidePosition) {
		return nil, false
	}
	return o.DualSidePosition, true
}

// HasDualSidePosition returns a boolean if a field has been set.
func (o *PmarginGetUmPositionSideDualV1Resp) HasDualSidePosition() bool {
	if o != nil && !IsNil(o.DualSidePosition) {
		return true
	}

	return false
}

// SetDualSidePosition gets a reference to the given bool and assigns it to the DualSidePosition field.
func (o *PmarginGetUmPositionSideDualV1Resp) SetDualSidePosition(v bool) {
	o.DualSidePosition = &v
}

func (o PmarginGetUmPositionSideDualV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetUmPositionSideDualV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DualSidePosition) {
		toSerialize["dualSidePosition"] = o.DualSidePosition
	}
	return toSerialize, nil
}

type NullablePmarginGetUmPositionSideDualV1Resp struct {
	value *PmarginGetUmPositionSideDualV1Resp
	isSet bool
}

func (v NullablePmarginGetUmPositionSideDualV1Resp) Get() *PmarginGetUmPositionSideDualV1Resp {
	return v.value
}

func (v *NullablePmarginGetUmPositionSideDualV1Resp) Set(val *PmarginGetUmPositionSideDualV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetUmPositionSideDualV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetUmPositionSideDualV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetUmPositionSideDualV1Resp(val *PmarginGetUmPositionSideDualV1Resp) *NullablePmarginGetUmPositionSideDualV1Resp {
	return &NullablePmarginGetUmPositionSideDualV1Resp{value: val, isSet: true}
}

func (v NullablePmarginGetUmPositionSideDualV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetUmPositionSideDualV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


