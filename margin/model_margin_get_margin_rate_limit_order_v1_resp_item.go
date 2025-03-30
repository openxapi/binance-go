/*
Binance Margin API

OpenAPI specification for Binance cryptocurrency exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginGetMarginRateLimitOrderV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginRateLimitOrderV1RespItem{}

// MarginGetMarginRateLimitOrderV1RespItem struct for MarginGetMarginRateLimitOrderV1RespItem
type MarginGetMarginRateLimitOrderV1RespItem struct {
	Count *int32 `json:"count,omitempty"`
	Interval *string `json:"interval,omitempty"`
	IntervalNum *int32 `json:"intervalNum,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	RateLimitType *string `json:"rateLimitType,omitempty"`
}

// NewMarginGetMarginRateLimitOrderV1RespItem instantiates a new MarginGetMarginRateLimitOrderV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginRateLimitOrderV1RespItem() *MarginGetMarginRateLimitOrderV1RespItem {
	this := MarginGetMarginRateLimitOrderV1RespItem{}
	return &this
}

// NewMarginGetMarginRateLimitOrderV1RespItemWithDefaults instantiates a new MarginGetMarginRateLimitOrderV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginRateLimitOrderV1RespItemWithDefaults() *MarginGetMarginRateLimitOrderV1RespItem {
	this := MarginGetMarginRateLimitOrderV1RespItem{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *MarginGetMarginRateLimitOrderV1RespItem) SetCount(v int32) {
	o.Count = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *MarginGetMarginRateLimitOrderV1RespItem) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetIntervalNum() int32 {
	if o == nil || IsNil(o.IntervalNum) {
		var ret int32
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetIntervalNumOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) HasIntervalNum() bool {
	if o != nil && !IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int32 and assigns it to the IntervalNum field.
func (o *MarginGetMarginRateLimitOrderV1RespItem) SetIntervalNum(v int32) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *MarginGetMarginRateLimitOrderV1RespItem) SetLimit(v int32) {
	o.Limit = &v
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetRateLimitType() string {
	if o == nil || IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *MarginGetMarginRateLimitOrderV1RespItem) HasRateLimitType() bool {
	if o != nil && !IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *MarginGetMarginRateLimitOrderV1RespItem) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

func (o MarginGetMarginRateLimitOrderV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginRateLimitOrderV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
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

type NullableMarginGetMarginRateLimitOrderV1RespItem struct {
	value *MarginGetMarginRateLimitOrderV1RespItem
	isSet bool
}

func (v NullableMarginGetMarginRateLimitOrderV1RespItem) Get() *MarginGetMarginRateLimitOrderV1RespItem {
	return v.value
}

func (v *NullableMarginGetMarginRateLimitOrderV1RespItem) Set(val *MarginGetMarginRateLimitOrderV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginRateLimitOrderV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginRateLimitOrderV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginRateLimitOrderV1RespItem(val *MarginGetMarginRateLimitOrderV1RespItem) *NullableMarginGetMarginRateLimitOrderV1RespItem {
	return &NullableMarginGetMarginRateLimitOrderV1RespItem{value: val, isSet: true}
}

func (v NullableMarginGetMarginRateLimitOrderV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginRateLimitOrderV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


