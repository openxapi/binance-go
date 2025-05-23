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

// checks if the GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp{}

// GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp struct for GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp
type GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp struct {
	Asset *string `json:"asset,omitempty"`
	SubAccountList []GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner `json:"subAccountList,omitempty"`
	TotalInitialMargin *string `json:"totalInitialMargin,omitempty"`
	TotalMaintenanceMargin *string `json:"totalMaintenanceMargin,omitempty"`
	TotalMarginBalance *string `json:"totalMarginBalance,omitempty"`
	TotalOpenOrderInitialMargin *string `json:"totalOpenOrderInitialMargin,omitempty"`
	TotalPositionInitialMargin *string `json:"totalPositionInitialMargin,omitempty"`
	TotalUnrealizedProfit *string `json:"totalUnrealizedProfit,omitempty"`
	TotalWalletBalance *string `json:"totalWalletBalance,omitempty"`
}

// NewGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp instantiates a new GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp() *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp {
	this := GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp{}
	return &this
}

// NewGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryRespWithDefaults instantiates a new GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryRespWithDefaults() *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp {
	this := GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) SetAsset(v string) {
	o.Asset = &v
}

// GetSubAccountList returns the SubAccountList field value if set, zero value otherwise.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetSubAccountList() []GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner {
	if o == nil || IsNil(o.SubAccountList) {
		var ret []GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner
		return ret
	}
	return o.SubAccountList
}

// GetSubAccountListOk returns a tuple with the SubAccountList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetSubAccountListOk() ([]GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner, bool) {
	if o == nil || IsNil(o.SubAccountList) {
		return nil, false
	}
	return o.SubAccountList, true
}

// HasSubAccountList returns a boolean if a field has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) HasSubAccountList() bool {
	if o != nil && !IsNil(o.SubAccountList) {
		return true
	}

	return false
}

// SetSubAccountList gets a reference to the given []GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner and assigns it to the SubAccountList field.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) SetSubAccountList(v []GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner) {
	o.SubAccountList = v
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalInitialMargin() string {
	if o == nil || IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) HasTotalInitialMargin() bool {
	if o != nil && !IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field value if set, zero value otherwise.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalMaintenanceMargin() string {
	if o == nil || IsNil(o.TotalMaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintenanceMargin
}

// GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalMaintenanceMarginOk() (*string, bool) {
	if o == nil || IsNil(o.TotalMaintenanceMargin) {
		return nil, false
	}
	return o.TotalMaintenanceMargin, true
}

// HasTotalMaintenanceMargin returns a boolean if a field has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) HasTotalMaintenanceMargin() bool {
	if o != nil && !IsNil(o.TotalMaintenanceMargin) {
		return true
	}

	return false
}

// SetTotalMaintenanceMargin gets a reference to the given string and assigns it to the TotalMaintenanceMargin field.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) SetTotalMaintenanceMargin(v string) {
	o.TotalMaintenanceMargin = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalMarginBalance() string {
	if o == nil || IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) HasTotalMarginBalance() bool {
	if o != nil && !IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalOpenOrderInitialMargin() string {
	if o == nil || IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalPositionInitialMargin() string {
	if o == nil || IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) HasTotalPositionInitialMargin() bool {
	if o != nil && !IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalUnrealizedProfit() string {
	if o == nil || IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) HasTotalUnrealizedProfit() bool {
	if o != nil && !IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalWalletBalance() string {
	if o == nil || IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) HasTotalWalletBalance() bool {
	if o != nil && !IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

func (o GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.SubAccountList) {
		toSerialize["subAccountList"] = o.SubAccountList
	}
	if !IsNil(o.TotalInitialMargin) {
		toSerialize["totalInitialMargin"] = o.TotalInitialMargin
	}
	if !IsNil(o.TotalMaintenanceMargin) {
		toSerialize["totalMaintenanceMargin"] = o.TotalMaintenanceMargin
	}
	if !IsNil(o.TotalMarginBalance) {
		toSerialize["totalMarginBalance"] = o.TotalMarginBalance
	}
	if !IsNil(o.TotalOpenOrderInitialMargin) {
		toSerialize["totalOpenOrderInitialMargin"] = o.TotalOpenOrderInitialMargin
	}
	if !IsNil(o.TotalPositionInitialMargin) {
		toSerialize["totalPositionInitialMargin"] = o.TotalPositionInitialMargin
	}
	if !IsNil(o.TotalUnrealizedProfit) {
		toSerialize["totalUnrealizedProfit"] = o.TotalUnrealizedProfit
	}
	if !IsNil(o.TotalWalletBalance) {
		toSerialize["totalWalletBalance"] = o.TotalWalletBalance
	}
	return toSerialize, nil
}

type NullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp struct {
	value *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp
	isSet bool
}

func (v NullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) Get() *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp {
	return v.value
}

func (v *NullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) Set(val *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp(val *GetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) *NullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp {
	return &NullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp{value: val, isSet: true}
}

func (v NullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountFuturesAccountSummaryV2RespFutureAccountSummaryResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


