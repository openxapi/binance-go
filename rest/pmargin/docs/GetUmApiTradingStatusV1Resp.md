# GetUmApiTradingStatusV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indicators** | Pointer to [**GetUmApiTradingStatusV1RespIndicators**](GetUmApiTradingStatusV1RespIndicators.md) |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetUmApiTradingStatusV1Resp

`func NewGetUmApiTradingStatusV1Resp() *GetUmApiTradingStatusV1Resp`

NewGetUmApiTradingStatusV1Resp instantiates a new GetUmApiTradingStatusV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmApiTradingStatusV1RespWithDefaults

`func NewGetUmApiTradingStatusV1RespWithDefaults() *GetUmApiTradingStatusV1Resp`

NewGetUmApiTradingStatusV1RespWithDefaults instantiates a new GetUmApiTradingStatusV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndicators

`func (o *GetUmApiTradingStatusV1Resp) GetIndicators() GetUmApiTradingStatusV1RespIndicators`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *GetUmApiTradingStatusV1Resp) GetIndicatorsOk() (*GetUmApiTradingStatusV1RespIndicators, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *GetUmApiTradingStatusV1Resp) SetIndicators(v GetUmApiTradingStatusV1RespIndicators)`

SetIndicators sets Indicators field to given value.

### HasIndicators

`func (o *GetUmApiTradingStatusV1Resp) HasIndicators() bool`

HasIndicators returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetUmApiTradingStatusV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetUmApiTradingStatusV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetUmApiTradingStatusV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetUmApiTradingStatusV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


