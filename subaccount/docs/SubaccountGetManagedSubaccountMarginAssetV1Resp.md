# SubaccountGetManagedSubaccountMarginAssetV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarginLevel** | Pointer to **string** |  | [optional] 
**TotalAssetOfBtc** | Pointer to **string** |  | [optional] 
**TotalLiabilityOfBtc** | Pointer to **string** |  | [optional] 
**TotalNetAssetOfBtc** | Pointer to **string** |  | [optional] 
**UserAssets** | Pointer to [**[]SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner**](SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner.md) |  | [optional] 

## Methods

### NewSubaccountGetManagedSubaccountMarginAssetV1Resp

`func NewSubaccountGetManagedSubaccountMarginAssetV1Resp() *SubaccountGetManagedSubaccountMarginAssetV1Resp`

NewSubaccountGetManagedSubaccountMarginAssetV1Resp instantiates a new SubaccountGetManagedSubaccountMarginAssetV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubaccountGetManagedSubaccountMarginAssetV1RespWithDefaults

`func NewSubaccountGetManagedSubaccountMarginAssetV1RespWithDefaults() *SubaccountGetManagedSubaccountMarginAssetV1Resp`

NewSubaccountGetManagedSubaccountMarginAssetV1RespWithDefaults instantiates a new SubaccountGetManagedSubaccountMarginAssetV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarginLevel

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetTotalAssetOfBtc

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.

### HasTotalAssetOfBtc

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasTotalAssetOfBtc() bool`

HasTotalAssetOfBtc returns a boolean if a field has been set.

### GetTotalLiabilityOfBtc

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalLiabilityOfBtc() string`

GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field if non-nil, zero value otherwise.

### GetTotalLiabilityOfBtcOk

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalLiabilityOfBtcOk() (*string, bool)`

GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilityOfBtc

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetTotalLiabilityOfBtc(v string)`

SetTotalLiabilityOfBtc sets TotalLiabilityOfBtc field to given value.

### HasTotalLiabilityOfBtc

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasTotalLiabilityOfBtc() bool`

HasTotalLiabilityOfBtc returns a boolean if a field has been set.

### GetTotalNetAssetOfBtc

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalNetAssetOfBtc() string`

GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalNetAssetOfBtcOk

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetTotalNetAssetOfBtcOk() (*string, bool)`

GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAssetOfBtc

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetTotalNetAssetOfBtc(v string)`

SetTotalNetAssetOfBtc sets TotalNetAssetOfBtc field to given value.

### HasTotalNetAssetOfBtc

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasTotalNetAssetOfBtc() bool`

HasTotalNetAssetOfBtc returns a boolean if a field has been set.

### GetUserAssets

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetUserAssets() []SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner`

GetUserAssets returns the UserAssets field if non-nil, zero value otherwise.

### GetUserAssetsOk

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) GetUserAssetsOk() (*[]SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner, bool)`

GetUserAssetsOk returns a tuple with the UserAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssets

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) SetUserAssets(v []SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner)`

SetUserAssets sets UserAssets field to given value.

### HasUserAssets

`func (o *SubaccountGetManagedSubaccountMarginAssetV1Resp) HasUserAssets() bool`

HasUserAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


