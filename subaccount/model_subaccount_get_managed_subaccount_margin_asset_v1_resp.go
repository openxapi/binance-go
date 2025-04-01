/*
Binance Sub Account API

OpenAPI specification for Binance exchange - Subaccount API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package subaccount

import (
	"encoding/json"
)

// checks if the SubaccountGetManagedSubaccountMarginAssetV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountGetManagedSubaccountMarginAssetV1Resp{}

// SubaccountGetManagedSubaccountMarginAssetV1Resp struct for SubaccountGetManagedSubaccountMarginAssetV1Resp
type SubaccountGetManagedSubaccountMarginAssetV1Resp struct {
	MarginLevel *string `json:"marginLevel,omitempty"`
	TotalAssetOfBtc *string `json:"totalAssetOfBtc,omitempty"`
	TotalLiabilityOfBtc *string `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc *string `json:"totalNetAssetOfBtc,omitempty"`
	UserAssets []SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner `json:"userAssets,omitempty"`
}

// NewSubaccountGetManagedSubaccountMarginAssetV1Resp instantiates a new SubaccountGetManagedSubaccountMarginAssetV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountGetManagedSubaccountMarginAssetV1Resp() *SubaccountGetManagedSubaccountMarginAssetV1Resp {
	this := SubaccountGetManagedSubaccountMarginAssetV1Resp{}
	return &this
}

// NewSubaccountGetManagedSubaccountMarginAssetV1RespWithDefaults instantiates a new SubaccountGetManagedSubaccountMarginAssetV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountGetManagedSubaccountMarginAssetV1RespWithDefaults() *SubaccountGetManagedSubaccountMarginAssetV1Resp {
	this := SubaccountGetManagedSubaccountMarginAssetV1Resp{}
	return &this
}

// GetMarginLevel returns the MarginLevel field value if set, zero value otherwise.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetMarginLevel() string {
	if o == nil || IsNil(o.MarginLevel) {
		var ret string
		return ret
	}
	return *o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetMarginLevelOk() (*string, bool) {
	if o == nil || IsNil(o.MarginLevel) {
		return nil, false
	}
	return o.MarginLevel, true
}

// HasMarginLevel returns a boolean if a field has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasMarginLevel() bool {
	if o != nil && !IsNil(o.MarginLevel) {
		return true
	}

	return false
}

// SetMarginLevel gets a reference to the given string and assigns it to the MarginLevel field.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetMarginLevel(v string) {
	o.MarginLevel = &v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalAssetOfBtc() string {
	if o == nil || IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasTotalAssetOfBtc() bool {
	if o != nil && !IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalLiabilityOfBtc() string {
	if o == nil || IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasTotalLiabilityOfBtc() bool {
	if o != nil && !IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalNetAssetOfBtc() string {
	if o == nil || IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasTotalNetAssetOfBtc() bool {
	if o != nil && !IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

// GetUserAssets returns the UserAssets field value if set, zero value otherwise.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetUserAssets() []SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner {
	if o == nil || IsNil(o.UserAssets) {
		var ret []SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner
		return ret
	}
	return o.UserAssets
}

// GetUserAssetsOk returns a tuple with the UserAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetUserAssetsOk() ([]SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner, bool) {
	if o == nil || IsNil(o.UserAssets) {
		return nil, false
	}
	return o.UserAssets, true
}

// HasUserAssets returns a boolean if a field has been set.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasUserAssets() bool {
	if o != nil && !IsNil(o.UserAssets) {
		return true
	}

	return false
}

// SetUserAssets gets a reference to the given []SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner and assigns it to the UserAssets field.
func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetUserAssets(v []SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner) {
	o.UserAssets = v
}

func (o SubaccountGetManagedSubaccountMarginAssetV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountGetManagedSubaccountMarginAssetV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarginLevel) {
		toSerialize["marginLevel"] = o.MarginLevel
	}
	if !IsNil(o.TotalAssetOfBtc) {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	if !IsNil(o.TotalLiabilityOfBtc) {
		toSerialize["totalLiabilityOfBtc"] = o.TotalLiabilityOfBtc
	}
	if !IsNil(o.TotalNetAssetOfBtc) {
		toSerialize["totalNetAssetOfBtc"] = o.TotalNetAssetOfBtc
	}
	if !IsNil(o.UserAssets) {
		toSerialize["userAssets"] = o.UserAssets
	}
	return toSerialize, nil
}

type NullableSubaccountGetManagedSubaccountMarginAssetV1Resp struct {
	value *SubaccountGetManagedSubaccountMarginAssetV1Resp
	isSet bool
}

func (v NullableSubaccountGetManagedSubaccountMarginAssetV1Resp) Get() *SubaccountGetManagedSubaccountMarginAssetV1Resp {
	return v.value
}

func (v *NullableSubaccountGetManagedSubaccountMarginAssetV1Resp) Set(val *SubaccountGetManagedSubaccountMarginAssetV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountGetManagedSubaccountMarginAssetV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountGetManagedSubaccountMarginAssetV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountGetManagedSubaccountMarginAssetV1Resp(val *SubaccountGetManagedSubaccountMarginAssetV1Resp) *NullableSubaccountGetManagedSubaccountMarginAssetV1Resp {
	return &NullableSubaccountGetManagedSubaccountMarginAssetV1Resp{value: val, isSet: true}
}

func (v NullableSubaccountGetManagedSubaccountMarginAssetV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountGetManagedSubaccountMarginAssetV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


