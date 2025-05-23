/*
Binance Options API

OpenAPI specification for Binance exchange - Options API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
)

// checks if the OptionsGetExchangeInfoV1RespRateLimitsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsGetExchangeInfoV1RespRateLimitsInner{}

// OptionsGetExchangeInfoV1RespRateLimitsInner struct for OptionsGetExchangeInfoV1RespRateLimitsInner
type OptionsGetExchangeInfoV1RespRateLimitsInner struct {
	Interval *string `json:"interval,omitempty"`
	IntervalNum *int32 `json:"intervalNum,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	RateLimitType *string `json:"rateLimitType,omitempty"`
}

// NewOptionsGetExchangeInfoV1RespRateLimitsInner instantiates a new OptionsGetExchangeInfoV1RespRateLimitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsGetExchangeInfoV1RespRateLimitsInner() *OptionsGetExchangeInfoV1RespRateLimitsInner {
	this := OptionsGetExchangeInfoV1RespRateLimitsInner{}
	return &this
}

// NewOptionsGetExchangeInfoV1RespRateLimitsInnerWithDefaults instantiates a new OptionsGetExchangeInfoV1RespRateLimitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsGetExchangeInfoV1RespRateLimitsInnerWithDefaults() *OptionsGetExchangeInfoV1RespRateLimitsInner {
	this := OptionsGetExchangeInfoV1RespRateLimitsInner{}
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) GetIntervalNum() int32 {
	if o == nil || IsNil(o.IntervalNum) {
		var ret int32
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) GetIntervalNumOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) HasIntervalNum() bool {
	if o != nil && !IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int32 and assigns it to the IntervalNum field.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) SetIntervalNum(v int32) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) SetLimit(v int32) {
	o.Limit = &v
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) GetRateLimitType() string {
	if o == nil || IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) HasRateLimitType() bool {
	if o != nil && !IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *OptionsGetExchangeInfoV1RespRateLimitsInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

func (o OptionsGetExchangeInfoV1RespRateLimitsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionsGetExchangeInfoV1RespRateLimitsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.IntervalNum) {
		toSerialize["intervalNum"] = o.IntervalNum
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.RateLimitType) {
		toSerialize["rateLimitType"] = o.RateLimitType
	}
	return toSerialize, nil
}

type NullableOptionsGetExchangeInfoV1RespRateLimitsInner struct {
	value *OptionsGetExchangeInfoV1RespRateLimitsInner
	isSet bool
}

func (v NullableOptionsGetExchangeInfoV1RespRateLimitsInner) Get() *OptionsGetExchangeInfoV1RespRateLimitsInner {
	return v.value
}

func (v *NullableOptionsGetExchangeInfoV1RespRateLimitsInner) Set(val *OptionsGetExchangeInfoV1RespRateLimitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsGetExchangeInfoV1RespRateLimitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsGetExchangeInfoV1RespRateLimitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsGetExchangeInfoV1RespRateLimitsInner(val *OptionsGetExchangeInfoV1RespRateLimitsInner) *NullableOptionsGetExchangeInfoV1RespRateLimitsInner {
	return &NullableOptionsGetExchangeInfoV1RespRateLimitsInner{value: val, isSet: true}
}

func (v NullableOptionsGetExchangeInfoV1RespRateLimitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsGetExchangeInfoV1RespRateLimitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


