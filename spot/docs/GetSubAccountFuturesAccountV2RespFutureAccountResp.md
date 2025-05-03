# GetSubAccountFuturesAccountV2RespFutureAccountResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]GetSubAccountFuturesAccountV1RespAssetsInner**](GetSubAccountFuturesAccountV1RespAssetsInner.md) |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FeeTier** | Pointer to **int32** |  | [optional] 
**MaxWithdrawAmount** | Pointer to **string** |  | [optional] 
**TotalInitialMargin** | Pointer to **string** |  | [optional] 
**TotalMaintenanceMargin** | Pointer to **string** |  | [optional] 
**TotalMarginBalance** | Pointer to **string** |  | [optional] 
**TotalOpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**TotalPositionInitialMargin** | Pointer to **string** |  | [optional] 
**TotalUnrealizedProfit** | Pointer to **string** |  | [optional] 
**TotalWalletBalance** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetSubAccountFuturesAccountV2RespFutureAccountResp

`func NewGetSubAccountFuturesAccountV2RespFutureAccountResp() *GetSubAccountFuturesAccountV2RespFutureAccountResp`

NewGetSubAccountFuturesAccountV2RespFutureAccountResp instantiates a new GetSubAccountFuturesAccountV2RespFutureAccountResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubAccountFuturesAccountV2RespFutureAccountRespWithDefaults

`func NewGetSubAccountFuturesAccountV2RespFutureAccountRespWithDefaults() *GetSubAccountFuturesAccountV2RespFutureAccountResp`

NewGetSubAccountFuturesAccountV2RespFutureAccountRespWithDefaults instantiates a new GetSubAccountFuturesAccountV2RespFutureAccountResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetAssets() []GetSubAccountFuturesAccountV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetAssetsOk() (*[]GetSubAccountFuturesAccountV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetAssets(v []GetSubAccountFuturesAccountV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCanDeposit

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetEmail

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFeeTier

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetFeeTier() int32`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetFeeTierOk() (*int32, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetFeeTier(v int32)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintenanceMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalMaintenanceMargin() string`

GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field if non-nil, zero value otherwise.

### GetTotalMaintenanceMarginOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalMaintenanceMarginOk() (*string, bool)`

GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintenanceMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetTotalMaintenanceMargin(v string)`

SetTotalMaintenanceMargin sets TotalMaintenanceMargin field to given value.

### HasTotalMaintenanceMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasTotalMaintenanceMargin() bool`

HasTotalMaintenanceMargin returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetSubAccountFuturesAccountV2RespFutureAccountResp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


