# GetPayTransactionsV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GetPayTransactionsV1RespDataInner**](GetPayTransactionsV1RespDataInner.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetPayTransactionsV1Resp

`func NewGetPayTransactionsV1Resp() *GetPayTransactionsV1Resp`

NewGetPayTransactionsV1Resp instantiates a new GetPayTransactionsV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayTransactionsV1RespWithDefaults

`func NewGetPayTransactionsV1RespWithDefaults() *GetPayTransactionsV1Resp`

NewGetPayTransactionsV1RespWithDefaults instantiates a new GetPayTransactionsV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetPayTransactionsV1Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetPayTransactionsV1Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetPayTransactionsV1Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetPayTransactionsV1Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetPayTransactionsV1Resp) GetData() []GetPayTransactionsV1RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPayTransactionsV1Resp) GetDataOk() (*[]GetPayTransactionsV1RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPayTransactionsV1Resp) SetData(v []GetPayTransactionsV1RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetPayTransactionsV1Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *GetPayTransactionsV1Resp) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetPayTransactionsV1Resp) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetPayTransactionsV1Resp) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetPayTransactionsV1Resp) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *GetPayTransactionsV1Resp) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetPayTransactionsV1Resp) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetPayTransactionsV1Resp) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetPayTransactionsV1Resp) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


