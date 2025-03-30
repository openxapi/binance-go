/*
Binance Margin API

OpenAPI specification for Binance cryptocurrency exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginGetMarginTradeCoeffV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginTradeCoeffV1Resp{}

// MarginGetMarginTradeCoeffV1Resp struct for MarginGetMarginTradeCoeffV1Resp
type MarginGetMarginTradeCoeffV1Resp struct {
	ForceLiquidationBar *string `json:"forceLiquidationBar,omitempty"`
	MarginCallBar *string `json:"marginCallBar,omitempty"`
	NormalBar *string `json:"normalBar,omitempty"`
}

// NewMarginGetMarginTradeCoeffV1Resp instantiates a new MarginGetMarginTradeCoeffV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginTradeCoeffV1Resp() *MarginGetMarginTradeCoeffV1Resp {
	this := MarginGetMarginTradeCoeffV1Resp{}
	return &this
}

// NewMarginGetMarginTradeCoeffV1RespWithDefaults instantiates a new MarginGetMarginTradeCoeffV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginTradeCoeffV1RespWithDefaults() *MarginGetMarginTradeCoeffV1Resp {
	this := MarginGetMarginTradeCoeffV1Resp{}
	return &this
}

// GetForceLiquidationBar returns the ForceLiquidationBar field value if set, zero value otherwise.
func (o *MarginGetMarginTradeCoeffV1Resp) GetForceLiquidationBar() string {
	if o == nil || IsNil(o.ForceLiquidationBar) {
		var ret string
		return ret
	}
	return *o.ForceLiquidationBar
}

// GetForceLiquidationBarOk returns a tuple with the ForceLiquidationBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTradeCoeffV1Resp) GetForceLiquidationBarOk() (*string, bool) {
	if o == nil || IsNil(o.ForceLiquidationBar) {
		return nil, false
	}
	return o.ForceLiquidationBar, true
}

// HasForceLiquidationBar returns a boolean if a field has been set.
func (o *MarginGetMarginTradeCoeffV1Resp) HasForceLiquidationBar() bool {
	if o != nil && !IsNil(o.ForceLiquidationBar) {
		return true
	}

	return false
}

// SetForceLiquidationBar gets a reference to the given string and assigns it to the ForceLiquidationBar field.
func (o *MarginGetMarginTradeCoeffV1Resp) SetForceLiquidationBar(v string) {
	o.ForceLiquidationBar = &v
}

// GetMarginCallBar returns the MarginCallBar field value if set, zero value otherwise.
func (o *MarginGetMarginTradeCoeffV1Resp) GetMarginCallBar() string {
	if o == nil || IsNil(o.MarginCallBar) {
		var ret string
		return ret
	}
	return *o.MarginCallBar
}

// GetMarginCallBarOk returns a tuple with the MarginCallBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTradeCoeffV1Resp) GetMarginCallBarOk() (*string, bool) {
	if o == nil || IsNil(o.MarginCallBar) {
		return nil, false
	}
	return o.MarginCallBar, true
}

// HasMarginCallBar returns a boolean if a field has been set.
func (o *MarginGetMarginTradeCoeffV1Resp) HasMarginCallBar() bool {
	if o != nil && !IsNil(o.MarginCallBar) {
		return true
	}

	return false
}

// SetMarginCallBar gets a reference to the given string and assigns it to the MarginCallBar field.
func (o *MarginGetMarginTradeCoeffV1Resp) SetMarginCallBar(v string) {
	o.MarginCallBar = &v
}

// GetNormalBar returns the NormalBar field value if set, zero value otherwise.
func (o *MarginGetMarginTradeCoeffV1Resp) GetNormalBar() string {
	if o == nil || IsNil(o.NormalBar) {
		var ret string
		return ret
	}
	return *o.NormalBar
}

// GetNormalBarOk returns a tuple with the NormalBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTradeCoeffV1Resp) GetNormalBarOk() (*string, bool) {
	if o == nil || IsNil(o.NormalBar) {
		return nil, false
	}
	return o.NormalBar, true
}

// HasNormalBar returns a boolean if a field has been set.
func (o *MarginGetMarginTradeCoeffV1Resp) HasNormalBar() bool {
	if o != nil && !IsNil(o.NormalBar) {
		return true
	}

	return false
}

// SetNormalBar gets a reference to the given string and assigns it to the NormalBar field.
func (o *MarginGetMarginTradeCoeffV1Resp) SetNormalBar(v string) {
	o.NormalBar = &v
}

func (o MarginGetMarginTradeCoeffV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginTradeCoeffV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForceLiquidationBar) {
		toSerialize["forceLiquidationBar"] = o.ForceLiquidationBar
	}
	if !IsNil(o.MarginCallBar) {
		toSerialize["marginCallBar"] = o.MarginCallBar
	}
	if !IsNil(o.NormalBar) {
		toSerialize["normalBar"] = o.NormalBar
	}
	return toSerialize, nil
}

type NullableMarginGetMarginTradeCoeffV1Resp struct {
	value *MarginGetMarginTradeCoeffV1Resp
	isSet bool
}

func (v NullableMarginGetMarginTradeCoeffV1Resp) Get() *MarginGetMarginTradeCoeffV1Resp {
	return v.value
}

func (v *NullableMarginGetMarginTradeCoeffV1Resp) Set(val *MarginGetMarginTradeCoeffV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginTradeCoeffV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginTradeCoeffV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginTradeCoeffV1Resp(val *MarginGetMarginTradeCoeffV1Resp) *NullableMarginGetMarginTradeCoeffV1Resp {
	return &NullableMarginGetMarginTradeCoeffV1Resp{value: val, isSet: true}
}

func (v NullableMarginGetMarginTradeCoeffV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginTradeCoeffV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


