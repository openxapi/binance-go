# GetFiatPaymentsV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GetFiatPaymentsV1RespDataInner**](GetFiatPaymentsV1RespDataInner.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetFiatPaymentsV1Resp

`func NewGetFiatPaymentsV1Resp() *GetFiatPaymentsV1Resp`

NewGetFiatPaymentsV1Resp instantiates a new GetFiatPaymentsV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatPaymentsV1RespWithDefaults

`func NewGetFiatPaymentsV1RespWithDefaults() *GetFiatPaymentsV1Resp`

NewGetFiatPaymentsV1RespWithDefaults instantiates a new GetFiatPaymentsV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetFiatPaymentsV1Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetFiatPaymentsV1Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetFiatPaymentsV1Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetFiatPaymentsV1Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetFiatPaymentsV1Resp) GetData() []GetFiatPaymentsV1RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFiatPaymentsV1Resp) GetDataOk() (*[]GetFiatPaymentsV1RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFiatPaymentsV1Resp) SetData(v []GetFiatPaymentsV1RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetFiatPaymentsV1Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *GetFiatPaymentsV1Resp) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetFiatPaymentsV1Resp) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetFiatPaymentsV1Resp) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetFiatPaymentsV1Resp) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *GetFiatPaymentsV1Resp) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetFiatPaymentsV1Resp) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetFiatPaymentsV1Resp) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetFiatPaymentsV1Resp) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTotal

`func (o *GetFiatPaymentsV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetFiatPaymentsV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetFiatPaymentsV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetFiatPaymentsV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


