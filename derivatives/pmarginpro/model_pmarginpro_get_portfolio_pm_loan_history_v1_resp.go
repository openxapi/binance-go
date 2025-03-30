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

// checks if the PmarginproGetPortfolioPmLoanHistoryV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginproGetPortfolioPmLoanHistoryV1Resp{}

// PmarginproGetPortfolioPmLoanHistoryV1Resp struct for PmarginproGetPortfolioPmLoanHistoryV1Resp
type PmarginproGetPortfolioPmLoanHistoryV1Resp struct {
	Rows []PmarginproGetPortfolioPmLoanHistoryV1RespRowsInner `json:"rows,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewPmarginproGetPortfolioPmLoanHistoryV1Resp instantiates a new PmarginproGetPortfolioPmLoanHistoryV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginproGetPortfolioPmLoanHistoryV1Resp() *PmarginproGetPortfolioPmLoanHistoryV1Resp {
	this := PmarginproGetPortfolioPmLoanHistoryV1Resp{}
	return &this
}

// NewPmarginproGetPortfolioPmLoanHistoryV1RespWithDefaults instantiates a new PmarginproGetPortfolioPmLoanHistoryV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginproGetPortfolioPmLoanHistoryV1RespWithDefaults() *PmarginproGetPortfolioPmLoanHistoryV1Resp {
	this := PmarginproGetPortfolioPmLoanHistoryV1Resp{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *PmarginproGetPortfolioPmLoanHistoryV1Resp) GetRows() []PmarginproGetPortfolioPmLoanHistoryV1RespRowsInner {
	if o == nil || IsNil(o.Rows) {
		var ret []PmarginproGetPortfolioPmLoanHistoryV1RespRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproGetPortfolioPmLoanHistoryV1Resp) GetRowsOk() ([]PmarginproGetPortfolioPmLoanHistoryV1RespRowsInner, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *PmarginproGetPortfolioPmLoanHistoryV1Resp) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []PmarginproGetPortfolioPmLoanHistoryV1RespRowsInner and assigns it to the Rows field.
func (o *PmarginproGetPortfolioPmLoanHistoryV1Resp) SetRows(v []PmarginproGetPortfolioPmLoanHistoryV1RespRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PmarginproGetPortfolioPmLoanHistoryV1Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproGetPortfolioPmLoanHistoryV1Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PmarginproGetPortfolioPmLoanHistoryV1Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PmarginproGetPortfolioPmLoanHistoryV1Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o PmarginproGetPortfolioPmLoanHistoryV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginproGetPortfolioPmLoanHistoryV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullablePmarginproGetPortfolioPmLoanHistoryV1Resp struct {
	value *PmarginproGetPortfolioPmLoanHistoryV1Resp
	isSet bool
}

func (v NullablePmarginproGetPortfolioPmLoanHistoryV1Resp) Get() *PmarginproGetPortfolioPmLoanHistoryV1Resp {
	return v.value
}

func (v *NullablePmarginproGetPortfolioPmLoanHistoryV1Resp) Set(val *PmarginproGetPortfolioPmLoanHistoryV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginproGetPortfolioPmLoanHistoryV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginproGetPortfolioPmLoanHistoryV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginproGetPortfolioPmLoanHistoryV1Resp(val *PmarginproGetPortfolioPmLoanHistoryV1Resp) *NullablePmarginproGetPortfolioPmLoanHistoryV1Resp {
	return &NullablePmarginproGetPortfolioPmLoanHistoryV1Resp{value: val, isSet: true}
}

func (v NullablePmarginproGetPortfolioPmLoanHistoryV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginproGetPortfolioPmLoanHistoryV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


