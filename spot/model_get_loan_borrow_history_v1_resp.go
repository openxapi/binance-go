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

// checks if the GetLoanBorrowHistoryV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLoanBorrowHistoryV1Resp{}

// GetLoanBorrowHistoryV1Resp struct for GetLoanBorrowHistoryV1Resp
type GetLoanBorrowHistoryV1Resp struct {
	Rows []GetLoanBorrowHistoryV1RespRowsInner `json:"rows,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewGetLoanBorrowHistoryV1Resp instantiates a new GetLoanBorrowHistoryV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanBorrowHistoryV1Resp() *GetLoanBorrowHistoryV1Resp {
	this := GetLoanBorrowHistoryV1Resp{}
	return &this
}

// NewGetLoanBorrowHistoryV1RespWithDefaults instantiates a new GetLoanBorrowHistoryV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanBorrowHistoryV1RespWithDefaults() *GetLoanBorrowHistoryV1Resp {
	this := GetLoanBorrowHistoryV1Resp{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryV1Resp) GetRows() []GetLoanBorrowHistoryV1RespRowsInner {
	if o == nil || IsNil(o.Rows) {
		var ret []GetLoanBorrowHistoryV1RespRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryV1Resp) GetRowsOk() ([]GetLoanBorrowHistoryV1RespRowsInner, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryV1Resp) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLoanBorrowHistoryV1RespRowsInner and assigns it to the Rows field.
func (o *GetLoanBorrowHistoryV1Resp) SetRows(v []GetLoanBorrowHistoryV1RespRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryV1Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryV1Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryV1Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetLoanBorrowHistoryV1Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o GetLoanBorrowHistoryV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanBorrowHistoryV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableGetLoanBorrowHistoryV1Resp struct {
	value *GetLoanBorrowHistoryV1Resp
	isSet bool
}

func (v NullableGetLoanBorrowHistoryV1Resp) Get() *GetLoanBorrowHistoryV1Resp {
	return v.value
}

func (v *NullableGetLoanBorrowHistoryV1Resp) Set(val *GetLoanBorrowHistoryV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanBorrowHistoryV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanBorrowHistoryV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanBorrowHistoryV1Resp(val *GetLoanBorrowHistoryV1Resp) *NullableGetLoanBorrowHistoryV1Resp {
	return &NullableGetLoanBorrowHistoryV1Resp{value: val, isSet: true}
}

func (v NullableGetLoanBorrowHistoryV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanBorrowHistoryV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


