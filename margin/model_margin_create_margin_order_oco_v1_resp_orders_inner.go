/*
Binance Margin Trading API

OpenAPI specification for Binance exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginCreateMarginOrderOcoV1RespOrdersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginCreateMarginOrderOcoV1RespOrdersInner{}

// MarginCreateMarginOrderOcoV1RespOrdersInner struct for MarginCreateMarginOrderOcoV1RespOrdersInner
type MarginCreateMarginOrderOcoV1RespOrdersInner struct {
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewMarginCreateMarginOrderOcoV1RespOrdersInner instantiates a new MarginCreateMarginOrderOcoV1RespOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCreateMarginOrderOcoV1RespOrdersInner() *MarginCreateMarginOrderOcoV1RespOrdersInner {
	this := MarginCreateMarginOrderOcoV1RespOrdersInner{}
	return &this
}

// NewMarginCreateMarginOrderOcoV1RespOrdersInnerWithDefaults instantiates a new MarginCreateMarginOrderOcoV1RespOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCreateMarginOrderOcoV1RespOrdersInnerWithDefaults() *MarginCreateMarginOrderOcoV1RespOrdersInner {
	this := MarginCreateMarginOrderOcoV1RespOrdersInner{}
	return &this
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginCreateMarginOrderOcoV1RespOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

func (o MarginCreateMarginOrderOcoV1RespOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginCreateMarginOrderOcoV1RespOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableMarginCreateMarginOrderOcoV1RespOrdersInner struct {
	value *MarginCreateMarginOrderOcoV1RespOrdersInner
	isSet bool
}

func (v NullableMarginCreateMarginOrderOcoV1RespOrdersInner) Get() *MarginCreateMarginOrderOcoV1RespOrdersInner {
	return v.value
}

func (v *NullableMarginCreateMarginOrderOcoV1RespOrdersInner) Set(val *MarginCreateMarginOrderOcoV1RespOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCreateMarginOrderOcoV1RespOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCreateMarginOrderOcoV1RespOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCreateMarginOrderOcoV1RespOrdersInner(val *MarginCreateMarginOrderOcoV1RespOrdersInner) *NullableMarginCreateMarginOrderOcoV1RespOrdersInner {
	return &NullableMarginCreateMarginOrderOcoV1RespOrdersInner{value: val, isSet: true}
}

func (v NullableMarginCreateMarginOrderOcoV1RespOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCreateMarginOrderOcoV1RespOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


