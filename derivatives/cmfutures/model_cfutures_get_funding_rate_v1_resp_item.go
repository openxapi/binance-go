/*
Binance Cfutures API

OpenAPI specification for Binance cryptocurrency exchange - Cfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
)

// checks if the CfuturesGetFundingRateV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CfuturesGetFundingRateV1RespItem{}

// CfuturesGetFundingRateV1RespItem struct for CfuturesGetFundingRateV1RespItem
type CfuturesGetFundingRateV1RespItem struct {
	FundingRate *string `json:"fundingRate,omitempty"`
	FundingTime *int64 `json:"fundingTime,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewCfuturesGetFundingRateV1RespItem instantiates a new CfuturesGetFundingRateV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCfuturesGetFundingRateV1RespItem() *CfuturesGetFundingRateV1RespItem {
	this := CfuturesGetFundingRateV1RespItem{}
	return &this
}

// NewCfuturesGetFundingRateV1RespItemWithDefaults instantiates a new CfuturesGetFundingRateV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCfuturesGetFundingRateV1RespItemWithDefaults() *CfuturesGetFundingRateV1RespItem {
	this := CfuturesGetFundingRateV1RespItem{}
	return &this
}

// GetFundingRate returns the FundingRate field value if set, zero value otherwise.
func (o *CfuturesGetFundingRateV1RespItem) GetFundingRate() string {
	if o == nil || IsNil(o.FundingRate) {
		var ret string
		return ret
	}
	return *o.FundingRate
}

// GetFundingRateOk returns a tuple with the FundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetFundingRateV1RespItem) GetFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.FundingRate) {
		return nil, false
	}
	return o.FundingRate, true
}

// HasFundingRate returns a boolean if a field has been set.
func (o *CfuturesGetFundingRateV1RespItem) HasFundingRate() bool {
	if o != nil && !IsNil(o.FundingRate) {
		return true
	}

	return false
}

// SetFundingRate gets a reference to the given string and assigns it to the FundingRate field.
func (o *CfuturesGetFundingRateV1RespItem) SetFundingRate(v string) {
	o.FundingRate = &v
}

// GetFundingTime returns the FundingTime field value if set, zero value otherwise.
func (o *CfuturesGetFundingRateV1RespItem) GetFundingTime() int64 {
	if o == nil || IsNil(o.FundingTime) {
		var ret int64
		return ret
	}
	return *o.FundingTime
}

// GetFundingTimeOk returns a tuple with the FundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetFundingRateV1RespItem) GetFundingTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.FundingTime) {
		return nil, false
	}
	return o.FundingTime, true
}

// HasFundingTime returns a boolean if a field has been set.
func (o *CfuturesGetFundingRateV1RespItem) HasFundingTime() bool {
	if o != nil && !IsNil(o.FundingTime) {
		return true
	}

	return false
}

// SetFundingTime gets a reference to the given int64 and assigns it to the FundingTime field.
func (o *CfuturesGetFundingRateV1RespItem) SetFundingTime(v int64) {
	o.FundingTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CfuturesGetFundingRateV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetFundingRateV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CfuturesGetFundingRateV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CfuturesGetFundingRateV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

func (o CfuturesGetFundingRateV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CfuturesGetFundingRateV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FundingRate) {
		toSerialize["fundingRate"] = o.FundingRate
	}
	if !IsNil(o.FundingTime) {
		toSerialize["fundingTime"] = o.FundingTime
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableCfuturesGetFundingRateV1RespItem struct {
	value *CfuturesGetFundingRateV1RespItem
	isSet bool
}

func (v NullableCfuturesGetFundingRateV1RespItem) Get() *CfuturesGetFundingRateV1RespItem {
	return v.value
}

func (v *NullableCfuturesGetFundingRateV1RespItem) Set(val *CfuturesGetFundingRateV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCfuturesGetFundingRateV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCfuturesGetFundingRateV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCfuturesGetFundingRateV1RespItem(val *CfuturesGetFundingRateV1RespItem) *NullableCfuturesGetFundingRateV1RespItem {
	return &NullableCfuturesGetFundingRateV1RespItem{value: val, isSet: true}
}

func (v NullableCfuturesGetFundingRateV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCfuturesGetFundingRateV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


