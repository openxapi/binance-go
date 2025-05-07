# GetEthStakingEthHistoryRewardsHistoryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetEthStakingEthHistoryRewardsHistoryV1RespRowsInner**](GetEthStakingEthHistoryRewardsHistoryV1RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetEthStakingEthHistoryRewardsHistoryV1Resp

`func NewGetEthStakingEthHistoryRewardsHistoryV1Resp() *GetEthStakingEthHistoryRewardsHistoryV1Resp`

NewGetEthStakingEthHistoryRewardsHistoryV1Resp instantiates a new GetEthStakingEthHistoryRewardsHistoryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthStakingEthHistoryRewardsHistoryV1RespWithDefaults

`func NewGetEthStakingEthHistoryRewardsHistoryV1RespWithDefaults() *GetEthStakingEthHistoryRewardsHistoryV1Resp`

NewGetEthStakingEthHistoryRewardsHistoryV1RespWithDefaults instantiates a new GetEthStakingEthHistoryRewardsHistoryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetEthStakingEthHistoryRewardsHistoryV1Resp) GetRows() []GetEthStakingEthHistoryRewardsHistoryV1RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetEthStakingEthHistoryRewardsHistoryV1Resp) GetRowsOk() (*[]GetEthStakingEthHistoryRewardsHistoryV1RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetEthStakingEthHistoryRewardsHistoryV1Resp) SetRows(v []GetEthStakingEthHistoryRewardsHistoryV1RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetEthStakingEthHistoryRewardsHistoryV1Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetEthStakingEthHistoryRewardsHistoryV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetEthStakingEthHistoryRewardsHistoryV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetEthStakingEthHistoryRewardsHistoryV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetEthStakingEthHistoryRewardsHistoryV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


