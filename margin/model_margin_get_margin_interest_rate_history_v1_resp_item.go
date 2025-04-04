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

// checks if the MarginGetMarginInterestRateHistoryV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginInterestRateHistoryV1RespItem{}

// MarginGetMarginInterestRateHistoryV1RespItem struct for MarginGetMarginInterestRateHistoryV1RespItem
type MarginGetMarginInterestRateHistoryV1RespItem struct {
	Asset *string `json:"asset,omitempty"`
	DailyInterestRate *string `json:"dailyInterestRate,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	VipLevel *int32 `json:"vipLevel,omitempty"`
}

// NewMarginGetMarginInterestRateHistoryV1RespItem instantiates a new MarginGetMarginInterestRateHistoryV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginInterestRateHistoryV1RespItem() *MarginGetMarginInterestRateHistoryV1RespItem {
	this := MarginGetMarginInterestRateHistoryV1RespItem{}
	return &this
}

// NewMarginGetMarginInterestRateHistoryV1RespItemWithDefaults instantiates a new MarginGetMarginInterestRateHistoryV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginInterestRateHistoryV1RespItemWithDefaults() *MarginGetMarginInterestRateHistoryV1RespItem {
	this := MarginGetMarginInterestRateHistoryV1RespItem{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) SetAsset(v string) {
	o.Asset = &v
}

// GetDailyInterestRate returns the DailyInterestRate field value if set, zero value otherwise.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) GetDailyInterestRate() string {
	if o == nil || IsNil(o.DailyInterestRate) {
		var ret string
		return ret
	}
	return *o.DailyInterestRate
}

// GetDailyInterestRateOk returns a tuple with the DailyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) GetDailyInterestRateOk() (*string, bool) {
	if o == nil || IsNil(o.DailyInterestRate) {
		return nil, false
	}
	return o.DailyInterestRate, true
}

// HasDailyInterestRate returns a boolean if a field has been set.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) HasDailyInterestRate() bool {
	if o != nil && !IsNil(o.DailyInterestRate) {
		return true
	}

	return false
}

// SetDailyInterestRate gets a reference to the given string and assigns it to the DailyInterestRate field.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) SetDailyInterestRate(v string) {
	o.DailyInterestRate = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetVipLevel returns the VipLevel field value if set, zero value otherwise.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) GetVipLevel() int32 {
	if o == nil || IsNil(o.VipLevel) {
		var ret int32
		return ret
	}
	return *o.VipLevel
}

// GetVipLevelOk returns a tuple with the VipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) GetVipLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.VipLevel) {
		return nil, false
	}
	return o.VipLevel, true
}

// HasVipLevel returns a boolean if a field has been set.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) HasVipLevel() bool {
	if o != nil && !IsNil(o.VipLevel) {
		return true
	}

	return false
}

// SetVipLevel gets a reference to the given int32 and assigns it to the VipLevel field.
func (o *MarginGetMarginInterestRateHistoryV1RespItem) SetVipLevel(v int32) {
	o.VipLevel = &v
}

func (o MarginGetMarginInterestRateHistoryV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginInterestRateHistoryV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.DailyInterestRate) {
		toSerialize["dailyInterestRate"] = o.DailyInterestRate
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.VipLevel) {
		toSerialize["vipLevel"] = o.VipLevel
	}
	return toSerialize, nil
}

type NullableMarginGetMarginInterestRateHistoryV1RespItem struct {
	value *MarginGetMarginInterestRateHistoryV1RespItem
	isSet bool
}

func (v NullableMarginGetMarginInterestRateHistoryV1RespItem) Get() *MarginGetMarginInterestRateHistoryV1RespItem {
	return v.value
}

func (v *NullableMarginGetMarginInterestRateHistoryV1RespItem) Set(val *MarginGetMarginInterestRateHistoryV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginInterestRateHistoryV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginInterestRateHistoryV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginInterestRateHistoryV1RespItem(val *MarginGetMarginInterestRateHistoryV1RespItem) *NullableMarginGetMarginInterestRateHistoryV1RespItem {
	return &NullableMarginGetMarginInterestRateHistoryV1RespItem{value: val, isSet: true}
}

func (v NullableMarginGetMarginInterestRateHistoryV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginInterestRateHistoryV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


