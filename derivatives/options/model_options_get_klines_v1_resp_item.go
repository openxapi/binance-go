/*
Binance Options API

OpenAPI specification for Binance cryptocurrency exchange - Options API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
)

// checks if the OptionsGetKlinesV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsGetKlinesV1RespItem{}

// OptionsGetKlinesV1RespItem struct for OptionsGetKlinesV1RespItem
type OptionsGetKlinesV1RespItem struct {
	Amount *string `json:"amount,omitempty"`
	Close *string `json:"close,omitempty"`
	CloseTime *int64 `json:"closeTime,omitempty"`
	High *string `json:"high,omitempty"`
	Interval *string `json:"interval,omitempty"`
	Low *string `json:"low,omitempty"`
	Open *string `json:"open,omitempty"`
	OpenTime *int64 `json:"openTime,omitempty"`
	TakerAmount *string `json:"takerAmount,omitempty"`
	TakerVolume *string `json:"takerVolume,omitempty"`
	TradeCount *int32 `json:"tradeCount,omitempty"`
	Volume *string `json:"volume,omitempty"`
}

// NewOptionsGetKlinesV1RespItem instantiates a new OptionsGetKlinesV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsGetKlinesV1RespItem() *OptionsGetKlinesV1RespItem {
	this := OptionsGetKlinesV1RespItem{}
	return &this
}

// NewOptionsGetKlinesV1RespItemWithDefaults instantiates a new OptionsGetKlinesV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsGetKlinesV1RespItemWithDefaults() *OptionsGetKlinesV1RespItem {
	this := OptionsGetKlinesV1RespItem{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *OptionsGetKlinesV1RespItem) SetAmount(v string) {
	o.Amount = &v
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetClose() string {
	if o == nil || IsNil(o.Close) {
		var ret string
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetCloseOk() (*string, bool) {
	if o == nil || IsNil(o.Close) {
		return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasClose() bool {
	if o != nil && !IsNil(o.Close) {
		return true
	}

	return false
}

// SetClose gets a reference to the given string and assigns it to the Close field.
func (o *OptionsGetKlinesV1RespItem) SetClose(v string) {
	o.Close = &v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetCloseTime() int64 {
	if o == nil || IsNil(o.CloseTime) {
		var ret int64
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetCloseTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasCloseTime() bool {
	if o != nil && !IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given int64 and assigns it to the CloseTime field.
func (o *OptionsGetKlinesV1RespItem) SetCloseTime(v int64) {
	o.CloseTime = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetHigh() string {
	if o == nil || IsNil(o.High) {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetHighOk() (*string, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *OptionsGetKlinesV1RespItem) SetHigh(v string) {
	o.High = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *OptionsGetKlinesV1RespItem) SetInterval(v string) {
	o.Interval = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetLow() string {
	if o == nil || IsNil(o.Low) {
		var ret string
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetLowOk() (*string, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given string and assigns it to the Low field.
func (o *OptionsGetKlinesV1RespItem) SetLow(v string) {
	o.Low = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetOpen() string {
	if o == nil || IsNil(o.Open) {
		var ret string
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetOpenOk() (*string, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given string and assigns it to the Open field.
func (o *OptionsGetKlinesV1RespItem) SetOpen(v string) {
	o.Open = &v
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetOpenTime() int64 {
	if o == nil || IsNil(o.OpenTime) {
		var ret int64
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetOpenTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasOpenTime() bool {
	if o != nil && !IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given int64 and assigns it to the OpenTime field.
func (o *OptionsGetKlinesV1RespItem) SetOpenTime(v int64) {
	o.OpenTime = &v
}

// GetTakerAmount returns the TakerAmount field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetTakerAmount() string {
	if o == nil || IsNil(o.TakerAmount) {
		var ret string
		return ret
	}
	return *o.TakerAmount
}

// GetTakerAmountOk returns a tuple with the TakerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetTakerAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TakerAmount) {
		return nil, false
	}
	return o.TakerAmount, true
}

// HasTakerAmount returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasTakerAmount() bool {
	if o != nil && !IsNil(o.TakerAmount) {
		return true
	}

	return false
}

// SetTakerAmount gets a reference to the given string and assigns it to the TakerAmount field.
func (o *OptionsGetKlinesV1RespItem) SetTakerAmount(v string) {
	o.TakerAmount = &v
}

// GetTakerVolume returns the TakerVolume field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetTakerVolume() string {
	if o == nil || IsNil(o.TakerVolume) {
		var ret string
		return ret
	}
	return *o.TakerVolume
}

// GetTakerVolumeOk returns a tuple with the TakerVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetTakerVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.TakerVolume) {
		return nil, false
	}
	return o.TakerVolume, true
}

// HasTakerVolume returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasTakerVolume() bool {
	if o != nil && !IsNil(o.TakerVolume) {
		return true
	}

	return false
}

// SetTakerVolume gets a reference to the given string and assigns it to the TakerVolume field.
func (o *OptionsGetKlinesV1RespItem) SetTakerVolume(v string) {
	o.TakerVolume = &v
}

// GetTradeCount returns the TradeCount field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetTradeCount() int32 {
	if o == nil || IsNil(o.TradeCount) {
		var ret int32
		return ret
	}
	return *o.TradeCount
}

// GetTradeCountOk returns a tuple with the TradeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetTradeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TradeCount) {
		return nil, false
	}
	return o.TradeCount, true
}

// HasTradeCount returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasTradeCount() bool {
	if o != nil && !IsNil(o.TradeCount) {
		return true
	}

	return false
}

// SetTradeCount gets a reference to the given int32 and assigns it to the TradeCount field.
func (o *OptionsGetKlinesV1RespItem) SetTradeCount(v int32) {
	o.TradeCount = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *OptionsGetKlinesV1RespItem) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetKlinesV1RespItem) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *OptionsGetKlinesV1RespItem) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *OptionsGetKlinesV1RespItem) SetVolume(v string) {
	o.Volume = &v
}

func (o OptionsGetKlinesV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionsGetKlinesV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Close) {
		toSerialize["close"] = o.Close
	}
	if !IsNil(o.CloseTime) {
		toSerialize["closeTime"] = o.CloseTime
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.OpenTime) {
		toSerialize["openTime"] = o.OpenTime
	}
	if !IsNil(o.TakerAmount) {
		toSerialize["takerAmount"] = o.TakerAmount
	}
	if !IsNil(o.TakerVolume) {
		toSerialize["takerVolume"] = o.TakerVolume
	}
	if !IsNil(o.TradeCount) {
		toSerialize["tradeCount"] = o.TradeCount
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableOptionsGetKlinesV1RespItem struct {
	value *OptionsGetKlinesV1RespItem
	isSet bool
}

func (v NullableOptionsGetKlinesV1RespItem) Get() *OptionsGetKlinesV1RespItem {
	return v.value
}

func (v *NullableOptionsGetKlinesV1RespItem) Set(val *OptionsGetKlinesV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsGetKlinesV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsGetKlinesV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsGetKlinesV1RespItem(val *OptionsGetKlinesV1RespItem) *NullableOptionsGetKlinesV1RespItem {
	return &NullableOptionsGetKlinesV1RespItem{value: val, isSet: true}
}

func (v NullableOptionsGetKlinesV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsGetKlinesV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


