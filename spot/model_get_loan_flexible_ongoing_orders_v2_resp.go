/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the GetLoanFlexibleOngoingOrdersV2Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLoanFlexibleOngoingOrdersV2Resp{}

// GetLoanFlexibleOngoingOrdersV2Resp struct for GetLoanFlexibleOngoingOrdersV2Resp
type GetLoanFlexibleOngoingOrdersV2Resp struct {
	Rows []GetLoanFlexibleOngoingOrdersV2RespRowsInner `json:"rows,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewGetLoanFlexibleOngoingOrdersV2Resp instantiates a new GetLoanFlexibleOngoingOrdersV2Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanFlexibleOngoingOrdersV2Resp() *GetLoanFlexibleOngoingOrdersV2Resp {
	this := GetLoanFlexibleOngoingOrdersV2Resp{}
	return &this
}

// NewGetLoanFlexibleOngoingOrdersV2RespWithDefaults instantiates a new GetLoanFlexibleOngoingOrdersV2Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanFlexibleOngoingOrdersV2RespWithDefaults() *GetLoanFlexibleOngoingOrdersV2Resp {
	this := GetLoanFlexibleOngoingOrdersV2Resp{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLoanFlexibleOngoingOrdersV2Resp) GetRows() []GetLoanFlexibleOngoingOrdersV2RespRowsInner {
	if o == nil || IsNil(o.Rows) {
		var ret []GetLoanFlexibleOngoingOrdersV2RespRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanFlexibleOngoingOrdersV2Resp) GetRowsOk() ([]GetLoanFlexibleOngoingOrdersV2RespRowsInner, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLoanFlexibleOngoingOrdersV2Resp) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLoanFlexibleOngoingOrdersV2RespRowsInner and assigns it to the Rows field.
func (o *GetLoanFlexibleOngoingOrdersV2Resp) SetRows(v []GetLoanFlexibleOngoingOrdersV2RespRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLoanFlexibleOngoingOrdersV2Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanFlexibleOngoingOrdersV2Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLoanFlexibleOngoingOrdersV2Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetLoanFlexibleOngoingOrdersV2Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o GetLoanFlexibleOngoingOrdersV2Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanFlexibleOngoingOrdersV2Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableGetLoanFlexibleOngoingOrdersV2Resp struct {
	value *GetLoanFlexibleOngoingOrdersV2Resp
	isSet bool
}

func (v NullableGetLoanFlexibleOngoingOrdersV2Resp) Get() *GetLoanFlexibleOngoingOrdersV2Resp {
	return v.value
}

func (v *NullableGetLoanFlexibleOngoingOrdersV2Resp) Set(val *GetLoanFlexibleOngoingOrdersV2Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanFlexibleOngoingOrdersV2Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanFlexibleOngoingOrdersV2Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanFlexibleOngoingOrdersV2Resp(val *GetLoanFlexibleOngoingOrdersV2Resp) *NullableGetLoanFlexibleOngoingOrdersV2Resp {
	return &NullableGetLoanFlexibleOngoingOrdersV2Resp{value: val, isSet: true}
}

func (v NullableGetLoanFlexibleOngoingOrdersV2Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanFlexibleOngoingOrdersV2Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


