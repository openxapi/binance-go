/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the GetSpotOpenSymbolListV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSpotOpenSymbolListV1RespItem{}

// GetSpotOpenSymbolListV1RespItem struct for GetSpotOpenSymbolListV1RespItem
type GetSpotOpenSymbolListV1RespItem struct {
	OpenTime *int64 `json:"openTime,omitempty"`
	Symbols []string `json:"symbols,omitempty"`
}

// NewGetSpotOpenSymbolListV1RespItem instantiates a new GetSpotOpenSymbolListV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpotOpenSymbolListV1RespItem() *GetSpotOpenSymbolListV1RespItem {
	this := GetSpotOpenSymbolListV1RespItem{}
	return &this
}

// NewGetSpotOpenSymbolListV1RespItemWithDefaults instantiates a new GetSpotOpenSymbolListV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpotOpenSymbolListV1RespItemWithDefaults() *GetSpotOpenSymbolListV1RespItem {
	this := GetSpotOpenSymbolListV1RespItem{}
	return &this
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *GetSpotOpenSymbolListV1RespItem) GetOpenTime() int64 {
	if o == nil || IsNil(o.OpenTime) {
		var ret int64
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotOpenSymbolListV1RespItem) GetOpenTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *GetSpotOpenSymbolListV1RespItem) HasOpenTime() bool {
	if o != nil && !IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given int64 and assigns it to the OpenTime field.
func (o *GetSpotOpenSymbolListV1RespItem) SetOpenTime(v int64) {
	o.OpenTime = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *GetSpotOpenSymbolListV1RespItem) GetSymbols() []string {
	if o == nil || IsNil(o.Symbols) {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotOpenSymbolListV1RespItem) GetSymbolsOk() ([]string, bool) {
	if o == nil || IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *GetSpotOpenSymbolListV1RespItem) HasSymbols() bool {
	if o != nil && !IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *GetSpotOpenSymbolListV1RespItem) SetSymbols(v []string) {
	o.Symbols = v
}

func (o GetSpotOpenSymbolListV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpotOpenSymbolListV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpenTime) {
		toSerialize["openTime"] = o.OpenTime
	}
	if !IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}
	return toSerialize, nil
}

type NullableGetSpotOpenSymbolListV1RespItem struct {
	value *GetSpotOpenSymbolListV1RespItem
	isSet bool
}

func (v NullableGetSpotOpenSymbolListV1RespItem) Get() *GetSpotOpenSymbolListV1RespItem {
	return v.value
}

func (v *NullableGetSpotOpenSymbolListV1RespItem) Set(val *GetSpotOpenSymbolListV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpotOpenSymbolListV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpotOpenSymbolListV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSpotOpenSymbolListV1RespItem(val *GetSpotOpenSymbolListV1RespItem) *NullableGetSpotOpenSymbolListV1RespItem {
	return &NullableGetSpotOpenSymbolListV1RespItem{value: val, isSet: true}
}

func (v NullableGetSpotOpenSymbolListV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpotOpenSymbolListV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


