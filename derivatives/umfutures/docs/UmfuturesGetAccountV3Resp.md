# UmfuturesGetAccountV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]UmfuturesGetAccountV3RespAssetsInner**](UmfuturesGetAccountV3RespAssetsInner.md) |  | [optional] 
**AvailableBalance** | Pointer to **string** |  | [optional] 
**MaxWithdrawAmount** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to [**[]UmfuturesGetAccountV3RespPositionsInner**](UmfuturesGetAccountV3RespPositionsInner.md) |  | [optional] 
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

### NewUmfuturesGetAccountV3Resp

`func NewUmfuturesGetAccountV3Resp() *UmfuturesGetAccountV3Resp`

NewUmfuturesGetAccountV3Resp instantiates a new UmfuturesGetAccountV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesGetAccountV3RespWithDefaults

`func NewUmfuturesGetAccountV3RespWithDefaults() *UmfuturesGetAccountV3Resp`

NewUmfuturesGetAccountV3RespWithDefaults instantiates a new UmfuturesGetAccountV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *UmfuturesGetAccountV3Resp) GetAssets() []UmfuturesGetAccountV3RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *UmfuturesGetAccountV3Resp) GetAssetsOk() (*[]UmfuturesGetAccountV3RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *UmfuturesGetAccountV3Resp) SetAssets(v []UmfuturesGetAccountV3RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *UmfuturesGetAccountV3Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *UmfuturesGetAccountV3Resp) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *UmfuturesGetAccountV3Resp) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *UmfuturesGetAccountV3Resp) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *UmfuturesGetAccountV3Resp) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *UmfuturesGetAccountV3Resp) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *UmfuturesGetAccountV3Resp) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *UmfuturesGetAccountV3Resp) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *UmfuturesGetAccountV3Resp) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetPositions

`func (o *UmfuturesGetAccountV3Resp) GetPositions() []UmfuturesGetAccountV3RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *UmfuturesGetAccountV3Resp) GetPositionsOk() (*[]UmfuturesGetAccountV3RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *UmfuturesGetAccountV3Resp) SetPositions(v []UmfuturesGetAccountV3RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *UmfuturesGetAccountV3Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetTotalCrossUnPnl

`func (o *UmfuturesGetAccountV3Resp) GetTotalCrossUnPnl() string`

GetTotalCrossUnPnl returns the TotalCrossUnPnl field if non-nil, zero value otherwise.

### GetTotalCrossUnPnlOk

`func (o *UmfuturesGetAccountV3Resp) GetTotalCrossUnPnlOk() (*string, bool)`

GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossUnPnl

`func (o *UmfuturesGetAccountV3Resp) SetTotalCrossUnPnl(v string)`

SetTotalCrossUnPnl sets TotalCrossUnPnl field to given value.

### HasTotalCrossUnPnl

`func (o *UmfuturesGetAccountV3Resp) HasTotalCrossUnPnl() bool`

HasTotalCrossUnPnl returns a boolean if a field has been set.

### GetTotalCrossWalletBalance

`func (o *UmfuturesGetAccountV3Resp) GetTotalCrossWalletBalance() string`

GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field if non-nil, zero value otherwise.

### GetTotalCrossWalletBalanceOk

`func (o *UmfuturesGetAccountV3Resp) GetTotalCrossWalletBalanceOk() (*string, bool)`

GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossWalletBalance

`func (o *UmfuturesGetAccountV3Resp) SetTotalCrossWalletBalance(v string)`

SetTotalCrossWalletBalance sets TotalCrossWalletBalance field to given value.

### HasTotalCrossWalletBalance

`func (o *UmfuturesGetAccountV3Resp) HasTotalCrossWalletBalance() bool`

HasTotalCrossWalletBalance returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *UmfuturesGetAccountV3Resp) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *UmfuturesGetAccountV3Resp) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *UmfuturesGetAccountV3Resp) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *UmfuturesGetAccountV3Resp) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintMargin

`func (o *UmfuturesGetAccountV3Resp) GetTotalMaintMargin() string`

GetTotalMaintMargin returns the TotalMaintMargin field if non-nil, zero value otherwise.

### GetTotalMaintMarginOk

`func (o *UmfuturesGetAccountV3Resp) GetTotalMaintMarginOk() (*string, bool)`

GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintMargin

`func (o *UmfuturesGetAccountV3Resp) SetTotalMaintMargin(v string)`

SetTotalMaintMargin sets TotalMaintMargin field to given value.

### HasTotalMaintMargin

`func (o *UmfuturesGetAccountV3Resp) HasTotalMaintMargin() bool`

HasTotalMaintMargin returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *UmfuturesGetAccountV3Resp) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *UmfuturesGetAccountV3Resp) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *UmfuturesGetAccountV3Resp) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *UmfuturesGetAccountV3Resp) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *UmfuturesGetAccountV3Resp) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *UmfuturesGetAccountV3Resp) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *UmfuturesGetAccountV3Resp) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *UmfuturesGetAccountV3Resp) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *UmfuturesGetAccountV3Resp) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *UmfuturesGetAccountV3Resp) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *UmfuturesGetAccountV3Resp) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *UmfuturesGetAccountV3Resp) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *UmfuturesGetAccountV3Resp) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *UmfuturesGetAccountV3Resp) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *UmfuturesGetAccountV3Resp) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *UmfuturesGetAccountV3Resp) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *UmfuturesGetAccountV3Resp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *UmfuturesGetAccountV3Resp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *UmfuturesGetAccountV3Resp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *UmfuturesGetAccountV3Resp) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


