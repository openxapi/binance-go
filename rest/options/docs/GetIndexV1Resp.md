# GetIndexV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexPrice** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetIndexV1Resp

`func NewGetIndexV1Resp() *GetIndexV1Resp`

NewGetIndexV1Resp instantiates a new GetIndexV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIndexV1RespWithDefaults

`func NewGetIndexV1RespWithDefaults() *GetIndexV1Resp`

NewGetIndexV1RespWithDefaults instantiates a new GetIndexV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexPrice

`func (o *GetIndexV1Resp) GetIndexPrice() string`

GetIndexPrice returns the IndexPrice field if non-nil, zero value otherwise.

### GetIndexPriceOk

`func (o *GetIndexV1Resp) GetIndexPriceOk() (*string, bool)`

GetIndexPriceOk returns a tuple with the IndexPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPrice

`func (o *GetIndexV1Resp) SetIndexPrice(v string)`

SetIndexPrice sets IndexPrice field to given value.

### HasIndexPrice

`func (o *GetIndexV1Resp) HasIndexPrice() bool`

HasIndexPrice returns a boolean if a field has been set.

### GetTime

`func (o *GetIndexV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetIndexV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetIndexV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetIndexV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


