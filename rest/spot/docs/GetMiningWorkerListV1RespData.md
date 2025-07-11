# GetMiningWorkerListV1RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalNum** | Pointer to **int32** |  | [optional] 
**WorkerDatas** | Pointer to [**[]GetMiningWorkerListV1RespDataWorkerDatasInner**](GetMiningWorkerListV1RespDataWorkerDatasInner.md) |  | [optional] 

## Methods

### NewGetMiningWorkerListV1RespData

`func NewGetMiningWorkerListV1RespData() *GetMiningWorkerListV1RespData`

NewGetMiningWorkerListV1RespData instantiates a new GetMiningWorkerListV1RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiningWorkerListV1RespDataWithDefaults

`func NewGetMiningWorkerListV1RespDataWithDefaults() *GetMiningWorkerListV1RespData`

NewGetMiningWorkerListV1RespDataWithDefaults instantiates a new GetMiningWorkerListV1RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *GetMiningWorkerListV1RespData) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetMiningWorkerListV1RespData) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetMiningWorkerListV1RespData) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetMiningWorkerListV1RespData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalNum

`func (o *GetMiningWorkerListV1RespData) GetTotalNum() int32`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *GetMiningWorkerListV1RespData) GetTotalNumOk() (*int32, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *GetMiningWorkerListV1RespData) SetTotalNum(v int32)`

SetTotalNum sets TotalNum field to given value.

### HasTotalNum

`func (o *GetMiningWorkerListV1RespData) HasTotalNum() bool`

HasTotalNum returns a boolean if a field has been set.

### GetWorkerDatas

`func (o *GetMiningWorkerListV1RespData) GetWorkerDatas() []GetMiningWorkerListV1RespDataWorkerDatasInner`

GetWorkerDatas returns the WorkerDatas field if non-nil, zero value otherwise.

### GetWorkerDatasOk

`func (o *GetMiningWorkerListV1RespData) GetWorkerDatasOk() (*[]GetMiningWorkerListV1RespDataWorkerDatasInner, bool)`

GetWorkerDatasOk returns a tuple with the WorkerDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerDatas

`func (o *GetMiningWorkerListV1RespData) SetWorkerDatas(v []GetMiningWorkerListV1RespDataWorkerDatasInner)`

SetWorkerDatas sets WorkerDatas field to given value.

### HasWorkerDatas

`func (o *GetMiningWorkerListV1RespData) HasWorkerDatas() bool`

HasWorkerDatas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


