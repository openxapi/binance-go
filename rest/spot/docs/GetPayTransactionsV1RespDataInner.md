# GetPayTransactionsV1RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**FundsDetail** | Pointer to [**[]GetPayTransactionsV1RespDataInnerFundsDetailInner**](GetPayTransactionsV1RespDataInnerFundsDetailInner.md) |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**PayerInfo** | Pointer to [**GetPayTransactionsV1RespDataInnerPayerInfo**](GetPayTransactionsV1RespDataInnerPayerInfo.md) |  | [optional] 
**ReceiverInfo** | Pointer to [**GetPayTransactionsV1RespDataInnerReceiverInfo**](GetPayTransactionsV1RespDataInnerReceiverInfo.md) |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 
**WalletType** | Pointer to **int32** |  | [optional] 
**WalletTypes** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewGetPayTransactionsV1RespDataInner

`func NewGetPayTransactionsV1RespDataInner() *GetPayTransactionsV1RespDataInner`

NewGetPayTransactionsV1RespDataInner instantiates a new GetPayTransactionsV1RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayTransactionsV1RespDataInnerWithDefaults

`func NewGetPayTransactionsV1RespDataInnerWithDefaults() *GetPayTransactionsV1RespDataInner`

NewGetPayTransactionsV1RespDataInnerWithDefaults instantiates a new GetPayTransactionsV1RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetPayTransactionsV1RespDataInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetPayTransactionsV1RespDataInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetPayTransactionsV1RespDataInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetPayTransactionsV1RespDataInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *GetPayTransactionsV1RespDataInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetPayTransactionsV1RespDataInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetPayTransactionsV1RespDataInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetPayTransactionsV1RespDataInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFundsDetail

`func (o *GetPayTransactionsV1RespDataInner) GetFundsDetail() []GetPayTransactionsV1RespDataInnerFundsDetailInner`

GetFundsDetail returns the FundsDetail field if non-nil, zero value otherwise.

### GetFundsDetailOk

`func (o *GetPayTransactionsV1RespDataInner) GetFundsDetailOk() (*[]GetPayTransactionsV1RespDataInnerFundsDetailInner, bool)`

GetFundsDetailOk returns a tuple with the FundsDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsDetail

`func (o *GetPayTransactionsV1RespDataInner) SetFundsDetail(v []GetPayTransactionsV1RespDataInnerFundsDetailInner)`

SetFundsDetail sets FundsDetail field to given value.

### HasFundsDetail

`func (o *GetPayTransactionsV1RespDataInner) HasFundsDetail() bool`

HasFundsDetail returns a boolean if a field has been set.

### GetOrderType

`func (o *GetPayTransactionsV1RespDataInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *GetPayTransactionsV1RespDataInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *GetPayTransactionsV1RespDataInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *GetPayTransactionsV1RespDataInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPayerInfo

`func (o *GetPayTransactionsV1RespDataInner) GetPayerInfo() GetPayTransactionsV1RespDataInnerPayerInfo`

GetPayerInfo returns the PayerInfo field if non-nil, zero value otherwise.

### GetPayerInfoOk

`func (o *GetPayTransactionsV1RespDataInner) GetPayerInfoOk() (*GetPayTransactionsV1RespDataInnerPayerInfo, bool)`

GetPayerInfoOk returns a tuple with the PayerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerInfo

`func (o *GetPayTransactionsV1RespDataInner) SetPayerInfo(v GetPayTransactionsV1RespDataInnerPayerInfo)`

SetPayerInfo sets PayerInfo field to given value.

### HasPayerInfo

`func (o *GetPayTransactionsV1RespDataInner) HasPayerInfo() bool`

HasPayerInfo returns a boolean if a field has been set.

### GetReceiverInfo

`func (o *GetPayTransactionsV1RespDataInner) GetReceiverInfo() GetPayTransactionsV1RespDataInnerReceiverInfo`

GetReceiverInfo returns the ReceiverInfo field if non-nil, zero value otherwise.

### GetReceiverInfoOk

`func (o *GetPayTransactionsV1RespDataInner) GetReceiverInfoOk() (*GetPayTransactionsV1RespDataInnerReceiverInfo, bool)`

GetReceiverInfoOk returns a tuple with the ReceiverInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverInfo

`func (o *GetPayTransactionsV1RespDataInner) SetReceiverInfo(v GetPayTransactionsV1RespDataInnerReceiverInfo)`

SetReceiverInfo sets ReceiverInfo field to given value.

### HasReceiverInfo

`func (o *GetPayTransactionsV1RespDataInner) HasReceiverInfo() bool`

HasReceiverInfo returns a boolean if a field has been set.

### GetTransactionId

`func (o *GetPayTransactionsV1RespDataInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetPayTransactionsV1RespDataInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetPayTransactionsV1RespDataInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *GetPayTransactionsV1RespDataInner) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *GetPayTransactionsV1RespDataInner) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *GetPayTransactionsV1RespDataInner) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *GetPayTransactionsV1RespDataInner) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *GetPayTransactionsV1RespDataInner) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetWalletType

`func (o *GetPayTransactionsV1RespDataInner) GetWalletType() int32`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *GetPayTransactionsV1RespDataInner) GetWalletTypeOk() (*int32, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *GetPayTransactionsV1RespDataInner) SetWalletType(v int32)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *GetPayTransactionsV1RespDataInner) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetWalletTypes

`func (o *GetPayTransactionsV1RespDataInner) GetWalletTypes() []int32`

GetWalletTypes returns the WalletTypes field if non-nil, zero value otherwise.

### GetWalletTypesOk

`func (o *GetPayTransactionsV1RespDataInner) GetWalletTypesOk() (*[]int32, bool)`

GetWalletTypesOk returns a tuple with the WalletTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletTypes

`func (o *GetPayTransactionsV1RespDataInner) SetWalletTypes(v []int32)`

SetWalletTypes sets WalletTypes field to given value.

### HasWalletTypes

`func (o *GetPayTransactionsV1RespDataInner) HasWalletTypes() bool`

HasWalletTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


