# GetMarginCrossMarginDataV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BorrowLimit** | Pointer to **string** |  | [optional] 
**Borrowable** | Pointer to **bool** |  | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**DailyInterest** | Pointer to **string** |  | [optional] 
**MarginablePairs** | Pointer to **[]string** |  | [optional] 
**TransferIn** | Pointer to **bool** |  | [optional] 
**VipLevel** | Pointer to **int32** |  | [optional] 
**YearlyInterest** | Pointer to **string** |  | [optional] 

## Methods

### NewGetMarginCrossMarginDataV1RespItem

`func NewGetMarginCrossMarginDataV1RespItem() *GetMarginCrossMarginDataV1RespItem`

NewGetMarginCrossMarginDataV1RespItem instantiates a new GetMarginCrossMarginDataV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarginCrossMarginDataV1RespItemWithDefaults

`func NewGetMarginCrossMarginDataV1RespItemWithDefaults() *GetMarginCrossMarginDataV1RespItem`

NewGetMarginCrossMarginDataV1RespItemWithDefaults instantiates a new GetMarginCrossMarginDataV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBorrowLimit

`func (o *GetMarginCrossMarginDataV1RespItem) GetBorrowLimit() string`

GetBorrowLimit returns the BorrowLimit field if non-nil, zero value otherwise.

### GetBorrowLimitOk

`func (o *GetMarginCrossMarginDataV1RespItem) GetBorrowLimitOk() (*string, bool)`

GetBorrowLimitOk returns a tuple with the BorrowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowLimit

`func (o *GetMarginCrossMarginDataV1RespItem) SetBorrowLimit(v string)`

SetBorrowLimit sets BorrowLimit field to given value.

### HasBorrowLimit

`func (o *GetMarginCrossMarginDataV1RespItem) HasBorrowLimit() bool`

HasBorrowLimit returns a boolean if a field has been set.

### GetBorrowable

`func (o *GetMarginCrossMarginDataV1RespItem) GetBorrowable() bool`

GetBorrowable returns the Borrowable field if non-nil, zero value otherwise.

### GetBorrowableOk

`func (o *GetMarginCrossMarginDataV1RespItem) GetBorrowableOk() (*bool, bool)`

GetBorrowableOk returns a tuple with the Borrowable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowable

`func (o *GetMarginCrossMarginDataV1RespItem) SetBorrowable(v bool)`

SetBorrowable sets Borrowable field to given value.

### HasBorrowable

`func (o *GetMarginCrossMarginDataV1RespItem) HasBorrowable() bool`

HasBorrowable returns a boolean if a field has been set.

### GetCoin

`func (o *GetMarginCrossMarginDataV1RespItem) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *GetMarginCrossMarginDataV1RespItem) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *GetMarginCrossMarginDataV1RespItem) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *GetMarginCrossMarginDataV1RespItem) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetDailyInterest

`func (o *GetMarginCrossMarginDataV1RespItem) GetDailyInterest() string`

GetDailyInterest returns the DailyInterest field if non-nil, zero value otherwise.

### GetDailyInterestOk

`func (o *GetMarginCrossMarginDataV1RespItem) GetDailyInterestOk() (*string, bool)`

GetDailyInterestOk returns a tuple with the DailyInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyInterest

`func (o *GetMarginCrossMarginDataV1RespItem) SetDailyInterest(v string)`

SetDailyInterest sets DailyInterest field to given value.

### HasDailyInterest

`func (o *GetMarginCrossMarginDataV1RespItem) HasDailyInterest() bool`

HasDailyInterest returns a boolean if a field has been set.

### GetMarginablePairs

`func (o *GetMarginCrossMarginDataV1RespItem) GetMarginablePairs() []string`

GetMarginablePairs returns the MarginablePairs field if non-nil, zero value otherwise.

### GetMarginablePairsOk

`func (o *GetMarginCrossMarginDataV1RespItem) GetMarginablePairsOk() (*[]string, bool)`

GetMarginablePairsOk returns a tuple with the MarginablePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginablePairs

`func (o *GetMarginCrossMarginDataV1RespItem) SetMarginablePairs(v []string)`

SetMarginablePairs sets MarginablePairs field to given value.

### HasMarginablePairs

`func (o *GetMarginCrossMarginDataV1RespItem) HasMarginablePairs() bool`

HasMarginablePairs returns a boolean if a field has been set.

### GetTransferIn

`func (o *GetMarginCrossMarginDataV1RespItem) GetTransferIn() bool`

GetTransferIn returns the TransferIn field if non-nil, zero value otherwise.

### GetTransferInOk

`func (o *GetMarginCrossMarginDataV1RespItem) GetTransferInOk() (*bool, bool)`

GetTransferInOk returns a tuple with the TransferIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferIn

`func (o *GetMarginCrossMarginDataV1RespItem) SetTransferIn(v bool)`

SetTransferIn sets TransferIn field to given value.

### HasTransferIn

`func (o *GetMarginCrossMarginDataV1RespItem) HasTransferIn() bool`

HasTransferIn returns a boolean if a field has been set.

### GetVipLevel

`func (o *GetMarginCrossMarginDataV1RespItem) GetVipLevel() int32`

GetVipLevel returns the VipLevel field if non-nil, zero value otherwise.

### GetVipLevelOk

`func (o *GetMarginCrossMarginDataV1RespItem) GetVipLevelOk() (*int32, bool)`

GetVipLevelOk returns a tuple with the VipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipLevel

`func (o *GetMarginCrossMarginDataV1RespItem) SetVipLevel(v int32)`

SetVipLevel sets VipLevel field to given value.

### HasVipLevel

`func (o *GetMarginCrossMarginDataV1RespItem) HasVipLevel() bool`

HasVipLevel returns a boolean if a field has been set.

### GetYearlyInterest

`func (o *GetMarginCrossMarginDataV1RespItem) GetYearlyInterest() string`

GetYearlyInterest returns the YearlyInterest field if non-nil, zero value otherwise.

### GetYearlyInterestOk

`func (o *GetMarginCrossMarginDataV1RespItem) GetYearlyInterestOk() (*string, bool)`

GetYearlyInterestOk returns a tuple with the YearlyInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyInterest

`func (o *GetMarginCrossMarginDataV1RespItem) SetYearlyInterest(v string)`

SetYearlyInterest sets YearlyInterest field to given value.

### HasYearlyInterest

`func (o *GetMarginCrossMarginDataV1RespItem) HasYearlyInterest() bool`

HasYearlyInterest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


