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

// checks if the PmarginCreateMarginRepayDebtV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginCreateMarginRepayDebtV1Resp{}

// PmarginCreateMarginRepayDebtV1Resp struct for PmarginCreateMarginRepayDebtV1Resp
type PmarginCreateMarginRepayDebtV1Resp struct {
	Amount *string `json:"amount,omitempty"`
	Asset *string `json:"asset,omitempty"`
	SpecifyRepayAssets []string `json:"specifyRepayAssets,omitempty"`
	Success *bool `json:"success,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewPmarginCreateMarginRepayDebtV1Resp instantiates a new PmarginCreateMarginRepayDebtV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginCreateMarginRepayDebtV1Resp() *PmarginCreateMarginRepayDebtV1Resp {
	this := PmarginCreateMarginRepayDebtV1Resp{}
	return &this
}

// NewPmarginCreateMarginRepayDebtV1RespWithDefaults instantiates a new PmarginCreateMarginRepayDebtV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginCreateMarginRepayDebtV1RespWithDefaults() *PmarginCreateMarginRepayDebtV1Resp {
	this := PmarginCreateMarginRepayDebtV1Resp{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *PmarginCreateMarginRepayDebtV1Resp) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *PmarginCreateMarginRepayDebtV1Resp) SetAsset(v string) {
	o.Asset = &v
}

// GetSpecifyRepayAssets returns the SpecifyRepayAssets field value if set, zero value otherwise.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetSpecifyRepayAssets() []string {
	if o == nil || IsNil(o.SpecifyRepayAssets) {
		var ret []string
		return ret
	}
	return o.SpecifyRepayAssets
}

// GetSpecifyRepayAssetsOk returns a tuple with the SpecifyRepayAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetSpecifyRepayAssetsOk() ([]string, bool) {
	if o == nil || IsNil(o.SpecifyRepayAssets) {
		return nil, false
	}
	return o.SpecifyRepayAssets, true
}

// HasSpecifyRepayAssets returns a boolean if a field has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) HasSpecifyRepayAssets() bool {
	if o != nil && !IsNil(o.SpecifyRepayAssets) {
		return true
	}

	return false
}

// SetSpecifyRepayAssets gets a reference to the given []string and assigns it to the SpecifyRepayAssets field.
func (o *PmarginCreateMarginRepayDebtV1Resp) SetSpecifyRepayAssets(v []string) {
	o.SpecifyRepayAssets = v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *PmarginCreateMarginRepayDebtV1Resp) SetSuccess(v bool) {
	o.Success = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PmarginCreateMarginRepayDebtV1Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *PmarginCreateMarginRepayDebtV1Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o PmarginCreateMarginRepayDebtV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginCreateMarginRepayDebtV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.SpecifyRepayAssets) {
		toSerialize["specifyRepayAssets"] = o.SpecifyRepayAssets
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullablePmarginCreateMarginRepayDebtV1Resp struct {
	value *PmarginCreateMarginRepayDebtV1Resp
	isSet bool
}

func (v NullablePmarginCreateMarginRepayDebtV1Resp) Get() *PmarginCreateMarginRepayDebtV1Resp {
	return v.value
}

func (v *NullablePmarginCreateMarginRepayDebtV1Resp) Set(val *PmarginCreateMarginRepayDebtV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginCreateMarginRepayDebtV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginCreateMarginRepayDebtV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginCreateMarginRepayDebtV1Resp(val *PmarginCreateMarginRepayDebtV1Resp) *NullablePmarginCreateMarginRepayDebtV1Resp {
	return &NullablePmarginCreateMarginRepayDebtV1Resp{value: val, isSet: true}
}

func (v NullablePmarginCreateMarginRepayDebtV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginCreateMarginRepayDebtV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


