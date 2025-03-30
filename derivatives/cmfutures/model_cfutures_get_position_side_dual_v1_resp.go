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

// checks if the CfuturesGetPositionSideDualV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CfuturesGetPositionSideDualV1Resp{}

// CfuturesGetPositionSideDualV1Resp struct for CfuturesGetPositionSideDualV1Resp
type CfuturesGetPositionSideDualV1Resp struct {
	DualSidePosition *bool `json:"dualSidePosition,omitempty"`
}

// NewCfuturesGetPositionSideDualV1Resp instantiates a new CfuturesGetPositionSideDualV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCfuturesGetPositionSideDualV1Resp() *CfuturesGetPositionSideDualV1Resp {
	this := CfuturesGetPositionSideDualV1Resp{}
	return &this
}

// NewCfuturesGetPositionSideDualV1RespWithDefaults instantiates a new CfuturesGetPositionSideDualV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCfuturesGetPositionSideDualV1RespWithDefaults() *CfuturesGetPositionSideDualV1Resp {
	this := CfuturesGetPositionSideDualV1Resp{}
	return &this
}

// GetDualSidePosition returns the DualSidePosition field value if set, zero value otherwise.
func (o *CfuturesGetPositionSideDualV1Resp) GetDualSidePosition() bool {
	if o == nil || IsNil(o.DualSidePosition) {
		var ret bool
		return ret
	}
	return *o.DualSidePosition
}

// GetDualSidePositionOk returns a tuple with the DualSidePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetPositionSideDualV1Resp) GetDualSidePositionOk() (*bool, bool) {
	if o == nil || IsNil(o.DualSidePosition) {
		return nil, false
	}
	return o.DualSidePosition, true
}

// HasDualSidePosition returns a boolean if a field has been set.
func (o *CfuturesGetPositionSideDualV1Resp) HasDualSidePosition() bool {
	if o != nil && !IsNil(o.DualSidePosition) {
		return true
	}

	return false
}

// SetDualSidePosition gets a reference to the given bool and assigns it to the DualSidePosition field.
func (o *CfuturesGetPositionSideDualV1Resp) SetDualSidePosition(v bool) {
	o.DualSidePosition = &v
}

func (o CfuturesGetPositionSideDualV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CfuturesGetPositionSideDualV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DualSidePosition) {
		toSerialize["dualSidePosition"] = o.DualSidePosition
	}
	return toSerialize, nil
}

type NullableCfuturesGetPositionSideDualV1Resp struct {
	value *CfuturesGetPositionSideDualV1Resp
	isSet bool
}

func (v NullableCfuturesGetPositionSideDualV1Resp) Get() *CfuturesGetPositionSideDualV1Resp {
	return v.value
}

func (v *NullableCfuturesGetPositionSideDualV1Resp) Set(val *CfuturesGetPositionSideDualV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCfuturesGetPositionSideDualV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCfuturesGetPositionSideDualV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCfuturesGetPositionSideDualV1Resp(val *CfuturesGetPositionSideDualV1Resp) *NullableCfuturesGetPositionSideDualV1Resp {
	return &NullableCfuturesGetPositionSideDualV1Resp{value: val, isSet: true}
}

func (v NullableCfuturesGetPositionSideDualV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCfuturesGetPositionSideDualV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


