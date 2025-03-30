/*
Binance Pmargin API

OpenAPI specification for Binance cryptocurrency exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginGetMarginForceOrdersV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetMarginForceOrdersV1Resp{}

// PmarginGetMarginForceOrdersV1Resp struct for PmarginGetMarginForceOrdersV1Resp
type PmarginGetMarginForceOrdersV1Resp struct {
	Rows []PmarginGetMarginForceOrdersV1RespRowsInner `json:"rows,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewPmarginGetMarginForceOrdersV1Resp instantiates a new PmarginGetMarginForceOrdersV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetMarginForceOrdersV1Resp() *PmarginGetMarginForceOrdersV1Resp {
	this := PmarginGetMarginForceOrdersV1Resp{}
	return &this
}

// NewPmarginGetMarginForceOrdersV1RespWithDefaults instantiates a new PmarginGetMarginForceOrdersV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetMarginForceOrdersV1RespWithDefaults() *PmarginGetMarginForceOrdersV1Resp {
	this := PmarginGetMarginForceOrdersV1Resp{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *PmarginGetMarginForceOrdersV1Resp) GetRows() []PmarginGetMarginForceOrdersV1RespRowsInner {
	if o == nil || IsNil(o.Rows) {
		var ret []PmarginGetMarginForceOrdersV1RespRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginForceOrdersV1Resp) GetRowsOk() ([]PmarginGetMarginForceOrdersV1RespRowsInner, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *PmarginGetMarginForceOrdersV1Resp) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []PmarginGetMarginForceOrdersV1RespRowsInner and assigns it to the Rows field.
func (o *PmarginGetMarginForceOrdersV1Resp) SetRows(v []PmarginGetMarginForceOrdersV1RespRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PmarginGetMarginForceOrdersV1Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginForceOrdersV1Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PmarginGetMarginForceOrdersV1Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PmarginGetMarginForceOrdersV1Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o PmarginGetMarginForceOrdersV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetMarginForceOrdersV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullablePmarginGetMarginForceOrdersV1Resp struct {
	value *PmarginGetMarginForceOrdersV1Resp
	isSet bool
}

func (v NullablePmarginGetMarginForceOrdersV1Resp) Get() *PmarginGetMarginForceOrdersV1Resp {
	return v.value
}

func (v *NullablePmarginGetMarginForceOrdersV1Resp) Set(val *PmarginGetMarginForceOrdersV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetMarginForceOrdersV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetMarginForceOrdersV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetMarginForceOrdersV1Resp(val *PmarginGetMarginForceOrdersV1Resp) *NullablePmarginGetMarginForceOrdersV1Resp {
	return &NullablePmarginGetMarginForceOrdersV1Resp{value: val, isSet: true}
}

func (v NullablePmarginGetMarginForceOrdersV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetMarginForceOrdersV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


