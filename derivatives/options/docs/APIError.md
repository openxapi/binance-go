# APIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewAPIError

`func NewAPIError() *APIError`

NewAPIError instantiates a new APIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIErrorWithDefaults

`func NewAPIErrorWithDefaults() *APIError`

NewAPIErrorWithDefaults instantiates a new APIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *APIError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *APIError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *APIError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *APIError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *APIError) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *APIError) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


