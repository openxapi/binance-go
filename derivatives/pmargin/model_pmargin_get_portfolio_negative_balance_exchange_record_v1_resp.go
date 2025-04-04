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

// checks if the PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp{}

// PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp struct for PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp
type PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp struct {
	Rows []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner `json:"rows,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewPmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp instantiates a new PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp() *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp {
	this := PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp{}
	return &this
}

// NewPmarginGetPortfolioNegativeBalanceExchangeRecordV1RespWithDefaults instantiates a new PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetPortfolioNegativeBalanceExchangeRecordV1RespWithDefaults() *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp {
	this := PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) GetRows() []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner {
	if o == nil || IsNil(o.Rows) {
		var ret []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) GetRowsOk() ([]PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner and assigns it to the Rows field.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) SetRows(v []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp struct {
	value *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp
	isSet bool
}

func (v NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) Get() *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp {
	return v.value
}

func (v *NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) Set(val *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp(val *PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) *NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp {
	return &NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp{value: val, isSet: true}
}

func (v NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


