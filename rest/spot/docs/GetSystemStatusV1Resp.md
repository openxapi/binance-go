# GetSystemStatusV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetSystemStatusV1Resp

`func NewGetSystemStatusV1Resp() *GetSystemStatusV1Resp`

NewGetSystemStatusV1Resp instantiates a new GetSystemStatusV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSystemStatusV1RespWithDefaults

`func NewGetSystemStatusV1RespWithDefaults() *GetSystemStatusV1Resp`

NewGetSystemStatusV1RespWithDefaults instantiates a new GetSystemStatusV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GetSystemStatusV1Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetSystemStatusV1Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetSystemStatusV1Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetSystemStatusV1Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *GetSystemStatusV1Resp) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSystemStatusV1Resp) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSystemStatusV1Resp) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSystemStatusV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


