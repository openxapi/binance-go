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

// checks if the CfuturesGetLeverageBracketV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CfuturesGetLeverageBracketV1RespItem{}

// CfuturesGetLeverageBracketV1RespItem struct for CfuturesGetLeverageBracketV1RespItem
type CfuturesGetLeverageBracketV1RespItem struct {
	Brackets []CfuturesGetLeverageBracketV1RespItemBracketsInner `json:"brackets,omitempty"`
	Pair *string `json:"pair,omitempty"`
}

// NewCfuturesGetLeverageBracketV1RespItem instantiates a new CfuturesGetLeverageBracketV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCfuturesGetLeverageBracketV1RespItem() *CfuturesGetLeverageBracketV1RespItem {
	this := CfuturesGetLeverageBracketV1RespItem{}
	return &this
}

// NewCfuturesGetLeverageBracketV1RespItemWithDefaults instantiates a new CfuturesGetLeverageBracketV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCfuturesGetLeverageBracketV1RespItemWithDefaults() *CfuturesGetLeverageBracketV1RespItem {
	this := CfuturesGetLeverageBracketV1RespItem{}
	return &this
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *CfuturesGetLeverageBracketV1RespItem) GetBrackets() []CfuturesGetLeverageBracketV1RespItemBracketsInner {
	if o == nil || IsNil(o.Brackets) {
		var ret []CfuturesGetLeverageBracketV1RespItemBracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetLeverageBracketV1RespItem) GetBracketsOk() ([]CfuturesGetLeverageBracketV1RespItemBracketsInner, bool) {
	if o == nil || IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *CfuturesGetLeverageBracketV1RespItem) HasBrackets() bool {
	if o != nil && !IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []CfuturesGetLeverageBracketV1RespItemBracketsInner and assigns it to the Brackets field.
func (o *CfuturesGetLeverageBracketV1RespItem) SetBrackets(v []CfuturesGetLeverageBracketV1RespItemBracketsInner) {
	o.Brackets = v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *CfuturesGetLeverageBracketV1RespItem) GetPair() string {
	if o == nil || IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesGetLeverageBracketV1RespItem) GetPairOk() (*string, bool) {
	if o == nil || IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *CfuturesGetLeverageBracketV1RespItem) HasPair() bool {
	if o != nil && !IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *CfuturesGetLeverageBracketV1RespItem) SetPair(v string) {
	o.Pair = &v
}

func (o CfuturesGetLeverageBracketV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CfuturesGetLeverageBracketV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Brackets) {
		toSerialize["brackets"] = o.Brackets
	}
	if !IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	return toSerialize, nil
}

type NullableCfuturesGetLeverageBracketV1RespItem struct {
	value *CfuturesGetLeverageBracketV1RespItem
	isSet bool
}

func (v NullableCfuturesGetLeverageBracketV1RespItem) Get() *CfuturesGetLeverageBracketV1RespItem {
	return v.value
}

func (v *NullableCfuturesGetLeverageBracketV1RespItem) Set(val *CfuturesGetLeverageBracketV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCfuturesGetLeverageBracketV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCfuturesGetLeverageBracketV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCfuturesGetLeverageBracketV1RespItem(val *CfuturesGetLeverageBracketV1RespItem) *NullableCfuturesGetLeverageBracketV1RespItem {
	return &NullableCfuturesGetLeverageBracketV1RespItem{value: val, isSet: true}
}

func (v NullableCfuturesGetLeverageBracketV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCfuturesGetLeverageBracketV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


