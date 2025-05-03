# CreateAssetDustV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalServiceCharge** | Pointer to **string** |  | [optional] 
**TotalTransfered** | Pointer to **string** |  | [optional] 
**TransferResult** | Pointer to [**[]CreateAssetDustV1RespTransferResultInner**](CreateAssetDustV1RespTransferResultInner.md) |  | [optional] 

## Methods

### NewCreateAssetDustV1Resp

`func NewCreateAssetDustV1Resp() *CreateAssetDustV1Resp`

NewCreateAssetDustV1Resp instantiates a new CreateAssetDustV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetDustV1RespWithDefaults

`func NewCreateAssetDustV1RespWithDefaults() *CreateAssetDustV1Resp`

NewCreateAssetDustV1RespWithDefaults instantiates a new CreateAssetDustV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalServiceCharge

`func (o *CreateAssetDustV1Resp) GetTotalServiceCharge() string`

GetTotalServiceCharge returns the TotalServiceCharge field if non-nil, zero value otherwise.

### GetTotalServiceChargeOk

`func (o *CreateAssetDustV1Resp) GetTotalServiceChargeOk() (*string, bool)`

GetTotalServiceChargeOk returns a tuple with the TotalServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceCharge

`func (o *CreateAssetDustV1Resp) SetTotalServiceCharge(v string)`

SetTotalServiceCharge sets TotalServiceCharge field to given value.

### HasTotalServiceCharge

`func (o *CreateAssetDustV1Resp) HasTotalServiceCharge() bool`

HasTotalServiceCharge returns a boolean if a field has been set.

### GetTotalTransfered

`func (o *CreateAssetDustV1Resp) GetTotalTransfered() string`

GetTotalTransfered returns the TotalTransfered field if non-nil, zero value otherwise.

### GetTotalTransferedOk

`func (o *CreateAssetDustV1Resp) GetTotalTransferedOk() (*string, bool)`

GetTotalTransferedOk returns a tuple with the TotalTransfered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransfered

`func (o *CreateAssetDustV1Resp) SetTotalTransfered(v string)`

SetTotalTransfered sets TotalTransfered field to given value.

### HasTotalTransfered

`func (o *CreateAssetDustV1Resp) HasTotalTransfered() bool`

HasTotalTransfered returns a boolean if a field has been set.

### GetTransferResult

`func (o *CreateAssetDustV1Resp) GetTransferResult() []CreateAssetDustV1RespTransferResultInner`

GetTransferResult returns the TransferResult field if non-nil, zero value otherwise.

### GetTransferResultOk

`func (o *CreateAssetDustV1Resp) GetTransferResultOk() (*[]CreateAssetDustV1RespTransferResultInner, bool)`

GetTransferResultOk returns a tuple with the TransferResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferResult

`func (o *CreateAssetDustV1Resp) SetTransferResult(v []CreateAssetDustV1RespTransferResultInner)`

SetTransferResult sets TransferResult field to given value.

### HasTransferResult

`func (o *CreateAssetDustV1Resp) HasTransferResult() bool`

HasTransferResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


