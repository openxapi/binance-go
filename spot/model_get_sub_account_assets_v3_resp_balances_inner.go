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

// checks if the GetSubAccountAssetsV3RespBalancesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubAccountAssetsV3RespBalancesInner{}

// GetSubAccountAssetsV3RespBalancesInner struct for GetSubAccountAssetsV3RespBalancesInner
type GetSubAccountAssetsV3RespBalancesInner struct {
	Asset *string `json:"asset,omitempty"`
	Free *int32 `json:"free,omitempty"`
	Freeze *int32 `json:"freeze,omitempty"`
	Locked *int32 `json:"locked,omitempty"`
	Withdrawing *int32 `json:"withdrawing,omitempty"`
}

// NewGetSubAccountAssetsV3RespBalancesInner instantiates a new GetSubAccountAssetsV3RespBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountAssetsV3RespBalancesInner() *GetSubAccountAssetsV3RespBalancesInner {
	this := GetSubAccountAssetsV3RespBalancesInner{}
	return &this
}

// NewGetSubAccountAssetsV3RespBalancesInnerWithDefaults instantiates a new GetSubAccountAssetsV3RespBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountAssetsV3RespBalancesInnerWithDefaults() *GetSubAccountAssetsV3RespBalancesInner {
	this := GetSubAccountAssetsV3RespBalancesInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSubAccountAssetsV3RespBalancesInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetFree() int32 {
	if o == nil || IsNil(o.Free) {
		var ret int32
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetFreeOk() (*int32, bool) {
	if o == nil || IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) HasFree() bool {
	if o != nil && !IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given int32 and assigns it to the Free field.
func (o *GetSubAccountAssetsV3RespBalancesInner) SetFree(v int32) {
	o.Free = &v
}

// GetFreeze returns the Freeze field value if set, zero value otherwise.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetFreeze() int32 {
	if o == nil || IsNil(o.Freeze) {
		var ret int32
		return ret
	}
	return *o.Freeze
}

// GetFreezeOk returns a tuple with the Freeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetFreezeOk() (*int32, bool) {
	if o == nil || IsNil(o.Freeze) {
		return nil, false
	}
	return o.Freeze, true
}

// HasFreeze returns a boolean if a field has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) HasFreeze() bool {
	if o != nil && !IsNil(o.Freeze) {
		return true
	}

	return false
}

// SetFreeze gets a reference to the given int32 and assigns it to the Freeze field.
func (o *GetSubAccountAssetsV3RespBalancesInner) SetFreeze(v int32) {
	o.Freeze = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetLocked() int32 {
	if o == nil || IsNil(o.Locked) {
		var ret int32
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetLockedOk() (*int32, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given int32 and assigns it to the Locked field.
func (o *GetSubAccountAssetsV3RespBalancesInner) SetLocked(v int32) {
	o.Locked = &v
}

// GetWithdrawing returns the Withdrawing field value if set, zero value otherwise.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetWithdrawing() int32 {
	if o == nil || IsNil(o.Withdrawing) {
		var ret int32
		return ret
	}
	return *o.Withdrawing
}

// GetWithdrawingOk returns a tuple with the Withdrawing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) GetWithdrawingOk() (*int32, bool) {
	if o == nil || IsNil(o.Withdrawing) {
		return nil, false
	}
	return o.Withdrawing, true
}

// HasWithdrawing returns a boolean if a field has been set.
func (o *GetSubAccountAssetsV3RespBalancesInner) HasWithdrawing() bool {
	if o != nil && !IsNil(o.Withdrawing) {
		return true
	}

	return false
}

// SetWithdrawing gets a reference to the given int32 and assigns it to the Withdrawing field.
func (o *GetSubAccountAssetsV3RespBalancesInner) SetWithdrawing(v int32) {
	o.Withdrawing = &v
}

func (o GetSubAccountAssetsV3RespBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountAssetsV3RespBalancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !IsNil(o.Freeze) {
		toSerialize["freeze"] = o.Freeze
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.Withdrawing) {
		toSerialize["withdrawing"] = o.Withdrawing
	}
	return toSerialize, nil
}

type NullableGetSubAccountAssetsV3RespBalancesInner struct {
	value *GetSubAccountAssetsV3RespBalancesInner
	isSet bool
}

func (v NullableGetSubAccountAssetsV3RespBalancesInner) Get() *GetSubAccountAssetsV3RespBalancesInner {
	return v.value
}

func (v *NullableGetSubAccountAssetsV3RespBalancesInner) Set(val *GetSubAccountAssetsV3RespBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountAssetsV3RespBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountAssetsV3RespBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountAssetsV3RespBalancesInner(val *GetSubAccountAssetsV3RespBalancesInner) *NullableGetSubAccountAssetsV3RespBalancesInner {
	return &NullableGetSubAccountAssetsV3RespBalancesInner{value: val, isSet: true}
}

func (v NullableGetSubAccountAssetsV3RespBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountAssetsV3RespBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


