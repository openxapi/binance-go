/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotExchangeFilterExchangeMaxNumAlgoOrders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotExchangeFilterExchangeMaxNumAlgoOrders{}

// SpotExchangeFilterExchangeMaxNumAlgoOrders struct for SpotExchangeFilterExchangeMaxNumAlgoOrders
type SpotExchangeFilterExchangeMaxNumAlgoOrders struct {
	FilterType *string `json:"filterType,omitempty"`
	MaxNumAlgoOrders *int32 `json:"maxNumAlgoOrders,omitempty"`
}

// NewSpotExchangeFilterExchangeMaxNumAlgoOrders instantiates a new SpotExchangeFilterExchangeMaxNumAlgoOrders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotExchangeFilterExchangeMaxNumAlgoOrders() *SpotExchangeFilterExchangeMaxNumAlgoOrders {
	this := SpotExchangeFilterExchangeMaxNumAlgoOrders{}
	return &this
}

// NewSpotExchangeFilterExchangeMaxNumAlgoOrdersWithDefaults instantiates a new SpotExchangeFilterExchangeMaxNumAlgoOrders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotExchangeFilterExchangeMaxNumAlgoOrdersWithDefaults() *SpotExchangeFilterExchangeMaxNumAlgoOrders {
	this := SpotExchangeFilterExchangeMaxNumAlgoOrders{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *SpotExchangeFilterExchangeMaxNumAlgoOrders) GetFilterType() string {
	if o == nil || IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotExchangeFilterExchangeMaxNumAlgoOrders) GetFilterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *SpotExchangeFilterExchangeMaxNumAlgoOrders) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *SpotExchangeFilterExchangeMaxNumAlgoOrders) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumAlgoOrders returns the MaxNumAlgoOrders field value if set, zero value otherwise.
func (o *SpotExchangeFilterExchangeMaxNumAlgoOrders) GetMaxNumAlgoOrders() int32 {
	if o == nil || IsNil(o.MaxNumAlgoOrders) {
		var ret int32
		return ret
	}
	return *o.MaxNumAlgoOrders
}

// GetMaxNumAlgoOrdersOk returns a tuple with the MaxNumAlgoOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotExchangeFilterExchangeMaxNumAlgoOrders) GetMaxNumAlgoOrdersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumAlgoOrders) {
		return nil, false
	}
	return o.MaxNumAlgoOrders, true
}

// HasMaxNumAlgoOrders returns a boolean if a field has been set.
func (o *SpotExchangeFilterExchangeMaxNumAlgoOrders) HasMaxNumAlgoOrders() bool {
	if o != nil && !IsNil(o.MaxNumAlgoOrders) {
		return true
	}

	return false
}

// SetMaxNumAlgoOrders gets a reference to the given int32 and assigns it to the MaxNumAlgoOrders field.
func (o *SpotExchangeFilterExchangeMaxNumAlgoOrders) SetMaxNumAlgoOrders(v int32) {
	o.MaxNumAlgoOrders = &v
}

func (o SpotExchangeFilterExchangeMaxNumAlgoOrders) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotExchangeFilterExchangeMaxNumAlgoOrders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !IsNil(o.MaxNumAlgoOrders) {
		toSerialize["maxNumAlgoOrders"] = o.MaxNumAlgoOrders
	}
	return toSerialize, nil
}

type NullableSpotExchangeFilterExchangeMaxNumAlgoOrders struct {
	value *SpotExchangeFilterExchangeMaxNumAlgoOrders
	isSet bool
}

func (v NullableSpotExchangeFilterExchangeMaxNumAlgoOrders) Get() *SpotExchangeFilterExchangeMaxNumAlgoOrders {
	return v.value
}

func (v *NullableSpotExchangeFilterExchangeMaxNumAlgoOrders) Set(val *SpotExchangeFilterExchangeMaxNumAlgoOrders) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotExchangeFilterExchangeMaxNumAlgoOrders) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotExchangeFilterExchangeMaxNumAlgoOrders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotExchangeFilterExchangeMaxNumAlgoOrders(val *SpotExchangeFilterExchangeMaxNumAlgoOrders) *NullableSpotExchangeFilterExchangeMaxNumAlgoOrders {
	return &NullableSpotExchangeFilterExchangeMaxNumAlgoOrders{value: val, isSet: true}
}

func (v NullableSpotExchangeFilterExchangeMaxNumAlgoOrders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotExchangeFilterExchangeMaxNumAlgoOrders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


