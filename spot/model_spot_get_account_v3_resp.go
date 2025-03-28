/*
Binance Spot API

OpenAPI specification for Binance cryptocurrency exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotGetAccountV3Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotGetAccountV3Resp{}

// SpotGetAccountV3Resp struct for SpotGetAccountV3Resp
type SpotGetAccountV3Resp struct {
	AccountType *string `json:"accountType,omitempty"`
	Balances []SpotGetAccountV3RespBalancesInner `json:"balances,omitempty"`
	Brokered *bool `json:"brokered,omitempty"`
	BuyerCommission *int32 `json:"buyerCommission,omitempty"`
	CanDeposit *bool `json:"canDeposit,omitempty"`
	CanTrade *bool `json:"canTrade,omitempty"`
	CanWithdraw *bool `json:"canWithdraw,omitempty"`
	CommissionRates *SpotGetAccountCommissionV3RespStandardCommission `json:"commissionRates,omitempty"`
	MakerCommission *int32 `json:"makerCommission,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	PreventSor *bool `json:"preventSor,omitempty"`
	RequireSelfTradePrevention *bool `json:"requireSelfTradePrevention,omitempty"`
	SellerCommission *int32 `json:"sellerCommission,omitempty"`
	TakerCommission *int32 `json:"takerCommission,omitempty"`
	Uid *int32 `json:"uid,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewSpotGetAccountV3Resp instantiates a new SpotGetAccountV3Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotGetAccountV3Resp() *SpotGetAccountV3Resp {
	this := SpotGetAccountV3Resp{}
	return &this
}

// NewSpotGetAccountV3RespWithDefaults instantiates a new SpotGetAccountV3Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotGetAccountV3RespWithDefaults() *SpotGetAccountV3Resp {
	this := SpotGetAccountV3Resp{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *SpotGetAccountV3Resp) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetBalances() []SpotGetAccountV3RespBalancesInner {
	if o == nil || IsNil(o.Balances) {
		var ret []SpotGetAccountV3RespBalancesInner
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetBalancesOk() ([]SpotGetAccountV3RespBalancesInner, bool) {
	if o == nil || IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasBalances() bool {
	if o != nil && !IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []SpotGetAccountV3RespBalancesInner and assigns it to the Balances field.
func (o *SpotGetAccountV3Resp) SetBalances(v []SpotGetAccountV3RespBalancesInner) {
	o.Balances = v
}

// GetBrokered returns the Brokered field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetBrokered() bool {
	if o == nil || IsNil(o.Brokered) {
		var ret bool
		return ret
	}
	return *o.Brokered
}

// GetBrokeredOk returns a tuple with the Brokered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetBrokeredOk() (*bool, bool) {
	if o == nil || IsNil(o.Brokered) {
		return nil, false
	}
	return o.Brokered, true
}

// HasBrokered returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasBrokered() bool {
	if o != nil && !IsNil(o.Brokered) {
		return true
	}

	return false
}

// SetBrokered gets a reference to the given bool and assigns it to the Brokered field.
func (o *SpotGetAccountV3Resp) SetBrokered(v bool) {
	o.Brokered = &v
}

// GetBuyerCommission returns the BuyerCommission field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetBuyerCommission() int32 {
	if o == nil || IsNil(o.BuyerCommission) {
		var ret int32
		return ret
	}
	return *o.BuyerCommission
}

// GetBuyerCommissionOk returns a tuple with the BuyerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetBuyerCommissionOk() (*int32, bool) {
	if o == nil || IsNil(o.BuyerCommission) {
		return nil, false
	}
	return o.BuyerCommission, true
}

// HasBuyerCommission returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasBuyerCommission() bool {
	if o != nil && !IsNil(o.BuyerCommission) {
		return true
	}

	return false
}

// SetBuyerCommission gets a reference to the given int32 and assigns it to the BuyerCommission field.
func (o *SpotGetAccountV3Resp) SetBuyerCommission(v int32) {
	o.BuyerCommission = &v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetCanDeposit() bool {
	if o == nil || IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetCanDepositOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasCanDeposit() bool {
	if o != nil && !IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *SpotGetAccountV3Resp) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetCanTrade() bool {
	if o == nil || IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetCanTradeOk() (*bool, bool) {
	if o == nil || IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasCanTrade() bool {
	if o != nil && !IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *SpotGetAccountV3Resp) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetCanWithdraw() bool {
	if o == nil || IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasCanWithdraw() bool {
	if o != nil && !IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *SpotGetAccountV3Resp) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetCommissionRates returns the CommissionRates field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetCommissionRates() SpotGetAccountCommissionV3RespStandardCommission {
	if o == nil || IsNil(o.CommissionRates) {
		var ret SpotGetAccountCommissionV3RespStandardCommission
		return ret
	}
	return *o.CommissionRates
}

// GetCommissionRatesOk returns a tuple with the CommissionRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetCommissionRatesOk() (*SpotGetAccountCommissionV3RespStandardCommission, bool) {
	if o == nil || IsNil(o.CommissionRates) {
		return nil, false
	}
	return o.CommissionRates, true
}

// HasCommissionRates returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasCommissionRates() bool {
	if o != nil && !IsNil(o.CommissionRates) {
		return true
	}

	return false
}

// SetCommissionRates gets a reference to the given SpotGetAccountCommissionV3RespStandardCommission and assigns it to the CommissionRates field.
func (o *SpotGetAccountV3Resp) SetCommissionRates(v SpotGetAccountCommissionV3RespStandardCommission) {
	o.CommissionRates = &v
}

// GetMakerCommission returns the MakerCommission field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetMakerCommission() int32 {
	if o == nil || IsNil(o.MakerCommission) {
		var ret int32
		return ret
	}
	return *o.MakerCommission
}

// GetMakerCommissionOk returns a tuple with the MakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetMakerCommissionOk() (*int32, bool) {
	if o == nil || IsNil(o.MakerCommission) {
		return nil, false
	}
	return o.MakerCommission, true
}

// HasMakerCommission returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasMakerCommission() bool {
	if o != nil && !IsNil(o.MakerCommission) {
		return true
	}

	return false
}

// SetMakerCommission gets a reference to the given int32 and assigns it to the MakerCommission field.
func (o *SpotGetAccountV3Resp) SetMakerCommission(v int32) {
	o.MakerCommission = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *SpotGetAccountV3Resp) SetPermissions(v []string) {
	o.Permissions = v
}

// GetPreventSor returns the PreventSor field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetPreventSor() bool {
	if o == nil || IsNil(o.PreventSor) {
		var ret bool
		return ret
	}
	return *o.PreventSor
}

// GetPreventSorOk returns a tuple with the PreventSor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetPreventSorOk() (*bool, bool) {
	if o == nil || IsNil(o.PreventSor) {
		return nil, false
	}
	return o.PreventSor, true
}

// HasPreventSor returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasPreventSor() bool {
	if o != nil && !IsNil(o.PreventSor) {
		return true
	}

	return false
}

// SetPreventSor gets a reference to the given bool and assigns it to the PreventSor field.
func (o *SpotGetAccountV3Resp) SetPreventSor(v bool) {
	o.PreventSor = &v
}

// GetRequireSelfTradePrevention returns the RequireSelfTradePrevention field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetRequireSelfTradePrevention() bool {
	if o == nil || IsNil(o.RequireSelfTradePrevention) {
		var ret bool
		return ret
	}
	return *o.RequireSelfTradePrevention
}

// GetRequireSelfTradePreventionOk returns a tuple with the RequireSelfTradePrevention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetRequireSelfTradePreventionOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSelfTradePrevention) {
		return nil, false
	}
	return o.RequireSelfTradePrevention, true
}

// HasRequireSelfTradePrevention returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasRequireSelfTradePrevention() bool {
	if o != nil && !IsNil(o.RequireSelfTradePrevention) {
		return true
	}

	return false
}

// SetRequireSelfTradePrevention gets a reference to the given bool and assigns it to the RequireSelfTradePrevention field.
func (o *SpotGetAccountV3Resp) SetRequireSelfTradePrevention(v bool) {
	o.RequireSelfTradePrevention = &v
}

// GetSellerCommission returns the SellerCommission field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetSellerCommission() int32 {
	if o == nil || IsNil(o.SellerCommission) {
		var ret int32
		return ret
	}
	return *o.SellerCommission
}

// GetSellerCommissionOk returns a tuple with the SellerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetSellerCommissionOk() (*int32, bool) {
	if o == nil || IsNil(o.SellerCommission) {
		return nil, false
	}
	return o.SellerCommission, true
}

// HasSellerCommission returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasSellerCommission() bool {
	if o != nil && !IsNil(o.SellerCommission) {
		return true
	}

	return false
}

// SetSellerCommission gets a reference to the given int32 and assigns it to the SellerCommission field.
func (o *SpotGetAccountV3Resp) SetSellerCommission(v int32) {
	o.SellerCommission = &v
}

// GetTakerCommission returns the TakerCommission field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetTakerCommission() int32 {
	if o == nil || IsNil(o.TakerCommission) {
		var ret int32
		return ret
	}
	return *o.TakerCommission
}

// GetTakerCommissionOk returns a tuple with the TakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetTakerCommissionOk() (*int32, bool) {
	if o == nil || IsNil(o.TakerCommission) {
		return nil, false
	}
	return o.TakerCommission, true
}

// HasTakerCommission returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasTakerCommission() bool {
	if o != nil && !IsNil(o.TakerCommission) {
		return true
	}

	return false
}

// SetTakerCommission gets a reference to the given int32 and assigns it to the TakerCommission field.
func (o *SpotGetAccountV3Resp) SetTakerCommission(v int32) {
	o.TakerCommission = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *SpotGetAccountV3Resp) SetUid(v int32) {
	o.Uid = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *SpotGetAccountV3Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetAccountV3Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *SpotGetAccountV3Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *SpotGetAccountV3Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o SpotGetAccountV3Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotGetAccountV3Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	if !IsNil(o.Brokered) {
		toSerialize["brokered"] = o.Brokered
	}
	if !IsNil(o.BuyerCommission) {
		toSerialize["buyerCommission"] = o.BuyerCommission
	}
	if !IsNil(o.CanDeposit) {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if !IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
	}
	if !IsNil(o.CanWithdraw) {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if !IsNil(o.CommissionRates) {
		toSerialize["commissionRates"] = o.CommissionRates
	}
	if !IsNil(o.MakerCommission) {
		toSerialize["makerCommission"] = o.MakerCommission
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.PreventSor) {
		toSerialize["preventSor"] = o.PreventSor
	}
	if !IsNil(o.RequireSelfTradePrevention) {
		toSerialize["requireSelfTradePrevention"] = o.RequireSelfTradePrevention
	}
	if !IsNil(o.SellerCommission) {
		toSerialize["sellerCommission"] = o.SellerCommission
	}
	if !IsNil(o.TakerCommission) {
		toSerialize["takerCommission"] = o.TakerCommission
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableSpotGetAccountV3Resp struct {
	value *SpotGetAccountV3Resp
	isSet bool
}

func (v NullableSpotGetAccountV3Resp) Get() *SpotGetAccountV3Resp {
	return v.value
}

func (v *NullableSpotGetAccountV3Resp) Set(val *SpotGetAccountV3Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotGetAccountV3Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotGetAccountV3Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotGetAccountV3Resp(val *SpotGetAccountV3Resp) *NullableSpotGetAccountV3Resp {
	return &NullableSpotGetAccountV3Resp{value: val, isSet: true}
}

func (v NullableSpotGetAccountV3Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotGetAccountV3Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


