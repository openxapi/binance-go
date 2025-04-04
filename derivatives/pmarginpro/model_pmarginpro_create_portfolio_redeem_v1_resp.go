/*
Binance Portfolio Margin Pro API

OpenAPI specification for Binance exchange - Pmarginpro API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmarginpro

import (
	"encoding/json"
)

// checks if the PmarginproCreatePortfolioRedeemV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginproCreatePortfolioRedeemV1Resp{}

// PmarginproCreatePortfolioRedeemV1Resp struct for PmarginproCreatePortfolioRedeemV1Resp
type PmarginproCreatePortfolioRedeemV1Resp struct {
	FromAsset *string `json:"fromAsset,omitempty"`
	FromAssetQty *float32 `json:"fromAssetQty,omitempty"`
	Rate *float32 `json:"rate,omitempty"`
	TargetAsset *string `json:"targetAsset,omitempty"`
	TargetAssetQty *float32 `json:"targetAssetQty,omitempty"`
}

// NewPmarginproCreatePortfolioRedeemV1Resp instantiates a new PmarginproCreatePortfolioRedeemV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginproCreatePortfolioRedeemV1Resp() *PmarginproCreatePortfolioRedeemV1Resp {
	this := PmarginproCreatePortfolioRedeemV1Resp{}
	return &this
}

// NewPmarginproCreatePortfolioRedeemV1RespWithDefaults instantiates a new PmarginproCreatePortfolioRedeemV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginproCreatePortfolioRedeemV1RespWithDefaults() *PmarginproCreatePortfolioRedeemV1Resp {
	this := PmarginproCreatePortfolioRedeemV1Resp{}
	return &this
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetFromAsset() string {
	if o == nil || IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetFromAssetOk() (*string, bool) {
	if o == nil || IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) HasFromAsset() bool {
	if o != nil && !IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *PmarginproCreatePortfolioRedeemV1Resp) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetFromAssetQty returns the FromAssetQty field value if set, zero value otherwise.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetFromAssetQty() float32 {
	if o == nil || IsNil(o.FromAssetQty) {
		var ret float32
		return ret
	}
	return *o.FromAssetQty
}

// GetFromAssetQtyOk returns a tuple with the FromAssetQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetFromAssetQtyOk() (*float32, bool) {
	if o == nil || IsNil(o.FromAssetQty) {
		return nil, false
	}
	return o.FromAssetQty, true
}

// HasFromAssetQty returns a boolean if a field has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) HasFromAssetQty() bool {
	if o != nil && !IsNil(o.FromAssetQty) {
		return true
	}

	return false
}

// SetFromAssetQty gets a reference to the given float32 and assigns it to the FromAssetQty field.
func (o *PmarginproCreatePortfolioRedeemV1Resp) SetFromAssetQty(v float32) {
	o.FromAssetQty = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *PmarginproCreatePortfolioRedeemV1Resp) SetRate(v float32) {
	o.Rate = &v
}

// GetTargetAsset returns the TargetAsset field value if set, zero value otherwise.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetTargetAsset() string {
	if o == nil || IsNil(o.TargetAsset) {
		var ret string
		return ret
	}
	return *o.TargetAsset
}

// GetTargetAssetOk returns a tuple with the TargetAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetTargetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAsset) {
		return nil, false
	}
	return o.TargetAsset, true
}

// HasTargetAsset returns a boolean if a field has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) HasTargetAsset() bool {
	if o != nil && !IsNil(o.TargetAsset) {
		return true
	}

	return false
}

// SetTargetAsset gets a reference to the given string and assigns it to the TargetAsset field.
func (o *PmarginproCreatePortfolioRedeemV1Resp) SetTargetAsset(v string) {
	o.TargetAsset = &v
}

// GetTargetAssetQty returns the TargetAssetQty field value if set, zero value otherwise.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetTargetAssetQty() float32 {
	if o == nil || IsNil(o.TargetAssetQty) {
		var ret float32
		return ret
	}
	return *o.TargetAssetQty
}

// GetTargetAssetQtyOk returns a tuple with the TargetAssetQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) GetTargetAssetQtyOk() (*float32, bool) {
	if o == nil || IsNil(o.TargetAssetQty) {
		return nil, false
	}
	return o.TargetAssetQty, true
}

// HasTargetAssetQty returns a boolean if a field has been set.
func (o *PmarginproCreatePortfolioRedeemV1Resp) HasTargetAssetQty() bool {
	if o != nil && !IsNil(o.TargetAssetQty) {
		return true
	}

	return false
}

// SetTargetAssetQty gets a reference to the given float32 and assigns it to the TargetAssetQty field.
func (o *PmarginproCreatePortfolioRedeemV1Resp) SetTargetAssetQty(v float32) {
	o.TargetAssetQty = &v
}

func (o PmarginproCreatePortfolioRedeemV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginproCreatePortfolioRedeemV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromAsset) {
		toSerialize["fromAsset"] = o.FromAsset
	}
	if !IsNil(o.FromAssetQty) {
		toSerialize["fromAssetQty"] = o.FromAssetQty
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.TargetAsset) {
		toSerialize["targetAsset"] = o.TargetAsset
	}
	if !IsNil(o.TargetAssetQty) {
		toSerialize["targetAssetQty"] = o.TargetAssetQty
	}
	return toSerialize, nil
}

type NullablePmarginproCreatePortfolioRedeemV1Resp struct {
	value *PmarginproCreatePortfolioRedeemV1Resp
	isSet bool
}

func (v NullablePmarginproCreatePortfolioRedeemV1Resp) Get() *PmarginproCreatePortfolioRedeemV1Resp {
	return v.value
}

func (v *NullablePmarginproCreatePortfolioRedeemV1Resp) Set(val *PmarginproCreatePortfolioRedeemV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginproCreatePortfolioRedeemV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginproCreatePortfolioRedeemV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginproCreatePortfolioRedeemV1Resp(val *PmarginproCreatePortfolioRedeemV1Resp) *NullablePmarginproCreatePortfolioRedeemV1Resp {
	return &NullablePmarginproCreatePortfolioRedeemV1Resp{value: val, isSet: true}
}

func (v NullablePmarginproCreatePortfolioRedeemV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginproCreatePortfolioRedeemV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


