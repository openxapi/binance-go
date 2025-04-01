# WalletGetAssetDribbletV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**UserAssetDribblets** | Pointer to [**[]WalletGetAssetDribbletV1RespUserAssetDribbletsInner**](WalletGetAssetDribbletV1RespUserAssetDribbletsInner.md) |  | [optional] 

## Methods

### NewWalletGetAssetDribbletV1Resp

`func NewWalletGetAssetDribbletV1Resp() *WalletGetAssetDribbletV1Resp`

NewWalletGetAssetDribbletV1Resp instantiates a new WalletGetAssetDribbletV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletGetAssetDribbletV1RespWithDefaults

`func NewWalletGetAssetDribbletV1RespWithDefaults() *WalletGetAssetDribbletV1Resp`

NewWalletGetAssetDribbletV1RespWithDefaults instantiates a new WalletGetAssetDribbletV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *WalletGetAssetDribbletV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WalletGetAssetDribbletV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WalletGetAssetDribbletV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *WalletGetAssetDribbletV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUserAssetDribblets

`func (o *WalletGetAssetDribbletV1Resp) GetUserAssetDribblets() []WalletGetAssetDribbletV1RespUserAssetDribbletsInner`

GetUserAssetDribblets returns the UserAssetDribblets field if non-nil, zero value otherwise.

### GetUserAssetDribbletsOk

`func (o *WalletGetAssetDribbletV1Resp) GetUserAssetDribbletsOk() (*[]WalletGetAssetDribbletV1RespUserAssetDribbletsInner, bool)`

GetUserAssetDribbletsOk returns a tuple with the UserAssetDribblets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssetDribblets

`func (o *WalletGetAssetDribbletV1Resp) SetUserAssetDribblets(v []WalletGetAssetDribbletV1RespUserAssetDribbletsInner)`

SetUserAssetDribblets sets UserAssetDribblets field to given value.

### HasUserAssetDribblets

`func (o *WalletGetAssetDribbletV1Resp) HasUserAssetDribblets() bool`

HasUserAssetDribblets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


