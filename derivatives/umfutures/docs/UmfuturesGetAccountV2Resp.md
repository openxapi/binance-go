# UmfuturesGetAccountV2Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]UmfuturesGetAccountV2RespAssetsInner**](UmfuturesGetAccountV2RespAssetsInner.md) |  | [optional] 
**AvailableBalance** | Pointer to **string** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**FeeBurn** | Pointer to **bool** |  | [optional] 
**FeeTier** | Pointer to **int32** |  | [optional] 
**MaxWithdrawAmount** | Pointer to **string** |  | [optional] 
**MultiAssetsMargin** | Pointer to **bool** |  | [optional] 
**Positions** | Pointer to [**[]UmfuturesGetAccountV2RespPositionsInner**](UmfuturesGetAccountV2RespPositionsInner.md) |  | [optional] 
**TotalCrossUnPnl** | Pointer to **string** |  | [optional] 
**TotalCrossWalletBalance** | Pointer to **string** |  | [optional] 
**TotalInitialMargin** | Pointer to **string** |  | [optional] 
**TotalMaintMargin** | Pointer to **string** |  | [optional] 
**TotalMarginBalance** | Pointer to **string** |  | [optional] 
**TotalOpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**TotalPositionInitialMargin** | Pointer to **string** |  | [optional] 
**TotalUnrealizedProfit** | Pointer to **string** |  | [optional] 
**TotalWalletBalance** | Pointer to **string** |  | [optional] 
**TradeGroupId** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewUmfuturesGetAccountV2Resp

`func NewUmfuturesGetAccountV2Resp() *UmfuturesGetAccountV2Resp`

NewUmfuturesGetAccountV2Resp instantiates a new UmfuturesGetAccountV2Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesGetAccountV2RespWithDefaults

`func NewUmfuturesGetAccountV2RespWithDefaults() *UmfuturesGetAccountV2Resp`

NewUmfuturesGetAccountV2RespWithDefaults instantiates a new UmfuturesGetAccountV2Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *UmfuturesGetAccountV2Resp) GetAssets() []UmfuturesGetAccountV2RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *UmfuturesGetAccountV2Resp) GetAssetsOk() (*[]UmfuturesGetAccountV2RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *UmfuturesGetAccountV2Resp) SetAssets(v []UmfuturesGetAccountV2RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *UmfuturesGetAccountV2Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *UmfuturesGetAccountV2Resp) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *UmfuturesGetAccountV2Resp) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *UmfuturesGetAccountV2Resp) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *UmfuturesGetAccountV2Resp) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetCanDeposit

`func (o *UmfuturesGetAccountV2Resp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *UmfuturesGetAccountV2Resp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *UmfuturesGetAccountV2Resp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *UmfuturesGetAccountV2Resp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *UmfuturesGetAccountV2Resp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *UmfuturesGetAccountV2Resp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *UmfuturesGetAccountV2Resp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *UmfuturesGetAccountV2Resp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetFeeBurn

`func (o *UmfuturesGetAccountV2Resp) GetFeeBurn() bool`

GetFeeBurn returns the FeeBurn field if non-nil, zero value otherwise.

### GetFeeBurnOk

`func (o *UmfuturesGetAccountV2Resp) GetFeeBurnOk() (*bool, bool)`

GetFeeBurnOk returns a tuple with the FeeBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBurn

`func (o *UmfuturesGetAccountV2Resp) SetFeeBurn(v bool)`

SetFeeBurn sets FeeBurn field to given value.

### HasFeeBurn

`func (o *UmfuturesGetAccountV2Resp) HasFeeBurn() bool`

HasFeeBurn returns a boolean if a field has been set.

### GetFeeTier

`func (o *UmfuturesGetAccountV2Resp) GetFeeTier() int32`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *UmfuturesGetAccountV2Resp) GetFeeTierOk() (*int32, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *UmfuturesGetAccountV2Resp) SetFeeTier(v int32)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *UmfuturesGetAccountV2Resp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *UmfuturesGetAccountV2Resp) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *UmfuturesGetAccountV2Resp) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *UmfuturesGetAccountV2Resp) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *UmfuturesGetAccountV2Resp) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetMultiAssetsMargin

`func (o *UmfuturesGetAccountV2Resp) GetMultiAssetsMargin() bool`

GetMultiAssetsMargin returns the MultiAssetsMargin field if non-nil, zero value otherwise.

### GetMultiAssetsMarginOk

`func (o *UmfuturesGetAccountV2Resp) GetMultiAssetsMarginOk() (*bool, bool)`

GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAssetsMargin

`func (o *UmfuturesGetAccountV2Resp) SetMultiAssetsMargin(v bool)`

SetMultiAssetsMargin sets MultiAssetsMargin field to given value.

### HasMultiAssetsMargin

`func (o *UmfuturesGetAccountV2Resp) HasMultiAssetsMargin() bool`

HasMultiAssetsMargin returns a boolean if a field has been set.

### GetPositions

`func (o *UmfuturesGetAccountV2Resp) GetPositions() []UmfuturesGetAccountV2RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *UmfuturesGetAccountV2Resp) GetPositionsOk() (*[]UmfuturesGetAccountV2RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *UmfuturesGetAccountV2Resp) SetPositions(v []UmfuturesGetAccountV2RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *UmfuturesGetAccountV2Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetTotalCrossUnPnl

`func (o *UmfuturesGetAccountV2Resp) GetTotalCrossUnPnl() string`

GetTotalCrossUnPnl returns the TotalCrossUnPnl field if non-nil, zero value otherwise.

### GetTotalCrossUnPnlOk

`func (o *UmfuturesGetAccountV2Resp) GetTotalCrossUnPnlOk() (*string, bool)`

GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossUnPnl

`func (o *UmfuturesGetAccountV2Resp) SetTotalCrossUnPnl(v string)`

SetTotalCrossUnPnl sets TotalCrossUnPnl field to given value.

### HasTotalCrossUnPnl

`func (o *UmfuturesGetAccountV2Resp) HasTotalCrossUnPnl() bool`

HasTotalCrossUnPnl returns a boolean if a field has been set.

### GetTotalCrossWalletBalance

`func (o *UmfuturesGetAccountV2Resp) GetTotalCrossWalletBalance() string`

GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field if non-nil, zero value otherwise.

### GetTotalCrossWalletBalanceOk

`func (o *UmfuturesGetAccountV2Resp) GetTotalCrossWalletBalanceOk() (*string, bool)`

GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossWalletBalance

`func (o *UmfuturesGetAccountV2Resp) SetTotalCrossWalletBalance(v string)`

SetTotalCrossWalletBalance sets TotalCrossWalletBalance field to given value.

### HasTotalCrossWalletBalance

`func (o *UmfuturesGetAccountV2Resp) HasTotalCrossWalletBalance() bool`

HasTotalCrossWalletBalance returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *UmfuturesGetAccountV2Resp) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *UmfuturesGetAccountV2Resp) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *UmfuturesGetAccountV2Resp) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *UmfuturesGetAccountV2Resp) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintMargin

