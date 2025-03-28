# UfuturesGetAccountV2Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]UfuturesGetAccountV2RespAssetsInner**](UfuturesGetAccountV2RespAssetsInner.md) |  | [optional] 
**AvailableBalance** | Pointer to **string** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**FeeBurn** | Pointer to **bool** |  | [optional] 
**FeeTier** | Pointer to **int32** |  | [optional] 
**MaxWithdrawAmount** | Pointer to **string** |  | [optional] 
**MultiAssetsMargin** | Pointer to **bool** |  | [optional] 
**Positions** | Pointer to [**[]UfuturesGetAccountV2RespPositionsInner**](UfuturesGetAccountV2RespPositionsInner.md) |  | [optional] 
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

### NewUfuturesGetAccountV2Resp

`func NewUfuturesGetAccountV2Resp() *UfuturesGetAccountV2Resp`

NewUfuturesGetAccountV2Resp instantiates a new UfuturesGetAccountV2Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesGetAccountV2RespWithDefaults

`func NewUfuturesGetAccountV2RespWithDefaults() *UfuturesGetAccountV2Resp`

NewUfuturesGetAccountV2RespWithDefaults instantiates a new UfuturesGetAccountV2Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *UfuturesGetAccountV2Resp) GetAssets() []UfuturesGetAccountV2RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *UfuturesGetAccountV2Resp) GetAssetsOk() (*[]UfuturesGetAccountV2RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *UfuturesGetAccountV2Resp) SetAssets(v []UfuturesGetAccountV2RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *UfuturesGetAccountV2Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *UfuturesGetAccountV2Resp) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *UfuturesGetAccountV2Resp) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *UfuturesGetAccountV2Resp) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *UfuturesGetAccountV2Resp) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetCanDeposit

`func (o *UfuturesGetAccountV2Resp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *UfuturesGetAccountV2Resp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *UfuturesGetAccountV2Resp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *UfuturesGetAccountV2Resp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *UfuturesGetAccountV2Resp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *UfuturesGetAccountV2Resp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *UfuturesGetAccountV2Resp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *UfuturesGetAccountV2Resp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetFeeBurn

`func (o *UfuturesGetAccountV2Resp) GetFeeBurn() bool`

GetFeeBurn returns the FeeBurn field if non-nil, zero value otherwise.

### GetFeeBurnOk

`func (o *UfuturesGetAccountV2Resp) GetFeeBurnOk() (*bool, bool)`

GetFeeBurnOk returns a tuple with the FeeBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBurn

`func (o *UfuturesGetAccountV2Resp) SetFeeBurn(v bool)`

SetFeeBurn sets FeeBurn field to given value.

### HasFeeBurn

`func (o *UfuturesGetAccountV2Resp) HasFeeBurn() bool`

HasFeeBurn returns a boolean if a field has been set.

### GetFeeTier

`func (o *UfuturesGetAccountV2Resp) GetFeeTier() int32`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *UfuturesGetAccountV2Resp) GetFeeTierOk() (*int32, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *UfuturesGetAccountV2Resp) SetFeeTier(v int32)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *UfuturesGetAccountV2Resp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *UfuturesGetAccountV2Resp) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *UfuturesGetAccountV2Resp) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *UfuturesGetAccountV2Resp) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *UfuturesGetAccountV2Resp) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetMultiAssetsMargin

`func (o *UfuturesGetAccountV2Resp) GetMultiAssetsMargin() bool`

GetMultiAssetsMargin returns the MultiAssetsMargin field if non-nil, zero value otherwise.

### GetMultiAssetsMarginOk

`func (o *UfuturesGetAccountV2Resp) GetMultiAssetsMarginOk() (*bool, bool)`

GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAssetsMargin

`func (o *UfuturesGetAccountV2Resp) SetMultiAssetsMargin(v bool)`

SetMultiAssetsMargin sets MultiAssetsMargin field to given value.

### HasMultiAssetsMargin

`func (o *UfuturesGetAccountV2Resp) HasMultiAssetsMargin() bool`

HasMultiAssetsMargin returns a boolean if a field has been set.

### GetPositions

`func (o *UfuturesGetAccountV2Resp) GetPositions() []UfuturesGetAccountV2RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *UfuturesGetAccountV2Resp) GetPositionsOk() (*[]UfuturesGetAccountV2RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *UfuturesGetAccountV2Resp) SetPositions(v []UfuturesGetAccountV2RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *UfuturesGetAccountV2Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetTotalCrossUnPnl

`func (o *UfuturesGetAccountV2Resp) GetTotalCrossUnPnl() string`

GetTotalCrossUnPnl returns the TotalCrossUnPnl field if non-nil, zero value otherwise.

### GetTotalCrossUnPnlOk

`func (o *UfuturesGetAccountV2Resp) GetTotalCrossUnPnlOk() (*string, bool)`

GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossUnPnl

`func (o *UfuturesGetAccountV2Resp) SetTotalCrossUnPnl(v string)`

SetTotalCrossUnPnl sets TotalCrossUnPnl field to given value.

### HasTotalCrossUnPnl

`func (o *UfuturesGetAccountV2Resp) HasTotalCrossUnPnl() bool`

HasTotalCrossUnPnl returns a boolean if a field has been set.

### GetTotalCrossWalletBalance

`func (o *UfuturesGetAccountV2Resp) GetTotalCrossWalletBalance() string`

GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field if non-nil, zero value otherwise.

### GetTotalCrossWalletBalanceOk

`func (o *UfuturesGetAccountV2Resp) GetTotalCrossWalletBalanceOk() (*string, bool)`

GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossWalletBalance

`func (o *UfuturesGetAccountV2Resp) SetTotalCrossWalletBalance(v string)`

SetTotalCrossWalletBalance sets TotalCrossWalletBalance field to given value.

### HasTotalCrossWalletBalance

`func (o *UfuturesGetAccountV2Resp) HasTotalCrossWalletBalance() bool`

HasTotalCrossWalletBalance returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *UfuturesGetAccountV2Resp) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *UfuturesGetAccountV2Resp) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *UfuturesGetAccountV2Resp) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *UfuturesGetAccountV2Resp) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintMargin

