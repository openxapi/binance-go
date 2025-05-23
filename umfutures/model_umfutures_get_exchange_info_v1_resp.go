/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the UmfuturesGetExchangeInfoV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmfuturesGetExchangeInfoV1Resp{}

// UmfuturesGetExchangeInfoV1Resp struct for UmfuturesGetExchangeInfoV1Resp
type UmfuturesGetExchangeInfoV1Resp struct {
	Assets []UmfuturesGetExchangeInfoV1RespAssetsInner `json:"assets,omitempty"`
	ExchangeFilters []string `json:"exchangeFilters,omitempty"`
	RateLimits []UmfuturesGetExchangeInfoV1RespRateLimitsInner `json:"rateLimits,omitempty"`
	ServerTime *int32 `json:"serverTime,omitempty"`
	Symbols []UmfuturesGetExchangeInfoV1RespSymbolsInner `json:"symbols,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// NewUmfuturesGetExchangeInfoV1Resp instantiates a new UmfuturesGetExchangeInfoV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmfuturesGetExchangeInfoV1Resp() *UmfuturesGetExchangeInfoV1Resp {
	this := UmfuturesGetExchangeInfoV1Resp{}
	return &this
}

// NewUmfuturesGetExchangeInfoV1RespWithDefaults instantiates a new UmfuturesGetExchangeInfoV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmfuturesGetExchangeInfoV1RespWithDefaults() *UmfuturesGetExchangeInfoV1Resp {
	this := UmfuturesGetExchangeInfoV1Resp{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *UmfuturesGetExchangeInfoV1Resp) GetAssets() []UmfuturesGetExchangeInfoV1RespAssetsInner {
	if o == nil || IsNil(o.Assets) {
		var ret []UmfuturesGetExchangeInfoV1RespAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) GetAssetsOk() ([]UmfuturesGetExchangeInfoV1RespAssetsInner, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []UmfuturesGetExchangeInfoV1RespAssetsInner and assigns it to the Assets field.
func (o *UmfuturesGetExchangeInfoV1Resp) SetAssets(v []UmfuturesGetExchangeInfoV1RespAssetsInner) {
	o.Assets = v
}

// GetExchangeFilters returns the ExchangeFilters field value if set, zero value otherwise.
func (o *UmfuturesGetExchangeInfoV1Resp) GetExchangeFilters() []string {
	if o == nil || IsNil(o.ExchangeFilters) {
		var ret []string
		return ret
	}
	return o.ExchangeFilters
}

// GetExchangeFiltersOk returns a tuple with the ExchangeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) GetExchangeFiltersOk() ([]string, bool) {
	if o == nil || IsNil(o.ExchangeFilters) {
		return nil, false
	}
	return o.ExchangeFilters, true
}

// HasExchangeFilters returns a boolean if a field has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) HasExchangeFilters() bool {
	if o != nil && !IsNil(o.ExchangeFilters) {
		return true
	}

	return false
}

// SetExchangeFilters gets a reference to the given []string and assigns it to the ExchangeFilters field.
func (o *UmfuturesGetExchangeInfoV1Resp) SetExchangeFilters(v []string) {
	o.ExchangeFilters = v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *UmfuturesGetExchangeInfoV1Resp) GetRateLimits() []UmfuturesGetExchangeInfoV1RespRateLimitsInner {
	if o == nil || IsNil(o.RateLimits) {
		var ret []UmfuturesGetExchangeInfoV1RespRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) GetRateLimitsOk() ([]UmfuturesGetExchangeInfoV1RespRateLimitsInner, bool) {
	if o == nil || IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) HasRateLimits() bool {
	if o != nil && !IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []UmfuturesGetExchangeInfoV1RespRateLimitsInner and assigns it to the RateLimits field.
func (o *UmfuturesGetExchangeInfoV1Resp) SetRateLimits(v []UmfuturesGetExchangeInfoV1RespRateLimitsInner) {
	o.RateLimits = v
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *UmfuturesGetExchangeInfoV1Resp) GetServerTime() int32 {
	if o == nil || IsNil(o.ServerTime) {
		var ret int32
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) GetServerTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerTime) {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) HasServerTime() bool {
	if o != nil && !IsNil(o.ServerTime) {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int32 and assigns it to the ServerTime field.
func (o *UmfuturesGetExchangeInfoV1Resp) SetServerTime(v int32) {
	o.ServerTime = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *UmfuturesGetExchangeInfoV1Resp) GetSymbols() []UmfuturesGetExchangeInfoV1RespSymbolsInner {
	if o == nil || IsNil(o.Symbols) {
		var ret []UmfuturesGetExchangeInfoV1RespSymbolsInner
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) GetSymbolsOk() ([]UmfuturesGetExchangeInfoV1RespSymbolsInner, bool) {
	if o == nil || IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) HasSymbols() bool {
	if o != nil && !IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []UmfuturesGetExchangeInfoV1RespSymbolsInner and assigns it to the Symbols field.
func (o *UmfuturesGetExchangeInfoV1Resp) SetSymbols(v []UmfuturesGetExchangeInfoV1RespSymbolsInner) {
	o.Symbols = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *UmfuturesGetExchangeInfoV1Resp) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UmfuturesGetExchangeInfoV1Resp) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *UmfuturesGetExchangeInfoV1Resp) SetTimezone(v string) {
	o.Timezone = &v
}

func (o UmfuturesGetExchangeInfoV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmfuturesGetExchangeInfoV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.ExchangeFilters) {
		toSerialize["exchangeFilters"] = o.ExchangeFilters
	}
	if !IsNil(o.RateLimits) {
		toSerialize["rateLimits"] = o.RateLimits
	}
	if !IsNil(o.ServerTime) {
		toSerialize["serverTime"] = o.ServerTime
	}
	if !IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	return toSerialize, nil
}

type NullableUmfuturesGetExchangeInfoV1Resp struct {
	value *UmfuturesGetExchangeInfoV1Resp
	isSet bool
}

func (v NullableUmfuturesGetExchangeInfoV1Resp) Get() *UmfuturesGetExchangeInfoV1Resp {
	return v.value
}

func (v *NullableUmfuturesGetExchangeInfoV1Resp) Set(val *UmfuturesGetExchangeInfoV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesGetExchangeInfoV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesGetExchangeInfoV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesGetExchangeInfoV1Resp(val *UmfuturesGetExchangeInfoV1Resp) *NullableUmfuturesGetExchangeInfoV1Resp {
	return &NullableUmfuturesGetExchangeInfoV1Resp{value: val, isSet: true}
}

func (v NullableUmfuturesGetExchangeInfoV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesGetExchangeInfoV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


