# GetBrokerInfoV1Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecvWindow** | Pointer to **int64** |  | [optional] 
**Timestamp** | **int64** |  | 

## Methods

### NewGetBrokerInfoV1Req

`func NewGetBrokerInfoV1Req(timestamp int64, ) *GetBrokerInfoV1Req`

NewGetBrokerInfoV1Req instantiates a new GetBrokerInfoV1Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBrokerInfoV1ReqWithDefaults

`func NewGetBrokerInfoV1ReqWithDefaults() *GetBrokerInfoV1Req`

NewGetBrokerInfoV1ReqWithDefaults instantiates a new GetBrokerInfoV1Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecvWindow

`func (o *GetBrokerInfoV1Req) GetRecvWindow() int64`

GetRecvWindow returns the RecvWindow field if non-nil, zero value otherwise.

### GetRecvWindowOk

`func (o *GetBrokerInfoV1Req) GetRecvWindowOk() (*int64, bool)`

GetRecvWindowOk returns a tuple with the RecvWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvWindow

`func (o *GetBrokerInfoV1Req) SetRecvWindow(v int64)`

SetRecvWindow sets RecvWindow field to given value.

### HasRecvWindow

`func (o *GetBrokerInfoV1Req) HasRecvWindow() bool`

HasRecvWindow returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetBrokerInfoV1Req) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetBrokerInfoV1Req) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetBrokerInfoV1Req) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


