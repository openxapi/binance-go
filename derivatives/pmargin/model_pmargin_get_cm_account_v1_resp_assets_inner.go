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

// checks if the PmarginGetCmAccountV1RespAssetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetCmAccountV1RespAssetsInner{}

// PmarginGetCmAccountV1RespAssetsInner struct for PmarginGetCmAccountV1RespAssetsInner
type PmarginGetCmAccountV1RespAssetsInner struct {
	Asset *string `json:"asset,omitempty"`
	CrossUnPnl *string `json:"crossUnPnl,omitempty"`
	CrossWalletBalance *string `json:"crossWalletBalance,omitempty"`
	InitialMargin *string `json:"initialMargin,omitempty"`
	MaintMargin *string `json:"maintMargin,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	PositionInitialMargin *string `json:"positionInitialMargin,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewPmarginGetCmAccountV1RespAssetsInner instantiates a new PmarginGetCmAccountV1RespAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetCmAccountV1RespAssetsInner() *PmarginGetCmAccountV1RespAssetsInner {
	this := PmarginGetCmAccountV1RespAssetsInner{}
	return &this
}

// NewPmarginGetCmAccountV1RespAssetsInnerWithDefaults instantiates a new PmarginGetCmAccountV1RespAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetCmAccountV1RespAssetsInnerWithDefaults() *PmarginGetCmAccountV1RespAssetsInner {
	this := PmarginGetCmAccountV1RespAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *PmarginGetCmAccountV1RespAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetCrossUnPnl returns the CrossUnPnl field value if set, zero value otherwise.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetCrossUnPnl() string {
	if o == nil || IsNil(o.CrossUnPnl) {
		var ret string
		return ret
	}
	return *o.CrossUnPnl
}

// GetCrossUnPnlOk returns a tuple with the CrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetCrossUnPnlOk() (*string, bool) {
	if o == nil || IsNil(o.CrossUnPnl) {
		return nil, false
	}
	return o.CrossUnPnl, true
}

// HasCrossUnPnl returns a boolean if a field has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) HasCrossUnPnl() bool {
	if o != nil && !IsNil(o.CrossUnPnl) {
		return true
	}

	return false
}

// SetCrossUnPnl gets a reference to the given string and assigns it to the CrossUnPnl field.
func (o *PmarginGetCmAccountV1RespAssetsInner) SetCrossUnPnl(v string) {
	o.CrossUnPnl = &v
}

// GetCrossWalletBalance returns the CrossWalletBalance field value if set, zero value otherwise.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetCrossWalletBalance() string {
	if o == nil || IsNil(o.CrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.CrossWalletBalance
}

// GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetCrossWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.CrossWalletBalance) {
		return nil, false
	}
	return o.CrossWalletBalance, true
}

// HasCrossWalletBalance returns a boolean if a field has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) HasCrossWalletBalance() bool {
	if o != nil && !IsNil(o.CrossWalletBalance) {
		return true
	}

	return false
}

// SetCrossWalletBalance gets a reference to the given string and assigns it to the CrossWalletBalance field.
func (o *PmarginGetCmAccountV1RespAssetsInner) SetCrossWalletBalance(v string) {
	o.CrossWalletBalance = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetInitialMargin() string {
	if o == nil || IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) HasInitialMargin() bool {
	if o != nil && !IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *PmarginGetCmAccountV1RespAssetsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetMaintMargin() string {
	if o == nil || IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) HasMaintMargin() bool {
	if o != nil && !IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *PmarginGetCmAccountV1RespAssetsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetOpenOrderInitialMargin() string {
	if o == nil || IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *PmarginGetCmAccountV1RespAssetsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetPositionInitialMargin() string {
	if o == nil || IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) HasPositionInitialMargin() bool {
	if o != nil && !IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *PmarginGetCmAccountV1RespAssetsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PmarginGetCmAccountV1RespAssetsInner) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *PmarginGetCmAccountV1RespAssetsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o PmarginGetCmAccountV1RespAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetCmAccountV1RespAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.CrossUnPnl) {
		toSerialize["crossUnPnl"] = o.CrossUnPnl
	}
	if !IsNil(o.CrossWalletBalance) {
		toSerialize["crossWalletBalance"] = o.CrossWalletBalance
	}
	if !IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !IsNil(o.OpenOrderInitialMargin) {
		toSerialize["openOrderInitialMargin"] = o.OpenOrderInitialMargin
	}
	if !IsNil(o.PositionInitialMargin) {
		toSerialize["positionInitialMargin"] = o.PositionInitialMargin
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullablePmarginGetCmAccountV1RespAssetsInner struct {
	value *PmarginGetCmAccountV1RespAssetsInner
	isSet bool
}

func (v NullablePmarginGetCmAccountV1RespAssetsInner) Get() *PmarginGetCmAccountV1RespAssetsInner {
	return v.value
}

func (v *NullablePmarginGetCmAccountV1RespAssetsInner) Set(val *PmarginGetCmAccountV1RespAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetCmAccountV1RespAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetCmAccountV1RespAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetCmAccountV1RespAssetsInner(val *PmarginGetCmAccountV1RespAssetsInner) *NullablePmarginGetCmAccountV1RespAssetsInner {
	return &NullablePmarginGetCmAccountV1RespAssetsInner{value: val, isSet: true}
}

func (v NullablePmarginGetCmAccountV1RespAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetCmAccountV1RespAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


