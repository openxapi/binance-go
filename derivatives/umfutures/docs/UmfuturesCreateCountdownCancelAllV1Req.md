# UmfuturesCreateCountdownCancelAllV1Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountdownTime** | **int64** |  | 
**RecvWindow** | Pointer to **int64** |  | [optional] 
**Symbol** | **string** |  | [default to ""]
**Timestamp** | **int64** |  | 

## Methods

### NewUmfuturesCreateCountdownCancelAllV1Req

`func NewUmfuturesCreateCountdownCancelAllV1Req(countdownTime int64, symbol string, timestamp int64, ) *UmfuturesCreateCountdownCancelAllV1Req`

NewUmfuturesCreateCountdownCancelAllV1Req instantiates a new UmfuturesCreateCountdownCancelAllV1Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesCreateCountdownCancelAllV1ReqWithDefaults

`func NewUmfuturesCreateCountdownCancelAllV1ReqWithDefaults() *UmfuturesCreateCountdownCancelAllV1Req`

NewUmfuturesCreateCountdownCancelAllV1ReqWithDefaults instantiates a new UmfuturesCreateCountdownCancelAllV1Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountdownTime

`func (o *UmfuturesCreateCountdownCancelAllV1Req) GetCountdownTime() int64`

GetCountdownTime returns the CountdownTime field if non-nil, zero value otherwise.

### GetCountdownTimeOk

`func (o *UmfuturesCreateCountdownCancelAllV1Req) GetCountdownTimeOk() (*int64, bool)`

GetCountdownTimeOk returns a tuple with the CountdownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountdownTime

`func (o *UmfuturesCreateCountdownCancelAllV1Req) SetCountdownTime(v int64)`

SetCountdownTime sets CountdownTime field to given value.


### GetRecvWindow

`func (o *UmfuturesCreateCountdownCancelAllV1Req) GetRecvWindow() int64`

GetRecvWindow returns the RecvWindow field if non-nil, zero value otherwise.

### GetRecvWindowOk

`func (o *UmfuturesCreateCountdownCancelAllV1Req) GetRecvWindowOk() (*int64, bool)`

GetRecvWindowOk returns a tuple with the RecvWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvWindow

`func (o *UmfuturesCreateCountdownCancelAllV1Req) SetRecvWindow(v int64)`

SetRecvWindow sets RecvWindow field to given value.

### HasRecvWindow

`func (o *UmfuturesCreateCountdownCancelAllV1Req) HasRecvWindow() bool`

HasRecvWindow returns a boolean if a field has been set.

### GetSymbol

`func (o *UmfuturesCreateCountdownCancelAllV1Req) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UmfuturesCreateCountdownCancelAllV1Req) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UmfuturesCreateCountdownCancelAllV1Req) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTimestamp

`func (o *UmfuturesCreateCountdownCancelAllV1Req) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UmfuturesCreateCountdownCancelAllV1Req) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UmfuturesCreateCountdownCancelAllV1Req) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


