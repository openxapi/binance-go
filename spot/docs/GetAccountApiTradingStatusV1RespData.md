# GetAccountApiTradingStatusV1RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLocked** | Pointer to **bool** |  | [optional] 
**PlannedRecoverTime** | Pointer to **int64** |  | [optional] 
**TriggerCondition** | Pointer to [**GetAccountApiTradingStatusV1RespDataTriggerCondition**](GetAccountApiTradingStatusV1RespDataTriggerCondition.md) |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetAccountApiTradingStatusV1RespData

`func NewGetAccountApiTradingStatusV1RespData() *GetAccountApiTradingStatusV1RespData`

NewGetAccountApiTradingStatusV1RespData instantiates a new GetAccountApiTradingStatusV1RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountApiTradingStatusV1RespDataWithDefaults

`func NewGetAccountApiTradingStatusV1RespDataWithDefaults() *GetAccountApiTradingStatusV1RespData`

NewGetAccountApiTradingStatusV1RespDataWithDefaults instantiates a new GetAccountApiTradingStatusV1RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLocked

`func (o *GetAccountApiTradingStatusV1RespData) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *GetAccountApiTradingStatusV1RespData) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *GetAccountApiTradingStatusV1RespData) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *GetAccountApiTradingStatusV1RespData) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetPlannedRecoverTime

`func (o *GetAccountApiTradingStatusV1RespData) GetPlannedRecoverTime() int64`

GetPlannedRecoverTime returns the PlannedRecoverTime field if non-nil, zero value otherwise.

### GetPlannedRecoverTimeOk

`func (o *GetAccountApiTradingStatusV1RespData) GetPlannedRecoverTimeOk() (*int64, bool)`

GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRecoverTime

`func (o *GetAccountApiTradingStatusV1RespData) SetPlannedRecoverTime(v int64)`

SetPlannedRecoverTime sets PlannedRecoverTime field to given value.

### HasPlannedRecoverTime

`func (o *GetAccountApiTradingStatusV1RespData) HasPlannedRecoverTime() bool`

HasPlannedRecoverTime returns a boolean if a field has been set.

### GetTriggerCondition

`func (o *GetAccountApiTradingStatusV1RespData) GetTriggerCondition() GetAccountApiTradingStatusV1RespDataTriggerCondition`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *GetAccountApiTradingStatusV1RespData) GetTriggerConditionOk() (*GetAccountApiTradingStatusV1RespDataTriggerCondition, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *GetAccountApiTradingStatusV1RespData) SetTriggerCondition(v GetAccountApiTradingStatusV1RespDataTriggerCondition)`

SetTriggerCondition sets TriggerCondition field to given value.

### HasTriggerCondition

`func (o *GetAccountApiTradingStatusV1RespData) HasTriggerCondition() bool`

HasTriggerCondition returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetAccountApiTradingStatusV1RespData) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetAccountApiTradingStatusV1RespData) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetAccountApiTradingStatusV1RespData) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetAccountApiTradingStatusV1RespData) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


