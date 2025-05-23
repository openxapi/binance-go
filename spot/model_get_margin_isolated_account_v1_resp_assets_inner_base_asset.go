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

// checks if the GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset{}

// GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset struct for GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset
type GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset struct {
	Asset *string `json:"asset,omitempty"`
	BorrowEnabled *bool `json:"borrowEnabled,omitempty"`
	Borrowed *string `json:"borrowed,omitempty"`
	Free *string `json:"free,omitempty"`
	Interest *string `json:"interest,omitempty"`
	Locked *string `json:"locked,omitempty"`
	NetAsset *string `json:"netAsset,omitempty"`
	NetAssetOfBtc *string `json:"netAssetOfBtc,omitempty"`
	RepayEnabled *bool `json:"repayEnabled,omitempty"`
	TotalAsset *string `json:"totalAsset,omitempty"`
}

// NewGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset instantiates a new GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset() *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset {
	this := GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset{}
	return &this
}

// NewGetMarginIsolatedAccountV1RespAssetsInnerBaseAssetWithDefaults instantiates a new GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginIsolatedAccountV1RespAssetsInnerBaseAssetWithDefaults() *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset {
	this := GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetAsset(v string) {
	o.Asset = &v
}

// GetBorrowEnabled returns the BorrowEnabled field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetBorrowEnabled() bool {
	if o == nil || IsNil(o.BorrowEnabled) {
		var ret bool
		return ret
	}
	return *o.BorrowEnabled
}

// GetBorrowEnabledOk returns a tuple with the BorrowEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetBorrowEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BorrowEnabled) {
		return nil, false
	}
	return o.BorrowEnabled, true
}

// HasBorrowEnabled returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasBorrowEnabled() bool {
	if o != nil && !IsNil(o.BorrowEnabled) {
		return true
	}

	return false
}

// SetBorrowEnabled gets a reference to the given bool and assigns it to the BorrowEnabled field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetBorrowEnabled(v bool) {
	o.BorrowEnabled = &v
}

// GetBorrowed returns the Borrowed field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetBorrowed() string {
	if o == nil || IsNil(o.Borrowed) {
		var ret string
		return ret
	}
	return *o.Borrowed
}

// GetBorrowedOk returns a tuple with the Borrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetBorrowedOk() (*string, bool) {
	if o == nil || IsNil(o.Borrowed) {
		return nil, false
	}
	return o.Borrowed, true
}

// HasBorrowed returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasBorrowed() bool {
	if o != nil && !IsNil(o.Borrowed) {
		return true
	}

	return false
}

// SetBorrowed gets a reference to the given string and assigns it to the Borrowed field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetBorrowed(v string) {
	o.Borrowed = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetFree() string {
	if o == nil || IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetFreeOk() (*string, bool) {
	if o == nil || IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasFree() bool {
	if o != nil && !IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetFree(v string) {
	o.Free = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetInterest() string {
	if o == nil || IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetInterestOk() (*string, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetInterest(v string) {
	o.Interest = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetLocked() string {
	if o == nil || IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetLockedOk() (*string, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetLocked(v string) {
	o.Locked = &v
}

// GetNetAsset returns the NetAsset field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetNetAsset() string {
	if o == nil || IsNil(o.NetAsset) {
		var ret string
		return ret
	}
	return *o.NetAsset
}

// GetNetAssetOk returns a tuple with the NetAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetNetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.NetAsset) {
		return nil, false
	}
	return o.NetAsset, true
}

// HasNetAsset returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasNetAsset() bool {
	if o != nil && !IsNil(o.NetAsset) {
		return true
	}

	return false
}

// SetNetAsset gets a reference to the given string and assigns it to the NetAsset field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetNetAsset(v string) {
	o.NetAsset = &v
}

// GetNetAssetOfBtc returns the NetAssetOfBtc field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetNetAssetOfBtc() string {
	if o == nil || IsNil(o.NetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.NetAssetOfBtc
}

// GetNetAssetOfBtcOk returns a tuple with the NetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetNetAssetOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.NetAssetOfBtc) {
		return nil, false
	}
	return o.NetAssetOfBtc, true
}

// HasNetAssetOfBtc returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasNetAssetOfBtc() bool {
	if o != nil && !IsNil(o.NetAssetOfBtc) {
		return true
	}

	return false
}

// SetNetAssetOfBtc gets a reference to the given string and assigns it to the NetAssetOfBtc field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetNetAssetOfBtc(v string) {
	o.NetAssetOfBtc = &v
}

// GetRepayEnabled returns the RepayEnabled field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetRepayEnabled() bool {
	if o == nil || IsNil(o.RepayEnabled) {
		var ret bool
		return ret
	}
	return *o.RepayEnabled
}

// GetRepayEnabledOk returns a tuple with the RepayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetRepayEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RepayEnabled) {
		return nil, false
	}
	return o.RepayEnabled, true
}

// HasRepayEnabled returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasRepayEnabled() bool {
	if o != nil && !IsNil(o.RepayEnabled) {
		return true
	}

	return false
}

// SetRepayEnabled gets a reference to the given bool and assigns it to the RepayEnabled field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetRepayEnabled(v bool) {
	o.RepayEnabled = &v
}

// GetTotalAsset returns the TotalAsset field value if set, zero value otherwise.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetTotalAsset() string {
	if o == nil || IsNil(o.TotalAsset) {
		var ret string
		return ret
	}
	return *o.TotalAsset
}

// GetTotalAssetOk returns a tuple with the TotalAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) GetTotalAssetOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAsset) {
		return nil, false
	}
	return o.TotalAsset, true
}

// HasTotalAsset returns a boolean if a field has been set.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) HasTotalAsset() bool {
	if o != nil && !IsNil(o.TotalAsset) {
		return true
	}

	return false
}

// SetTotalAsset gets a reference to the given string and assigns it to the TotalAsset field.
func (o *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) SetTotalAsset(v string) {
	o.TotalAsset = &v
}

func (o GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.BorrowEnabled) {
		toSerialize["borrowEnabled"] = o.BorrowEnabled
	}
	if !IsNil(o.Borrowed) {
		toSerialize["borrowed"] = o.Borrowed
	}
	if !IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.NetAsset) {
		toSerialize["netAsset"] = o.NetAsset
	}
	if !IsNil(o.NetAssetOfBtc) {
		toSerialize["netAssetOfBtc"] = o.NetAssetOfBtc
	}
	if !IsNil(o.RepayEnabled) {
		toSerialize["repayEnabled"] = o.RepayEnabled
	}
	if !IsNil(o.TotalAsset) {
		toSerialize["totalAsset"] = o.TotalAsset
	}
	return toSerialize, nil
}

type NullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset struct {
	value *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset
	isSet bool
}

func (v NullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) Get() *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset {
	return v.value
}

func (v *NullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) Set(val *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset(val *GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) *NullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset {
	return &NullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset{value: val, isSet: true}
}

func (v NullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginIsolatedAccountV1RespAssetsInnerBaseAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


