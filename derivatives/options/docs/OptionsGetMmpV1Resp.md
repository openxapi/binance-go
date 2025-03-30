# OptionsGetMmpV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeltaLimit** | Pointer to **string** |  | [optional] 
**FrozenTimeInMilliseconds** | Pointer to **int32** |  | [optional] 
**LastTriggerTime** | Pointer to **int64** |  | [optional] 
**QtyLimit** | Pointer to **string** |  | [optional] 
**Underlying** | Pointer to **string** |  | [optional] 
**UnderlyingId** | Pointer to **int64** |  | [optional] 
**WindowTimeInMilliseconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewOptionsGetMmpV1Resp

`func NewOptionsGetMmpV1Resp() *OptionsGetMmpV1Resp`

NewOptionsGetMmpV1Resp instantiates a new OptionsGetMmpV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsGetMmpV1RespWithDefaults

`func NewOptionsGetMmpV1RespWithDefaults() *OptionsGetMmpV1Resp`

NewOptionsGetMmpV1RespWithDefaults instantiates a new OptionsGetMmpV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeltaLimit

`func (o *OptionsGetMmpV1Resp) GetDeltaLimit() string`

GetDeltaLimit returns the DeltaLimit field if non-nil, zero value otherwise.

### GetDeltaLimitOk

`func (o *OptionsGetMmpV1Resp) GetDeltaLimitOk() (*string, bool)`

GetDeltaLimitOk returns a tuple with the DeltaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaLimit

`func (o *OptionsGetMmpV1Resp) SetDeltaLimit(v string)`

SetDeltaLimit sets DeltaLimit field to given value.

### HasDeltaLimit

`func (o *OptionsGetMmpV1Resp) HasDeltaLimit() bool`

HasDeltaLimit returns a boolean if a field has been set.

### GetFrozenTimeInMilliseconds

`func (o *OptionsGetMmpV1Resp) GetFrozenTimeInMilliseconds() int32`

GetFrozenTimeInMilliseconds returns the FrozenTimeInMilliseconds field if non-nil, zero value otherwise.

### GetFrozenTimeInMillisecondsOk

`func (o *OptionsGetMmpV1Resp) GetFrozenTimeInMillisecondsOk() (*int32, bool)`

GetFrozenTimeInMillisecondsOk returns a tuple with the FrozenTimeInMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenTimeInMilliseconds

`func (o *OptionsGetMmpV1Resp) SetFrozenTimeInMilliseconds(v int32)`

SetFrozenTimeInMilliseconds sets FrozenTimeInMilliseconds field to given value.

### HasFrozenTimeInMilliseconds

`func (o *OptionsGetMmpV1Resp) HasFrozenTimeInMilliseconds() bool`

HasFrozenTimeInMilliseconds returns a boolean if a field has been set.

### GetLastTriggerTime

`func (o *OptionsGetMmpV1Resp) GetLastTriggerTime() int64`

GetLastTriggerTime returns the LastTriggerTime field if non-nil, zero value otherwise.

### GetLastTriggerTimeOk

`func (o *OptionsGetMmpV1Resp) GetLastTriggerTimeOk() (*int64, bool)`

GetLastTriggerTimeOk returns a tuple with the LastTriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggerTime

`func (o *OptionsGetMmpV1Resp) SetLastTriggerTime(v int64)`

SetLastTriggerTime sets LastTriggerTime field to given value.

### HasLastTriggerTime

`func (o *OptionsGetMmpV1Resp) HasLastTriggerTime() bool`

HasLastTriggerTime returns a boolean if a field has been set.

### GetQtyLimit

`func (o *OptionsGetMmpV1Resp) GetQtyLimit() string`

GetQtyLimit returns the QtyLimit field if non-nil, zero value otherwise.

### GetQtyLimitOk

`func (o *OptionsGetMmpV1Resp) GetQtyLimitOk() (*string, bool)`

GetQtyLimitOk returns a tuple with the QtyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyLimit

`func (o *OptionsGetMmpV1Resp) SetQtyLimit(v string)`

SetQtyLimit sets QtyLimit field to given value.

### HasQtyLimit

`func (o *OptionsGetMmpV1Resp) HasQtyLimit() bool`

HasQtyLimit returns a boolean if a field has been set.

### GetUnderlying

`func (o *OptionsGetMmpV1Resp) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *OptionsGetMmpV1Resp) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *OptionsGetMmpV1Resp) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *OptionsGetMmpV1Resp) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetUnderlyingId

`func (o *OptionsGetMmpV1Resp) GetUnderlyingId() int64`

GetUnderlyingId returns the UnderlyingId field if non-nil, zero value otherwise.

### GetUnderlyingIdOk

`func (o *OptionsGetMmpV1Resp) GetUnderlyingIdOk() (*int64, bool)`

GetUnderlyingIdOk returns a tuple with the UnderlyingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingId

`func (o *OptionsGetMmpV1Resp) SetUnderlyingId(v int64)`

SetUnderlyingId sets UnderlyingId field to given value.

### HasUnderlyingId

`func (o *OptionsGetMmpV1Resp) HasUnderlyingId() bool`

HasUnderlyingId returns a boolean if a field has been set.

### GetWindowTimeInMilliseconds

`func (o *OptionsGetMmpV1Resp) GetWindowTimeInMilliseconds() int32`

GetWindowTimeInMilliseconds returns the WindowTimeInMilliseconds field if non-nil, zero value otherwise.

### GetWindowTimeInMillisecondsOk

`func (o *OptionsGetMmpV1Resp) GetWindowTimeInMillisecondsOk() (*int32, bool)`

GetWindowTimeInMillisecondsOk returns a tuple with the WindowTimeInMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowTimeInMilliseconds

`func (o *OptionsGetMmpV1Resp) SetWindowTimeInMilliseconds(v int32)`

SetWindowTimeInMilliseconds sets WindowTimeInMilliseconds field to given value.

### HasWindowTimeInMilliseconds

`func (o *OptionsGetMmpV1Resp) HasWindowTimeInMilliseconds() bool`

HasWindowTimeInMilliseconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


