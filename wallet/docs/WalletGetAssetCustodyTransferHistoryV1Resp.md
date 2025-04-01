# WalletGetAssetCustodyTransferHistoryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]WalletGetAssetCustodyTransferHistoryV1RespRowsInner**](WalletGetAssetCustodyTransferHistoryV1RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewWalletGetAssetCustodyTransferHistoryV1Resp

`func NewWalletGetAssetCustodyTransferHistoryV1Resp() *WalletGetAssetCustodyTransferHistoryV1Resp`

NewWalletGetAssetCustodyTransferHistoryV1Resp instantiates a new WalletGetAssetCustodyTransferHistoryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletGetAssetCustodyTransferHistoryV1RespWithDefaults

`func NewWalletGetAssetCustodyTransferHistoryV1RespWithDefaults() *WalletGetAssetCustodyTransferHistoryV1Resp`

NewWalletGetAssetCustodyTransferHistoryV1RespWithDefaults instantiates a new WalletGetAssetCustodyTransferHistoryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *WalletGetAssetCustodyTransferHistoryV1Resp) GetRows() []WalletGetAssetCustodyTransferHistoryV1RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WalletGetAssetCustodyTransferHistoryV1Resp) GetRowsOk() (*[]WalletGetAssetCustodyTransferHistoryV1RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *WalletGetAssetCustodyTransferHistoryV1Resp) SetRows(v []WalletGetAssetCustodyTransferHistoryV1RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *WalletGetAssetCustodyTransferHistoryV1Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *WalletGetAssetCustodyTransferHistoryV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WalletGetAssetCustodyTransferHistoryV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WalletGetAssetCustodyTransferHistoryV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *WalletGetAssetCustodyTransferHistoryV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


