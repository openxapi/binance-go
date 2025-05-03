# GetTimeV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetTimeV3Resp

`func NewGetTimeV3Resp() *GetTimeV3Resp`

NewGetTimeV3Resp instantiates a new GetTimeV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeV3RespWithDefaults

`func NewGetTimeV3RespWithDefaults() *GetTimeV3Resp`

NewGetTimeV3RespWithDefaults instantiates a new GetTimeV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerTime

`func (o *GetTimeV3Resp) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *GetTimeV3Resp) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *GetTimeV3Resp) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *GetTimeV3Resp) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


