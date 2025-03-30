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

// checks if the MarginGetMarginCrossMarginCollateralRatioV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginCrossMarginCollateralRatioV1RespItem{}

// MarginGetMarginCrossMarginCollateralRatioV1RespItem struct for MarginGetMarginCrossMarginCollateralRatioV1RespItem
type MarginGetMarginCrossMarginCollateralRatioV1RespItem struct {
	AssetNames []string `json:"assetNames,omitempty"`
	Collaterals []MarginGetMarginCrossMarginCollateralRatioV1RespItemCollateralsInner `json:"collaterals,omitempty"`
}

// NewMarginGetMarginCrossMarginCollateralRatioV1RespItem instantiates a new MarginGetMarginCrossMarginCollateralRatioV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginCrossMarginCollateralRatioV1RespItem() *MarginGetMarginCrossMarginCollateralRatioV1RespItem {
	this := MarginGetMarginCrossMarginCollateralRatioV1RespItem{}
	return &this
}

// NewMarginGetMarginCrossMarginCollateralRatioV1RespItemWithDefaults instantiates a new MarginGetMarginCrossMarginCollateralRatioV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginCrossMarginCollateralRatioV1RespItemWithDefaults() *MarginGetMarginCrossMarginCollateralRatioV1RespItem {
	this := MarginGetMarginCrossMarginCollateralRatioV1RespItem{}
	return &this
}

// GetAssetNames returns the AssetNames field value if set, zero value otherwise.
func (o *MarginGetMarginCrossMarginCollateralRatioV1RespItem) GetAssetNames() []string {
	if o == nil || IsNil(o.AssetNames) {
		var ret []string
		return ret
	}
	return o.AssetNames
}

// GetAssetNamesOk returns a tuple with the AssetNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginCrossMarginCollateralRatioV1RespItem) GetAssetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AssetNames) {
		return nil, false
	}
	return o.AssetNames, true
}

// HasAssetNames returns a boolean if a field has been set.
func (o *MarginGetMarginCrossMarginCollateralRatioV1RespItem) HasAssetNames() bool {
	if o != nil && !IsNil(o.AssetNames) {
		return true
	}

	return false
}

// SetAssetNames gets a reference to the given []string and assigns it to the AssetNames field.
func (o *MarginGetMarginCrossMarginCollateralRatioV1RespItem) SetAssetNames(v []string) {
	o.AssetNames = v
}

// GetCollaterals returns the Collaterals field value if set, zero value otherwise.
func (o *MarginGetMarginCrossMarginCollateralRatioV1RespItem) GetCollaterals() []MarginGetMarginCrossMarginCollateralRatioV1RespItemCollateralsInner {
	if o == nil || IsNil(o.Collaterals) {
		var ret []MarginGetMarginCrossMarginCollateralRatioV1RespItemCollateralsInner
		return ret
	}
	return o.Collaterals
}

// GetCollateralsOk returns a tuple with the Collaterals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginCrossMarginCollateralRatioV1RespItem) GetCollateralsOk() ([]MarginGetMarginCrossMarginCollateralRatioV1RespItemCollateralsInner, bool) {
	if o == nil || IsNil(o.Collaterals) {
		return nil, false
	}
	return o.Collaterals, true
}

// HasCollaterals returns a boolean if a field has been set.
func (o *MarginGetMarginCrossMarginCollateralRatioV1RespItem) HasCollaterals() bool {
	if o != nil && !IsNil(o.Collaterals) {
		return true
	}

	return false
}

// SetCollaterals gets a reference to the given []MarginGetMarginCrossMarginCollateralRatioV1RespItemCollateralsInner and assigns it to the Collaterals field.
func (o *MarginGetMarginCrossMarginCollateralRatioV1RespItem) SetCollaterals(v []MarginGetMarginCrossMarginCollateralRatioV1RespItemCollateralsInner) {
	o.Collaterals = v
}

func (o MarginGetMarginCrossMarginCollateralRatioV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginCrossMarginCollateralRatioV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetNames) {
		toSerialize["assetNames"] = o.AssetNames
	}
	if !IsNil(o.Collaterals) {
		toSerialize["collaterals"] = o.Collaterals
	}
	return toSerialize, nil
}

type NullableMarginGetMarginCrossMarginCollateralRatioV1RespItem struct {
	value *MarginGetMarginCrossMarginCollateralRatioV1RespItem
	isSet bool
}

func (v NullableMarginGetMarginCrossMarginCollateralRatioV1RespItem) Get() *MarginGetMarginCrossMarginCollateralRatioV1RespItem {
	return v.value
}

func (v *NullableMarginGetMarginCrossMarginCollateralRatioV1RespItem) Set(val *MarginGetMarginCrossMarginCollateralRatioV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginCrossMarginCollateralRatioV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginCrossMarginCollateralRatioV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginCrossMarginCollateralRatioV1RespItem(val *MarginGetMarginCrossMarginCollateralRatioV1RespItem) *NullableMarginGetMarginCrossMarginCollateralRatioV1RespItem {
	return &NullableMarginGetMarginCrossMarginCollateralRatioV1RespItem{value: val, isSet: true}
}

func (v NullableMarginGetMarginCrossMarginCollateralRatioV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginCrossMarginCollateralRatioV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


