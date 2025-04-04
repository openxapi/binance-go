/*
Binance Margin Trading API

OpenAPI specification for Binance exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginGetMarginTransferV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginTransferV1Resp{}

// MarginGetMarginTransferV1Resp struct for MarginGetMarginTransferV1Resp
type MarginGetMarginTransferV1Resp struct {
	Rows []MarginGetMarginTransferV1RespRowsInner `json:"rows,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewMarginGetMarginTransferV1Resp instantiates a new MarginGetMarginTransferV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginTransferV1Resp() *MarginGetMarginTransferV1Resp {
	this := MarginGetMarginTransferV1Resp{}
	return &this
}

// NewMarginGetMarginTransferV1RespWithDefaults instantiates a new MarginGetMarginTransferV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginTransferV1RespWithDefaults() *MarginGetMarginTransferV1Resp {
	this := MarginGetMarginTransferV1Resp{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1Resp) GetRows() []MarginGetMarginTransferV1RespRowsInner {
	if o == nil || IsNil(o.Rows) {
		var ret []MarginGetMarginTransferV1RespRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1Resp) GetRowsOk() ([]MarginGetMarginTransferV1RespRowsInner, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1Resp) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []MarginGetMarginTransferV1RespRowsInner and assigns it to the Rows field.
func (o *MarginGetMarginTransferV1Resp) SetRows(v []MarginGetMarginTransferV1RespRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MarginGetMarginTransferV1Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o MarginGetMarginTransferV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginTransferV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMarginGetMarginTransferV1Resp struct {
	value *MarginGetMarginTransferV1Resp
	isSet bool
}

func (v NullableMarginGetMarginTransferV1Resp) Get() *MarginGetMarginTransferV1Resp {
	return v.value
}

func (v *NullableMarginGetMarginTransferV1Resp) Set(val *MarginGetMarginTransferV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginTransferV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginTransferV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginTransferV1Resp(val *MarginGetMarginTransferV1Resp) *NullableMarginGetMarginTransferV1Resp {
	return &NullableMarginGetMarginTransferV1Resp{value: val, isSet: true}
}

func (v NullableMarginGetMarginTransferV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginTransferV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