`func (o *UfuturesGetAccountV2Resp) GetTotalMaintMargin() string`

GetTotalMaintMargin returns the TotalMaintMargin field if non-nil, zero value otherwise.

### GetTotalMaintMarginOk

`func (o *UfuturesGetAccountV2Resp) GetTotalMaintMarginOk() (*string, bool)`

GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintMargin

`func (o *UfuturesGetAccountV2Resp) SetTotalMaintMargin(v string)`

SetTotalMaintMargin sets TotalMaintMargin field to given value.

### HasTotalMaintMargin

`func (o *UfuturesGetAccountV2Resp) HasTotalMaintMargin() bool`

HasTotalMaintMargin returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *UfuturesGetAccountV2Resp) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *UfuturesGetAccountV2Resp) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *UfuturesGetAccountV2Resp) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *UfuturesGetAccountV2Resp) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *UfuturesGetAccountV2Resp) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *UfuturesGetAccountV2Resp) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *UfuturesGetAccountV2Resp) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *UfuturesGetAccountV2Resp) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *UfuturesGetAccountV2Resp) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *UfuturesGetAccountV2Resp) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *UfuturesGetAccountV2Resp) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *UfuturesGetAccountV2Resp) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *UfuturesGetAccountV2Resp) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *UfuturesGetAccountV2Resp) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *UfuturesGetAccountV2Resp) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *UfuturesGetAccountV2Resp) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *UfuturesGetAccountV2Resp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *UfuturesGetAccountV2Resp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *UfuturesGetAccountV2Resp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *UfuturesGetAccountV2Resp) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.

### GetTradeGroupId

`func (o *UfuturesGetAccountV2Resp) GetTradeGroupId() int64`

GetTradeGroupId returns the TradeGroupId field if non-nil, zero value otherwise.

### GetTradeGroupIdOk

`func (o *UfuturesGetAccountV2Resp) GetTradeGroupIdOk() (*int64, bool)`

GetTradeGroupIdOk returns a tuple with the TradeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeGroupId

`func (o *UfuturesGetAccountV2Resp) SetTradeGroupId(v int64)`

SetTradeGroupId sets TradeGroupId field to given value.

### HasTradeGroupId

`func (o *UfuturesGetAccountV2Resp) HasTradeGroupId() bool`

HasTradeGroupId returns a boolean if a field has been set.

### GetUpdateTime

`func (o *UfuturesGetAccountV2Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *UfuturesGetAccountV2Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *UfuturesGetAccountV2Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *UfuturesGetAccountV2Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


