# GetAssetTransferV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetAssetTransferV1RespRowsInner**](GetAssetTransferV1RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAssetTransferV1Resp

`func NewGetAssetTransferV1Resp() *GetAssetTransferV1Resp`

NewGetAssetTransferV1Resp instantiates a new GetAssetTransferV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetTransferV1RespWithDefaults

`func NewGetAssetTransferV1RespWithDefaults() *GetAssetTransferV1Resp`

NewGetAssetTransferV1RespWithDefaults instantiates a new GetAssetTransferV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetAssetTransferV1Resp) GetRows() []GetAssetTransferV1RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetAssetTransferV1Resp) GetRowsOk() (*[]GetAssetTransferV1RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetAssetTransferV1Resp) SetRows(v []GetAssetTransferV1RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetAssetTransferV1Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetAssetTransferV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAssetTransferV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAssetTransferV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAssetTransferV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


