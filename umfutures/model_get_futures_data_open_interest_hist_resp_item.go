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

// checks if the GetFuturesDataOpenInterestHistRespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFuturesDataOpenInterestHistRespItem{}

// GetFuturesDataOpenInterestHistRespItem struct for GetFuturesDataOpenInterestHistRespItem
type GetFuturesDataOpenInterestHistRespItem struct {
	SumOpenInterest *string `json:"sumOpenInterest,omitempty"`
	SumOpenInterestValue *string `json:"sumOpenInterestValue,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewGetFuturesDataOpenInterestHistRespItem instantiates a new GetFuturesDataOpenInterestHistRespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesDataOpenInterestHistRespItem() *GetFuturesDataOpenInterestHistRespItem {
	this := GetFuturesDataOpenInterestHistRespItem{}
	return &this
}

// NewGetFuturesDataOpenInterestHistRespItemWithDefaults instantiates a new GetFuturesDataOpenInterestHistRespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesDataOpenInterestHistRespItemWithDefaults() *GetFuturesDataOpenInterestHistRespItem {
	this := GetFuturesDataOpenInterestHistRespItem{}
	return &this
}

// GetSumOpenInterest returns the SumOpenInterest field value if set, zero value otherwise.
func (o *GetFuturesDataOpenInterestHistRespItem) GetSumOpenInterest() string {
	if o == nil || IsNil(o.SumOpenInterest) {
		var ret string
		return ret
	}
	return *o.SumOpenInterest
}

// GetSumOpenInterestOk returns a tuple with the SumOpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataOpenInterestHistRespItem) GetSumOpenInterestOk() (*string, bool) {
	if o == nil || IsNil(o.SumOpenInterest) {
		return nil, false
	}
	return o.SumOpenInterest, true
}

// HasSumOpenInterest returns a boolean if a field has been set.
func (o *GetFuturesDataOpenInterestHistRespItem) HasSumOpenInterest() bool {
	if o != nil && !IsNil(o.SumOpenInterest) {
		return true
	}

	return false
}

// SetSumOpenInterest gets a reference to the given string and assigns it to the SumOpenInterest field.
func (o *GetFuturesDataOpenInterestHistRespItem) SetSumOpenInterest(v string) {
	o.SumOpenInterest = &v
}

// GetSumOpenInterestValue returns the SumOpenInterestValue field value if set, zero value otherwise.
func (o *GetFuturesDataOpenInterestHistRespItem) GetSumOpenInterestValue() string {
	if o == nil || IsNil(o.SumOpenInterestValue) {
		var ret string
		return ret
	}
	return *o.SumOpenInterestValue
}

// GetSumOpenInterestValueOk returns a tuple with the SumOpenInterestValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataOpenInterestHistRespItem) GetSumOpenInterestValueOk() (*string, bool) {
	if o == nil || IsNil(o.SumOpenInterestValue) {
		return nil, false
	}
	return o.SumOpenInterestValue, true
}

// HasSumOpenInterestValue returns a boolean if a field has been set.
func (o *GetFuturesDataOpenInterestHistRespItem) HasSumOpenInterestValue() bool {
	if o != nil && !IsNil(o.SumOpenInterestValue) {
		return true
	}

	return false
}

// SetSumOpenInterestValue gets a reference to the given string and assigns it to the SumOpenInterestValue field.
func (o *GetFuturesDataOpenInterestHistRespItem) SetSumOpenInterestValue(v string) {
	o.SumOpenInterestValue = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetFuturesDataOpenInterestHistRespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataOpenInterestHistRespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetFuturesDataOpenInterestHistRespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetFuturesDataOpenInterestHistRespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetFuturesDataOpenInterestHistRespItem) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataOpenInterestHistRespItem) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetFuturesDataOpenInterestHistRespItem) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *GetFuturesDataOpenInterestHistRespItem) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o GetFuturesDataOpenInterestHistRespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesDataOpenInterestHistRespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SumOpenInterest) {
		toSerialize["sumOpenInterest"] = o.SumOpenInterest
	}
	if !IsNil(o.SumOpenInterestValue) {
		toSerialize["sumOpenInterestValue"] = o.SumOpenInterestValue
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableGetFuturesDataOpenInterestHistRespItem struct {
	value *GetFuturesDataOpenInterestHistRespItem
	isSet bool
}

func (v NullableGetFuturesDataOpenInterestHistRespItem) Get() *GetFuturesDataOpenInterestHistRespItem {
	return v.value
}

func (v *NullableGetFuturesDataOpenInterestHistRespItem) Set(val *GetFuturesDataOpenInterestHistRespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesDataOpenInterestHistRespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesDataOpenInterestHistRespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesDataOpenInterestHistRespItem(val *GetFuturesDataOpenInterestHistRespItem) *NullableGetFuturesDataOpenInterestHistRespItem {
	return &NullableGetFuturesDataOpenInterestHistRespItem{value: val, isSet: true}
}

func (v NullableGetFuturesDataOpenInterestHistRespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesDataOpenInterestHistRespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


