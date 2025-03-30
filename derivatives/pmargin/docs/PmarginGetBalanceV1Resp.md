# PmarginGetBalanceV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** |  | [optional] 
**CmUnrealizedPNL** | Pointer to **string** |  | [optional] 
**CmWalletBalance** | Pointer to **string** |  | [optional] 
**CrossMarginAsset** | Pointer to **string** |  | [optional] 
**CrossMarginBorrowed** | Pointer to **string** |  | [optional] 
**CrossMarginFree** | Pointer to **string** |  | [optional] 
**CrossMarginInterest** | Pointer to **string** |  | [optional] 
**CrossMarginLocked** | Pointer to **string** |  | [optional] 
**NegativeBalance** | Pointer to **string** |  | [optional] 
**TotalWalletBalance** | Pointer to **string** |  | [optional] 
**UmUnrealizedPNL** | Pointer to **string** |  | [optional] 
**UmWalletBalance** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPmarginGetBalanceV1Resp

`func NewPmarginGetBalanceV1Resp() *PmarginGetBalanceV1Resp`

NewPmarginGetBalanceV1Resp instantiates a new PmarginGetBalanceV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetBalanceV1RespWithDefaults

`func NewPmarginGetBalanceV1RespWithDefaults() *PmarginGetBalanceV1Resp`

NewPmarginGetBalanceV1RespWithDefaults instantiates a new PmarginGetBalanceV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *PmarginGetBalanceV1Resp) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *PmarginGetBalanceV1Resp) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *PmarginGetBalanceV1Resp) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *PmarginGetBalanceV1Resp) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetCmUnrealizedPNL

`func (o *PmarginGetBalanceV1Resp) GetCmUnrealizedPNL() string`

GetCmUnrealizedPNL returns the CmUnrealizedPNL field if non-nil, zero value otherwise.

### GetCmUnrealizedPNLOk

`func (o *PmarginGetBalanceV1Resp) GetCmUnrealizedPNLOk() (*string, bool)`

GetCmUnrealizedPNLOk returns a tuple with the CmUnrealizedPNL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmUnrealizedPNL

`func (o *PmarginGetBalanceV1Resp) SetCmUnrealizedPNL(v string)`

SetCmUnrealizedPNL sets CmUnrealizedPNL field to given value.

### HasCmUnrealizedPNL

`func (o *PmarginGetBalanceV1Resp) HasCmUnrealizedPNL() bool`

HasCmUnrealizedPNL returns a boolean if a field has been set.

### GetCmWalletBalance

`func (o *PmarginGetBalanceV1Resp) GetCmWalletBalance() string`

GetCmWalletBalance returns the CmWalletBalance field if non-nil, zero value otherwise.

### GetCmWalletBalanceOk

`func (o *PmarginGetBalanceV1Resp) GetCmWalletBalanceOk() (*string, bool)`

GetCmWalletBalanceOk returns a tuple with the CmWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmWalletBalance

`func (o *PmarginGetBalanceV1Resp) SetCmWalletBalance(v string)`

SetCmWalletBalance sets CmWalletBalance field to given value.

### HasCmWalletBalance

`func (o *PmarginGetBalanceV1Resp) HasCmWalletBalance() bool`

HasCmWalletBalance returns a boolean if a field has been set.

### GetCrossMarginAsset

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginAsset() string`

GetCrossMarginAsset returns the CrossMarginAsset field if non-nil, zero value otherwise.

### GetCrossMarginAssetOk

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginAssetOk() (*string, bool)`

GetCrossMarginAssetOk returns a tuple with the CrossMarginAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMarginAsset

`func (o *PmarginGetBalanceV1Resp) SetCrossMarginAsset(v string)`

SetCrossMarginAsset sets CrossMarginAsset field to given value.

### HasCrossMarginAsset

`func (o *PmarginGetBalanceV1Resp) HasCrossMarginAsset() bool`

HasCrossMarginAsset returns a boolean if a field has been set.

### GetCrossMarginBorrowed

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginBorrowed() string`

GetCrossMarginBorrowed returns the CrossMarginBorrowed field if non-nil, zero value otherwise.

### GetCrossMarginBorrowedOk

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginBorrowedOk() (*string, bool)`

GetCrossMarginBorrowedOk returns a tuple with the CrossMarginBorrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMarginBorrowed

`func (o *PmarginGetBalanceV1Resp) SetCrossMarginBorrowed(v string)`

SetCrossMarginBorrowed sets CrossMarginBorrowed field to given value.

### HasCrossMarginBorrowed

`func (o *PmarginGetBalanceV1Resp) HasCrossMarginBorrowed() bool`

HasCrossMarginBorrowed returns a boolean if a field has been set.

