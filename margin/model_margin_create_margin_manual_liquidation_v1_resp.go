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

// checks if the MarginCreateMarginManualLiquidationV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginCreateMarginManualLiquidationV1Resp{}

// MarginCreateMarginManualLiquidationV1Resp struct for MarginCreateMarginManualLiquidationV1Resp
type MarginCreateMarginManualLiquidationV1Resp struct {
	Asset *string `json:"asset,omitempty"`
	Interest *string `json:"interest,omitempty"`
	LiabilityAsset *string `json:"liabilityAsset,omitempty"`
	LiabilityQty *float32 `json:"liabilityQty,omitempty"`
	Principal *string `json:"principal,omitempty"`
}

// NewMarginCreateMarginManualLiquidationV1Resp instantiates a new MarginCreateMarginManualLiquidationV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCreateMarginManualLiquidationV1Resp() *MarginCreateMarginManualLiquidationV1Resp {
	this := MarginCreateMarginManualLiquidationV1Resp{}
	return &this
}

// NewMarginCreateMarginManualLiquidationV1RespWithDefaults instantiates a new MarginCreateMarginManualLiquidationV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCreateMarginManualLiquidationV1RespWithDefaults() *MarginCreateMarginManualLiquidationV1Resp {
	this := MarginCreateMarginManualLiquidationV1Resp{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *MarginCreateMarginManualLiquidationV1Resp) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetInterest() string {
	if o == nil || IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetInterestOk() (*string, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *MarginCreateMarginManualLiquidationV1Resp) SetInterest(v string) {
	o.Interest = &v
}

// GetLiabilityAsset returns the LiabilityAsset field value if set, zero value otherwise.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetLiabilityAsset() string {
	if o == nil || IsNil(o.LiabilityAsset) {
		var ret string
		return ret
	}
	return *o.LiabilityAsset
}

// GetLiabilityAssetOk returns a tuple with the LiabilityAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetLiabilityAssetOk() (*string, bool) {
	if o == nil || IsNil(o.LiabilityAsset) {
		return nil, false
	}
	return o.LiabilityAsset, true
}

// HasLiabilityAsset returns a boolean if a field has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) HasLiabilityAsset() bool {
	if o != nil && !IsNil(o.LiabilityAsset) {
		return true
	}

	return false
}

// SetLiabilityAsset gets a reference to the given string and assigns it to the LiabilityAsset field.
func (o *MarginCreateMarginManualLiquidationV1Resp) SetLiabilityAsset(v string) {
	o.LiabilityAsset = &v
}

// GetLiabilityQty returns the LiabilityQty field value if set, zero value otherwise.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetLiabilityQty() float32 {
	if o == nil || IsNil(o.LiabilityQty) {
		var ret float32
		return ret
	}
	return *o.LiabilityQty
}

// GetLiabilityQtyOk returns a tuple with the LiabilityQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetLiabilityQtyOk() (*float32, bool) {
	if o == nil || IsNil(o.LiabilityQty) {
		return nil, false
	}
	return o.LiabilityQty, true
}

// HasLiabilityQty returns a boolean if a field has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) HasLiabilityQty() bool {
	if o != nil && !IsNil(o.LiabilityQty) {
		return true
	}

	return false
}

// SetLiabilityQty gets a reference to the given float32 and assigns it to the LiabilityQty field.
func (o *MarginCreateMarginManualLiquidationV1Resp) SetLiabilityQty(v float32) {
	o.LiabilityQty = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetPrincipal() string {
	if o == nil || IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) GetPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *MarginCreateMarginManualLiquidationV1Resp) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *MarginCreateMarginManualLiquidationV1Resp) SetPrincipal(v string) {
	o.Principal = &v
}

func (o MarginCreateMarginManualLiquidationV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginCreateMarginManualLiquidationV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.LiabilityAsset) {
		toSerialize["liabilityAsset"] = o.LiabilityAsset
	}
	if !IsNil(o.LiabilityQty) {
		toSerialize["liabilityQty"] = o.LiabilityQty
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	return toSerialize, nil
}

type NullableMarginCreateMarginManualLiquidationV1Resp struct {
	value *MarginCreateMarginManualLiquidationV1Resp
	isSet bool
}

func (v NullableMarginCreateMarginManualLiquidationV1Resp) Get() *MarginCreateMarginManualLiquidationV1Resp {
	return v.value
}

func (v *NullableMarginCreateMarginManualLiquidationV1Resp) Set(val *MarginCreateMarginManualLiquidationV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCreateMarginManualLiquidationV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCreateMarginManualLiquidationV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCreateMarginManualLiquidationV1Resp(val *MarginCreateMarginManualLiquidationV1Resp) *NullableMarginCreateMarginManualLiquidationV1Resp {
	return &NullableMarginCreateMarginManualLiquidationV1Resp{value: val, isSet: true}
}

func (v NullableMarginCreateMarginManualLiquidationV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCreateMarginManualLiquidationV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


