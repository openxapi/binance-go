/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginGetCmAdlQuantileV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetCmAdlQuantileV1RespItem{}

// PmarginGetCmAdlQuantileV1RespItem struct for PmarginGetCmAdlQuantileV1RespItem
type PmarginGetCmAdlQuantileV1RespItem struct {
	AdlQuantile *PmarginGetCmAdlQuantileV1RespItemAdlQuantile `json:"adlQuantile,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewPmarginGetCmAdlQuantileV1RespItem instantiates a new PmarginGetCmAdlQuantileV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetCmAdlQuantileV1RespItem() *PmarginGetCmAdlQuantileV1RespItem {
	this := PmarginGetCmAdlQuantileV1RespItem{}
	return &this
}

// NewPmarginGetCmAdlQuantileV1RespItemWithDefaults instantiates a new PmarginGetCmAdlQuantileV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetCmAdlQuantileV1RespItemWithDefaults() *PmarginGetCmAdlQuantileV1RespItem {
	this := PmarginGetCmAdlQuantileV1RespItem{}
	return &this
}

// GetAdlQuantile returns the AdlQuantile field value if set, zero value otherwise.
func (o *PmarginGetCmAdlQuantileV1RespItem) GetAdlQuantile() PmarginGetCmAdlQuantileV1RespItemAdlQuantile {
	if o == nil || IsNil(o.AdlQuantile) {
		var ret PmarginGetCmAdlQuantileV1RespItemAdlQuantile
		return ret
	}
	return *o.AdlQuantile
}

// GetAdlQuantileOk returns a tuple with the AdlQuantile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAdlQuantileV1RespItem) GetAdlQuantileOk() (*PmarginGetCmAdlQuantileV1RespItemAdlQuantile, bool) {
	if o == nil || IsNil(o.AdlQuantile) {
		return nil, false
	}
	return o.AdlQuantile, true
}

// HasAdlQuantile returns a boolean if a field has been set.
func (o *PmarginGetCmAdlQuantileV1RespItem) HasAdlQuantile() bool {
	if o != nil && !IsNil(o.AdlQuantile) {
		return true
	}

	return false
}

// SetAdlQuantile gets a reference to the given PmarginGetCmAdlQuantileV1RespItemAdlQuantile and assigns it to the AdlQuantile field.
func (o *PmarginGetCmAdlQuantileV1RespItem) SetAdlQuantile(v PmarginGetCmAdlQuantileV1RespItemAdlQuantile) {
	o.AdlQuantile = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PmarginGetCmAdlQuantileV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAdlQuantileV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PmarginGetCmAdlQuantileV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PmarginGetCmAdlQuantileV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

func (o PmarginGetCmAdlQuantileV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetCmAdlQuantileV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdlQuantile) {
		toSerialize["adlQuantile"] = o.AdlQuantile
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullablePmarginGetCmAdlQuantileV1RespItem struct {
	value *PmarginGetCmAdlQuantileV1RespItem
	isSet bool
}

func (v NullablePmarginGetCmAdlQuantileV1RespItem) Get() *PmarginGetCmAdlQuantileV1RespItem {
	return v.value
}

func (v *NullablePmarginGetCmAdlQuantileV1RespItem) Set(val *PmarginGetCmAdlQuantileV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetCmAdlQuantileV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetCmAdlQuantileV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetCmAdlQuantileV1RespItem(val *PmarginGetCmAdlQuantileV1RespItem) *NullablePmarginGetCmAdlQuantileV1RespItem {
	return &NullablePmarginGetCmAdlQuantileV1RespItem{value: val, isSet: true}
}

func (v NullablePmarginGetCmAdlQuantileV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetCmAdlQuantileV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


