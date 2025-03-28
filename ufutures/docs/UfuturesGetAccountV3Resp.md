# UfuturesGetAccountV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]UfuturesGetAccountV3RespAssetsInner**](UfuturesGetAccountV3RespAssetsInner.md) |  | [optional] 
**AvailableBalance** | Pointer to **string** |  | [optional] 
**MaxWithdrawAmount** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to [**[]UfuturesGetAccountV3RespPositionsInner**](UfuturesGetAccountV3RespPositionsInner.md) |  | [optional] 
**TotalCrossUnPnl** | Pointer to **string** |  | [optional] 
**TotalCrossWalletBalance** | Pointer to **string** |  | [optional] 
**TotalInitialMargin** | Pointer to **string** |  | [optional] 
**TotalMaintMargin** | Pointer to **string** |  | [optional] 
**TotalMarginBalance** | Pointer to **string** |  | [optional] 
**TotalOpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**TotalPositionInitialMargin** | Pointer to **string** |  | [optional] 
**TotalUnrealizedProfit** | Pointer to **string** |  | [optional] 
**TotalWalletBalance** | Pointer to **string** |  | [optional] 

## Methods

### NewUfuturesGetAccountV3Resp

`func NewUfuturesGetAccountV3Resp() *UfuturesGetAccountV3Resp`

NewUfuturesGetAccountV3Resp instantiates a new UfuturesGetAccountV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesGetAccountV3RespWithDefaults

`func NewUfuturesGetAccountV3RespWithDefaults() *UfuturesGetAccountV3Resp`

NewUfuturesGetAccountV3RespWithDefaults instantiates a new UfuturesGetAccountV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *UfuturesGetAccountV3Resp) GetAssets() []UfuturesGetAccountV3RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *UfuturesGetAccountV3Resp) GetAssetsOk() (*[]UfuturesGetAccountV3RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *UfuturesGetAccountV3Resp) SetAssets(v []UfuturesGetAccountV3RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *UfuturesGetAccountV3Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *UfuturesGetAccountV3Resp) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *UfuturesGetAccountV3Resp) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *UfuturesGetAccountV3Resp) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *UfuturesGetAccountV3Resp) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *UfuturesGetAccountV3Resp) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *UfuturesGetAccountV3Resp) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *UfuturesGetAccountV3Resp) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *UfuturesGetAccountV3Resp) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetPositions

`func (o *UfuturesGetAccountV3Resp) GetPositions() []UfuturesGetAccountV3RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *UfuturesGetAccountV3Resp) GetPositionsOk() (*[]UfuturesGetAccountV3RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *UfuturesGetAccountV3Resp) SetPositions(v []UfuturesGetAccountV3RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *UfuturesGetAccountV3Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetTotalCrossUnPnl

`func (o *UfuturesGetAccountV3Resp) GetTotalCrossUnPnl() string`

GetTotalCrossUnPnl returns the TotalCrossUnPnl field if non-nil, zero value otherwise.

### GetTotalCrossUnPnlOk

`func (o *UfuturesGetAccountV3Resp) GetTotalCrossUnPnlOk() (*string, bool)`

GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossUnPnl

`func (o *UfuturesGetAccountV3Resp) SetTotalCrossUnPnl(v string)`

SetTotalCrossUnPnl sets TotalCrossUnPnl field to given value.

### HasTotalCrossUnPnl

`func (o *UfuturesGetAccountV3Resp) HasTotalCrossUnPnl() bool`

HasTotalCrossUnPnl returns a boolean if a field has been set.

### GetTotalCrossWalletBalance

`func (o *UfuturesGetAccountV3Resp) GetTotalCrossWalletBalance() string`

GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field if non-nil, zero value otherwise.

### GetTotalCrossWalletBalanceOk

`func (o *UfuturesGetAccountV3Resp) GetTotalCrossWalletBalanceOk() (*string, bool)`

GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossWalletBalance

`func (o *UfuturesGetAccountV3Resp) SetTotalCrossWalletBalance(v string)`

SetTotalCrossWalletBalance sets TotalCrossWalletBalance field to given value.

### HasTotalCrossWalletBalance

`func (o *UfuturesGetAccountV3Resp) HasTotalCrossWalletBalance() bool`

HasTotalCrossWalletBalance returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *UfuturesGetAccountV3Resp) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *UfuturesGetAccountV3Resp) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *UfuturesGetAccountV3Resp) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *UfuturesGetAccountV3Resp) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintMargin

`func (o *UfuturesGetAccountV3Resp) GetTotalMaintMargin() string`

GetTotalMaintMargin returns the TotalMaintMargin field if non-nil, zero value otherwise.

### GetTotalMaintMarginOk

`func (o *UfuturesGetAccountV3Resp) GetTotalMaintMarginOk() (*string, bool)`

GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintMargin

`func (o *UfuturesGetAccountV3Resp) SetTotalMaintMargin(v string)`

SetTotalMaintMargin sets TotalMaintMargin field to given value.

### HasTotalMaintMargin

`func (o *UfuturesGetAccountV3Resp) HasTotalMaintMargin() bool`

HasTotalMaintMargin returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *UfuturesGetAccountV3Resp) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *UfuturesGetAccountV3Resp) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *UfuturesGetAccountV3Resp) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *UfuturesGetAccountV3Resp) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *UfuturesGetAccountV3Resp) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *UfuturesGetAccountV3Resp) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *UfuturesGetAccountV3Resp) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *UfuturesGetAccountV3Resp) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *UfuturesGetAccountV3Resp) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *UfuturesGetAccountV3Resp) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *UfuturesGetAccountV3Resp) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *UfuturesGetAccountV3Resp) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *UfuturesGetAccountV3Resp) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *UfuturesGetAccountV3Resp) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *UfuturesGetAccountV3Resp) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *UfuturesGetAccountV3Resp) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *UfuturesGetAccountV3Resp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *UfuturesGetAccountV3Resp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *UfuturesGetAccountV3Resp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *UfuturesGetAccountV3Resp) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


