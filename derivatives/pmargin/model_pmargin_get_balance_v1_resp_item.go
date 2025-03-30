/*
Binance Pmargin API

OpenAPI specification for Binance cryptocurrency exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginGetBalanceV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetBalanceV1RespItem{}

// PmarginGetBalanceV1RespItem struct for PmarginGetBalanceV1RespItem
type PmarginGetBalanceV1RespItem struct {
	Asset *string `json:"asset,omitempty"`
	CmUnrealizedPNL *string `json:"cmUnrealizedPNL,omitempty"`
	CmWalletBalance *string `json:"cmWalletBalance,omitempty"`
	CrossMarginAsset *string `json:"crossMarginAsset,omitempty"`
	CrossMarginBorrowed *string `json:"crossMarginBorrowed,omitempty"`
	CrossMarginFree *string `json:"crossMarginFree,omitempty"`
	CrossMarginInterest *string `json:"crossMarginInterest,omitempty"`
	CrossMarginLocked *string `json:"crossMarginLocked,omitempty"`
	NegativeBalance *string `json:"negativeBalance,omitempty"`
	TotalWalletBalance *string `json:"totalWalletBalance,omitempty"`
	UmUnrealizedPNL *string `json:"umUnrealizedPNL,omitempty"`
	UmWalletBalance *string `json:"umWalletBalance,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewPmarginGetBalanceV1RespItem instantiates a new PmarginGetBalanceV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetBalanceV1RespItem() *PmarginGetBalanceV1RespItem {
	this := PmarginGetBalanceV1RespItem{}
	return &this
}

// NewPmarginGetBalanceV1RespItemWithDefaults instantiates a new PmarginGetBalanceV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetBalanceV1RespItemWithDefaults() *PmarginGetBalanceV1RespItem {
	this := PmarginGetBalanceV1RespItem{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *PmarginGetBalanceV1RespItem) SetAsset(v string) {
	o.Asset = &v
}

// GetCmUnrealizedPNL returns the CmUnrealizedPNL field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetCmUnrealizedPNL() string {
	if o == nil || IsNil(o.CmUnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.CmUnrealizedPNL
}

// GetCmUnrealizedPNLOk returns a tuple with the CmUnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetCmUnrealizedPNLOk() (*string, bool) {
	if o == nil || IsNil(o.CmUnrealizedPNL) {
		return nil, false
	}
	return o.CmUnrealizedPNL, true
}

// HasCmUnrealizedPNL returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasCmUnrealizedPNL() bool {
	if o != nil && !IsNil(o.CmUnrealizedPNL) {
		return true
	}

	return false
}

// SetCmUnrealizedPNL gets a reference to the given string and assigns it to the CmUnrealizedPNL field.
func (o *PmarginGetBalanceV1RespItem) SetCmUnrealizedPNL(v string) {
	o.CmUnrealizedPNL = &v
}

// GetCmWalletBalance returns the CmWalletBalance field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetCmWalletBalance() string {
	if o == nil || IsNil(o.CmWalletBalance) {
		var ret string
		return ret
	}
	return *o.CmWalletBalance
}

// GetCmWalletBalanceOk returns a tuple with the CmWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetCmWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.CmWalletBalance) {
		return nil, false
	}
	return o.CmWalletBalance, true
}

// HasCmWalletBalance returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasCmWalletBalance() bool {
	if o != nil && !IsNil(o.CmWalletBalance) {
		return true
	}

	return false
}

// SetCmWalletBalance gets a reference to the given string and assigns it to the CmWalletBalance field.
func (o *PmarginGetBalanceV1RespItem) SetCmWalletBalance(v string) {
	o.CmWalletBalance = &v
}

// GetCrossMarginAsset returns the CrossMarginAsset field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginAsset() string {
	if o == nil || IsNil(o.CrossMarginAsset) {
		var ret string
		return ret
	}
	return *o.CrossMarginAsset
}

// GetCrossMarginAssetOk returns a tuple with the CrossMarginAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginAssetOk() (*string, bool) {
	if o == nil || IsNil(o.CrossMarginAsset) {
		return nil, false
	}
	return o.CrossMarginAsset, true
}

// HasCrossMarginAsset returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasCrossMarginAsset() bool {
	if o != nil && !IsNil(o.CrossMarginAsset) {
		return true
	}

	return false
}

// SetCrossMarginAsset gets a reference to the given string and assigns it to the CrossMarginAsset field.
func (o *PmarginGetBalanceV1RespItem) SetCrossMarginAsset(v string) {
	o.CrossMarginAsset = &v
}

// GetCrossMarginBorrowed returns the CrossMarginBorrowed field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginBorrowed() string {
	if o == nil || IsNil(o.CrossMarginBorrowed) {
		var ret string
		return ret
	}
	return *o.CrossMarginBorrowed
}

// GetCrossMarginBorrowedOk returns a tuple with the CrossMarginBorrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginBorrowedOk() (*string, bool) {
	if o == nil || IsNil(o.CrossMarginBorrowed) {
		return nil, false
	}
	return o.CrossMarginBorrowed, true
}

// HasCrossMarginBorrowed returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasCrossMarginBorrowed() bool {
	if o != nil && !IsNil(o.CrossMarginBorrowed) {
		return true
	}

	return false
}

// SetCrossMarginBorrowed gets a reference to the given string and assigns it to the CrossMarginBorrowed field.
func (o *PmarginGetBalanceV1RespItem) SetCrossMarginBorrowed(v string) {
	o.CrossMarginBorrowed = &v
}

// GetCrossMarginFree returns the CrossMarginFree field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginFree() string {
	if o == nil || IsNil(o.CrossMarginFree) {
		var ret string
		return ret
	}
	return *o.CrossMarginFree
}

// GetCrossMarginFreeOk returns a tuple with the CrossMarginFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginFreeOk() (*string, bool) {
	if o == nil || IsNil(o.CrossMarginFree) {
		return nil, false
	}
	return o.CrossMarginFree, true
}

// HasCrossMarginFree returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasCrossMarginFree() bool {
	if o != nil && !IsNil(o.CrossMarginFree) {
		return true
	}

	return false
}

// SetCrossMarginFree gets a reference to the given string and assigns it to the CrossMarginFree field.
func (o *PmarginGetBalanceV1RespItem) SetCrossMarginFree(v string) {
	o.CrossMarginFree = &v
}

// GetCrossMarginInterest returns the CrossMarginInterest field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginInterest() string {
	if o == nil || IsNil(o.CrossMarginInterest) {
		var ret string
		return ret
	}
	return *o.CrossMarginInterest
}

// GetCrossMarginInterestOk returns a tuple with the CrossMarginInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginInterestOk() (*string, bool) {
	if o == nil || IsNil(o.CrossMarginInterest) {
		return nil, false
	}
	return o.CrossMarginInterest, true
}

// HasCrossMarginInterest returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasCrossMarginInterest() bool {
	if o != nil && !IsNil(o.CrossMarginInterest) {
		return true
	}

	return false
}

// SetCrossMarginInterest gets a reference to the given string and assigns it to the CrossMarginInterest field.
func (o *PmarginGetBalanceV1RespItem) SetCrossMarginInterest(v string) {
	o.CrossMarginInterest = &v
}

// GetCrossMarginLocked returns the CrossMarginLocked field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginLocked() string {
	if o == nil || IsNil(o.CrossMarginLocked) {
		var ret string
		return ret
	}
	return *o.CrossMarginLocked
}

// GetCrossMarginLockedOk returns a tuple with the CrossMarginLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetCrossMarginLockedOk() (*string, bool) {
	if o == nil || IsNil(o.CrossMarginLocked) {
		return nil, false
	}
	return o.CrossMarginLocked, true
}

// HasCrossMarginLocked returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasCrossMarginLocked() bool {
	if o != nil && !IsNil(o.CrossMarginLocked) {
		return true
	}

	return false
}

// SetCrossMarginLocked gets a reference to the given string and assigns it to the CrossMarginLocked field.
func (o *PmarginGetBalanceV1RespItem) SetCrossMarginLocked(v string) {
	o.CrossMarginLocked = &v
}

// GetNegativeBalance returns the NegativeBalance field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetNegativeBalance() string {
	if o == nil || IsNil(o.NegativeBalance) {
		var ret string
		return ret
	}
	return *o.NegativeBalance
}

// GetNegativeBalanceOk returns a tuple with the NegativeBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetNegativeBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.NegativeBalance) {
		return nil, false
	}
	return o.NegativeBalance, true
}

// HasNegativeBalance returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasNegativeBalance() bool {
	if o != nil && !IsNil(o.NegativeBalance) {
		return true
	}

	return false
}

// SetNegativeBalance gets a reference to the given string and assigns it to the NegativeBalance field.
func (o *PmarginGetBalanceV1RespItem) SetNegativeBalance(v string) {
	o.NegativeBalance = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetTotalWalletBalance() string {
	if o == nil || IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasTotalWalletBalance() bool {
	if o != nil && !IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *PmarginGetBalanceV1RespItem) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetUmUnrealizedPNL returns the UmUnrealizedPNL field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetUmUnrealizedPNL() string {
	if o == nil || IsNil(o.UmUnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.UmUnrealizedPNL
}

// GetUmUnrealizedPNLOk returns a tuple with the UmUnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetUmUnrealizedPNLOk() (*string, bool) {
	if o == nil || IsNil(o.UmUnrealizedPNL) {
		return nil, false
	}
	return o.UmUnrealizedPNL, true
}

// HasUmUnrealizedPNL returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasUmUnrealizedPNL() bool {
	if o != nil && !IsNil(o.UmUnrealizedPNL) {
		return true
	}

	return false
}

// SetUmUnrealizedPNL gets a reference to the given string and assigns it to the UmUnrealizedPNL field.
func (o *PmarginGetBalanceV1RespItem) SetUmUnrealizedPNL(v string) {
	o.UmUnrealizedPNL = &v
}

// GetUmWalletBalance returns the UmWalletBalance field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetUmWalletBalance() string {
	if o == nil || IsNil(o.UmWalletBalance) {
		var ret string
		return ret
	}
	return *o.UmWalletBalance
}

// GetUmWalletBalanceOk returns a tuple with the UmWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetUmWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.UmWalletBalance) {
		return nil, false
	}
	return o.UmWalletBalance, true
}

// HasUmWalletBalance returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasUmWalletBalance() bool {
	if o != nil && !IsNil(o.UmWalletBalance) {
		return true
	}

	return false
}

// SetUmWalletBalance gets a reference to the given string and assigns it to the UmWalletBalance field.
func (o *PmarginGetBalanceV1RespItem) SetUmWalletBalance(v string) {
	o.UmWalletBalance = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PmarginGetBalanceV1RespItem) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetBalanceV1RespItem) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PmarginGetBalanceV1RespItem) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *PmarginGetBalanceV1RespItem) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o PmarginGetBalanceV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetBalanceV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.CmUnrealizedPNL) {
		toSerialize["cmUnrealizedPNL"] = o.CmUnrealizedPNL
	}
	if !IsNil(o.CmWalletBalance) {
		toSerialize["cmWalletBalance"] = o.CmWalletBalance
	}
	if !IsNil(o.CrossMarginAsset) {
		toSerialize["crossMarginAsset"] = o.CrossMarginAsset
	}
	if !IsNil(o.CrossMarginBorrowed) {
		toSerialize["crossMarginBorrowed"] = o.CrossMarginBorrowed
	}
	if !IsNil(o.CrossMarginFree) {
		toSerialize["crossMarginFree"] = o.CrossMarginFree
	}
	if !IsNil(o.CrossMarginInterest) {
		toSerialize["crossMarginInterest"] = o.CrossMarginInterest
	}
	if !IsNil(o.CrossMarginLocked) {
		toSerialize["crossMarginLocked"] = o.CrossMarginLocked
	}
	if !IsNil(o.NegativeBalance) {
		toSerialize["negativeBalance"] = o.NegativeBalance
	}
	if !IsNil(o.TotalWalletBalance) {
		toSerialize["totalWalletBalance"] = o.TotalWalletBalance
	}
	if !IsNil(o.UmUnrealizedPNL) {
		toSerialize["umUnrealizedPNL"] = o.UmUnrealizedPNL
	}
	if !IsNil(o.UmWalletBalance) {
		toSerialize["umWalletBalance"] = o.UmWalletBalance
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullablePmarginGetBalanceV1RespItem struct {
	value *PmarginGetBalanceV1RespItem
	isSet bool
}

func (v NullablePmarginGetBalanceV1RespItem) Get() *PmarginGetBalanceV1RespItem {
	return v.value
}

func (v *NullablePmarginGetBalanceV1RespItem) Set(val *PmarginGetBalanceV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetBalanceV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetBalanceV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetBalanceV1RespItem(val *PmarginGetBalanceV1RespItem) *NullablePmarginGetBalanceV1RespItem {
	return &NullablePmarginGetBalanceV1RespItem{value: val, isSet: true}
}

func (v NullablePmarginGetBalanceV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetBalanceV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


