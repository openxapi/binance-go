/*
Binance Ufutures API

OpenAPI specification for Binance cryptocurrency exchange - Ufutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ufutures

import (
	"encoding/json"
)

// checks if the UfuturesGetAccountV3RespAssetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetAccountV3RespAssetsInner{}

// UfuturesGetAccountV3RespAssetsInner struct for UfuturesGetAccountV3RespAssetsInner
type UfuturesGetAccountV3RespAssetsInner struct {
	Asset *string `json:"asset,omitempty"`
	AvailableBalance *string `json:"availableBalance,omitempty"`
	CrossUnPnl *string `json:"crossUnPnl,omitempty"`
	CrossWalletBalance *string `json:"crossWalletBalance,omitempty"`
	InitialMargin *string `json:"initialMargin,omitempty"`
	MaintMargin *string `json:"maintMargin,omitempty"`
	MarginBalance *string `json:"marginBalance,omitempty"`
	MaxWithdrawAmount *string `json:"maxWithdrawAmount,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	PositionInitialMargin *string `json:"positionInitialMargin,omitempty"`
	UnrealizedProfit *string `json:"unrealizedProfit,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
	WalletBalance *string `json:"walletBalance,omitempty"`
}

// NewUfuturesGetAccountV3RespAssetsInner instantiates a new UfuturesGetAccountV3RespAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetAccountV3RespAssetsInner() *UfuturesGetAccountV3RespAssetsInner {
	this := UfuturesGetAccountV3RespAssetsInner{}
	return &this
}

// NewUfuturesGetAccountV3RespAssetsInnerWithDefaults instantiates a new UfuturesGetAccountV3RespAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetAccountV3RespAssetsInnerWithDefaults() *UfuturesGetAccountV3RespAssetsInner {
	this := UfuturesGetAccountV3RespAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetAvailableBalance() string {
	if o == nil || IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasAvailableBalance() bool {
	if o != nil && !IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetCrossUnPnl returns the CrossUnPnl field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetCrossUnPnl() string {
	if o == nil || IsNil(o.CrossUnPnl) {
		var ret string
		return ret
	}
	return *o.CrossUnPnl
}

// GetCrossUnPnlOk returns a tuple with the CrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetCrossUnPnlOk() (*string, bool) {
	if o == nil || IsNil(o.CrossUnPnl) {
		return nil, false
	}
	return o.CrossUnPnl, true
}

// HasCrossUnPnl returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasCrossUnPnl() bool {
	if o != nil && !IsNil(o.CrossUnPnl) {
		return true
	}

	return false
}

// SetCrossUnPnl gets a reference to the given string and assigns it to the CrossUnPnl field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetCrossUnPnl(v string) {
	o.CrossUnPnl = &v
}

// GetCrossWalletBalance returns the CrossWalletBalance field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetCrossWalletBalance() string {
	if o == nil || IsNil(o.CrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.CrossWalletBalance
}

// GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetCrossWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.CrossWalletBalance) {
		return nil, false
	}
	return o.CrossWalletBalance, true
}

// HasCrossWalletBalance returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasCrossWalletBalance() bool {
	if o != nil && !IsNil(o.CrossWalletBalance) {
		return true
	}

	return false
}

// SetCrossWalletBalance gets a reference to the given string and assigns it to the CrossWalletBalance field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetCrossWalletBalance(v string) {
	o.CrossWalletBalance = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetInitialMargin() string {
	if o == nil || IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasInitialMargin() bool {
	if o != nil && !IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetMaintMargin() string {
	if o == nil || IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasMaintMargin() bool {
	if o != nil && !IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetMarginBalance() string {
	if o == nil || IsNil(o.MarginBalance) {
		var ret string
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetMarginBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasMarginBalance() bool {
	if o != nil && !IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given string and assigns it to the MarginBalance field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetMarginBalance(v string) {
	o.MarginBalance = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetMaxWithdrawAmount() string {
	if o == nil || IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasMaxWithdrawAmount() bool {
	if o != nil && !IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetOpenOrderInitialMargin() string {
	if o == nil || IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetPositionInitialMargin() string {
	if o == nil || IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasPositionInitialMargin() bool {
	if o != nil && !IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetUnrealizedProfit() string {
	if o == nil || IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasUnrealizedProfit() bool {
	if o != nil && !IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetWalletBalance returns the WalletBalance field value if set, zero value otherwise.
func (o *UfuturesGetAccountV3RespAssetsInner) GetWalletBalance() string {
	if o == nil || IsNil(o.WalletBalance) {
		var ret string
		return ret
	}
	return *o.WalletBalance
}

// GetWalletBalanceOk returns a tuple with the WalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) GetWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.WalletBalance) {
		return nil, false
	}
	return o.WalletBalance, true
}

// HasWalletBalance returns a boolean if a field has been set.
func (o *UfuturesGetAccountV3RespAssetsInner) HasWalletBalance() bool {
	if o != nil && !IsNil(o.WalletBalance) {
		return true
	}

	return false
}

// SetWalletBalance gets a reference to the given string and assigns it to the WalletBalance field.
func (o *UfuturesGetAccountV3RespAssetsInner) SetWalletBalance(v string) {
	o.WalletBalance = &v
}

func (o UfuturesGetAccountV3RespAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetAccountV3RespAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.AvailableBalance) {
		toSerialize["availableBalance"] = o.AvailableBalance
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
	if !IsNil(o.MarginBalance) {
		toSerialize["marginBalance"] = o.MarginBalance
	}
	if !IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}
	if !IsNil(o.OpenOrderInitialMargin) {
		toSerialize["openOrderInitialMargin"] = o.OpenOrderInitialMargin
	}
	if !IsNil(o.PositionInitialMargin) {
		toSerialize["positionInitialMargin"] = o.PositionInitialMargin
	}
	if !IsNil(o.UnrealizedProfit) {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.WalletBalance) {
		toSerialize["walletBalance"] = o.WalletBalance
	}
	return toSerialize, nil
}

type NullableUfuturesGetAccountV3RespAssetsInner struct {
	value *UfuturesGetAccountV3RespAssetsInner
	isSet bool
}

func (v NullableUfuturesGetAccountV3RespAssetsInner) Get() *UfuturesGetAccountV3RespAssetsInner {
	return v.value
}

func (v *NullableUfuturesGetAccountV3RespAssetsInner) Set(val *UfuturesGetAccountV3RespAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetAccountV3RespAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetAccountV3RespAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetAccountV3RespAssetsInner(val *UfuturesGetAccountV3RespAssetsInner) *NullableUfuturesGetAccountV3RespAssetsInner {
	return &NullableUfuturesGetAccountV3RespAssetsInner{value: val, isSet: true}
}

func (v NullableUfuturesGetAccountV3RespAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetAccountV3RespAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


