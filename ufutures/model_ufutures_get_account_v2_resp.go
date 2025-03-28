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

// checks if the UfuturesGetAccountV2Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetAccountV2Resp{}

// UfuturesGetAccountV2Resp struct for UfuturesGetAccountV2Resp
type UfuturesGetAccountV2Resp struct {
	Assets []UfuturesGetAccountV2RespAssetsInner `json:"assets,omitempty"`
	AvailableBalance *string `json:"availableBalance,omitempty"`
	CanDeposit *bool `json:"canDeposit,omitempty"`
	CanWithdraw *bool `json:"canWithdraw,omitempty"`
	FeeBurn *bool `json:"feeBurn,omitempty"`
	FeeTier *int32 `json:"feeTier,omitempty"`
	MaxWithdrawAmount *string `json:"maxWithdrawAmount,omitempty"`
	MultiAssetsMargin *bool `json:"multiAssetsMargin,omitempty"`
	Positions []UfuturesGetAccountV2RespPositionsInner `json:"positions,omitempty"`
	TotalCrossUnPnl *string `json:"totalCrossUnPnl,omitempty"`
	TotalCrossWalletBalance *string `json:"totalCrossWalletBalance,omitempty"`
	TotalInitialMargin *string `json:"totalInitialMargin,omitempty"`
	TotalMaintMargin *string `json:"totalMaintMargin,omitempty"`
	TotalMarginBalance *string `json:"totalMarginBalance,omitempty"`
	TotalOpenOrderInitialMargin *string `json:"totalOpenOrderInitialMargin,omitempty"`
	TotalPositionInitialMargin *string `json:"totalPositionInitialMargin,omitempty"`
	TotalUnrealizedProfit *string `json:"totalUnrealizedProfit,omitempty"`
	TotalWalletBalance *string `json:"totalWalletBalance,omitempty"`
	TradeGroupId *int64 `json:"tradeGroupId,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewUfuturesGetAccountV2Resp instantiates a new UfuturesGetAccountV2Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetAccountV2Resp() *UfuturesGetAccountV2Resp {
	this := UfuturesGetAccountV2Resp{}
	return &this
}

// NewUfuturesGetAccountV2RespWithDefaults instantiates a new UfuturesGetAccountV2Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetAccountV2RespWithDefaults() *UfuturesGetAccountV2Resp {
	this := UfuturesGetAccountV2Resp{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetAssets() []UfuturesGetAccountV2RespAssetsInner {
	if o == nil || IsNil(o.Assets) {
		var ret []UfuturesGetAccountV2RespAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetAssetsOk() ([]UfuturesGetAccountV2RespAssetsInner, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []UfuturesGetAccountV2RespAssetsInner and assigns it to the Assets field.
func (o *UfuturesGetAccountV2Resp) SetAssets(v []UfuturesGetAccountV2RespAssetsInner) {
	o.Assets = v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetAvailableBalance() string {
	if o == nil || IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasAvailableBalance() bool {
	if o != nil && !IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *UfuturesGetAccountV2Resp) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetCanDeposit() bool {
	if o == nil || IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetCanDepositOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasCanDeposit() bool {
	if o != nil && !IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *UfuturesGetAccountV2Resp) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetCanWithdraw() bool {
	if o == nil || IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasCanWithdraw() bool {
	if o != nil && !IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *UfuturesGetAccountV2Resp) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetFeeBurn returns the FeeBurn field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetFeeBurn() bool {
	if o == nil || IsNil(o.FeeBurn) {
		var ret bool
		return ret
	}
	return *o.FeeBurn
}

// GetFeeBurnOk returns a tuple with the FeeBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetFeeBurnOk() (*bool, bool) {
	if o == nil || IsNil(o.FeeBurn) {
		return nil, false
	}
	return o.FeeBurn, true
}

// HasFeeBurn returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasFeeBurn() bool {
	if o != nil && !IsNil(o.FeeBurn) {
		return true
	}

	return false
}

// SetFeeBurn gets a reference to the given bool and assigns it to the FeeBurn field.
func (o *UfuturesGetAccountV2Resp) SetFeeBurn(v bool) {
	o.FeeBurn = &v
}

// GetFeeTier returns the FeeTier field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetFeeTier() int32 {
	if o == nil || IsNil(o.FeeTier) {
		var ret int32
		return ret
	}
	return *o.FeeTier
}

// GetFeeTierOk returns a tuple with the FeeTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetFeeTierOk() (*int32, bool) {
	if o == nil || IsNil(o.FeeTier) {
		return nil, false
	}
	return o.FeeTier, true
}

// HasFeeTier returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasFeeTier() bool {
	if o != nil && !IsNil(o.FeeTier) {
		return true
	}

	return false
}

// SetFeeTier gets a reference to the given int32 and assigns it to the FeeTier field.
func (o *UfuturesGetAccountV2Resp) SetFeeTier(v int32) {
	o.FeeTier = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetMaxWithdrawAmount() string {
	if o == nil || IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasMaxWithdrawAmount() bool {
	if o != nil && !IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *UfuturesGetAccountV2Resp) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetMultiAssetsMargin returns the MultiAssetsMargin field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetMultiAssetsMargin() bool {
	if o == nil || IsNil(o.MultiAssetsMargin) {
		var ret bool
		return ret
	}
	return *o.MultiAssetsMargin
}

// GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetMultiAssetsMarginOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiAssetsMargin) {
		return nil, false
	}
	return o.MultiAssetsMargin, true
}

// HasMultiAssetsMargin returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasMultiAssetsMargin() bool {
	if o != nil && !IsNil(o.MultiAssetsMargin) {
		return true
	}

	return false
}

// SetMultiAssetsMargin gets a reference to the given bool and assigns it to the MultiAssetsMargin field.
func (o *UfuturesGetAccountV2Resp) SetMultiAssetsMargin(v bool) {
	o.MultiAssetsMargin = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetPositions() []UfuturesGetAccountV2RespPositionsInner {
	if o == nil || IsNil(o.Positions) {
		var ret []UfuturesGetAccountV2RespPositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetPositionsOk() ([]UfuturesGetAccountV2RespPositionsInner, bool) {
	if o == nil || IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasPositions() bool {
	if o != nil && !IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []UfuturesGetAccountV2RespPositionsInner and assigns it to the Positions field.
func (o *UfuturesGetAccountV2Resp) SetPositions(v []UfuturesGetAccountV2RespPositionsInner) {
	o.Positions = v
}

// GetTotalCrossUnPnl returns the TotalCrossUnPnl field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTotalCrossUnPnl() string {
	if o == nil || IsNil(o.TotalCrossUnPnl) {
		var ret string
		return ret
	}
	return *o.TotalCrossUnPnl
}

// GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTotalCrossUnPnlOk() (*string, bool) {
	if o == nil || IsNil(o.TotalCrossUnPnl) {
		return nil, false
	}
	return o.TotalCrossUnPnl, true
}

// HasTotalCrossUnPnl returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTotalCrossUnPnl() bool {
	if o != nil && !IsNil(o.TotalCrossUnPnl) {
		return true
	}

	return false
}

// SetTotalCrossUnPnl gets a reference to the given string and assigns it to the TotalCrossUnPnl field.
func (o *UfuturesGetAccountV2Resp) SetTotalCrossUnPnl(v string) {
	o.TotalCrossUnPnl = &v
}

// GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTotalCrossWalletBalance() string {
	if o == nil || IsNil(o.TotalCrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalCrossWalletBalance
}

// GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTotalCrossWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.TotalCrossWalletBalance) {
		return nil, false
	}
	return o.TotalCrossWalletBalance, true
}

// HasTotalCrossWalletBalance returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTotalCrossWalletBalance() bool {
	if o != nil && !IsNil(o.TotalCrossWalletBalance) {
		return true
	}

	return false
}

// SetTotalCrossWalletBalance gets a reference to the given string and assigns it to the TotalCrossWalletBalance field.
func (o *UfuturesGetAccountV2Resp) SetTotalCrossWalletBalance(v string) {
	o.TotalCrossWalletBalance = &v
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTotalInitialMargin() string {
	if o == nil || IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTotalInitialMargin() bool {
	if o != nil && !IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *UfuturesGetAccountV2Resp) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintMargin returns the TotalMaintMargin field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTotalMaintMargin() string {
	if o == nil || IsNil(o.TotalMaintMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintMargin
}

// GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTotalMaintMarginOk() (*string, bool) {
	if o == nil || IsNil(o.TotalMaintMargin) {
		return nil, false
	}
	return o.TotalMaintMargin, true
}

// HasTotalMaintMargin returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTotalMaintMargin() bool {
	if o != nil && !IsNil(o.TotalMaintMargin) {
		return true
	}

	return false
}

// SetTotalMaintMargin gets a reference to the given string and assigns it to the TotalMaintMargin field.
func (o *UfuturesGetAccountV2Resp) SetTotalMaintMargin(v string) {
	o.TotalMaintMargin = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTotalMarginBalance() string {
	if o == nil || IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTotalMarginBalance() bool {
	if o != nil && !IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *UfuturesGetAccountV2Resp) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTotalOpenOrderInitialMargin() string {
	if o == nil || IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *UfuturesGetAccountV2Resp) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTotalPositionInitialMargin() string {
	if o == nil || IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTotalPositionInitialMargin() bool {
	if o != nil && !IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *UfuturesGetAccountV2Resp) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTotalUnrealizedProfit() string {
	if o == nil || IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTotalUnrealizedProfit() bool {
	if o != nil && !IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *UfuturesGetAccountV2Resp) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTotalWalletBalance() string {
	if o == nil || IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTotalWalletBalance() bool {
	if o != nil && !IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *UfuturesGetAccountV2Resp) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetTradeGroupId returns the TradeGroupId field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetTradeGroupId() int64 {
	if o == nil || IsNil(o.TradeGroupId) {
		var ret int64
		return ret
	}
	return *o.TradeGroupId
}

// GetTradeGroupIdOk returns a tuple with the TradeGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetTradeGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TradeGroupId) {
		return nil, false
	}
	return o.TradeGroupId, true
}

// HasTradeGroupId returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasTradeGroupId() bool {
	if o != nil && !IsNil(o.TradeGroupId) {
		return true
	}

	return false
}

// SetTradeGroupId gets a reference to the given int64 and assigns it to the TradeGroupId field.
func (o *UfuturesGetAccountV2Resp) SetTradeGroupId(v int64) {
	o.TradeGroupId = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *UfuturesGetAccountV2Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAccountV2Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *UfuturesGetAccountV2Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *UfuturesGetAccountV2Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o UfuturesGetAccountV2Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetAccountV2Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.AvailableBalance) {
		toSerialize["availableBalance"] = o.AvailableBalance
	}
	if !IsNil(o.CanDeposit) {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if !IsNil(o.CanWithdraw) {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if !IsNil(o.FeeBurn) {
		toSerialize["feeBurn"] = o.FeeBurn
	}
	if !IsNil(o.FeeTier) {
		toSerialize["feeTier"] = o.FeeTier
	}
	if !IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}
	if !IsNil(o.MultiAssetsMargin) {
		toSerialize["multiAssetsMargin"] = o.MultiAssetsMargin
	}
	if !IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}
	if !IsNil(o.TotalCrossUnPnl) {
		toSerialize["totalCrossUnPnl"] = o.TotalCrossUnPnl
	}
	if !IsNil(o.TotalCrossWalletBalance) {
		toSerialize["totalCrossWalletBalance"] = o.TotalCrossWalletBalance
	}
	if !IsNil(o.TotalInitialMargin) {
		toSerialize["totalInitialMargin"] = o.TotalInitialMargin
	}
	if !IsNil(o.TotalMaintMargin) {
		toSerialize["totalMaintMargin"] = o.TotalMaintMargin
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
	if !IsNil(o.TradeGroupId) {
		toSerialize["tradeGroupId"] = o.TradeGroupId
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableUfuturesGetAccountV2Resp struct {
	value *UfuturesGetAccountV2Resp
	isSet bool
}

func (v NullableUfuturesGetAccountV2Resp) Get() *UfuturesGetAccountV2Resp {
	return v.value
}

func (v *NullableUfuturesGetAccountV2Resp) Set(val *UfuturesGetAccountV2Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetAccountV2Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetAccountV2Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetAccountV2Resp(val *UfuturesGetAccountV2Resp) *NullableUfuturesGetAccountV2Resp {
	return &NullableUfuturesGetAccountV2Resp{value: val, isSet: true}
}

func (v NullableUfuturesGetAccountV2Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetAccountV2Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


