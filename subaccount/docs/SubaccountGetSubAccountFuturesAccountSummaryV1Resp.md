# SubaccountGetSubAccountFuturesAccountSummaryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** |  | [optional] 
**SubAccountList** | Pointer to [**[]SubaccountGetSubAccountFuturesAccountSummaryV1RespSubAccountListInner**](SubaccountGetSubAccountFuturesAccountSummaryV1RespSubAccountListInner.md) |  | [optional] 
**TotalInitialMargin** | Pointer to **string** |  | [optional] 
**TotalMaintenanceMargin** | Pointer to **string** |  | [optional] 
**TotalMarginBalance** | Pointer to **string** |  | [optional] 
**TotalOpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**TotalPositionInitialMargin** | Pointer to **string** |  | [optional] 
**TotalUnrealizedProfit** | Pointer to **string** |  | [optional] 
**TotalWalletBalance** | Pointer to **string** |  | [optional] 

## Methods

### NewSubaccountGetSubAccountFuturesAccountSummaryV1Resp

`func NewSubaccountGetSubAccountFuturesAccountSummaryV1Resp() *SubaccountGetSubAccountFuturesAccountSummaryV1Resp`

NewSubaccountGetSubAccountFuturesAccountSummaryV1Resp instantiates a new SubaccountGetSubAccountFuturesAccountSummaryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubaccountGetSubAccountFuturesAccountSummaryV1RespWithDefaults

`func NewSubaccountGetSubAccountFuturesAccountSummaryV1RespWithDefaults() *SubaccountGetSubAccountFuturesAccountSummaryV1Resp`

NewSubaccountGetSubAccountFuturesAccountSummaryV1RespWithDefaults instantiates a new SubaccountGetSubAccountFuturesAccountSummaryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetSubAccountList

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetSubAccountList() []SubaccountGetSubAccountFuturesAccountSummaryV1RespSubAccountListInner`

GetSubAccountList returns the SubAccountList field if non-nil, zero value otherwise.

### GetSubAccountListOk

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetSubAccountListOk() (*[]SubaccountGetSubAccountFuturesAccountSummaryV1RespSubAccountListInner, bool)`

GetSubAccountListOk returns a tuple with the SubAccountList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountList

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) SetSubAccountList(v []SubaccountGetSubAccountFuturesAccountSummaryV1RespSubAccountListInner)`

SetSubAccountList sets SubAccountList field to given value.

### HasSubAccountList

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) HasSubAccountList() bool`

HasSubAccountList returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintenanceMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalMaintenanceMargin() string`

GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field if non-nil, zero value otherwise.

### GetTotalMaintenanceMarginOk

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalMaintenanceMarginOk() (*string, bool)`

GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintenanceMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) SetTotalMaintenanceMargin(v string)`

SetTotalMaintenanceMargin sets TotalMaintenanceMargin field to given value.

### HasTotalMaintenanceMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) HasTotalMaintenanceMargin() bool`

HasTotalMaintenanceMargin returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *SubaccountGetSubAccountFuturesAccountSummaryV1Resp) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


