# GetAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountEquity** | Pointer to **string** |  | [optional] 
**AccountInitialMargin** | Pointer to **string** |  | [optional] 
**AccountMaintMargin** | Pointer to **string** |  | [optional] 
**AccountStatus** | Pointer to **string** |  | [optional] 
**ActualEquity** | Pointer to **string** |  | [optional] 
**TotalAvailableBalance** | Pointer to **string** |  | [optional] 
**TotalMarginOpenLoss** | Pointer to **string** |  | [optional] 
**UniMMR** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**VirtualMaxWithdrawAmount** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAccountV1Resp

`func NewGetAccountV1Resp() *GetAccountV1Resp`

NewGetAccountV1Resp instantiates a new GetAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountV1RespWithDefaults

`func NewGetAccountV1RespWithDefaults() *GetAccountV1Resp`

NewGetAccountV1RespWithDefaults instantiates a new GetAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountEquity

`func (o *GetAccountV1Resp) GetAccountEquity() string`

GetAccountEquity returns the AccountEquity field if non-nil, zero value otherwise.

### GetAccountEquityOk

`func (o *GetAccountV1Resp) GetAccountEquityOk() (*string, bool)`

GetAccountEquityOk returns a tuple with the AccountEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEquity

`func (o *GetAccountV1Resp) SetAccountEquity(v string)`

SetAccountEquity sets AccountEquity field to given value.

### HasAccountEquity

`func (o *GetAccountV1Resp) HasAccountEquity() bool`

HasAccountEquity returns a boolean if a field has been set.

### GetAccountInitialMargin

`func (o *GetAccountV1Resp) GetAccountInitialMargin() string`

GetAccountInitialMargin returns the AccountInitialMargin field if non-nil, zero value otherwise.

### GetAccountInitialMarginOk

`func (o *GetAccountV1Resp) GetAccountInitialMarginOk() (*string, bool)`

GetAccountInitialMarginOk returns a tuple with the AccountInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInitialMargin

`func (o *GetAccountV1Resp) SetAccountInitialMargin(v string)`

SetAccountInitialMargin sets AccountInitialMargin field to given value.

### HasAccountInitialMargin

`func (o *GetAccountV1Resp) HasAccountInitialMargin() bool`

HasAccountInitialMargin returns a boolean if a field has been set.

### GetAccountMaintMargin

`func (o *GetAccountV1Resp) GetAccountMaintMargin() string`

GetAccountMaintMargin returns the AccountMaintMargin field if non-nil, zero value otherwise.

### GetAccountMaintMarginOk

`func (o *GetAccountV1Resp) GetAccountMaintMarginOk() (*string, bool)`

GetAccountMaintMarginOk returns a tuple with the AccountMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMaintMargin

`func (o *GetAccountV1Resp) SetAccountMaintMargin(v string)`

SetAccountMaintMargin sets AccountMaintMargin field to given value.

### HasAccountMaintMargin

`func (o *GetAccountV1Resp) HasAccountMaintMargin() bool`

HasAccountMaintMargin returns a boolean if a field has been set.

### GetAccountStatus

`func (o *GetAccountV1Resp) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *GetAccountV1Resp) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *GetAccountV1Resp) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *GetAccountV1Resp) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetActualEquity

`func (o *GetAccountV1Resp) GetActualEquity() string`

GetActualEquity returns the ActualEquity field if non-nil, zero value otherwise.

### GetActualEquityOk

`func (o *GetAccountV1Resp) GetActualEquityOk() (*string, bool)`

GetActualEquityOk returns a tuple with the ActualEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEquity

`func (o *GetAccountV1Resp) SetActualEquity(v string)`

SetActualEquity sets ActualEquity field to given value.

### HasActualEquity

`func (o *GetAccountV1Resp) HasActualEquity() bool`

HasActualEquity returns a boolean if a field has been set.

### GetTotalAvailableBalance

`func (o *GetAccountV1Resp) GetTotalAvailableBalance() string`

GetTotalAvailableBalance returns the TotalAvailableBalance field if non-nil, zero value otherwise.

### GetTotalAvailableBalanceOk

`func (o *GetAccountV1Resp) GetTotalAvailableBalanceOk() (*string, bool)`

GetTotalAvailableBalanceOk returns a tuple with the TotalAvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvailableBalance

`func (o *GetAccountV1Resp) SetTotalAvailableBalance(v string)`

SetTotalAvailableBalance sets TotalAvailableBalance field to given value.

### HasTotalAvailableBalance

`func (o *GetAccountV1Resp) HasTotalAvailableBalance() bool`

HasTotalAvailableBalance returns a boolean if a field has been set.

### GetTotalMarginOpenLoss

`func (o *GetAccountV1Resp) GetTotalMarginOpenLoss() string`

GetTotalMarginOpenLoss returns the TotalMarginOpenLoss field if non-nil, zero value otherwise.

### GetTotalMarginOpenLossOk

`func (o *GetAccountV1Resp) GetTotalMarginOpenLossOk() (*string, bool)`

GetTotalMarginOpenLossOk returns a tuple with the TotalMarginOpenLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginOpenLoss

`func (o *GetAccountV1Resp) SetTotalMarginOpenLoss(v string)`

SetTotalMarginOpenLoss sets TotalMarginOpenLoss field to given value.

### HasTotalMarginOpenLoss

`func (o *GetAccountV1Resp) HasTotalMarginOpenLoss() bool`

HasTotalMarginOpenLoss returns a boolean if a field has been set.

### GetUniMMR

`func (o *GetAccountV1Resp) GetUniMMR() string`

GetUniMMR returns the UniMMR field if non-nil, zero value otherwise.

### GetUniMMROk

`func (o *GetAccountV1Resp) GetUniMMROk() (*string, bool)`

GetUniMMROk returns a tuple with the UniMMR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniMMR

`func (o *GetAccountV1Resp) SetUniMMR(v string)`

SetUniMMR sets UniMMR field to given value.

### HasUniMMR

`func (o *GetAccountV1Resp) HasUniMMR() bool`

HasUniMMR returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetAccountV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetAccountV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetAccountV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetAccountV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVirtualMaxWithdrawAmount

`func (o *GetAccountV1Resp) GetVirtualMaxWithdrawAmount() string`

GetVirtualMaxWithdrawAmount returns the VirtualMaxWithdrawAmount field if non-nil, zero value otherwise.

### GetVirtualMaxWithdrawAmountOk

`func (o *GetAccountV1Resp) GetVirtualMaxWithdrawAmountOk() (*string, bool)`

GetVirtualMaxWithdrawAmountOk returns a tuple with the VirtualMaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMaxWithdrawAmount

`func (o *GetAccountV1Resp) SetVirtualMaxWithdrawAmount(v string)`

SetVirtualMaxWithdrawAmount sets VirtualMaxWithdrawAmount field to given value.

### HasVirtualMaxWithdrawAmount

`func (o *GetAccountV1Resp) HasVirtualMaxWithdrawAmount() bool`

HasVirtualMaxWithdrawAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


