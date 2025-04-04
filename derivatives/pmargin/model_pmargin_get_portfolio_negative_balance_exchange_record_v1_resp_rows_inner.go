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

// checks if the PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner{}

// PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner struct for PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner
type PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner struct {
	Details []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInnerDetailsInner `json:"details,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
}

// NewPmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner instantiates a new PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner() *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner {
	this := PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner{}
	return &this
}

// NewPmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInnerWithDefaults instantiates a new PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInnerWithDefaults() *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner {
	this := PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) GetDetails() []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInnerDetailsInner {
	if o == nil || IsNil(o.Details) {
		var ret []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInnerDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) GetDetailsOk() ([]PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInnerDetailsInner, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInnerDetailsInner and assigns it to the Details field.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) SetDetails(v []PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInnerDetailsInner) {
	o.Details = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) GetEndTime() int64 {
	if o == nil || IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) GetEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) SetStartTime(v int64) {
	o.StartTime = &v
}

func (o PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	return toSerialize, nil
}

type NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner struct {
	value *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner
	isSet bool
}

func (v NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) Get() *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner {
	return v.value
}

func (v *NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) Set(val *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner(val *PmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) *NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner {
	return &NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner{value: val, isSet: true}
}

func (v NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetPortfolioNegativeBalanceExchangeRecordV1RespRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


