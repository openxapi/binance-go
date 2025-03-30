# UfuturesCreateCountdownCancelAllV1Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountdownTime** | **int64** |  | 
**RecvWindow** | Pointer to **int64** |  | [optional] 
**Symbol** | **string** |  | [default to ""]
**Timestamp** | **int64** |  | 

## Methods

### NewUfuturesCreateCountdownCancelAllV1Req

`func NewUfuturesCreateCountdownCancelAllV1Req(countdownTime int64, symbol string, timestamp int64, ) *UfuturesCreateCountdownCancelAllV1Req`

NewUfuturesCreateCountdownCancelAllV1Req instantiates a new UfuturesCreateCountdownCancelAllV1Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesCreateCountdownCancelAllV1ReqWithDefaults

`func NewUfuturesCreateCountdownCancelAllV1ReqWithDefaults() *UfuturesCreateCountdownCancelAllV1Req`

NewUfuturesCreateCountdownCancelAllV1ReqWithDefaults instantiates a new UfuturesCreateCountdownCancelAllV1Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountdownTime

`func (o *UfuturesCreateCountdownCancelAllV1Req) GetCountdownTime() int64`

GetCountdownTime returns the CountdownTime field if non-nil, zero value otherwise.

### GetCountdownTimeOk

`func (o *UfuturesCreateCountdownCancelAllV1Req) GetCountdownTimeOk() (*int64, bool)`

GetCountdownTimeOk returns a tuple with the CountdownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountdownTime

`func (o *UfuturesCreateCountdownCancelAllV1Req) SetCountdownTime(v int64)`

SetCountdownTime sets CountdownTime field to given value.


### GetRecvWindow

`func (o *UfuturesCreateCountdownCancelAllV1Req) GetRecvWindow() int64`

GetRecvWindow returns the RecvWindow field if non-nil, zero value otherwise.

### GetRecvWindowOk

`func (o *UfuturesCreateCountdownCancelAllV1Req) GetRecvWindowOk() (*int64, bool)`

GetRecvWindowOk returns a tuple with the RecvWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvWindow

`func (o *UfuturesCreateCountdownCancelAllV1Req) SetRecvWindow(v int64)`

SetRecvWindow sets RecvWindow field to given value.

### HasRecvWindow

`func (o *UfuturesCreateCountdownCancelAllV1Req) HasRecvWindow() bool`

HasRecvWindow returns a boolean if a field has been set.

### GetSymbol

`func (o *UfuturesCreateCountdownCancelAllV1Req) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UfuturesCreateCountdownCancelAllV1Req) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UfuturesCreateCountdownCancelAllV1Req) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTimestamp

`func (o *UfuturesCreateCountdownCancelAllV1Req) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UfuturesCreateCountdownCancelAllV1Req) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UfuturesCreateCountdownCancelAllV1Req) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


