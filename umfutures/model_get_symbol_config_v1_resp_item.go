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

// checks if the GetSymbolConfigV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSymbolConfigV1RespItem{}

// GetSymbolConfigV1RespItem struct for GetSymbolConfigV1RespItem
type GetSymbolConfigV1RespItem struct {
	IsAutoAddMargin *string `json:"isAutoAddMargin,omitempty"`
	Leverage *int32 `json:"leverage,omitempty"`
	MarginType *string `json:"marginType,omitempty"`
	MaxNotionalValue *string `json:"maxNotionalValue,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewGetSymbolConfigV1RespItem instantiates a new GetSymbolConfigV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSymbolConfigV1RespItem() *GetSymbolConfigV1RespItem {
	this := GetSymbolConfigV1RespItem{}
	return &this
}

// NewGetSymbolConfigV1RespItemWithDefaults instantiates a new GetSymbolConfigV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSymbolConfigV1RespItemWithDefaults() *GetSymbolConfigV1RespItem {
	this := GetSymbolConfigV1RespItem{}
	return &this
}

// GetIsAutoAddMargin returns the IsAutoAddMargin field value if set, zero value otherwise.
func (o *GetSymbolConfigV1RespItem) GetIsAutoAddMargin() string {
	if o == nil || IsNil(o.IsAutoAddMargin) {
		var ret string
		return ret
	}
	return *o.IsAutoAddMargin
}

// GetIsAutoAddMarginOk returns a tuple with the IsAutoAddMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSymbolConfigV1RespItem) GetIsAutoAddMarginOk() (*string, bool) {
	if o == nil || IsNil(o.IsAutoAddMargin) {
		return nil, false
	}
	return o.IsAutoAddMargin, true
}

// HasIsAutoAddMargin returns a boolean if a field has been set.
func (o *GetSymbolConfigV1RespItem) HasIsAutoAddMargin() bool {
	if o != nil && !IsNil(o.IsAutoAddMargin) {
		return true
	}

	return false
}

// SetIsAutoAddMargin gets a reference to the given string and assigns it to the IsAutoAddMargin field.
func (o *GetSymbolConfigV1RespItem) SetIsAutoAddMargin(v string) {
	o.IsAutoAddMargin = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *GetSymbolConfigV1RespItem) GetLeverage() int32 {
	if o == nil || IsNil(o.Leverage) {
		var ret int32
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSymbolConfigV1RespItem) GetLeverageOk() (*int32, bool) {
	if o == nil || IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *GetSymbolConfigV1RespItem) HasLeverage() bool {
	if o != nil && !IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given int32 and assigns it to the Leverage field.
func (o *GetSymbolConfigV1RespItem) SetLeverage(v int32) {
	o.Leverage = &v
}

// GetMarginType returns the MarginType field value if set, zero value otherwise.
func (o *GetSymbolConfigV1RespItem) GetMarginType() string {
	if o == nil || IsNil(o.MarginType) {
		var ret string
		return ret
	}
	return *o.MarginType
}

// GetMarginTypeOk returns a tuple with the MarginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSymbolConfigV1RespItem) GetMarginTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MarginType) {
		return nil, false
	}
	return o.MarginType, true
}

// HasMarginType returns a boolean if a field has been set.
func (o *GetSymbolConfigV1RespItem) HasMarginType() bool {
	if o != nil && !IsNil(o.MarginType) {
		return true
	}

	return false
}

// SetMarginType gets a reference to the given string and assigns it to the MarginType field.
func (o *GetSymbolConfigV1RespItem) SetMarginType(v string) {
	o.MarginType = &v
}

// GetMaxNotionalValue returns the MaxNotionalValue field value if set, zero value otherwise.
func (o *GetSymbolConfigV1RespItem) GetMaxNotionalValue() string {
	if o == nil || IsNil(o.MaxNotionalValue) {
		var ret string
		return ret
	}
	return *o.MaxNotionalValue
}

// GetMaxNotionalValueOk returns a tuple with the MaxNotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSymbolConfigV1RespItem) GetMaxNotionalValueOk() (*string, bool) {
	if o == nil || IsNil(o.MaxNotionalValue) {
		return nil, false
	}
	return o.MaxNotionalValue, true
}

// HasMaxNotionalValue returns a boolean if a field has been set.
func (o *GetSymbolConfigV1RespItem) HasMaxNotionalValue() bool {
	if o != nil && !IsNil(o.MaxNotionalValue) {
		return true
	}

	return false
}

// SetMaxNotionalValue gets a reference to the given string and assigns it to the MaxNotionalValue field.
func (o *GetSymbolConfigV1RespItem) SetMaxNotionalValue(v string) {
	o.MaxNotionalValue = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetSymbolConfigV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSymbolConfigV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetSymbolConfigV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetSymbolConfigV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

func (o GetSymbolConfigV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSymbolConfigV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAutoAddMargin) {
		toSerialize["isAutoAddMargin"] = o.IsAutoAddMargin
	}
	if !IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !IsNil(o.MarginType) {
		toSerialize["marginType"] = o.MarginType
	}
	if !IsNil(o.MaxNotionalValue) {
		toSerialize["maxNotionalValue"] = o.MaxNotionalValue
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableGetSymbolConfigV1RespItem struct {
	value *GetSymbolConfigV1RespItem
	isSet bool
}

func (v NullableGetSymbolConfigV1RespItem) Get() *GetSymbolConfigV1RespItem {
	return v.value
}

func (v *NullableGetSymbolConfigV1RespItem) Set(val *GetSymbolConfigV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSymbolConfigV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSymbolConfigV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSymbolConfigV1RespItem(val *GetSymbolConfigV1RespItem) *NullableGetSymbolConfigV1RespItem {
	return &NullableGetSymbolConfigV1RespItem{value: val, isSet: true}
}

func (v NullableGetSymbolConfigV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSymbolConfigV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


