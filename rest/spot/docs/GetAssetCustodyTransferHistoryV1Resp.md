# GetAssetCustodyTransferHistoryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetAssetCustodyTransferHistoryV1RespRowsInner**](GetAssetCustodyTransferHistoryV1RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAssetCustodyTransferHistoryV1Resp

`func NewGetAssetCustodyTransferHistoryV1Resp() *GetAssetCustodyTransferHistoryV1Resp`

NewGetAssetCustodyTransferHistoryV1Resp instantiates a new GetAssetCustodyTransferHistoryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetCustodyTransferHistoryV1RespWithDefaults

`func NewGetAssetCustodyTransferHistoryV1RespWithDefaults() *GetAssetCustodyTransferHistoryV1Resp`

NewGetAssetCustodyTransferHistoryV1RespWithDefaults instantiates a new GetAssetCustodyTransferHistoryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetAssetCustodyTransferHistoryV1Resp) GetRows() []GetAssetCustodyTransferHistoryV1RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetAssetCustodyTransferHistoryV1Resp) GetRowsOk() (*[]GetAssetCustodyTransferHistoryV1RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetAssetCustodyTransferHistoryV1Resp) SetRows(v []GetAssetCustodyTransferHistoryV1RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetAssetCustodyTransferHistoryV1Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetAssetCustodyTransferHistoryV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAssetCustodyTransferHistoryV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAssetCustodyTransferHistoryV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAssetCustodyTransferHistoryV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