### GetCrossMarginFree

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginFree() string`

GetCrossMarginFree returns the CrossMarginFree field if non-nil, zero value otherwise.

### GetCrossMarginFreeOk

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginFreeOk() (*string, bool)`

GetCrossMarginFreeOk returns a tuple with the CrossMarginFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMarginFree

`func (o *PmarginGetBalanceV1Resp) SetCrossMarginFree(v string)`

SetCrossMarginFree sets CrossMarginFree field to given value.

### HasCrossMarginFree

`func (o *PmarginGetBalanceV1Resp) HasCrossMarginFree() bool`

HasCrossMarginFree returns a boolean if a field has been set.

### GetCrossMarginInterest

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginInterest() string`

GetCrossMarginInterest returns the CrossMarginInterest field if non-nil, zero value otherwise.

### GetCrossMarginInterestOk

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginInterestOk() (*string, bool)`

GetCrossMarginInterestOk returns a tuple with the CrossMarginInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMarginInterest

`func (o *PmarginGetBalanceV1Resp) SetCrossMarginInterest(v string)`

SetCrossMarginInterest sets CrossMarginInterest field to given value.

### HasCrossMarginInterest

`func (o *PmarginGetBalanceV1Resp) HasCrossMarginInterest() bool`

HasCrossMarginInterest returns a boolean if a field has been set.

### GetCrossMarginLocked

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginLocked() string`

GetCrossMarginLocked returns the CrossMarginLocked field if non-nil, zero value otherwise.

### GetCrossMarginLockedOk

`func (o *PmarginGetBalanceV1Resp) GetCrossMarginLockedOk() (*string, bool)`

GetCrossMarginLockedOk returns a tuple with the CrossMarginLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMarginLocked

`func (o *PmarginGetBalanceV1Resp) SetCrossMarginLocked(v string)`

SetCrossMarginLocked sets CrossMarginLocked field to given value.

### HasCrossMarginLocked

`func (o *PmarginGetBalanceV1Resp) HasCrossMarginLocked() bool`

HasCrossMarginLocked returns a boolean if a field has been set.

### GetNegativeBalance

`func (o *PmarginGetBalanceV1Resp) GetNegativeBalance() string`

GetNegativeBalance returns the NegativeBalance field if non-nil, zero value otherwise.

### GetNegativeBalanceOk

`func (o *PmarginGetBalanceV1Resp) GetNegativeBalanceOk() (*string, bool)`

GetNegativeBalanceOk returns a tuple with the NegativeBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeBalance

`func (o *PmarginGetBalanceV1Resp) SetNegativeBalance(v string)`

SetNegativeBalance sets NegativeBalance field to given value.

### HasNegativeBalance

`func (o *PmarginGetBalanceV1Resp) HasNegativeBalance() bool`

HasNegativeBalance returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *PmarginGetBalanceV1Resp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *PmarginGetBalanceV1Resp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *PmarginGetBalanceV1Resp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *PmarginGetBalanceV1Resp) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.

### GetUmUnrealizedPNL

`func (o *PmarginGetBalanceV1Resp) GetUmUnrealizedPNL() string`

GetUmUnrealizedPNL returns the UmUnrealizedPNL field if non-nil, zero value otherwise.

### GetUmUnrealizedPNLOk

`func (o *PmarginGetBalanceV1Resp) GetUmUnrealizedPNLOk() (*string, bool)`

GetUmUnrealizedPNLOk returns a tuple with the UmUnrealizedPNL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmUnrealizedPNL

`func (o *PmarginGetBalanceV1Resp) SetUmUnrealizedPNL(v string)`

SetUmUnrealizedPNL sets UmUnrealizedPNL field to given value.

### HasUmUnrealizedPNL

`func (o *PmarginGetBalanceV1Resp) HasUmUnrealizedPNL() bool`

HasUmUnrealizedPNL returns a boolean if a field has been set.

### GetUmWalletBalance

`func (o *PmarginGetBalanceV1Resp) GetUmWalletBalance() string`

GetUmWalletBalance returns the UmWalletBalance field if non-nil, zero value otherwise.

### GetUmWalletBalanceOk

`func (o *PmarginGetBalanceV1Resp) GetUmWalletBalanceOk() (*string, bool)`

GetUmWalletBalanceOk returns a tuple with the UmWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmWalletBalance

`func (o *PmarginGetBalanceV1Resp) SetUmWalletBalance(v string)`

SetUmWalletBalance sets UmWalletBalance field to given value.

### HasUmWalletBalance

`func (o *PmarginGetBalanceV1Resp) HasUmWalletBalance() bool`

HasUmWalletBalance returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PmarginGetBalanceV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PmarginGetBalanceV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PmarginGetBalanceV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PmarginGetBalanceV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


