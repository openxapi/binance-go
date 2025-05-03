# GetSubAccountFuturesInternalTransferV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuturesType** | Pointer to **int32** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Transfers** | Pointer to [**[]GetSubAccountFuturesInternalTransferV1RespTransfersInner**](GetSubAccountFuturesInternalTransferV1RespTransfersInner.md) |  | [optional] 

## Methods

### NewGetSubAccountFuturesInternalTransferV1Resp

`func NewGetSubAccountFuturesInternalTransferV1Resp() *GetSubAccountFuturesInternalTransferV1Resp`

NewGetSubAccountFuturesInternalTransferV1Resp instantiates a new GetSubAccountFuturesInternalTransferV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubAccountFuturesInternalTransferV1RespWithDefaults

`func NewGetSubAccountFuturesInternalTransferV1RespWithDefaults() *GetSubAccountFuturesInternalTransferV1Resp`

NewGetSubAccountFuturesInternalTransferV1RespWithDefaults instantiates a new GetSubAccountFuturesInternalTransferV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuturesType

`func (o *GetSubAccountFuturesInternalTransferV1Resp) GetFuturesType() int32`

GetFuturesType returns the FuturesType field if non-nil, zero value otherwise.

### GetFuturesTypeOk

`func (o *GetSubAccountFuturesInternalTransferV1Resp) GetFuturesTypeOk() (*int32, bool)`

GetFuturesTypeOk returns a tuple with the FuturesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuturesType

`func (o *GetSubAccountFuturesInternalTransferV1Resp) SetFuturesType(v int32)`

SetFuturesType sets FuturesType field to given value.

### HasFuturesType

`func (o *GetSubAccountFuturesInternalTransferV1Resp) HasFuturesType() bool`

HasFuturesType returns a boolean if a field has been set.

### GetSuccess

`func (o *GetSubAccountFuturesInternalTransferV1Resp) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetSubAccountFuturesInternalTransferV1Resp) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetSubAccountFuturesInternalTransferV1Resp) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetSubAccountFuturesInternalTransferV1Resp) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTransfers

`func (o *GetSubAccountFuturesInternalTransferV1Resp) GetTransfers() []GetSubAccountFuturesInternalTransferV1RespTransfersInner`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *GetSubAccountFuturesInternalTransferV1Resp) GetTransfersOk() (*[]GetSubAccountFuturesInternalTransferV1RespTransfersInner, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *GetSubAccountFuturesInternalTransferV1Resp) SetTransfers(v []GetSubAccountFuturesInternalTransferV1RespTransfersInner)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *GetSubAccountFuturesInternalTransferV1Resp) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


