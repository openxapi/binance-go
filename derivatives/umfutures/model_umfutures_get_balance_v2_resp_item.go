/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the UmfuturesGetBalanceV2RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmfuturesGetBalanceV2RespItem{}

// UmfuturesGetBalanceV2RespItem struct for UmfuturesGetBalanceV2RespItem
type UmfuturesGetBalanceV2RespItem struct {
	AccountAlias *string `json:"accountAlias,omitempty"`
	Asset *string `json:"asset,omitempty"`
	AvailableBalance *string `json:"availableBalance,omitempty"`
	Balance *string `json:"balance,omitempty"`
	CrossUnPnl *string `json:"crossUnPnl,omitempty"`
	CrossWalletBalance *string `json:"crossWalletBalance,omitempty"`
	MarginAvailable *bool `json:"marginAvailable,omitempty"`
	MaxWithdrawAmount *string `json:"maxWithdrawAmount,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewUmfuturesGetBalanceV2RespItem instantiates a new UmfuturesGetBalanceV2RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmfuturesGetBalanceV2RespItem() *UmfuturesGetBalanceV2RespItem {
	this := UmfuturesGetBalanceV2RespItem{}
	return &this
}

// NewUmfuturesGetBalanceV2RespItemWithDefaults instantiates a new UmfuturesGetBalanceV2RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmfuturesGetBalanceV2RespItemWithDefaults() *UmfuturesGetBalanceV2RespItem {
	this := UmfuturesGetBalanceV2RespItem{}
	return &this
}

// GetAccountAlias returns the AccountAlias field value if set, zero value otherwise.
func (o *UmfuturesGetBalanceV2RespItem) GetAccountAlias() string {
	if o == nil || IsNil(o.AccountAlias) {
		var ret string
		return ret
	}
	return *o.AccountAlias
}

// GetAccountAliasOk returns a tuple with the AccountAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetBalanceV2RespItem) GetAccountAliasOk() (*string, bool) {
	if o == nil || IsNil(o.AccountAlias) {
		return nil, false
	}
	return o.AccountAlias, true
}

// HasAccountAlias returns a boolean if a field has been set.
func (o *UmfuturesGetBalanceV2RespItem) HasAccountAlias() bool {
	if o != nil && !IsNil(o.AccountAlias) {
		return true
	}

	return false
}

// SetAccountAlias gets a reference to the given string and assigns it to the AccountAlias field.
func (o *UmfuturesGetBalanceV2RespItem) SetAccountAlias(v string) {
	o.AccountAlias = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *UmfuturesGetBalanceV2RespItem) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetBalanceV2RespItem) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *UmfuturesGetBalanceV2RespItem) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *UmfuturesGetBalanceV2RespItem) SetAsset(v string) {
	o.Asset = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *UmfuturesGetBalanceV2RespItem) GetAvailableBalance() string {
	if o == nil || IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetBalanceV2RespItem) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *UmfuturesGetBalanceV2RespItem) HasAvailableBalance() bool {
	if o != nil && !IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *UmfuturesGetBalanceV2RespItem) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *UmfuturesGetBalanceV2RespItem) GetBalance() string {
	if o == nil || IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetBalanceV2RespItem) GetBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *UmfuturesGetBalanceV2RespItem) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *UmfuturesGetBalanceV2RespItem) SetBalance(v string) {
	o.Balance = &v
}

// GetCrossUnPnl returns the CrossUnPnl field value if set, zero value otherwise.
func (o *UmfuturesGetBalanceV2RespItem) GetCrossUnPnl() string {
	if o == nil || IsNil(o.CrossUnPnl) {
		var ret string
		return ret
	}
	return *o.CrossUnPnl
}

// GetCrossUnPnlOk returns a tuple with the CrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetBalanceV2RespItem) GetCrossUnPnlOk() (*string, bool) {
	if o == nil || IsNil(o.CrossUnPnl) {
		return nil, false
	}
	return o.CrossUnPnl, true
}

// HasCrossUnPnl returns a boolean if a field has been set.
func (o *UmfuturesGetBalanceV2RespItem) HasCrossUnPnl() bool {
	if o != nil && !IsNil(o.CrossUnPnl) {
		return true
	}

	return false
}

// SetCrossUnPnl gets a reference to the given string and assigns it to the CrossUnPnl field.
func (o *UmfuturesGetBalanceV2RespItem) SetCrossUnPnl(v string) {
	o.CrossUnPnl = &v
}

// GetCrossWalletBalance returns the CrossWalletBalance field value if set, zero value otherwise.
func (o *UmfuturesGetBalanceV2RespItem) GetCrossWalletBalance() string {
	if o == nil || IsNil(o.CrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.CrossWalletBalance
}

// GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetBalanceV2RespItem) GetCrossWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.CrossWalletBalance) {
		return nil, false
	}
	return o.CrossWalletBalance, true
}

// HasCrossWalletBalance returns a boolean if a field has been set.
func (o *UmfuturesGetBalanceV2RespItem) HasCrossWalletBalance() bool {
	if o != nil && !IsNil(o.CrossWalletBalance) {
		return true
	}

	return false
}

// SetCrossWalletBalance gets a reference to the given string and assigns it to the CrossWalletBalance field.
func (o *UmfuturesGetBalanceV2RespItem) SetCrossWalletBalance(v string) {
	o.CrossWalletBalance = &v
}

// GetMarginAvailable returns the MarginAvailable field value if set, zero value otherwise.
func (o *UmfuturesGetBalanceV2RespItem) GetMarginAvailable() bool {
	if o == nil || IsNil(o.MarginAvailable) {
		var ret bool
		return ret
	}
	return *o.MarginAvailable
}

// GetMarginAvailableOk returns a tuple with the MarginAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetBalanceV2RespItem) GetMarginAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.MarginAvailable) {
		return nil, false
	}
	return o.MarginAvailable, true
}

// HasMarginAvailable returns a boolean if a field has been set.
func (o *UmfuturesGetBalanceV2RespItem) HasMarginAvailable() bool {
	if o != nil && !IsNil(o.MarginAvailable) {
		return true
	}

	return false
}

// SetMarginAvailable gets a reference to the given bool and assigns it to the MarginAvailable field.
func (o *UmfuturesGetBalanceV2RespItem) SetMarginAvailable(v bool) {
	o.MarginAvailable = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *UmfuturesGetBalanceV2RespItem) GetMaxWithdrawAmount() string {
	if o == nil || IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetBalanceV2RespItem) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *UmfuturesGetBalanceV2RespItem) HasMaxWithdrawAmount() bool {
	if o != nil && !IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *UmfuturesGetBalanceV2RespItem) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *UmfuturesGetBalanceV2RespItem) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetBalanceV2RespItem) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *UmfuturesGetBalanceV2RespItem) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *UmfuturesGetBalanceV2RespItem) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o UmfuturesGetBalanceV2RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmfuturesGetBalanceV2RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountAlias) {
		toSerialize["accountAlias"] = o.AccountAlias
	}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.AvailableBalance) {
		toSerialize["availableBalance"] = o.AvailableBalance
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.CrossUnPnl) {
		toSerialize["crossUnPnl"] = o.CrossUnPnl
	}
	if !IsNil(o.CrossWalletBalance) {
		toSerialize["crossWalletBalance"] = o.CrossWalletBalance
	}
	if !IsNil(o.MarginAvailable) {
		toSerialize["marginAvailable"] = o.MarginAvailable
	}
	if !IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableUmfuturesGetBalanceV2RespItem struct {
	value *UmfuturesGetBalanceV2RespItem
	isSet bool
}

func (v NullableUmfuturesGetBalanceV2RespItem) Get() *UmfuturesGetBalanceV2RespItem {
	return v.value
}

func (v *NullableUmfuturesGetBalanceV2RespItem) Set(val *UmfuturesGetBalanceV2RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesGetBalanceV2RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesGetBalanceV2RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesGetBalanceV2RespItem(val *UmfuturesGetBalanceV2RespItem) *NullableUmfuturesGetBalanceV2RespItem {
	return &NullableUmfuturesGetBalanceV2RespItem{value: val, isSet: true}
}

func (v NullableUmfuturesGetBalanceV2RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesGetBalanceV2RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


