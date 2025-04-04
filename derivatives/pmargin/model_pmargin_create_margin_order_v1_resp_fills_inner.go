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

// checks if the PmarginCreateMarginOrderV1RespFillsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginCreateMarginOrderV1RespFillsInner{}

// PmarginCreateMarginOrderV1RespFillsInner struct for PmarginCreateMarginOrderV1RespFillsInner
type PmarginCreateMarginOrderV1RespFillsInner struct {
	Commission *string `json:"commission,omitempty"`
	CommissionAsset *string `json:"commissionAsset,omitempty"`
	Price *string `json:"price,omitempty"`
	Qty *string `json:"qty,omitempty"`
}

// NewPmarginCreateMarginOrderV1RespFillsInner instantiates a new PmarginCreateMarginOrderV1RespFillsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginCreateMarginOrderV1RespFillsInner() *PmarginCreateMarginOrderV1RespFillsInner {
	this := PmarginCreateMarginOrderV1RespFillsInner{}
	return &this
}

// NewPmarginCreateMarginOrderV1RespFillsInnerWithDefaults instantiates a new PmarginCreateMarginOrderV1RespFillsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginCreateMarginOrderV1RespFillsInnerWithDefaults() *PmarginCreateMarginOrderV1RespFillsInner {
	this := PmarginCreateMarginOrderV1RespFillsInner{}
	return &this
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *PmarginCreateMarginOrderV1RespFillsInner) GetCommission() string {
	if o == nil || IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateMarginOrderV1RespFillsInner) GetCommissionOk() (*string, bool) {
	if o == nil || IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *PmarginCreateMarginOrderV1RespFillsInner) HasCommission() bool {
	if o != nil && !IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *PmarginCreateMarginOrderV1RespFillsInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *PmarginCreateMarginOrderV1RespFillsInner) GetCommissionAsset() string {
	if o == nil || IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateMarginOrderV1RespFillsInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *PmarginCreateMarginOrderV1RespFillsInner) HasCommissionAsset() bool {
	if o != nil && !IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *PmarginCreateMarginOrderV1RespFillsInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PmarginCreateMarginOrderV1RespFillsInner) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateMarginOrderV1RespFillsInner) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PmarginCreateMarginOrderV1RespFillsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *PmarginCreateMarginOrderV1RespFillsInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *PmarginCreateMarginOrderV1RespFillsInner) GetQty() string {
	if o == nil || IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateMarginOrderV1RespFillsInner) GetQtyOk() (*string, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *PmarginCreateMarginOrderV1RespFillsInner) HasQty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *PmarginCreateMarginOrderV1RespFillsInner) SetQty(v string) {
	o.Qty = &v
}

func (o PmarginCreateMarginOrderV1RespFillsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginCreateMarginOrderV1RespFillsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !IsNil(o.CommissionAsset) {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	return toSerialize, nil
}

type NullablePmarginCreateMarginOrderV1RespFillsInner struct {
	value *PmarginCreateMarginOrderV1RespFillsInner
	isSet bool
}

func (v NullablePmarginCreateMarginOrderV1RespFillsInner) Get() *PmarginCreateMarginOrderV1RespFillsInner {
	return v.value
}

func (v *NullablePmarginCreateMarginOrderV1RespFillsInner) Set(val *PmarginCreateMarginOrderV1RespFillsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginCreateMarginOrderV1RespFillsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginCreateMarginOrderV1RespFillsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginCreateMarginOrderV1RespFillsInner(val *PmarginCreateMarginOrderV1RespFillsInner) *NullablePmarginCreateMarginOrderV1RespFillsInner {
	return &NullablePmarginCreateMarginOrderV1RespFillsInner{value: val, isSet: true}
}

func (v NullablePmarginCreateMarginOrderV1RespFillsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginCreateMarginOrderV1RespFillsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


