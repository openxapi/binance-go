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

// checks if the GetExerciseHistoryV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExerciseHistoryV1RespItem{}

// GetExerciseHistoryV1RespItem struct for GetExerciseHistoryV1RespItem
type GetExerciseHistoryV1RespItem struct {
	ExpiryDate *int32 `json:"expiryDate,omitempty"`
	RealStrikePrice *string `json:"realStrikePrice,omitempty"`
	StrikePrice *string `json:"strikePrice,omitempty"`
	StrikeResult *string `json:"strikeResult,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewGetExerciseHistoryV1RespItem instantiates a new GetExerciseHistoryV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExerciseHistoryV1RespItem() *GetExerciseHistoryV1RespItem {
	this := GetExerciseHistoryV1RespItem{}
	return &this
}

// NewGetExerciseHistoryV1RespItemWithDefaults instantiates a new GetExerciseHistoryV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExerciseHistoryV1RespItemWithDefaults() *GetExerciseHistoryV1RespItem {
	this := GetExerciseHistoryV1RespItem{}
	return &this
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *GetExerciseHistoryV1RespItem) GetExpiryDate() int32 {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret int32
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExerciseHistoryV1RespItem) GetExpiryDateOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *GetExerciseHistoryV1RespItem) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given int32 and assigns it to the ExpiryDate field.
func (o *GetExerciseHistoryV1RespItem) SetExpiryDate(v int32) {
	o.ExpiryDate = &v
}

// GetRealStrikePrice returns the RealStrikePrice field value if set, zero value otherwise.
func (o *GetExerciseHistoryV1RespItem) GetRealStrikePrice() string {
	if o == nil || IsNil(o.RealStrikePrice) {
		var ret string
		return ret
	}
	return *o.RealStrikePrice
}

// GetRealStrikePriceOk returns a tuple with the RealStrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExerciseHistoryV1RespItem) GetRealStrikePriceOk() (*string, bool) {
	if o == nil || IsNil(o.RealStrikePrice) {
		return nil, false
	}
	return o.RealStrikePrice, true
}

// HasRealStrikePrice returns a boolean if a field has been set.
func (o *GetExerciseHistoryV1RespItem) HasRealStrikePrice() bool {
	if o != nil && !IsNil(o.RealStrikePrice) {
		return true
	}

	return false
}

// SetRealStrikePrice gets a reference to the given string and assigns it to the RealStrikePrice field.
func (o *GetExerciseHistoryV1RespItem) SetRealStrikePrice(v string) {
	o.RealStrikePrice = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *GetExerciseHistoryV1RespItem) GetStrikePrice() string {
	if o == nil || IsNil(o.StrikePrice) {
		var ret string
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExerciseHistoryV1RespItem) GetStrikePriceOk() (*string, bool) {
	if o == nil || IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *GetExerciseHistoryV1RespItem) HasStrikePrice() bool {
	if o != nil && !IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given string and assigns it to the StrikePrice field.
func (o *GetExerciseHistoryV1RespItem) SetStrikePrice(v string) {
	o.StrikePrice = &v
}

// GetStrikeResult returns the StrikeResult field value if set, zero value otherwise.
func (o *GetExerciseHistoryV1RespItem) GetStrikeResult() string {
	if o == nil || IsNil(o.StrikeResult) {
		var ret string
		return ret
	}
	return *o.StrikeResult
}

// GetStrikeResultOk returns a tuple with the StrikeResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExerciseHistoryV1RespItem) GetStrikeResultOk() (*string, bool) {
	if o == nil || IsNil(o.StrikeResult) {
		return nil, false
	}
	return o.StrikeResult, true
}

// HasStrikeResult returns a boolean if a field has been set.
func (o *GetExerciseHistoryV1RespItem) HasStrikeResult() bool {
	if o != nil && !IsNil(o.StrikeResult) {
		return true
	}

	return false
}

// SetStrikeResult gets a reference to the given string and assigns it to the StrikeResult field.
func (o *GetExerciseHistoryV1RespItem) SetStrikeResult(v string) {
	o.StrikeResult = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetExerciseHistoryV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExerciseHistoryV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetExerciseHistoryV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetExerciseHistoryV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

func (o GetExerciseHistoryV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExerciseHistoryV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !IsNil(o.RealStrikePrice) {
		toSerialize["realStrikePrice"] = o.RealStrikePrice
	}
	if !IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
	}
	if !IsNil(o.StrikeResult) {
		toSerialize["strikeResult"] = o.StrikeResult
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableGetExerciseHistoryV1RespItem struct {
	value *GetExerciseHistoryV1RespItem
	isSet bool
}

func (v NullableGetExerciseHistoryV1RespItem) Get() *GetExerciseHistoryV1RespItem {
	return v.value
}

func (v *NullableGetExerciseHistoryV1RespItem) Set(val *GetExerciseHistoryV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExerciseHistoryV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExerciseHistoryV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExerciseHistoryV1RespItem(val *GetExerciseHistoryV1RespItem) *NullableGetExerciseHistoryV1RespItem {
	return &NullableGetExerciseHistoryV1RespItem{value: val, isSet: true}
}

func (v NullableGetExerciseHistoryV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExerciseHistoryV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


