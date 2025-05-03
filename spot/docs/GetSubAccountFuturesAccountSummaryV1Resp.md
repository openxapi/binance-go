# GetSubAccountFuturesAccountSummaryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** |  | [optional] 
**SubAccountList** | Pointer to [**[]GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner**](GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner.md) |  | [optional] 
**TotalInitialMargin** | Pointer to **string** |  | [optional] 
**TotalMaintenanceMargin** | Pointer to **string** |  | [optional] 
**TotalMarginBalance** | Pointer to **string** |  | [optional] 
**TotalOpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**TotalPositionInitialMargin** | Pointer to **string** |  | [optional] 
**TotalUnrealizedProfit** | Pointer to **string** |  | [optional] 
**TotalWalletBalance** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSubAccountFuturesAccountSummaryV1Resp

`func NewGetSubAccountFuturesAccountSummaryV1Resp() *GetSubAccountFuturesAccountSummaryV1Resp`

NewGetSubAccountFuturesAccountSummaryV1Resp instantiates a new GetSubAccountFuturesAccountSummaryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubAccountFuturesAccountSummaryV1RespWithDefaults

`func NewGetSubAccountFuturesAccountSummaryV1RespWithDefaults() *GetSubAccountFuturesAccountSummaryV1Resp`

NewGetSubAccountFuturesAccountSummaryV1RespWithDefaults instantiates a new GetSubAccountFuturesAccountSummaryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetSubAccountList

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetSubAccountList() []GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner`

GetSubAccountList returns the SubAccountList field if non-nil, zero value otherwise.

### GetSubAccountListOk

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetSubAccountListOk() (*[]GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner, bool)`

GetSubAccountListOk returns a tuple with the SubAccountList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountList

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) SetSubAccountList(v []GetSubAccountFuturesAccountSummaryV1RespSubAccountListInner)`

SetSubAccountList sets SubAccountList field to given value.

### HasSubAccountList

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) HasSubAccountList() bool`

HasSubAccountList returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintenanceMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalMaintenanceMargin() string`

GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field if non-nil, zero value otherwise.

### GetTotalMaintenanceMarginOk

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalMaintenanceMarginOk() (*string, bool)`

GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintenanceMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) SetTotalMaintenanceMargin(v string)`

SetTotalMaintenanceMargin sets TotalMaintenanceMargin field to given value.

### HasTotalMaintenanceMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) HasTotalMaintenanceMargin() bool`

HasTotalMaintenanceMargin returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *GetSubAccountFuturesAccountSummaryV1Resp) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


