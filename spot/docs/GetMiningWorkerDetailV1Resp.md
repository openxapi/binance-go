# GetMiningWorkerDetailV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]GetMiningWorkerDetailV1RespDataInner**](GetMiningWorkerDetailV1RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewGetMiningWorkerDetailV1Resp

`func NewGetMiningWorkerDetailV1Resp() *GetMiningWorkerDetailV1Resp`

NewGetMiningWorkerDetailV1Resp instantiates a new GetMiningWorkerDetailV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiningWorkerDetailV1RespWithDefaults

`func NewGetMiningWorkerDetailV1RespWithDefaults() *GetMiningWorkerDetailV1Resp`

NewGetMiningWorkerDetailV1RespWithDefaults instantiates a new GetMiningWorkerDetailV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetMiningWorkerDetailV1Resp) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetMiningWorkerDetailV1Resp) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetMiningWorkerDetailV1Resp) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetMiningWorkerDetailV1Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetMiningWorkerDetailV1Resp) GetData() []GetMiningWorkerDetailV1RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMiningWorkerDetailV1Resp) GetDataOk() (*[]GetMiningWorkerDetailV1RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMiningWorkerDetailV1Resp) SetData(v []GetMiningWorkerDetailV1RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetMiningWorkerDetailV1Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetMiningWorkerDetailV1Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetMiningWorkerDetailV1Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetMiningWorkerDetailV1Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetMiningWorkerDetailV1Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


