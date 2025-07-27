# GetBrokerSubAccountSpotSummaryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetBrokerSubAccountSpotSummaryV1RespDataInner**](GetBrokerSubAccountSpotSummaryV1RespDataInner.md) |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetBrokerSubAccountSpotSummaryV1Resp

`func NewGetBrokerSubAccountSpotSummaryV1Resp() *GetBrokerSubAccountSpotSummaryV1Resp`

NewGetBrokerSubAccountSpotSummaryV1Resp instantiates a new GetBrokerSubAccountSpotSummaryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBrokerSubAccountSpotSummaryV1RespWithDefaults

`func NewGetBrokerSubAccountSpotSummaryV1RespWithDefaults() *GetBrokerSubAccountSpotSummaryV1Resp`

NewGetBrokerSubAccountSpotSummaryV1RespWithDefaults instantiates a new GetBrokerSubAccountSpotSummaryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetBrokerSubAccountSpotSummaryV1Resp) GetData() []GetBrokerSubAccountSpotSummaryV1RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBrokerSubAccountSpotSummaryV1Resp) GetDataOk() (*[]GetBrokerSubAccountSpotSummaryV1RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBrokerSubAccountSpotSummaryV1Resp) SetData(v []GetBrokerSubAccountSpotSummaryV1RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetBrokerSubAccountSpotSummaryV1Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetBrokerSubAccountSpotSummaryV1Resp) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetBrokerSubAccountSpotSummaryV1Resp) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetBrokerSubAccountSpotSummaryV1Resp) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetBrokerSubAccountSpotSummaryV1Resp) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


