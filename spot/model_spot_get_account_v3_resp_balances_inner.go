/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotGetAccountV3RespBalancesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotGetAccountV3RespBalancesInner{}

// SpotGetAccountV3RespBalancesInner struct for SpotGetAccountV3RespBalancesInner
type SpotGetAccountV3RespBalancesInner struct {
	Asset *string `json:"asset,omitempty"`
	Free *string `json:"free,omitempty"`
	Locked *string `json:"locked,omitempty"`
}

// NewSpotGetAccountV3RespBalancesInner instantiates a new SpotGetAccountV3RespBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotGetAccountV3RespBalancesInner() *SpotGetAccountV3RespBalancesInner {
	this := SpotGetAccountV3RespBalancesInner{}
	return &this
}

// NewSpotGetAccountV3RespBalancesInnerWithDefaults instantiates a new SpotGetAccountV3RespBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotGetAccountV3RespBalancesInnerWithDefaults() *SpotGetAccountV3RespBalancesInner {
	this := SpotGetAccountV3RespBalancesInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *SpotGetAccountV3RespBalancesInner) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3RespBalancesInner) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *SpotGetAccountV3RespBalancesInner) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *SpotGetAccountV3RespBalancesInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *SpotGetAccountV3RespBalancesInner) GetFree() string {
	if o == nil || IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3RespBalancesInner) GetFreeOk() (*string, bool) {
	if o == nil || IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *SpotGetAccountV3RespBalancesInner) HasFree() bool {
	if o != nil && !IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *SpotGetAccountV3RespBalancesInner) SetFree(v string) {
	o.Free = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *SpotGetAccountV3RespBalancesInner) GetLocked() string {
	if o == nil || IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3RespBalancesInner) GetLockedOk() (*string, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *SpotGetAccountV3RespBalancesInner) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *SpotGetAccountV3RespBalancesInner) SetLocked(v string) {
	o.Locked = &v
}

func (o SpotGetAccountV3RespBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotGetAccountV3RespBalancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	return toSerialize, nil
}

type NullableSpotGetAccountV3RespBalancesInner struct {
	value *SpotGetAccountV3RespBalancesInner
	isSet bool
}

func (v NullableSpotGetAccountV3RespBalancesInner) Get() *SpotGetAccountV3RespBalancesInner {
	return v.value
}

func (v *NullableSpotGetAccountV3RespBalancesInner) Set(val *SpotGetAccountV3RespBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotGetAccountV3RespBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotGetAccountV3RespBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotGetAccountV3RespBalancesInner(val *SpotGetAccountV3RespBalancesInner) *NullableSpotGetAccountV3RespBalancesInner {
	return &NullableSpotGetAccountV3RespBalancesInner{value: val, isSet: true}
}

func (v NullableSpotGetAccountV3RespBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotGetAccountV3RespBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


