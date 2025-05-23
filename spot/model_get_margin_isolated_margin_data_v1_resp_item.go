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

// checks if the GetMarginIsolatedMarginDataV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarginIsolatedMarginDataV1RespItem{}

// GetMarginIsolatedMarginDataV1RespItem struct for GetMarginIsolatedMarginDataV1RespItem
type GetMarginIsolatedMarginDataV1RespItem struct {
	Data []GetMarginIsolatedMarginDataV1RespItemDataInner `json:"data,omitempty"`
	Leverage *string `json:"leverage,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	VipLevel *int32 `json:"vipLevel,omitempty"`
}

// NewGetMarginIsolatedMarginDataV1RespItem instantiates a new GetMarginIsolatedMarginDataV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginIsolatedMarginDataV1RespItem() *GetMarginIsolatedMarginDataV1RespItem {
	this := GetMarginIsolatedMarginDataV1RespItem{}
	return &this
}

// NewGetMarginIsolatedMarginDataV1RespItemWithDefaults instantiates a new GetMarginIsolatedMarginDataV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginIsolatedMarginDataV1RespItemWithDefaults() *GetMarginIsolatedMarginDataV1RespItem {
	this := GetMarginIsolatedMarginDataV1RespItem{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetMarginIsolatedMarginDataV1RespItem) GetData() []GetMarginIsolatedMarginDataV1RespItemDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetMarginIsolatedMarginDataV1RespItemDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedMarginDataV1RespItem) GetDataOk() ([]GetMarginIsolatedMarginDataV1RespItemDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetMarginIsolatedMarginDataV1RespItem) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetMarginIsolatedMarginDataV1RespItemDataInner and assigns it to the Data field.
func (o *GetMarginIsolatedMarginDataV1RespItem) SetData(v []GetMarginIsolatedMarginDataV1RespItemDataInner) {
	o.Data = v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *GetMarginIsolatedMarginDataV1RespItem) GetLeverage() string {
	if o == nil || IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedMarginDataV1RespItem) GetLeverageOk() (*string, bool) {
	if o == nil || IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *GetMarginIsolatedMarginDataV1RespItem) HasLeverage() bool {
	if o != nil && !IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *GetMarginIsolatedMarginDataV1RespItem) SetLeverage(v string) {
	o.Leverage = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetMarginIsolatedMarginDataV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedMarginDataV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetMarginIsolatedMarginDataV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetMarginIsolatedMarginDataV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetVipLevel returns the VipLevel field value if set, zero value otherwise.
func (o *GetMarginIsolatedMarginDataV1RespItem) GetVipLevel() int32 {
	if o == nil || IsNil(o.VipLevel) {
		var ret int32
		return ret
	}
	return *o.VipLevel
}

// GetVipLevelOk returns a tuple with the VipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedMarginDataV1RespItem) GetVipLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.VipLevel) {
		return nil, false
	}
	return o.VipLevel, true
}

// HasVipLevel returns a boolean if a field has been set.
func (o *GetMarginIsolatedMarginDataV1RespItem) HasVipLevel() bool {
	if o != nil && !IsNil(o.VipLevel) {
		return true
	}

	return false
}

// SetVipLevel gets a reference to the given int32 and assigns it to the VipLevel field.
func (o *GetMarginIsolatedMarginDataV1RespItem) SetVipLevel(v int32) {
	o.VipLevel = &v
}

func (o GetMarginIsolatedMarginDataV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginIsolatedMarginDataV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.VipLevel) {
		toSerialize["vipLevel"] = o.VipLevel
	}
	return toSerialize, nil
}

type NullableGetMarginIsolatedMarginDataV1RespItem struct {
	value *GetMarginIsolatedMarginDataV1RespItem
	isSet bool
}

func (v NullableGetMarginIsolatedMarginDataV1RespItem) Get() *GetMarginIsolatedMarginDataV1RespItem {
	return v.value
}

func (v *NullableGetMarginIsolatedMarginDataV1RespItem) Set(val *GetMarginIsolatedMarginDataV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginIsolatedMarginDataV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginIsolatedMarginDataV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginIsolatedMarginDataV1RespItem(val *GetMarginIsolatedMarginDataV1RespItem) *NullableGetMarginIsolatedMarginDataV1RespItem {
	return &NullableGetMarginIsolatedMarginDataV1RespItem{value: val, isSet: true}
}

func (v NullableGetMarginIsolatedMarginDataV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginIsolatedMarginDataV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


