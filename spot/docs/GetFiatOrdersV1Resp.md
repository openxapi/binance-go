# GetFiatOrdersV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GetFiatOrdersV1RespDataInner**](GetFiatOrdersV1RespDataInner.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetFiatOrdersV1Resp

`func NewGetFiatOrdersV1Resp() *GetFiatOrdersV1Resp`

NewGetFiatOrdersV1Resp instantiates a new GetFiatOrdersV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatOrdersV1RespWithDefaults

`func NewGetFiatOrdersV1RespWithDefaults() *GetFiatOrdersV1Resp`

NewGetFiatOrdersV1RespWithDefaults instantiates a new GetFiatOrdersV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetFiatOrdersV1Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetFiatOrdersV1Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetFiatOrdersV1Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetFiatOrdersV1Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetFiatOrdersV1Resp) GetData() []GetFiatOrdersV1RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFiatOrdersV1Resp) GetDataOk() (*[]GetFiatOrdersV1RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFiatOrdersV1Resp) SetData(v []GetFiatOrdersV1RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetFiatOrdersV1Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *GetFiatOrdersV1Resp) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetFiatOrdersV1Resp) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetFiatOrdersV1Resp) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetFiatOrdersV1Resp) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *GetFiatOrdersV1Resp) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetFiatOrdersV1Resp) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetFiatOrdersV1Resp) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetFiatOrdersV1Resp) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTotal

`func (o *GetFiatOrdersV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetFiatOrdersV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetFiatOrdersV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetFiatOrdersV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


