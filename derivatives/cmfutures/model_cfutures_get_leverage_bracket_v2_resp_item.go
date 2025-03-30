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

// checks if the CfuturesGetLeverageBracketV2RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CfuturesGetLeverageBracketV2RespItem{}

// CfuturesGetLeverageBracketV2RespItem struct for CfuturesGetLeverageBracketV2RespItem
type CfuturesGetLeverageBracketV2RespItem struct {
	Brackets []CfuturesGetLeverageBracketV1RespItemBracketsInner `json:"brackets,omitempty"`
	NotionalCoef *float32 `json:"notionalCoef,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewCfuturesGetLeverageBracketV2RespItem instantiates a new CfuturesGetLeverageBracketV2RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCfuturesGetLeverageBracketV2RespItem() *CfuturesGetLeverageBracketV2RespItem {
	this := CfuturesGetLeverageBracketV2RespItem{}
	return &this
}

// NewCfuturesGetLeverageBracketV2RespItemWithDefaults instantiates a new CfuturesGetLeverageBracketV2RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCfuturesGetLeverageBracketV2RespItemWithDefaults() *CfuturesGetLeverageBracketV2RespItem {
	this := CfuturesGetLeverageBracketV2RespItem{}
	return &this
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *CfuturesGetLeverageBracketV2RespItem) GetBrackets() []CfuturesGetLeverageBracketV1RespItemBracketsInner {
	if o == nil || IsNil(o.Brackets) {
		var ret []CfuturesGetLeverageBracketV1RespItemBracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetLeverageBracketV2RespItem) GetBracketsOk() ([]CfuturesGetLeverageBracketV1RespItemBracketsInner, bool) {
	if o == nil || IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *CfuturesGetLeverageBracketV2RespItem) HasBrackets() bool {
	if o != nil && !IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []CfuturesGetLeverageBracketV1RespItemBracketsInner and assigns it to the Brackets field.
func (o *CfuturesGetLeverageBracketV2RespItem) SetBrackets(v []CfuturesGetLeverageBracketV1RespItemBracketsInner) {
	o.Brackets = v
}

// GetNotionalCoef returns the NotionalCoef field value if set, zero value otherwise.
func (o *CfuturesGetLeverageBracketV2RespItem) GetNotionalCoef() float32 {
	if o == nil || IsNil(o.NotionalCoef) {
		var ret float32
		return ret
	}
	return *o.NotionalCoef
}

// GetNotionalCoefOk returns a tuple with the NotionalCoef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetLeverageBracketV2RespItem) GetNotionalCoefOk() (*float32, bool) {
	if o == nil || IsNil(o.NotionalCoef) {
		return nil, false
	}
	return o.NotionalCoef, true
}

// HasNotionalCoef returns a boolean if a field has been set.
func (o *CfuturesGetLeverageBracketV2RespItem) HasNotionalCoef() bool {
	if o != nil && !IsNil(o.NotionalCoef) {
		return true
	}

	return false
}

// SetNotionalCoef gets a reference to the given float32 and assigns it to the NotionalCoef field.
func (o *CfuturesGetLeverageBracketV2RespItem) SetNotionalCoef(v float32) {
	o.NotionalCoef = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CfuturesGetLeverageBracketV2RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetLeverageBracketV2RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CfuturesGetLeverageBracketV2RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CfuturesGetLeverageBracketV2RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

func (o CfuturesGetLeverageBracketV2RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CfuturesGetLeverageBracketV2RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Brackets) {
		toSerialize["brackets"] = o.Brackets
	}
	if !IsNil(o.NotionalCoef) {
		toSerialize["notionalCoef"] = o.NotionalCoef
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableCfuturesGetLeverageBracketV2RespItem struct {
	value *CfuturesGetLeverageBracketV2RespItem
	isSet bool
}

func (v NullableCfuturesGetLeverageBracketV2RespItem) Get() *CfuturesGetLeverageBracketV2RespItem {
	return v.value
}

func (v *NullableCfuturesGetLeverageBracketV2RespItem) Set(val *CfuturesGetLeverageBracketV2RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCfuturesGetLeverageBracketV2RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCfuturesGetLeverageBracketV2RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCfuturesGetLeverageBracketV2RespItem(val *CfuturesGetLeverageBracketV2RespItem) *NullableCfuturesGetLeverageBracketV2RespItem {
	return &NullableCfuturesGetLeverageBracketV2RespItem{value: val, isSet: true}
}

func (v NullableCfuturesGetLeverageBracketV2RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCfuturesGetLeverageBracketV2RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


