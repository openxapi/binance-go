# WalletCreateAssetDustV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalServiceCharge** | Pointer to **string** |  | [optional] 
**TotalTransfered** | Pointer to **string** |  | [optional] 
**TransferResult** | Pointer to [**[]WalletCreateAssetDustV1RespTransferResultInner**](WalletCreateAssetDustV1RespTransferResultInner.md) |  | [optional] 

## Methods

### NewWalletCreateAssetDustV1Resp

`func NewWalletCreateAssetDustV1Resp() *WalletCreateAssetDustV1Resp`

NewWalletCreateAssetDustV1Resp instantiates a new WalletCreateAssetDustV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletCreateAssetDustV1RespWithDefaults

`func NewWalletCreateAssetDustV1RespWithDefaults() *WalletCreateAssetDustV1Resp`

NewWalletCreateAssetDustV1RespWithDefaults instantiates a new WalletCreateAssetDustV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalServiceCharge

`func (o *WalletCreateAssetDustV1Resp) GetTotalServiceCharge() string`

GetTotalServiceCharge returns the TotalServiceCharge field if non-nil, zero value otherwise.

### GetTotalServiceChargeOk

`func (o *WalletCreateAssetDustV1Resp) GetTotalServiceChargeOk() (*string, bool)`

GetTotalServiceChargeOk returns a tuple with the TotalServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceCharge

`func (o *WalletCreateAssetDustV1Resp) SetTotalServiceCharge(v string)`

SetTotalServiceCharge sets TotalServiceCharge field to given value.

### HasTotalServiceCharge

`func (o *WalletCreateAssetDustV1Resp) HasTotalServiceCharge() bool`

HasTotalServiceCharge returns a boolean if a field has been set.

### GetTotalTransfered

`func (o *WalletCreateAssetDustV1Resp) GetTotalTransfered() string`

GetTotalTransfered returns the TotalTransfered field if non-nil, zero value otherwise.

### GetTotalTransferedOk

`func (o *WalletCreateAssetDustV1Resp) GetTotalTransferedOk() (*string, bool)`

GetTotalTransferedOk returns a tuple with the TotalTransfered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransfered

`func (o *WalletCreateAssetDustV1Resp) SetTotalTransfered(v string)`

SetTotalTransfered sets TotalTransfered field to given value.

### HasTotalTransfered

`func (o *WalletCreateAssetDustV1Resp) HasTotalTransfered() bool`

HasTotalTransfered returns a boolean if a field has been set.

### GetTransferResult

`func (o *WalletCreateAssetDustV1Resp) GetTransferResult() []WalletCreateAssetDustV1RespTransferResultInner`

GetTransferResult returns the TransferResult field if non-nil, zero value otherwise.

### GetTransferResultOk

`func (o *WalletCreateAssetDustV1Resp) GetTransferResultOk() (*[]WalletCreateAssetDustV1RespTransferResultInner, bool)`

GetTransferResultOk returns a tuple with the TransferResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferResult

`func (o *WalletCreateAssetDustV1Resp) SetTransferResult(v []WalletCreateAssetDustV1RespTransferResultInner)`

SetTransferResult sets TransferResult field to given value.

### HasTransferResult

`func (o *WalletCreateAssetDustV1Resp) HasTransferResult() bool`

HasTransferResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


