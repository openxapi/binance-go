/*
Binance Margin API

OpenAPI specification for Binance cryptocurrency exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginGetMarginAccountV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginAccountV1Resp{}

// MarginGetMarginAccountV1Resp struct for MarginGetMarginAccountV1Resp
type MarginGetMarginAccountV1Resp struct {
	TotalCollateralValueInUSDT *string `json:"TotalCollateralValueInUSDT,omitempty"`
	AccountType *string `json:"accountType,omitempty"`
	BorrowEnabled *bool `json:"borrowEnabled,omitempty"`
	CollateralMarginLevel *string `json:"collateralMarginLevel,omitempty"`
	Created *bool `json:"created,omitempty"`
	MarginLevel *string `json:"marginLevel,omitempty"`
	TotalAssetOfBtc *string `json:"totalAssetOfBtc,omitempty"`
	TotalLiabilityOfBtc *string `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc *string `json:"totalNetAssetOfBtc,omitempty"`
	TotalOpenOrderLossInUSDT *string `json:"totalOpenOrderLossInUSDT,omitempty"`
	TradeEnabled *bool `json:"tradeEnabled,omitempty"`
	TransferInEnabled *bool `json:"transferInEnabled,omitempty"`
	TransferOutEnabled *bool `json:"transferOutEnabled,omitempty"`
	UserAssets []MarginGetMarginAccountV1RespUserAssetsInner `json:"userAssets,omitempty"`
}

// NewMarginGetMarginAccountV1Resp instantiates a new MarginGetMarginAccountV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginAccountV1Resp() *MarginGetMarginAccountV1Resp {
	this := MarginGetMarginAccountV1Resp{}
	return &this
}

// NewMarginGetMarginAccountV1RespWithDefaults instantiates a new MarginGetMarginAccountV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginAccountV1RespWithDefaults() *MarginGetMarginAccountV1Resp {
	this := MarginGetMarginAccountV1Resp{}
	return &this
}

// GetTotalCollateralValueInUSDT returns the TotalCollateralValueInUSDT field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetTotalCollateralValueInUSDT() string {
	if o == nil || IsNil(o.TotalCollateralValueInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalCollateralValueInUSDT
}

// GetTotalCollateralValueInUSDTOk returns a tuple with the TotalCollateralValueInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetTotalCollateralValueInUSDTOk() (*string, bool) {
	if o == nil || IsNil(o.TotalCollateralValueInUSDT) {
		return nil, false
	}
	return o.TotalCollateralValueInUSDT, true
}

// HasTotalCollateralValueInUSDT returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasTotalCollateralValueInUSDT() bool {
	if o != nil && !IsNil(o.TotalCollateralValueInUSDT) {
		return true
	}

	return false
}

// SetTotalCollateralValueInUSDT gets a reference to the given string and assigns it to the TotalCollateralValueInUSDT field.
func (o *MarginGetMarginAccountV1Resp) SetTotalCollateralValueInUSDT(v string) {
	o.TotalCollateralValueInUSDT = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *MarginGetMarginAccountV1Resp) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBorrowEnabled returns the BorrowEnabled field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetBorrowEnabled() bool {
	if o == nil || IsNil(o.BorrowEnabled) {
		var ret bool
		return ret
	}
	return *o.BorrowEnabled
}

// GetBorrowEnabledOk returns a tuple with the BorrowEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetBorrowEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BorrowEnabled) {
		return nil, false
	}
	return o.BorrowEnabled, true
}

// HasBorrowEnabled returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasBorrowEnabled() bool {
	if o != nil && !IsNil(o.BorrowEnabled) {
		return true
	}

	return false
}

// SetBorrowEnabled gets a reference to the given bool and assigns it to the BorrowEnabled field.
func (o *MarginGetMarginAccountV1Resp) SetBorrowEnabled(v bool) {
	o.BorrowEnabled = &v
}

// GetCollateralMarginLevel returns the CollateralMarginLevel field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetCollateralMarginLevel() string {
	if o == nil || IsNil(o.CollateralMarginLevel) {
		var ret string
		return ret
	}
	return *o.CollateralMarginLevel
}

// GetCollateralMarginLevelOk returns a tuple with the CollateralMarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetCollateralMarginLevelOk() (*string, bool) {
	if o == nil || IsNil(o.CollateralMarginLevel) {
		return nil, false
	}
	return o.CollateralMarginLevel, true
}

// HasCollateralMarginLevel returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasCollateralMarginLevel() bool {
	if o != nil && !IsNil(o.CollateralMarginLevel) {
		return true
	}

	return false
}

// SetCollateralMarginLevel gets a reference to the given string and assigns it to the CollateralMarginLevel field.
func (o *MarginGetMarginAccountV1Resp) SetCollateralMarginLevel(v string) {
	o.CollateralMarginLevel = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetCreated() bool {
	if o == nil || IsNil(o.Created) {
		var ret bool
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given bool and assigns it to the Created field.
func (o *MarginGetMarginAccountV1Resp) SetCreated(v bool) {
	o.Created = &v
}

// GetMarginLevel returns the MarginLevel field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetMarginLevel() string {
	if o == nil || IsNil(o.MarginLevel) {
		var ret string
		return ret
	}
	return *o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetMarginLevelOk() (*string, bool) {
	if o == nil || IsNil(o.MarginLevel) {
		return nil, false
	}
	return o.MarginLevel, true
}

// HasMarginLevel returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasMarginLevel() bool {
	if o != nil && !IsNil(o.MarginLevel) {
		return true
	}

	return false
}

// SetMarginLevel gets a reference to the given string and assigns it to the MarginLevel field.
func (o *MarginGetMarginAccountV1Resp) SetMarginLevel(v string) {
	o.MarginLevel = &v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetTotalAssetOfBtc() string {
	if o == nil || IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasTotalAssetOfBtc() bool {
	if o != nil && !IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *MarginGetMarginAccountV1Resp) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetTotalLiabilityOfBtc() string {
	if o == nil || IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasTotalLiabilityOfBtc() bool {
	if o != nil && !IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *MarginGetMarginAccountV1Resp) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetTotalNetAssetOfBtc() string {
	if o == nil || IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasTotalNetAssetOfBtc() bool {
	if o != nil && !IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *MarginGetMarginAccountV1Resp) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

// GetTotalOpenOrderLossInUSDT returns the TotalOpenOrderLossInUSDT field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetTotalOpenOrderLossInUSDT() string {
	if o == nil || IsNil(o.TotalOpenOrderLossInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderLossInUSDT
}

// GetTotalOpenOrderLossInUSDTOk returns a tuple with the TotalOpenOrderLossInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetTotalOpenOrderLossInUSDTOk() (*string, bool) {
	if o == nil || IsNil(o.TotalOpenOrderLossInUSDT) {
		return nil, false
	}
	return o.TotalOpenOrderLossInUSDT, true
}

// HasTotalOpenOrderLossInUSDT returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasTotalOpenOrderLossInUSDT() bool {
	if o != nil && !IsNil(o.TotalOpenOrderLossInUSDT) {
		return true
	}

	return false
}

// SetTotalOpenOrderLossInUSDT gets a reference to the given string and assigns it to the TotalOpenOrderLossInUSDT field.
func (o *MarginGetMarginAccountV1Resp) SetTotalOpenOrderLossInUSDT(v string) {
	o.TotalOpenOrderLossInUSDT = &v
}

// GetTradeEnabled returns the TradeEnabled field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetTradeEnabled() bool {
	if o == nil || IsNil(o.TradeEnabled) {
		var ret bool
		return ret
	}
	return *o.TradeEnabled
}

// GetTradeEnabledOk returns a tuple with the TradeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetTradeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TradeEnabled) {
		return nil, false
	}
	return o.TradeEnabled, true
}

// HasTradeEnabled returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasTradeEnabled() bool {
	if o != nil && !IsNil(o.TradeEnabled) {
		return true
	}

	return false
}

// SetTradeEnabled gets a reference to the given bool and assigns it to the TradeEnabled field.
func (o *MarginGetMarginAccountV1Resp) SetTradeEnabled(v bool) {
	o.TradeEnabled = &v
}

// GetTransferInEnabled returns the TransferInEnabled field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetTransferInEnabled() bool {
	if o == nil || IsNil(o.TransferInEnabled) {
		var ret bool
		return ret
	}
	return *o.TransferInEnabled
}

// GetTransferInEnabledOk returns a tuple with the TransferInEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetTransferInEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TransferInEnabled) {
		return nil, false
	}
	return o.TransferInEnabled, true
}

// HasTransferInEnabled returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasTransferInEnabled() bool {
	if o != nil && !IsNil(o.TransferInEnabled) {
		return true
	}

	return false
}

// SetTransferInEnabled gets a reference to the given bool and assigns it to the TransferInEnabled field.
func (o *MarginGetMarginAccountV1Resp) SetTransferInEnabled(v bool) {
	o.TransferInEnabled = &v
}

// GetTransferOutEnabled returns the TransferOutEnabled field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetTransferOutEnabled() bool {
	if o == nil || IsNil(o.TransferOutEnabled) {
		var ret bool
		return ret
	}
	return *o.TransferOutEnabled
}

// GetTransferOutEnabledOk returns a tuple with the TransferOutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetTransferOutEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TransferOutEnabled) {
		return nil, false
	}
	return o.TransferOutEnabled, true
}

// HasTransferOutEnabled returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasTransferOutEnabled() bool {
	if o != nil && !IsNil(o.TransferOutEnabled) {
		return true
	}

	return false
}

// SetTransferOutEnabled gets a reference to the given bool and assigns it to the TransferOutEnabled field.
func (o *MarginGetMarginAccountV1Resp) SetTransferOutEnabled(v bool) {
	o.TransferOutEnabled = &v
}

// GetUserAssets returns the UserAssets field value if set, zero value otherwise.
func (o *MarginGetMarginAccountV1Resp) GetUserAssets() []MarginGetMarginAccountV1RespUserAssetsInner {
	if o == nil || IsNil(o.UserAssets) {
		var ret []MarginGetMarginAccountV1RespUserAssetsInner
		return ret
	}
	return o.UserAssets
}

// GetUserAssetsOk returns a tuple with the UserAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAccountV1Resp) GetUserAssetsOk() ([]MarginGetMarginAccountV1RespUserAssetsInner, bool) {
	if o == nil || IsNil(o.UserAssets) {
		return nil, false
	}
	return o.UserAssets, true
}

// HasUserAssets returns a boolean if a field has been set.
func (o *MarginGetMarginAccountV1Resp) HasUserAssets() bool {
	if o != nil && !IsNil(o.UserAssets) {
		return true
	}

	return false
}

// SetUserAssets gets a reference to the given []MarginGetMarginAccountV1RespUserAssetsInner and assigns it to the UserAssets field.
func (o *MarginGetMarginAccountV1Resp) SetUserAssets(v []MarginGetMarginAccountV1RespUserAssetsInner) {
	o.UserAssets = v
}

func (o MarginGetMarginAccountV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginAccountV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCollateralValueInUSDT) {
		toSerialize["TotalCollateralValueInUSDT"] = o.TotalCollateralValueInUSDT
	}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.BorrowEnabled) {
		toSerialize["borrowEnabled"] = o.BorrowEnabled
	}
	if !IsNil(o.CollateralMarginLevel) {
		toSerialize["collateralMarginLevel"] = o.CollateralMarginLevel
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
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
	if !IsNil(o.TotalOpenOrderLossInUSDT) {
		toSerialize["totalOpenOrderLossInUSDT"] = o.TotalOpenOrderLossInUSDT
	}
	if !IsNil(o.TradeEnabled) {
		toSerialize["tradeEnabled"] = o.TradeEnabled
	}
	if !IsNil(o.TransferInEnabled) {
		toSerialize["transferInEnabled"] = o.TransferInEnabled
	}
	if !IsNil(o.TransferOutEnabled) {
		toSerialize["transferOutEnabled"] = o.TransferOutEnabled
	}
	if !IsNil(o.UserAssets) {
		toSerialize["userAssets"] = o.UserAssets
	}
	return toSerialize, nil
}

type NullableMarginGetMarginAccountV1Resp struct {
	value *MarginGetMarginAccountV1Resp
	isSet bool
}

func (v NullableMarginGetMarginAccountV1Resp) Get() *MarginGetMarginAccountV1Resp {
	return v.value
}

func (v *NullableMarginGetMarginAccountV1Resp) Set(val *MarginGetMarginAccountV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginAccountV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginAccountV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginAccountV1Resp(val *MarginGetMarginAccountV1Resp) *NullableMarginGetMarginAccountV1Resp {
	return &NullableMarginGetMarginAccountV1Resp{value: val, isSet: true}
}

func (v NullableMarginGetMarginAccountV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginAccountV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


