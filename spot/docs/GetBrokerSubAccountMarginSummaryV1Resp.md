# GetBrokerSubAccountMarginSummaryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetBrokerSubAccountMarginSummaryV1RespDataInner**](GetBrokerSubAccountMarginSummaryV1RespDataInner.md) |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetBrokerSubAccountMarginSummaryV1Resp

`func NewGetBrokerSubAccountMarginSummaryV1Resp() *GetBrokerSubAccountMarginSummaryV1Resp`

NewGetBrokerSubAccountMarginSummaryV1Resp instantiates a new GetBrokerSubAccountMarginSummaryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBrokerSubAccountMarginSummaryV1RespWithDefaults

`func NewGetBrokerSubAccountMarginSummaryV1RespWithDefaults() *GetBrokerSubAccountMarginSummaryV1Resp`

NewGetBrokerSubAccountMarginSummaryV1RespWithDefaults instantiates a new GetBrokerSubAccountMarginSummaryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetBrokerSubAccountMarginSummaryV1Resp) GetData() []GetBrokerSubAccountMarginSummaryV1RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBrokerSubAccountMarginSummaryV1Resp) GetDataOk() (*[]GetBrokerSubAccountMarginSummaryV1RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBrokerSubAccountMarginSummaryV1Resp) SetData(v []GetBrokerSubAccountMarginSummaryV1RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetBrokerSubAccountMarginSummaryV1Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetBrokerSubAccountMarginSummaryV1Resp) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetBrokerSubAccountMarginSummaryV1Resp) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetBrokerSubAccountMarginSummaryV1Resp) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetBrokerSubAccountMarginSummaryV1Resp) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