`func (o *UmfuturesGetAccountV2Resp) GetTotalMaintMargin() string`

GetTotalMaintMargin returns the TotalMaintMargin field if non-nil, zero value otherwise.

### GetTotalMaintMarginOk

`func (o *UmfuturesGetAccountV2Resp) GetTotalMaintMarginOk() (*string, bool)`

GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintMargin

`func (o *UmfuturesGetAccountV2Resp) SetTotalMaintMargin(v string)`

SetTotalMaintMargin sets TotalMaintMargin field to given value.

### HasTotalMaintMargin

`func (o *UmfuturesGetAccountV2Resp) HasTotalMaintMargin() bool`

HasTotalMaintMargin returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *UmfuturesGetAccountV2Resp) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *UmfuturesGetAccountV2Resp) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *UmfuturesGetAccountV2Resp) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *UmfuturesGetAccountV2Resp) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *UmfuturesGetAccountV2Resp) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *UmfuturesGetAccountV2Resp) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *UmfuturesGetAccountV2Resp) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *UmfuturesGetAccountV2Resp) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *UmfuturesGetAccountV2Resp) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *UmfuturesGetAccountV2Resp) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *UmfuturesGetAccountV2Resp) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *UmfuturesGetAccountV2Resp) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *UmfuturesGetAccountV2Resp) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *UmfuturesGetAccountV2Resp) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *UmfuturesGetAccountV2Resp) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *UmfuturesGetAccountV2Resp) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *UmfuturesGetAccountV2Resp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *UmfuturesGetAccountV2Resp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *UmfuturesGetAccountV2Resp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *UmfuturesGetAccountV2Resp) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.

### GetTradeGroupId

`func (o *UmfuturesGetAccountV2Resp) GetTradeGroupId() int64`

GetTradeGroupId returns the TradeGroupId field if non-nil, zero value otherwise.

### GetTradeGroupIdOk

`func (o *UmfuturesGetAccountV2Resp) GetTradeGroupIdOk() (*int64, bool)`

GetTradeGroupIdOk returns a tuple with the TradeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeGroupId

`func (o *UmfuturesGetAccountV2Resp) SetTradeGroupId(v int64)`

SetTradeGroupId sets TradeGroupId field to given value.

### HasTradeGroupId

`func (o *UmfuturesGetAccountV2Resp) HasTradeGroupId() bool`

HasTradeGroupId returns a boolean if a field has been set.

### GetUpdateTime

`func (o *UmfuturesGetAccountV2Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *UmfuturesGetAccountV2Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *UmfuturesGetAccountV2Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *UmfuturesGetAccountV2Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


